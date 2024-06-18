package deploytf

import (
	"fmt"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	deploymentspb "github.com/nitrictech/nitric/core/pkg/proto/deployments/v1"
)

// Policy - Deploy a Policy
func (a *NitricAzureTerraformProvider) Policy(stack cdktf.TerraformStack, name string, config *deploymentspb.Policy) error {
	return fmt.Errorf("Not implemented")
}
