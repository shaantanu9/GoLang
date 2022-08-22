package main

import "fmt"

func main() {
	s := "sfdfasdfabcdefghijkmnlopqrst"
	funcResult := firstRepetedChar(s)
	fmt.Println(funcResult)
}

// create a empty slice and fill it with zero values
// a:= make([1]int, 0)


func firstRepetedChar(s string) string {
	ans := make(map[string]int)
	for i := 0; i < len(s); i++ {
		if _, ok := ans[string(s[i])]; ok {
			ans[string(s[i])]++
			} else {
				ans[string(s[i])] = 1
			}
		}
		fmt.Println(ans)
	return ""
}