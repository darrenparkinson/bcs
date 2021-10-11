package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi2conv"
)

//go:generate oapi-codegen -package bcs-openapi -o ./bcs-openapi.gen.go ./openapi.json
func main() {
	// Load v2 openapi
	input, err := ioutil.ReadFile("swagger.json")
	if err != nil {
		log.Fatal(err)
	}
	var doc openapi2.T
	if err = json.Unmarshal(input, &doc); err != nil {
		log.Fatal(err)
	}

	// Convert to v3 openapi
	v3, err := openapi2conv.ToV3(&doc)
	if err != nil {
		log.Fatal(err)
	}

	// Validate and write
	err = v3.Validate(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	j, err := v3.MarshalJSON()
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("openapi.json", j, 0644)
}
