package roles

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RolesConfig struct {
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Providers *[]interface{} `field:"optional" json:"providers" yaml:"providers"`
	// Experimental.
	SkipAssetCreationFromLocalModules *bool `field:"optional" json:"skipAssetCreationFromLocalModules" yaml:"skipAssetCreationFromLocalModules"`
	// The Azure resource group name.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// The id of the nitric stack.
	StackId *string `field:"required" json:"stackId" yaml:"stackId"`
	// The Azure subscription id.
	SubscriptionId *string `field:"required" json:"subscriptionId" yaml:"subscriptionId"`
}

