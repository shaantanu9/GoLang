package main

import (
	"fmt"
	"go/format"
)

func main() {

	data,err := fmt.Scanf(format string,a)
	if(err!=nil){ fmt.Println(err)}
	fmt.Println(data)
}
