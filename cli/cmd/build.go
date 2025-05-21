package cmd

import (
	"fmt"

	"github.com/nitrictech/nitric/cli/pkg/schema"
	"github.com/nitrictech/nitric/engines/terraform"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type MockTerraformPluginRepository struct {
	plugins map[string]*terraform.PluginManifest
}

func (r *MockTerraformPluginRepository) GetPlugin(name string) (*terraform.PluginManifest, error) {
	return r.plugins[name], nil
}

func createMockTerraformPluginRepository() *MockTerraformPluginRepository {
	return &MockTerraformPluginRepository{
		plugins: map[string]*terraform.PluginManifest{
			"nitric-aws-lambda": {
				Name: "nitric-aws-lambda",
				Deployment: terraform.DeploymentModule{
					Terraform: "terraform-aws-modules/lambda/aws",
				},
			},
			"nitric-aws-cloudfront": {
				Name: "nitric-aws-cloudfront",
				Deployment: terraform.DeploymentModule{
					Terraform: "terraform-aws-modules/cloudfront/aws",
				},
			},
			"nitric-aws-vpc": {
				Name: "nitric-aws-vpc",
				Deployment: terraform.DeploymentModule{
					Terraform: "terraform-aws-modules/vpc/aws",
				},
			},
		},
	}
}

func writeExampleNitricYaml(fs afero.Fs) {
	example := &schema.Application{
		Name: "test",
		Resources: map[string]schema.Resource{
			"service": {
				Type: "service",
				ServiceResource: &schema.ServiceResource{
					Port: 8080,
					Env: map[string]string{
						"TEST": "test",
					},
					Container: schema.Container{
						Image: &schema.DockerImage{
							ID: "test",
						},
					},
				},
			},
			"ingress": {
				Type: "entrypoint",
				EntrypointResource: &schema.EntrypointResource{
					Routes: map[string]schema.Route{
						"/": {
							TargetName: "service",
						},
					},
				},
			},
		},
	}

	yamlBytes, err := yaml.Marshal(example)
	if err != nil {
		fmt.Println("Error marshalling example to YAML:", err)
	}

	afero.WriteFile(fs, "nitric.yaml", yamlBytes, 0644)
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Builds the nitric application",
	Long:  `Builds an application using the nitric.yaml application spec and referenced platform.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Read the nitric.yaml file
		fs := afero.NewOsFs()

		writeExampleNitricYaml(fs)

		yamlData, err := afero.ReadFile(fs, "nitric.yaml")
		if err != nil {
			fmt.Print("Error opening nitric.yaml file: ", err)
			return
		}

		appSpec, results, err := schema.ApplicationFromYaml(string(yamlData))
		if err != nil {
			fmt.Print("Error parsing nitric.yaml file: ", err)
			return
		}

		if results != nil && !results.Valid() {
			fmt.Println("Error validating nitric.yaml file:")
			for _, err := range results.Errors() {
				fmt.Println(err)
			}
			return
		}

		// TODO: 912 repository
		mockRepository := createMockTerraformPluginRepository()

		mockPlatformRepository := terraform.NewMockPlatformRepository()

		platform := terraform.New(mockPlatformRepository.GetPlatform(appSpec.Platform), terraform.WithRepository(mockRepository))
		// Parse the application spec
		// Validate the application spec
		// Build the application using the specified platform
		// Handle any errors that occur during the build process

		err = platform.Apply(appSpec)
		if err != nil {
			fmt.Print("Error applying platform: ", err)
			return
		}

		fmt.Println("Build completed successfully.")

	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
