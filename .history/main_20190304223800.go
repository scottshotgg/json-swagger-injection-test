package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type swagger struct {
	Paths map[string]map[string]Operation
}

type Operation struct {
	Parameters []Param
}

type Param struct {
	Name     string
	In       string
	Required bool
}

func main() {
	var data, err = ioutil.ReadFile("swagger.json")
	if err != nil {
		log.Fatal(err)
	}

	var swag swagger

	err = json.Unmarshal(data, &swag)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("swag", swag)

	for en := range swag.Paths {
		for op := range swag.Paths[en] {
			// op.Parameters
			// swag.Paths[endpoint]
		}
	}
}
