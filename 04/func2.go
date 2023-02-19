package main

import (
	"fmt"
	"strings"
)

func lenAndUooer(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")

	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	totalLength, upperName := lenAndUooer("summer")
	fmt.Println(totalLength, upperName)

}
