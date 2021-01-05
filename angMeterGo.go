package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello world!")
	if len(os.Args) != 2 {
		fmt.Println("Wrong arguments")
		return
	}
	selfName := os.Args[0]
	figPath := os.Args[1]
	fmt.Println(selfName, figPath)
}
