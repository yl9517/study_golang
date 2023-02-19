package main

import "fmt"

//기본
func canIDrink(age int) bool {
	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	}
	return false
}

//기본
func canIDrink2(age int) bool {
	switch age {
	case 18:
		return false
	case 20:
		return true
	}
	return false
}

//내부 변수 선언
func canIDrink3(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 18:
		return false
	case 20:
		return true
	}
	return false
}
func main() {
	fmt.Println(canIDrink(18))
}
