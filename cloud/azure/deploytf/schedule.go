package deploytf

import (
	"fmt"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	deploymentspb "github.com/nitrictech/nitric/core/pkg/proto/deployments/v1"
)

// Schedule - Deploy a Schedule
func (a *NitricAzureTerraformProvider) Schedule(stack cdktf.TerraformStack, name string, config *deploymentspb.Schedule) error {
	return fmt.Errorf("Not implemented")
}
