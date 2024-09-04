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
	x, y, x1, y1 := 0, 0, 0, 0
	mymap := make(map[[2]int]int)
	mymap[[2]int{x, y}]++
	fmt.Println(mymap)
	for i, v := range string(file) {
		if i%2 == 0 {
			if v == '^' {
				y++
			} else if v == 'v' {
				y--
			} else if v == '>' {
				x++
			} else if v == '<' {
				x--
			}
			mymap[[2]int{x, y}]++
		} else {
			if v == '^' {
				y1++
			} else if v == 'v' {
				y1--
			} else if v == '>' {
				x1++
			} else if v == '<' {
				x1--
			}
			mymap[[2]int{x1, y1}]++
		}

	}
	fmt.Println(len(mymap))

}
