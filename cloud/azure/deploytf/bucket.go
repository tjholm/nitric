package deploytf

import (
	"fmt"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	deploymentspb "github.com/nitrictech/nitric/core/pkg/proto/deployments/v1"
)

// Bucket - Deploy a Storage Bucket
func (a *NitricAzureTerraformProvider) Bucket(stack cdktf.TerraformStack, name string, config *deploymentspb.Bucket) error {
	return fmt.Errorf("Not implemented")
}
