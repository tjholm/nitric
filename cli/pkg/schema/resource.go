package schema

import "github.com/invopop/jsonschema"

type Resource struct {
	Type    string `json:"type" yaml:"type" jsonschema:"-"`
	SubType string `json:"sub-type,omitempty" yaml:"sub-type,omitempty"`

	// A resource can contain oneof the following sets of keys (see JSONSchemaExtended)
	*ServiceResource `json:",inline,omitempty" yaml:",inline,omitempty" jsonschema:"-"`
	*BucketResource  `json:",inline,omitempty" yaml:",inline,omitempty" jsonschema:"-"`
}

// schema types defined for the output schema
var schemaTypes = map[string]interface{}{
	"ServiceResource": ServiceResource{},
	"BucketResource":  BucketResource{},
}

func (Resource) JSONSchemaExtend(schema *jsonschema.Schema) {
	if schema.Definitions == nil {
		schema.Definitions = map[string]*jsonschema.Schema{}
	}

	subSchemas := []*jsonschema.Schema{}
	for _, res := range schemaTypes {
		s := jsonschema.Reflect(res)

		s.AdditionalProperties = nil
		s.Properties = nil
		// TODO: Make sure sub definitions are also collected
		// subSchemas = append(subSchemas, s.Definitions[name])
		subSchemas = append(subSchemas, s)

		// for n, def := range s.Definitions {
		// 	if n != name {
		// 		schema.Definitions[n] = def
		// 	}
		// }
	}

	schema.Properties = nil
	schema.AdditionalProperties = nil
	schema.OneOf = subSchemas
}
