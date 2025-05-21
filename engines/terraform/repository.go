package terraform

import "github.com/nitrictech/nitric/engines/terraform/schema"

type TerraformPluginRepository interface {
	GetPlugin(name string) (*schema.TerraformPluginManifest, error)
}
