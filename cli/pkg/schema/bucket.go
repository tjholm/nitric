package schema

type BucketResource struct {
	// Only used for schema generation, will always be nil. Do not use or remove.
	BucketSchemaOnlyHackType string `json:"type" jsonschema:"type,enum=bucket"`
}
