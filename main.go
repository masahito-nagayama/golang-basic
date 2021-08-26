package main

import (
	"fmt"
	"time"
)

func outer() {
	var i4 = 10000000000
	fmt.Println(i4)
}

func main() {
	fmt.Println("Hello World")

	fmt.Println(time.Now())

	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	var t,f bool = true,false
	fmt.Println(t,f)

	var (
		si = 100
		i2 = "Hello Golang"
	)
	fmt.Println(si,i2)

	i = 150
	fmt.Println(i)

	outer()
}