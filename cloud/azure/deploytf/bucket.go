// Copyright 2021 Nitric Technologies Pty Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package deploytf

import (
	"fmt"
	"strings"

	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/nitrictech/nitric/cloud/azure/deploytf/generated/bucket"
	deploymentspb "github.com/nitrictech/nitric/core/pkg/proto/deployments/v1"
	storagepb "github.com/nitrictech/nitric/core/pkg/proto/storage/v1"
)

type BucketSubscriber struct {
	HostUrl      string   `json:"host_url"`
	EventToken   string   `json:"event_token"`
	SpClientId   string   `json:"sp_client_id"`
	SpTenantId   string   `json:"sp_tenant_id"`
	PrefixFilter string   `json:"prefix_filter"`
	EventTypes   []string `json:"event_types"`
}

// removeWildcard - Remove the trailing wildcard from a prefix filter, they're not supported by Azure
func removeWildcard(prefixFilter string) string {
	return strings.TrimRight(prefixFilter, "*")
}

func eventTypeToStorageEventType(eventType storagepb.BlobEventType) []string {
	switch eventType {
	case storagepb.BlobEventType_Created:
		return []string{"Microsoft.Storage.BlobCreated"}
	case storagepb.BlobEventType_Deleted:
		return []string{"Microsoft.Storage.BlobDeleted"}
	default:
		return []string{}
	}
}

// Bucket - Deploy a Storage Bucket
func (a *NitricAzureTerraformProvider) Bucket(stack cdktf.TerraformStack, name string, config *deploymentspb.Bucket) error {
	subscribers := map[string]*BucketSubscriber{}
	for _, v := range config.GetListeners() {
		deployedService := a.Services[v.GetService()]

		subscribers[v.GetService()] = &BucketSubscriber{
			HostUrl:      *deployedService.HostUrlOutput(),
			EventToken:   *deployedService.EventTokenOutput(),
			SpClientId:   *deployedService.ServicePrincipalClientIdOutput(),
			SpTenantId:   *deployedService.ServicePrincipalTenantIdOutput(),
			PrefixFilter: removeWildcard(v.Config.KeyPrefixFilter),
			EventTypes:   eventTypeToStorageEventType(v.Config.GetBlobEventType()),
		}
	}

	a.Buckets[name] = bucket.NewBucket(stack, jsii.String(name), &bucket.BucketConfig{
		BucketName:         jsii.String(name),
		StorageAccountName: a.Stack.StorageAccountNameOutput(),
		BucketSubcribers:   subscribers,
	})

	return fmt.Errorf("Not implemented")
}
