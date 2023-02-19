package main

import "fmt"

func main() {
	//크기 제한
	names := [5]string{"summer", "autumn", "winter"}
	names[3] = "spring"
	names[4] = "other"

	fmt.Println(names)

	//크기제한 X => slice (length 없음)
	//item추가하기 위해서 append 사용
	season := []string{"summer", "autumn", "winter"}
	season = append(season, "spring")
	season = append(season, "push")
	season = append(season, "pull")

	fmt.Println(season)
}
