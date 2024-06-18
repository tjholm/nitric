package deploytf

import (
	"fmt"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	deploymentspb "github.com/nitrictech/nitric/core/pkg/proto/deployments/v1"
)

// Websocket - Deploy a Websocket Gateway
func (a *NitricAzureTerraformProvider) Websocket(stack cdktf.TerraformStack, name string, config *deploymentspb.Websocket) error {
	return fmt.Errorf("Not implemented")
}
