package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("puzzle.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
	}
	x, y := 0, 0
	mymap := make(map[[2]int]int)
	mymap[[2]int{x, y}]++
	fmt.Println(mymap)
	for _, v := range string(file) {
		if v == '^' {
			y ++
		} else if v == 'v' {
			y--
		}else if v == '>'{
			x++
		} else if v== '<' {
			x--
		}
		mymap[[2]int{x, y}]++

	}
	fmt.Println(len(mymap))

}
