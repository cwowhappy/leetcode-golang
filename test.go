package main

import "fmt"

func main() {
	numMap := make(map[int]int)
	freq, ok := numMap[0]
	fmt.Println(freq, ok)
}
