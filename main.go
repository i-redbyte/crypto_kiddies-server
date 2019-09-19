package main

import (
	"cryptokiddies-server/crypt"
	"fmt"
)

func main() {
	// TODO: Red_byte for test transposition crypt
	crypt.SetStringKey("java and kotlin")
	text := crypt.Encrypt("I love coding on go language")
	fmt.Println(text)
	fmt.Println(crypt.Decrypt(text))
}
