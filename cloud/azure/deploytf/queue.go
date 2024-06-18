package deploytf

import (
	"fmt"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	deploymentspb "github.com/nitrictech/nitric/core/pkg/proto/deployments/v1"
)

// Queue - Deploy a Queue
func (a *NitricAzureTerraformProvider) Queue(stack cdktf.TerraformStack, name string, config *deploymentspb.Queue) error {
	return fmt.Errorf("Not implemented")
}
