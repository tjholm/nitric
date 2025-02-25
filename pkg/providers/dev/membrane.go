// Copyright 2021 Nitric Pty Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nitrictech/nitric/pkg/membrane"
	boltdb_service "github.com/nitrictech/nitric/pkg/plugins/document/boltdb"
	events_service "github.com/nitrictech/nitric/pkg/plugins/events/dev"
	gateway_plugin "github.com/nitrictech/nitric/pkg/plugins/gateway/dev"
	queue_service "github.com/nitrictech/nitric/pkg/plugins/queue/dev"
	secret_service "github.com/nitrictech/nitric/pkg/plugins/secret/dev"
	minio_storage_service "github.com/nitrictech/nitric/pkg/plugins/storage/minio"
)

func main() {
	// Setup signal interrupt handling for graceful shutdown
	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	signal.Notify(term, os.Interrupt, syscall.SIGINT)

	membraneOpts := membrane.DefaultMembraneOptions()

	membraneOpts.SecretPlugin, _ = secret_service.New()
	membraneOpts.DocumentPlugin, _ = boltdb_service.New()
	membraneOpts.EventsPlugin, _ = events_service.New()
	membraneOpts.GatewayPlugin, _ = gateway_plugin.New()
	membraneOpts.QueuePlugin, _ = queue_service.New()
	membraneOpts.StoragePlugin, _ = minio_storage_service.New()

	m, err := membrane.New(membraneOpts)

	if err != nil {
		log.Fatalf("There was an error initialising the membraneServer server: %v", err)
	}

	errChan := make(chan error)
	// Start the Membrane server
	go func(chan error) {
		errChan <- m.Start()
	}(errChan)

	select {
	case membraneError := <-errChan:
		log.Default().Println(fmt.Sprintf("Membrane Error: %v, exiting", membraneError))
	case sigTerm := <-term:
		log.Default().Println(fmt.Sprintf("Received %v, exiting", sigTerm))
	}

	m.Stop()
}
