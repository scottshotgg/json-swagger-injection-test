package main

import (
	"fmt"
	"log"
	"os"

	"github.com/buger/jsonparser"
)

func main() {
	var swagger, err = os.ReadFile("swagger.json")
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
