package main

import (
	"io/ioutil"
	"log"

	"github.com/buger/jsonparser"
)

var (
	myHeader = `{
		Name: "stuff",
		In:   "header",
	}`
)

type swagger struct {
	Paths map[string]map[string]Operation `json:"paths"`
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

	paths, _, _, _ := jsonparser.Get(data, "paths")

	jsonparser.ObjectEach(paths, func(key, value []byte, _ ValueType, _ int) error {
		return nil
	})
}
