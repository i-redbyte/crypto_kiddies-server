package caesar

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

//noinspection GoSnakeCaseUsage
const (
	RUS            = 32
	ENG            = 26
	RUS_LOWER_CASE = 1072
	ENG_LOWER_CASE = 97
)

func CipherCaesar(text []rune, key string) (string, error) {
	shift, err := strconv.Atoi(key)
	if err != nil {
		return string(text), err
	}
	if len(text) == 0 {
		return string(text), errors.New("отсутствует текст шифрования")
	}
	for i, r := range text {
		text[i] = unicode.ToLower(r)
	}
	result := make([]rune, len(text))

	for i, r := range text {
		if r >= 'а' && r <= 'я' {
			c := (int(r) + shift - RUS_LOWER_CASE) % RUS
			result[i] = rune(c + RUS_LOWER_CASE)
		} else if r >= 'a' && r <= 'z' {
			c := (int(r) + shift - ENG_LOWER_CASE) % ENG
			if shift < 0 && (c+ENG_LOWER_CASE) < 'a' {
				c += ENG
			}
			result[i] = rune(c + ENG_LOWER_CASE)
		} else {
			result[i] = r
		}
	}
	return strings.ToLower(string(result)), nil
}
