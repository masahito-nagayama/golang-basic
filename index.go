package main

import (
	"fmt"
)

func main() {
	var i int = 100
	
	fmt.Println(i + 50)

	var t, f bool = true,false
	fmt.Println(t, f)

	var s string = "HELLO GO LANG"
	fmt.Println(s)

	byteA := []byte{72, 73}
	fmt.Println(byteA)
}