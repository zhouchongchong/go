package main

import (
	"fmt"
)

func main() {
	fmt.Println("Base grammer is began!!")
	b, a := getData()
	fmt.Println(b, a)

	fmt.Println("Base grammer is end!!")
}

func getData() (int16, int16) {
	return 134, 145
}
