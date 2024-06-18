package deploytf

import (
	"fmt"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/nitrictech/nitric/cloud/common/deploy/provider"
	deploymentspb "github.com/nitrictech/nitric/core/pkg/proto/deployments/v1"
)

// Service - Deploy an service (Service)
func (a *NitricAzureTerraformProvider) Service(stack cdktf.TerraformStack, name string, config *deploymentspb.Service, runtimeProvider provider.RuntimeProvider) error {
	return fmt.Errorf("Not implemented")
}
