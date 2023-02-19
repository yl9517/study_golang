package main

import "fmt"

func main() {
	a := 2
	b := a //값 복사
	a = 10 //b에 영향 없음

	fmt.Println(a, b) //10, 2
}

//내가 원하는 것 : 값 복사 말고 메모리주소 참조
// 메모리주소 보는 법 : &
// 메모리 주소의 값 보는 법 : *
func main2() {
	a := 2
	b := &a //a의 메모리주소 복사 (&a == b)

	a = 10
	fmt.Println(a, *b) // 10, 10
}

//b를 가지고 a의 값을 변경 시킬수도 있음
func main3() {
	a := 2
	b := &a //a의 메모리주소 복사 (&a == b)

	*b = 3
	fmt.Println(a, *b) // 10, 10
}
