package main

import (
	_ "embed"
	"encoding/json"
	"fmt"

	"github.com/nitrictech/nitric/cli/pkg/schema"
)

//go:embed test.json
var testSchema []byte

func main() {
	var schema schema.Schema

	json.Unmarshal(testSchema, &schema)
	fmt.Printf("%+v\n", schema)

	for _, res := range schema.Resources {
		if res.Type == "service" {
			fmt.Printf("%+v\n", res.SchemaServiceONLY)
		} else {
			fmt.Printf("%+v\n", res.SchemaBucketONLY)
		}

	}

	data, _ := json.MarshalIndent(schema, "", "  ")

	fmt.Printf("%s\n", string(data))
}
