package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	shasha()
	//fmt.Printf("%x", )

}

func shasha() {
	myPassword := "123456"
	bv := []byte(myPassword)
	hash := sha256.New()
	hash.Write(bv)
	md := hash.Sum(nil)
	fmt.Println("=============================================")
	fmt.Println(md)
	mdStr := hex.EncodeToString(md)
	fmt.Println("=============================================")
	fmt.Println(mdStr)
}
