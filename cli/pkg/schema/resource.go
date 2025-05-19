package schema

import "github.com/invopop/jsonschema"

type Resource struct {
	Type             string `json:"type" jsonschema:"-"`
	*ServiceResource `json:",inline,omitempty" jsonschema:"-"`
	*BucketResource  `json:",inline,omitempty" jsonschema:"-"`
}

func (Resource) JSONSchemaExtend(schema *jsonschema.Schema) {
	serviceSchema := jsonschema.Reflect(ServiceResource{})
	bucketSchema := jsonschema.Reflect(BucketResource{})

	schema.Properties = nil
	schema.AdditionalProperties = nil
	schema.OneOf = []*jsonschema.Schema{serviceSchema.Definitions["ServiceResource"], bucketSchema.Definitions["BucketResource"]}
}

type ServiceResource struct {
	Test string `json:"test"`
	// Only used for schema generation, will always be nil. Do not use or remove.
	SchemaOnlyServiceTypeHack string `json:"type" jsonschema:"type,enum=service"`
}

type BucketResource struct {
	TestTwo string `json:"test-two"`
	// Only used for schema generation, will always be nil. Do not use or remove.
	SchemaOnlyBucketTypeHack string `json:"type" jsonschema:"type,enum=bucket"`
}
