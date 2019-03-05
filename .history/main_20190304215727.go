package main

import (
	"fmt"
	"log"
	"os"

	"github.com/buger/jsonparser"
)

func main() {
	ioutil.ReadAll()

	var swagger, err = os. ("swagger.json")
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
