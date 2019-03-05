package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/buger/jsonparser"
)

func main() {
	var swagger, err = ioutil.ReadAll("swagger.json")
	if err != nil {
		log.Fatal(err)
	}

	var value = `
		{
			"name": "stuff",
			"in": "header",
			"description" : "things"
		}
	`

	paths, err := jsonparser.Get(swagger, "paths")
	fmt.Println(string(paths))

	// jsonparser.EachKey()

	// jsonparser.Set(swagger, )
}
