package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type swagger struct {
	Paths struct {
		Array map[string]map[string][]Param
	}
}

// type PathObject struct {
// 	Paths map[string]map[string][]Param
// }

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

	// // var value = `
	// // 	{
	// // 		"name": "stuff",
	// // 		"in": "header",
	// // 		"description" : "things"
	// // 	}
	// // `

	// paths, _, _, _ := jsonparser.Get(swagger, "paths")

	// fmt.Println(jsonparser.ObjectEach(paths, func(key []byte, value []byte, t jsonparser.ValueType, offset int) error {
	// 	fmt.Println(string(key), string(value), t.String(), offset)

	// 	fmt.Println(jsonparser.ObjectEach(value, func(key []byte, value []byte, t jsonparser.ValueType, offset int) error {
	// 		fmt.Println(string(key), string(value), t.String(), offset)

	// 		params, _, _, _ := jsonparser.Get(value, "parameters")

	// 		fmt.Println("params2", string(params))

	// 		type param struct {
	// 			Name     string
	// 			In       string
	// 			Required bool
	// 		}

	// 		var paramsArray []param

	// 		err = json.Unmarshal(params, &paramsArray)

	// 		paramsArray = append(paramsArray, param{
	// 			Name: "stuff",
	// 			In:   "header",
	// 		})

	// 		params, err = json.Marshal(paramsArray)

	// 		jsonparser.Set(value, params, "parameters")

	// 		return nil
	// 	}))

	// 	jsonparser.Set(paths, params, string(key))

	// 	return nil
	// }))

	// ioutil.WriteFile("mod_swagger.json", swagger, 0666)
}
