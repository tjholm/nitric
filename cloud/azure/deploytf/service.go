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
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/nitrictech/nitric/cloud/azure/deploytf/generated/service"
	"github.com/nitrictech/nitric/cloud/common/deploy/provider"
	deploymentspb "github.com/nitrictech/nitric/core/pkg/proto/deployments/v1"
)

// Service - Deploy an service (Service)
func (a *NitricAzureTerraformProvider) Service(stack cdktf.TerraformStack, name string, config *deploymentspb.Service, runtimeProvider provider.RuntimeProvider) error {
	a.Services[name] = service.NewService(stack, jsii.String(name), &service.ServiceConfig{
		Name:                      jsii.String(name),
		ApplicationClientId:       a.Stack.WebhookApplicationIdOutput(),
		ContainerAppEnvironmentId: a.Stack.ContainerAppEnvironmentIdOutput(),
		ResourceGroupName:         a.Stack.ResourceGroupNameOutput(),
		RegistryServer:            a.Stack.RegistryLoginServerOutput(),
		RegistryUsername:          a.Stack.RegistryUserNameOutput(),
		RegistryPassword:          a.Stack.RegistryPasswordOutput(),
	})

	return nil
}
