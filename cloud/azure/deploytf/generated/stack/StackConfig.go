package stack

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StackConfig struct {
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Providers *[]interface{} `field:"optional" json:"providers" yaml:"providers"`
	// Experimental.
	SkipAssetCreationFromLocalModules *bool `field:"optional" json:"skipAssetCreationFromLocalModules" yaml:"skipAssetCreationFromLocalModules"`
	// Whether to deploy an Azure Key Vault for this stack.
	DeployKeyVault *bool `field:"required" json:"deployKeyVault" yaml:"deployKeyVault"`
	// Whether to deploy an Azure Storage Account for this stack.
	DeployStorage *bool `field:"required" json:"deployStorage" yaml:"deployStorage"`
	// The Azure region to deploy resources to.
	Region *string `field:"required" json:"region" yaml:"region"`
	// The name of the nitric stack.
	StackName *string `field:"required" json:"stackName" yaml:"stackName"`
}

