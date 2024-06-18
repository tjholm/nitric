package deploytf

import (
	"fmt"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	deploymentspb "github.com/nitrictech/nitric/core/pkg/proto/deployments/v1"
)

// Api - Deploy an API Gateway
func (a *NitricAzureTerraformProvider) Api(tack cdktf.TerraformStack, name string, config *deploymentspb.Api) error {
	return fmt.Errorf("Not implemented")
}
