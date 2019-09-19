package main

import (
	"cryptokiddies-server/crypt"
	"fmt"
)

func main() {
	//x := []int{2, 4, 1, 5, 3}
	//crypt.SetKey(x)
	crypt.SetStringKey("zabeba")
	text := crypt.Encrypt("I love coding on go language")
	fmt.Println(text)
	fmt.Println(crypt.Decrypt(text))
}
