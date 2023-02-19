package main

import "fmt"

func main() {
	summer := map[string]string{"name": "summer", "age": "26"} //key = string, value = string

	//object처럼 여러 타입을 값으로 넣고싶으면 struct 이용해야함

	for key, value := range summer {
		fmt.Println(key, value)
	}
}
