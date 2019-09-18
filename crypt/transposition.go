package crypt

/*
Реализация шифра перестановки
*/

var key []int

func SetKey(intKey []int) {
	key = make([]int, len(intKey))
	copy(key, intKey)
}

func Encrypt(text string) string {
	r := []rune(text)
	textLength := len(r)
	keyLength := len(key)
	n := textLength % keyLength
	for i := 0; i < keyLength-n; i++ {
		text += " "
	}
	r = []rune(text)
	textLength = len(r)
	result := ""

	for i := 0; i < textLength; i += keyLength {
		transposition := make([]byte, keyLength)
		for j := 0; j < keyLength; j++ {
			transposition[key[j]-1] = text[i+j]
		}
		result += string(transposition)
	}
	return result
}

func Decrypt(text string) string {
	result := ""
	r := []rune(text)
	textLength := len(r)
	keyLength := len(key)
	for i := 0; i < textLength; i += keyLength {
		transposition := make([]byte, keyLength)
		for j := 0; j < keyLength; j++ {
			transposition[j] = text[i+key[j]-1]
		}
		result += string(transposition)
	}
	return result
}
