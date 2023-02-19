package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUooer(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}
func main() {
	fmt.Println(multiply(2, 2))

	totalLength, upperName := lenAndUooer("summer")
	fmt.Println(totalLength, upperName)

	repeatMe("summer", "spring", "winter", "authmn")
}
