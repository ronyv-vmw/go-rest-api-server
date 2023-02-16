package main

import (
	"fmt"

	"example.com/backend"
)

func main() {
	fmt.Println("<<<<<<<<<<Frontend App>>>>>>>>>>>>")
	a := backend.App{}
	a.Port = ":9002"
	a.Initialize()
	a.Run()
}
