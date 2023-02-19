package main

import "fmt"

func superAdd(numbers ...int) {
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}

func superAdd2(numbers ...int) {
	for index, number := range numbers { //range는 인덱스를 줌
		fmt.Println(index, number)
	}
}

func superAdd3(numbers ...int) {
	total := 0
	for _, number := range numbers { //index _로 무시가능
		total += number
	}
	fmt.Println(total)
}

func main() {
	superAdd3(1, 2, 3, 4, 5, 6, 7)
}
