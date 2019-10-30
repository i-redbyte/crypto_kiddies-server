package crypt

import (
	trans "cryptokiddies-server/crypt/transposition"
	"errors"
)

const (
	transparent = "transposition"
	caesar      = "ciphercaesar"
)

func GetCryptoText(slug string, text string, key string) (*string, error) {
	switch slug {
	case transparent:
		text, err := trans.Encrypt([]rune(text), key)
		if err != nil {
			return nil, err
		}
		return &text, nil
	case caesar:
		return nil, errors.New("шифр не найден")
	default:
		return nil, errors.New("шифр не найден")
	}
}

func GetDecryptText(slug string, text string, key string) (*string, error) {
	switch slug {
	case transparent:
		text, err := trans.Decrypt([]rune(text), key)
		if err != nil {
			return nil, err
		}
		return &text, nil
	case caesar:
		return nil, nil
	default:
		return nil, errors.New("текст не найден")
	}

}
