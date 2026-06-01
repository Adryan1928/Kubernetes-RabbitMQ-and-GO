package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("GO Runtime")
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}