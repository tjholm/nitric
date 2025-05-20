package terraform

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/nitrictech/nitric/cli/pkg/schema"
	"github.com/nitrictech/nitric/engines"
)

type TerraformEngine struct {
	platform interface{}
}

// Apply the engine to the target environment
func (e *TerraformEngine) Apply(application *schema.Application, environment map[string]interface{}) error {
	app := cdktf.NewApp(&cdktf.AppConfig{})

	stack := cdktf.NewTerraformStack(app, jsii.String(application.Name))
	terraformResources := map[string]cdktf.TerraformHclModule{}

	// 1. Start deploying the platform
	for resourceName, resource := range application.Resources {
		switch resource.Type {
		case "service":
			// Locate the plugin for the service from platform
			// Deploy the service
			terraformResources[resourceName] = cdktf.NewTerraformHclModule(stack, jsii.String(resourceName), &cdktf.TerraformHclModuleConfig{})

		case "entrypoint":
			// Locate the plugin for the entrypoint from platform
			// Deploy the entrypoint
			terraformResources[resourceName] = cdktf.NewTerraformHclModule(stack, jsii.String(resourceName), &cdktf.TerraformHclModuleConfig{})
		}
	}

	// 2. Deploy platform infra resources

	// 3. Map inputs and outputs from the platform and environment to the stack resources

	app.Synth()

	return nil
}

var _ engines.Engine = &TerraformEngine{}

func New(platform interface{}) *TerraformEngine {
	return &TerraformEngine{
		platform: platform,
	}
}
