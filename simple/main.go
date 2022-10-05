package main

import (
	"fmt"

	"github.com/goreq/goreq"
)

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	resp, err := goreq.Get("https://my-json-server.typicode.com/hadihammurabi/flutter-webservice/contacts")
	panicOnError(err)
	defer resp.Body.Close()

	var data interface{}
	err = resp.Json(&data)
	panicOnError(err)

	fmt.Println(data)
}
