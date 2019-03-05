package main

import (
	"encoding/json"
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

		fmt.Println(jsonparser.ObjectEach(value, func(key []byte, value []byte, t jsonparser.ValueType, offset int) error {
			fmt.Println(string(key), string(value), t.String(), offset)

			params, _, _, _ := jsonparser.Get(value, "parameters")

			fmt.Println("params2", string(params))

			type param struct {
				Name     string
				In       string
				Required bool
			}

			var paramsArray []param

			err = json.Unmarshal(params, &paramsArray)

			paramsArray = append(paramsArray, param{
				Name: "stuff",
				In:   "header",
			})

			params, err = json.Marshal(paramsArray)

			jsonparser.Set(value, params, "parameters")

			return nil
		}))

		return nil
	}))

	ioutil.WriteFile("mod_swagger.json", swagger)
}
