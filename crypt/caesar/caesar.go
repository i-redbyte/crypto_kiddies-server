package caesar

import (
	"errors"
	"unicode"
)

//noinspection GoSnakeCaseUsage
const (
	RUS            = 32
	ENG            = 25
	RUS_LOWER_CASE = 1072
	ENG_LOWER_CASE = 97
)

func Encrypt(text []rune, shift int) (string, error) {
	if len(text) == 0 {
		return "", errors.New("отсутствует текст шифрования")
	}
	for i, r := range text {
		text[i] = unicode.ToLower(r)
	}
	result := make([]rune, len(text))

	for i, r := range text {

		if r >= 'а' && r <= 'я' {
			c := (int(r) + shift - RUS_LOWER_CASE) % RUS
			result[i] = rune(c + RUS_LOWER_CASE)
		}

		if r >= 'a' && r <= 'z' {
			c := (int(r) + shift - ENG_LOWER_CASE) % ENG
			result[i] = rune(c + ENG_LOWER_CASE)
		}
	}
	return string(result), nil
}
