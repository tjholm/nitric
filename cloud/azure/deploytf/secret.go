package deploytf

import (
	"fmt"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	deploymentspb "github.com/nitrictech/nitric/core/pkg/proto/deployments/v1"
)

// Secret - Deploy a Secret
func (a *NitricAzureTerraformProvider) Secret(stack cdktf.TerraformStack, name string, config *deploymentspb.Secret) error {
	return fmt.Errorf("Not implemented")
}
