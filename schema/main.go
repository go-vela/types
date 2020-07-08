//+build ignore

package main

import (
	"encoding/json"
	"fmt"

	"github.com/alecthomas/jsonschema"

	"github.com/go-vela/types/yaml"
)

func main() {
	r := jsonschema.Reflector{
		ExpandedStruct:             true,
		RequiredFromJSONSchemaTags: true,
	}
	s := r.Reflect(&yaml.Build{})
	

	j, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%s\n", j)
}
