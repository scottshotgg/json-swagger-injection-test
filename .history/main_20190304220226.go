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
	fmt.Println(string(paths) + "\n")
	fmt.Println(jsonparser.EachKey(swagger, func(_ int, _ []byte, _ jsonparser.ValueType, _ error) {
		fmt.Println("hi")
	}, []string{"/paths"}))

	// jsonparser.Set(swagger, )
}
