// Copyright Nitric Pty Ltd.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
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
	"embed"
	_ "embed"
	"io/fs"

	"github.com/aws/jsii-runtime-go"
	// azureprovider "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/provider"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/nitrictech/nitric/cloud/azure/common"
	"github.com/nitrictech/nitric/cloud/azure/deploytf/generated/roles"
	azstack "github.com/nitrictech/nitric/cloud/azure/deploytf/generated/stack"
	"github.com/nitrictech/nitric/cloud/common/deploy"
	"github.com/nitrictech/nitric/cloud/common/deploy/provider"
	deploymentspb "github.com/nitrictech/nitric/core/pkg/proto/deployments/v1"
	apimanagement "github.com/pulumi/pulumi-azure-native-sdk/apimanagement"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ApiResources struct {
	ApiManagementService *apimanagement.ApiManagementService
	Api                  *apimanagement.Api
}

type NitricAzureTerraformProvider struct {
	*deploy.CommonStackDetails

	AzureConfig *common.AzureConfig

	StackId string

	Stack azstack.Stack

	Roles roles.Roles

	provider.NitricDefaultOrder
}

// embed the modules directory here
//
//go:embed .nitric/modules/**/*
var modules embed.FS

func (a *NitricAzureTerraformProvider) CdkTfModules() (string, fs.FS, error) {
	return ".nitric/modules", modules, nil
}

var _ provider.NitricTerraformProvider = (*NitricAzureTerraformProvider)(nil)

const (
	pulumiAzureNativeVersion = "2.40.0"
	pulumiAzureVersion       = "5.52.0"
)

func (a *NitricAzureTerraformProvider) Config() (auto.ConfigMap, error) {
	return auto.ConfigMap{
		"azure-native:location": auto.ConfigValue{Value: a.Region},
		"azure:location":        auto.ConfigValue{Value: a.Region},
		"azure-native:version":  auto.ConfigValue{Value: pulumiAzureNativeVersion},
		"azure:version":         auto.ConfigValue{Value: pulumiAzureVersion},
		"docker:version":        auto.ConfigValue{Value: deploy.PulumiDockerVersion},
	}, nil
}

func (a *NitricAzureTerraformProvider) Init(attributes map[string]interface{}) error {
	var err error

	a.CommonStackDetails, err = deploy.CommonStackDetailsFromAttributes(attributes)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, err.Error())
	}

	a.AzureConfig, err = common.ConfigFromAttributes(attributes)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Bad stack configuration: %s", err)
	}

	return nil
}

func (a *NitricAzureTerraformProvider) Pre(stack cdktf.TerraformStack, resources []*deploymentspb.Resource) error {
	tfRegion := cdktf.NewTerraformVariable(stack, jsii.String("region"), &cdktf.TerraformVariableConfig{
		Type:        jsii.String("string"),
		Default:     jsii.String(a.Region),
		Description: jsii.String("The AWS region to deploy resources to"),
	})

	// Deploy the stack and necessary resources
	a.Stack = azstack.NewStack(stack, jsii.String("nitric-azure-stack"), &azstack.StackConfig{
		Region:    tfRegion.StringValue(),
		StackName: jsii.String(a.StackName),
	})

	// Create the roles module
	a.Roles = roles.NewRoles(stack, jsii.String("nitric-azure-roles"), &roles.RolesConfig{})

	return nil
}

func (a *NitricAzureTerraformProvider) Post(stack cdktf.TerraformStack) error {
	return nil
}

func NewNitricAzurePulumiProvider() *NitricAzureTerraformProvider {
	return &NitricAzureTerraformProvider{}
}
