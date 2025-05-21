package main

import (
	"bytes"
	"encoding/json"
	"log"

	coreschema "github.com/nitrictech/nitric/cli/pkg/schema"
	"github.com/nitrictech/nitric/engines/terraform"
	"github.com/nitrictech/nitric/engines/terraform/schema"
)

type MockTerraformPluginRepository struct {
	plugins map[string]*schema.TerraformPluginManifest
}

func (r *MockTerraformPluginRepository) GetPlugin(name string) (*schema.TerraformPluginManifest, error) {
	return r.plugins[name], nil
}

func createMockTerraformPluginRepository() *MockTerraformPluginRepository {
	return &MockTerraformPluginRepository{
		plugins: map[string]*schema.TerraformPluginManifest{
			"nitric-aws-lambda": {
				Name: "nitric-aws-lambda",
				Deployment: schema.TerraformDeploymentModule{
					Terraform: "terraform-aws-modules/lambda/aws",
				},
			},
			"nitric-aws-cloudfront": {
				Name: "nitric-aws-cloudfront",
				Deployment: schema.TerraformDeploymentModule{
					Terraform: "terraform-aws-modules/cloudfront/aws",
				},
			},
			"nitric-aws-vpc": {
				Name: "nitric-aws-vpc",
				Deployment: schema.TerraformDeploymentModule{
					Terraform: "terraform-aws-modules/vpc/aws",
				},
			},
		},
	}
}

func main() {
	platformConfig := schema.TerraformPlatform{
		Name: "aws",
		Services: schema.TerraformPlatformResource{
			BaseTerraformResource: schema.BaseTerraformResource{
				Plugin: "nitric-aws-lambda",
				Properties: map[string]interface{}{
					"test": "${infra.vpc.id}",
				},
			},
		},
		Entrypoints: schema.TerraformPlatformResource{
			BaseTerraformResource: schema.BaseTerraformResource{
				Plugin: "nitric-aws-cloudfront",
				Properties: map[string]interface{}{
					"region": "us-east-1",
				},
			},
		},
		Infra: map[string]schema.BaseTerraformResource{
			"vpc": {
				Plugin: "nitric-aws-vpc",
				Properties: map[string]interface{}{
					"region": "us-east-1",
				},
			},
		},
	}

	// serialize the platform config to json
	platformConfigJSON, err := json.Marshal(platformConfig)
	if err != nil {
		log.Fatalf("failed to marshal platform config: %v", err)
	}

	mockRepository := createMockTerraformPluginRepository()

	// provide a bytes reader to the terraform engine
	platform := terraform.New(bytes.NewReader(platformConfigJSON), terraform.WithRepository(mockRepository))

	err = platform.Apply(&coreschema.Application{
		Name: "test",
		Resources: map[string]coreschema.Resource{
			"service": {
				Type: "service",
			},
		},
	}, map[string]interface{}{})

	if err != nil {
		log.Fatalf("failed to apply platform: %v", err)
	}
}
