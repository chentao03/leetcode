package main

import (
	"fmt"
	"unsafe"
)

func main() {
	array := [3][3]int16{}
	var sint int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(i, j, &array[i][j], unsafe.Sizeof(array[i][j]))
		}

	}
	fmt.Printf("sint:%d\n", &sint)
}
