package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/buger/jsonparser"
)

func main() {
	var swagger, err = ioutil.ReadFile("swagger.json")
	if err != nil {
		log.Fatal(err)
	}

	// var value = `
	// 	{
	// 		"name": "stuff",
	// 		"in": "header",
	// 		"description" : "things"
	// 	}
	// `

	paths, _, _, _ := jsonparser.Get(swagger, "paths")

	fmt.Println(jsonparser.ObjectEach(paths, func(key []byte, value []byte, t jsonparser.ValueType, offset int) error {
		fmt.Println(string(key), string(value), t.String(), offset)

		return nil
	}))

	// jsonparser.Set(swagger, )
}
