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
        
        // if hashString[:5] == "00000" {
        //     fmt.Println("The lowest number to produce the hashstring:", i)
        //     break
        // }
		 if hashString[:6] == "000000" {
			fmt.Println("Starting with 6 zeroes is: ", i)
			
		 }

        i++
    }
}
