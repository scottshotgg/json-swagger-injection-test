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
	fmt.Println(string(paths))
	fmt.Println(jsonparser.EachKey(paths, func(_ int, _ []byte, _ jsonparser.ValueType, _ error) {}, []string{"*"}))

	// jsonparser.Set(swagger, )
}
