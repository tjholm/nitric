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
	"github.com/nitrictech/nitric/pkg/plugins/document"
	boltdb_service "github.com/nitrictech/nitric/pkg/plugins/document/boltdb"
	"github.com/nitrictech/nitric/pkg/plugins/events"
	events_service "github.com/nitrictech/nitric/pkg/plugins/events/dev"
	"github.com/nitrictech/nitric/pkg/plugins/gateway"
	gateway_plugin "github.com/nitrictech/nitric/pkg/plugins/gateway/dev"
	"github.com/nitrictech/nitric/pkg/plugins/queue"
	queue_service "github.com/nitrictech/nitric/pkg/plugins/queue/dev"
	"github.com/nitrictech/nitric/pkg/plugins/storage"
	minio_storage_service "github.com/nitrictech/nitric/pkg/plugins/storage/minio"
)

type DevServiceFactory struct {
}

// NewDocumentService - Returns local dev document plugin
func (p *DevServiceFactory) NewDocumentService() (document.DocumentService, error) {
	return boltdb_service.New()
}

// NewEventService - Returns local dev events plugin
func (p *DevServiceFactory) NewEventService() (events.EventService, error) {
	return events_service.New()
}

// NewGatewayService - Returns local dev Gateway plugin
func (p *DevServiceFactory) NewGatewayService() (gateway.GatewayService, error) {
	return gateway_plugin.New()
}

// NewQueueService - Returns local dev queue plugin
func (p *DevServiceFactory) NewQueueService() (queue.QueueService, error) {
	return queue_service.New()
}

// NewStorageService - Returns local dev storage plugin
func (p *DevServiceFactory) NewStorageService() (storage.StorageService, error) {
	return minio_storage_service.New()
}
