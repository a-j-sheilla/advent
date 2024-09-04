package main

import (
    "strconv"
    "fmt"
)

func simpleHash(s string) string {
    var hash int
    for _, char := range s {
        hash = (hash + int(char)) % 100000
    }

    result := ""
    for hash > 0 {
        digit := hash % 16
        if digit < 10 {
            result = string(rune('0'+digit)) + result
        } else {
            result = string(rune('a'+digit-10)) + result
        }
        hash /= 16
    }

    for len(result) < 5 {
        result = "0" + result
    }
    return result
}

func main() {
    secretKey := "yzbqklnj"
    i := 1

    for {
        data := secretKey + strconv.Itoa(i)
        hash := simpleHash(data)

        if hash[:5] == "00000" {
            fmt.Println("The lowest number that produces a hash starting with five zeroes is:", i)
            break
        }

        i++
    }
}
