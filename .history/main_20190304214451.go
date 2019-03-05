package main

import "io/ioutil"

func main() {
	var swagger, err = ioutil.ReadAll("swagger.json")
}
