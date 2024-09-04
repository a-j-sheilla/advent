package main

import (
	"os"
	"fmt"
)

func main(){
	start := 0
	//directionUp := "("
	//directionDown := ")"

	if len(os.Args) != 2 {
		return
	}
	directions := os.Args[1]

	for i, ch := range directions {
		if  ch == '(' {
			start += 1
		} else if ch == ')' {
			start -= 1
		}
		if start == -1 {
			fmt.Println("First basement occurance is: ", i+1)
			
		}
			
	}
	fmt.Println("floor number is: ", start)
}