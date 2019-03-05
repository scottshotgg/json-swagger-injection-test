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

	jsonparser
}
