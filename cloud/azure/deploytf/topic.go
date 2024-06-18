package deploytf

import (
	"fmt"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	deploymentspb "github.com/nitrictech/nitric/core/pkg/proto/deployments/v1"
)

// Topic - Deploy a Pub/Sub Topic
func (a *NitricAzureTerraformProvider) Topic(stack cdktf.TerraformStack, name string, config *deploymentspb.Topic) error {
	return fmt.Errorf("Not implemented")
}
