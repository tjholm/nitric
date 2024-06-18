package deploytf

import (
	"fmt"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	deploymentspb "github.com/nitrictech/nitric/core/pkg/proto/deployments/v1"
)

// Http - Deploy a HTTP Proxy
func (a *NitricAzureTerraformProvider) Http(tack cdktf.TerraformStack, name string, config *deploymentspb.Http) error {
	return fmt.Errorf("Not implemented")
}
