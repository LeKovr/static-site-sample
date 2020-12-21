package main

import (
	"fmt"
)

func main() {

	var s1 string
	if 2 == 2 {
		s1, s2 := "1", "2"
		fmt.Printf("Ret: %s - %s\n", s1, s2)

	}
	fmt.Printf("Ret: %s\n", s1)
	return
}
