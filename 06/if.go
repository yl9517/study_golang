package main

import "fmt"

func canIDrink(age int) bool {

	// if age < 18 {
	// 	return false
	// }
	// return true

	//내부에변수 선언가능 (if-else 조건에만 사용하기 위해 변수를 생성했구나)
	if koreanAge := age + 2; koreanAge < 20 {
		return false
	}
	return true
}
func main() {
	fmt.Println(canIDrink(18))
}
