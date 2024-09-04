package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("stringFile.txt")
	if err != nil {
		fmt.Println("error reading file:",err)
		return
	}
	checkString := strings.Split(string(file), "\n")
	count := 0
	for _, checkLine := range checkString{
		if niceString(checkLine) {
			count ++
		}
	}
	fmt.Println("Nice strings are: ", count)
}

func niceString(str string) bool {
	vowels := "aeiou"
	countVow := 0
	doubleLetter := false
	notAllowed := []string{"ab", "cd", "pq", "xy"}

	for i := 0; i < +len(str)-1; i++ {
		if strings.ContainsRune(vowels, rune(str[i])) {
			countVow++
		}
		if i > 0 && str[i] == str[i-1] {
			doubleLetter = true
		}
		if i > 0 {
			substring := str[i-1 : i+1]
			for _, ch := range notAllowed {
				if substring == ch {
					return false
				}
			}
		}
	}
	return countVow >= 3 && doubleLetter
}
