package caesar

import (
	"strings"
	"testing"
)

func TestCaesar(t *testing.T) {
	targetTextRus := strings.ToLower("Ленин! Партия! Комсомол!")
	encryptText, _ := CipherCaesar([]rune("ленин! партия! комсомол!"), 7)
	text, _ := CipherCaesar([]rune(encryptText), -7)
	if text != targetTextRus {
		t.Error("Сломался алгоритм шифрования. ", targetTextRus, " зашифрован не правильно")
	}
}
