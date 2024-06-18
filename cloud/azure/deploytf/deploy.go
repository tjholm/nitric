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
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/nitrictech/nitric/cloud/azure/common"
	"github.com/nitrictech/nitric/cloud/common/deploy"
	"github.com/nitrictech/nitric/cloud/common/deploy/provider"
	"github.com/nitrictech/nitric/cloud/common/deploy/pulumix"
	deploymentspb "github.com/nitrictech/nitric/core/pkg/proto/deployments/v1"
	resourcespb "github.com/nitrictech/nitric/core/pkg/proto/resources/v1"
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

	StackId   string
	resources []*pulumix.NitricPulumiResource[any]

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

// func createKeyVault(ctx *pulumi.Context, group *resources.ResourceGroup, tenantId string, tags map[string]string) (*keyvault.Vault, error) {
// 	// Create a stack level keyvault if secrets are enabled
// 	// At the moment secrets have no config level setting
// 	vaultName := ResourceName(ctx, "", KeyVaultRT)

// 	keyVault, err := keyvault.NewVault(ctx, vaultName, &keyvault.VaultArgs{
// 		Location:          group.Location,
// 		ResourceGroupName: group.Name,
// 		Properties: &keyvault.VaultPropertiesArgs{
// 			EnableSoftDelete:        pulumi.Bool(false),
// 			EnableRbacAuthorization: pulumi.Bool(true),
// 			Sku: &keyvault.SkuArgs{
// 				Family: pulumi.String("A"),
// 				Name:   keyvault.SkuNameStandard,
// 			},
// 			TenantId: pulumi.String(tenantId),
// 		},
// 		Tags: pulumi.ToStringMap(tags),
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return keyVault, nil
// }

// func createStorageAccount(ctx *pulumi.Context, group *resources.ResourceGroup, tags map[string]string) (*storage.StorageAccount, error) {
// 	accName := ResourceName(ctx, "", StorageAccountRT)
// 	storageAccount, err := storage.NewStorageAccount(ctx, accName, &storage.StorageAccountArgs{
// 		AccessTier:        storage.AccessTierHot,
// 		ResourceGroupName: group.Name,
// 		Kind:              pulumi.String("StorageV2"),
// 		Sku: storage.SkuArgs{
// 			Name: pulumi.String(storage.SkuName_Standard_LRS),
// 		},
// 		Tags: pulumi.ToStringMap(tags),
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return storageAccount, nil
// }

func hasResourceType(resources []*pulumix.NitricPulumiResource[any], resourceType resourcespb.ResourceType) bool {
	for _, r := range resources {
		if r.Id.GetType() == resourceType {
			return true
		}
	}

	return false
}

func (a *NitricAzureTerraformProvider) Pre(stack cdktf.TerraformStack, resources []*deploymentspb.Resource) error {
	tfRegion := cdktf.NewTerraformVariable(stack, jsii.String("region"), &cdktf.TerraformVariableConfig{
		Type:        jsii.String("string"),
		Default:     jsii.String(a.Region),
		Description: jsii.String("The AWS region to deploy resources to"),
	})

	return nil
}

func (a *NitricAzureTerraformProvider) Post(stack cdktf.TerraformStack) error {
	return nil
}

func NewNitricAzurePulumiProvider() *NitricAzureTerraformProvider {
	return &NitricAzureTerraformProvider{}
}
