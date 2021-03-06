package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

var (
	myHeader interface{} = Param{
		Name: "stuff",
		In:   "header",
	}
)

type swagger struct {
	Paths map[string]map[string]map[string]interface{}
}

type Operation struct {
	Parameters []Param `json:"parameters"`
}

type Param struct {
	Name     string `json:"name"`
	In       string `json:"in"`
	Required bool   `json:"required"`
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
			// swag.Paths[en][op].Parameters = append(swag.Paths[en][op].Parameters, myHeader)
			ops := swag.Paths[en][op]["parameters"].([]interface{})
			ops = append(ops, myHeader)
			swag.Paths[en][op]["parameters"] = ops
		}
	}

	swagData, err := json.Marshal(swag)
	if err != nil {
		log.Fatal(err)
	}

	// newSwagger, err := jsonparser.Set(data, swagData, "paths")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	ioutil.WriteFile("mod_swagger.json", swagData, 0666)
}
