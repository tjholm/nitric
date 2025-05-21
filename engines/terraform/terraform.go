package terraform

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	coreschema "github.com/nitrictech/nitric/cli/pkg/schema"
	"github.com/nitrictech/nitric/engines"
	"github.com/nitrictech/nitric/engines/terraform/schema"
)

type TerraformEngine struct {
	platform   *schema.TerraformPlatform
	repository TerraformPluginRepository
}

// func resolvePlugin(pluginName string) (*schema.TerraformPluginManifest, error) {
// 	// return nil, fmt.Errorf("plugin %s not found", pluginName)

// 	// just resolve a known plugin for now
// 	return &schema.TerraformPluginManifest{
// 		Deployment: schema.TerraformDeploymentModule{
// 			Terraform: "terraform-aws-modules/s3-bucket/aws",
// 		},
// 	}, nil
// }

func resolvePluginName(resource schema.TerraformPlatformResource, subtype string) (string, error) {
	plugin := resource.Plugin
	if subtype != "" {
		if _, ok := resource.Subtypes[subtype]; !ok {
			return "", fmt.Errorf("subtype %s not found", subtype)
		}

		plugin = resource.Subtypes[subtype].Plugin
	}

	// resolve the plugins manifest and locate the deployment module
	return plugin, nil
}

func (e *TerraformEngine) getPlatformResourceForType(resourceType string) (schema.TerraformPlatformResource, error) {
	switch resourceType {
	case "service":
		return e.platform.Services, nil
	case "entrypoint":
		return e.platform.Entrypoints, nil
	}
	return schema.TerraformPlatformResource{}, fmt.Errorf("resource type %s not found", resourceType)
}

func (e *TerraformEngine) getPlatformResourceProperties() map[string]map[string]interface{} {
	propertyMappings := map[string]map[string]interface{}{}

	propertyMappings["service"] = e.platform.Services.Properties
	for _, subtype := range e.platform.Services.Subtypes {
		propertyMappings[fmt.Sprintf("service.%s", subtype)] = subtype.Properties
	}

	propertyMappings["entrypoint"] = e.platform.Entrypoints.Properties
	for _, subtype := range e.platform.Entrypoints.Subtypes {
		propertyMappings[fmt.Sprintf("entrypoint.%s", subtype)] = subtype.Properties
	}

	return propertyMappings
}

// extractTokenContents extracts the contents between ${} from a token string
func extractTokenContents(token string) (string, bool) {
	if matches := tokenPattern.FindStringSubmatch(token); len(matches) == 2 {
		return matches[1], true
	}
	return "", false
}

var tokenPattern = regexp.MustCompile(`^\${([^}]+)}$`)

// Apply the engine to the target environment
func (e *TerraformEngine) Apply(application *coreschema.Application, environment map[string]interface{}) error {
	app := cdktf.NewApp(&cdktf.AppConfig{})

	stack := cdktf.NewTerraformStack(app, jsii.String(application.Name))
	terraformResources := map[string]cdktf.TerraformHclModule{}
	terraformInfraResources := map[string]cdktf.TerraformHclModule{}

	// 1. Start deploying the platform
	for resourceName, resource := range application.Resources {
		terraformPlatformResource, err := e.getPlatformResourceForType(resource.Type)
		if err != nil {
			return err
		}

		pluginName, err := resolvePluginName(terraformPlatformResource, resource.SubType)
		if err != nil {
			return err
		}

		plugin, err := e.repository.GetPlugin(pluginName)
		if err != nil {
			return err
		}

		terraformResources[resourceName] = cdktf.NewTerraformHclModule(stack, jsii.String(resourceName), &cdktf.TerraformHclModuleConfig{
			// This assumes that the plugin is resolvable as a URI
			Source: jsii.String(plugin.Deployment.Terraform),
		})
	}

	// 2. Deploy platform infra resources
	for infraName, infra := range e.platform.Infra {
		// Locate the plugin for the infra from platform
		plugin, err := e.repository.GetPlugin(infra.Plugin)
		if err != nil {
			return err
		}

		terraformInfraResources[infraName] = cdktf.NewTerraformHclModule(stack, jsii.String(infraName), &cdktf.TerraformHclModuleConfig{
			// This assumes that the plugin is resolvable as a URI
			Source: jsii.String(plugin.Deployment.Terraform),
		})
	}

	// 3. Map inputs and outputs from the platform and environment to the stack resources
	resourceProperties := e.getPlatformResourceProperties()

	fmt.Println("resourceProperties", resourceProperties)

	for resourceName, resource := range application.Resources {
		// get its plugin properties
		pluginProperties := resourceProperties[resource.Type]
		if resource.SubType != "" {
			pluginProperties = resourceProperties[fmt.Sprintf("%s.%s", resource.Type, resource.SubType)]
		}

		// for each property in the plugin, map it to its respective value
		for property, value := range pluginProperties {
			fmt.Println("setting resource", resourceName, property, value)
			// If the property represents a token that needs to be mapped then do so
			if token, ok := value.(string); ok {
				if contents, ok := extractTokenContents(token); ok {
					// The contents can be processed by the caller by splitting on '.'
					// For example: "infra.vpc.id" or "stage.variable_name"
					parts := strings.Split(contents, ".")
					source := parts[0]

					if source == "infra" {
						fmt.Println("setting infra resource", parts)

						refName := parts[1]
						propertyName := parts[2]
						// map the variable output to the infra resource
						refProperty := terraformInfraResources[refName].Get(jsii.String(propertyName))

						terraformResources[resourceName].Set(jsii.String(property), refProperty)
					} else if source == "stage" {
						// TODO: Implement stage variable mapping
					} else {
						return fmt.Errorf("unknown variable mapping")
						// Invalid variable mapping for now
					}

					_ = parts // TODO: Handle the token parts based on your needs
				}

				continue
			}

			// otherwise, just set the value
			terraformResources[resourceName].Set(jsii.String(property), value)
		}
	}

	app.Synth()

	return nil
}

var _ engines.Engine = &TerraformEngine{}

type terraformEngineOption func(*TerraformEngine)

func WithRepository(repository TerraformPluginRepository) terraformEngineOption {
	return func(engine *TerraformEngine) {
		engine.repository = repository
	}
}

func New(platformFile io.Reader, opts ...terraformEngineOption) *TerraformEngine {
	platform := &schema.TerraformPlatform{}

	json.NewDecoder(platformFile).Decode(platform)

	engine := &TerraformEngine{
		platform: platform,
	}

	for _, opt := range opts {
		opt(engine)
	}

	return engine
}
