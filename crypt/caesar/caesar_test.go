package caesar

import (
	"strings"
	"testing"
)

func TestCaesar(t *testing.T) {
	targetTextRus := "Строка будет находиться в сегменте кода, который доступен только на чтение и хранит весь Ваш код и такие мелочи, как, например, строковые литералы."
	targetTextEng := "“WebSocket” is an app protocol built on top of TCP that permits full duplex communication between client and server."
	targetText := "a zoри здесь tihie, а zoni здесь dikie!"
	if !compareTexts(targetTextRus, 7) || !compareTexts(targetTextRus, 37) ||
		!compareTexts(targetTextRus, 9) || !compareTexts(targetTextRus, 29) {
		t.Error("Сломался алгоритм шифрования. Текст:", targetTextRus, " зашифрован не правильно")
	}
	if !compareTexts(targetTextEng, 20) || !compareTexts(targetTextEng, 2) ||
		!compareTexts(targetTextEng, 31) || !compareTexts(targetTextEng, 12) {
		t.Error("Сломался алгоритм шифрования. Текст:", targetTextEng, " зашифрован не правильно")
	}
	if !compareTexts(targetText, 55) || !compareTexts(targetText, 5) ||
		!compareTexts(targetText, 68) || !compareTexts(targetText, 17) {
		t.Error("Сломался алгоритм шифрования. Текст:", targetText, " зашифрован не правильно")
	}
}

func compareTexts(targetText string, shift int) bool {
	targetText = strings.ToLower(targetText)
	encryptText, _ := CipherCaesar([]rune(targetText), shift)
	text, _ := CipherCaesar([]rune(encryptText), -shift)
	return text == targetText
}
