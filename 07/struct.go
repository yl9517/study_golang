package main

import "fmt"

//1. 어떤 struct인지 정의

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"hamberger", "remen"}
	summer := person{
		name:    "summer",
		age:     26,
		favFood: favFood,
	}
	fmt.Println(summer)

}
