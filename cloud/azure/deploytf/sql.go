package deploytf

import (
	"fmt"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	deploymentspb "github.com/nitrictech/nitric/core/pkg/proto/deployments/v1"
)

// SqlDatabase - Deploy a SQL Database
func (a *NitricAzureTerraformProvider) SqlDatabase(stack cdktf.TerraformStack, name string, config *deploymentspb.SqlDatabase) error {
	return fmt.Errorf("The Nitric Azure Terraform provider does not currently support sql databases")
}
