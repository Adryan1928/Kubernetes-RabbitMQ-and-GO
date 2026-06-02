package main

import (
	"fmt"
	"runtime"
)

const (
	_  = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
	GB
	TB
)

const (
	x int = 1
	y int = iota
	_
	_
	_
	z 
	
)

func main() {

	fmt.Println("binary\t\t\t\tdecimal")
	fmt.Printf("%b\t\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)


	fmt.Println("GO Runtime")
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	fmt.Printf("%v - %b - %X \n", x, x, x)
	fmt.Printf("%v - %b - %X \n", x << 2, x << 2, x << 2)

	fmt.Println(y)
	fmt.Println(z)
}