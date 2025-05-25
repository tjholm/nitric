package main

import (
	_ "embed"
	"log"
	"os"
	"text/template"
)

//go:embed main.tmpl
var mainTmpl string

type goPlugin struct {
	Alias  string `json:"Alias"`
	Name   string `json:"Name"`
	Import string `json:"Import"`
}

type pluginDefintion struct {
	Pubsub  []goPlugin `json:"Pubsub"`
	Storage []goPlugin `json:"Storage"`
	Service goPlugin   `json:"Service"`
}

func main() {
	// template our main.go by injecting the plugin name and known plugin constructor
	tmpl, err := template.New("main").Parse(mainTmpl)
	if err != nil {
		log.Fatalf("error parsing template: %v", err)
	}

	// NOTE: The plugin definitions will come from an externally provided config file
	// This is hardcoded here as a demonstration
	err = tmpl.Execute(os.Stdout, pluginDefintion{
		Service: goPlugin{
			Alias:  "awslambda",
			Name:   "default",
			Import: "github.com/nitrictech/nitric/engines/terraform/plugins/awslambda",
		},
		Pubsub:  []goPlugin{},
		Storage: []goPlugin{},
	})
	if err != nil {
		log.Fatalf("error executing template: %v", err)
	}
}
