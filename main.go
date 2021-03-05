package main

import (
	"fmt"
	"working-with-errors/httpRequest"
)

func main() {
	fmt.Println("Hello error's treatment!")
	fmt.Println("...........")
	fmt.Println(httpRequest.GetRequest("http://google.com.br"))
}
