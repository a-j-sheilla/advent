package main

import (
    "crypto/md5"
    "encoding/hex"
    "fmt"
    "strconv"
)

func main() {
    secretKey := "yzbqklnj"
    i := 1
// yzbqklnj1
    for {
		numString := strconv.Itoa(i)
        data := secretKey + numString
        
        hash := md5.Sum([]byte(data))
        hashString := hex.EncodeToString(hash[:])
		//fmt.Println(numString,secretKey,data,hash,hashString)
        //part 1 solution of day 4
        if hashString[:5] == "00000" {
            fmt.Println("The lowest number to produce the hashstring:", i)
            break
        }
		// part 2 solution of day 4
		 if hashString[:6] == "000000" {
			fmt.Println("Starting with 6 zeroes is: ", i)
			break

		 }
		 //********comment out part 1 solution to find answer for part 2**********

        i++
    }
}
