package main

import (
	"bytes"
	"encoding/json"
	"log"

	coreschema "github.com/nitrictech/nitric/cli/pkg/schema"
	"github.com/nitrictech/nitric/engines/terraform"
	"github.com/nitrictech/nitric/engines/terraform/schema"
)

func main() {
	platformConfig := schema.TerraformPlatform{
		Name: "aws",
		Services: schema.TerraformPlatformResource{
			BaseTerraformResource: schema.BaseTerraformResource{
				Plugin: "nitric-aws",
				Properties: map[string]interface{}{
					"test": "${infra.vpc.id}",
				},
			},
		},
		Entrypoints: schema.TerraformPlatformResource{
			BaseTerraformResource: schema.BaseTerraformResource{
				Plugin: "nitric-aws",
				Properties: map[string]interface{}{
					"region": "us-east-1",
				},
			},
		},
		Infra: map[string]schema.BaseTerraformResource{
			"vpc": {
				Plugin: "nitric-aws",
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

	// provide a bytes reader to the terraform engine
	platform := terraform.New(bytes.NewReader(platformConfigJSON))

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
