package main

import "strings"

func main() {

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
