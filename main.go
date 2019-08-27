package main

import (
	"fmt"
	"wire_test/application/handler"
	"wire_test/configuration"
)

func main() {
	configuration.Init()

	fmt.Println(handler.EditExam())
	fmt.Println("111")
}
