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

package service

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServiceConfig struct {
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Providers *[]interface{} `field:"optional" json:"providers" yaml:"providers"`
	// Experimental.
	SkipAssetCreationFromLocalModules *bool `field:"optional" json:"skipAssetCreationFromLocalModules" yaml:"skipAssetCreationFromLocalModules"`
	// The client ID of the application for which to create this services service principal.
	ApplicationClientId *string `field:"required" json:"applicationClientId" yaml:"applicationClientId"`
	// The client secret of the application for which to create this services service principal.
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// The ID of the container app environment.
	ContainerAppEnvironmentId *string `field:"required" json:"containerAppEnvironmentId" yaml:"containerAppEnvironmentId"`
	// The name of the service.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The password of the container registry.
	RegistryPassword *string `field:"required" json:"registryPassword" yaml:"registryPassword"`
	// The server of the container registry.
	RegistryServer *string `field:"required" json:"registryServer" yaml:"registryServer"`
	// The username of the container registry.
	RegistryUsername *string `field:"required" json:"registryUsername" yaml:"registryUsername"`
	// The name of the resource group.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// The tenant ID of the application for which to create this services service principal.
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
}
