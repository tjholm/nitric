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

package api

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiConfig struct {
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Providers *[]interface{} `field:"optional" json:"providers" yaml:"providers"`
	// Experimental.
	SkipAssetCreationFromLocalModules *bool `field:"optional" json:"skipAssetCreationFromLocalModules" yaml:"skipAssetCreationFromLocalModules"`
	// The location of the API Gateway.
	Location *string `field:"required" json:"location" yaml:"location"`
	// The name of the API Gateway.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The email of the publisher.
	PublisherEmail *string `field:"required" json:"publisherEmail" yaml:"publisherEmail"`
	// The name of the publisher.
	PublisherName *string `field:"required" json:"publisherName" yaml:"publisherName"`
	// The name of the resource group.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Open API spec.
	Spec *string `field:"required" json:"spec" yaml:"spec"`
	// The ID of the stack.
	StackId *string `field:"required" json:"stackId" yaml:"stackId"`
}
