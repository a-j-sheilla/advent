package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// l*w*h
func main() {
	file := "juma.txt"
	filecontent, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("SJuma doesnt exist")
	}

	arr := strings.Split(string(filecontent), "\n")

	totalArea := 0
	totalSlack := 0
	totalVolume := 0
	totalLength := 0


	for _, line := range arr {
		slice := strings.Split(line, "x")
		l, err := strconv.Atoi(slice[0])
		if err != nil {
			fmt.Println("Error converting l to int: ", err)
		}
		w, err := strconv.Atoi(slice[1])
		if err != nil {
			fmt.Println("Error converting w to int: ", err)
		}
		h, err := strconv.Atoi(slice[2])
		if err != nil {
			fmt.Println("Error converting h to int: ", err)
		}
		surfaceArea := 2*(l*w) + 2*(w*h) + 2*(l*h)
		totalArea += surfaceArea
		slack := slackArea(l, w, h)
		totalSlack += slack

		volume := l*w*h
		totalVolume += volume

		length := ribbon(l, w, h)
		totalLength += length



	}
newLength := totalVolume + totalLength
fmt.Println("Foolish moses: ", newLength)

	newArea := totalArea + totalSlack
	fmt.Println("Moses is a big fool: ", newArea)
}

func slackArea(l, w, h int) int {
	side1 := l * w
	side2 := l * h
	side3 := h * w
	if side1 <= side2 && side1 <= side3 {
		return side1
	}
	if side2 <= side1 && side2 <= side3 {
		return side2
	}
	return side3
}

func ribbon(l, w, h int) int {

	length1 := 2*(l+w)
	length2 := 2*(l+h)
	length3 := 2*(w+h)
	if length1 <= length2 && length1 <= length3{
		return length1
	}
	if length2 <= length1 && length2 <= length3{
		return length2
	}
	return length3
}

