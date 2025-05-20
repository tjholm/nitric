package terraform

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/nitrictech/nitric/cli/pkg/schema"
	"github.com/nitrictech/nitric/engines"
)

type TerraformEngine struct {
}

// Apply the engine to the target environment
func (e *TerraformEngine) Apply(application *schema.Application, platform map[string]interface{}, environment map[string]interface{}) error {
	app := cdktf.NewApp(&cdktf.AppConfig{})

	stack := cdktf.NewTerraformStack(app, jsii.String(application.Name))

	terraformResources := map[string]cdktf.TerraformHclModule{}

	// Start deploying the platform
	for resourceName, resource := range application.Resources {
		switch resource.Type {
		case "service":
			// Deploy the service
			// Locate the plugin for the service from platform
			terraformResources[resourceName] = cdktf.NewTerraformHclModule(stack, jsii.String(resourceName), &cdktf.TerraformHclModuleConfig{})

		case "entrypoint":
			// Deploy the entrypoint
			// Locate the plugin for the entrypoint from platform
			terraformResources[resourceName] = cdktf.NewTerraformHclModule(stack, jsii.String(resourceName), &cdktf.TerraformHclModuleConfig{})
		}
	}

	// Map inputs and outputs from the platform and environment to the stack resources

	app.Synth()

	return nil
}

var _ engines.Engine = &TerraformEngine{}

func New() *TerraformEngine {
	return &TerraformEngine{}
}
