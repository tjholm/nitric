package api

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiConfig struct {
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Providers *[]interface{} `field:"optional" json:"providers" yaml:"providers"`
	// Experimental.
	SkipAssetCreationFromLocalModules *bool `field:"optional" json:"skipAssetCreationFromLocalModules" yaml:"skipAssetCreationFromLocalModules"`
	// The name of the API Gateway.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Open API spec.
	Spec *string `field:"required" json:"spec" yaml:"spec"`
	// The ID of the stack.
	StackId *string `field:"required" json:"stackId" yaml:"stackId"`
}

