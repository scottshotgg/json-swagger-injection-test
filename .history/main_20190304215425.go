package main

import (
	"io/ioutil"
	"log"
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

	// jsonparser.Set(swagger, )
}
