package crypt

import (
	csr "cryptokiddies-server/crypt/caesar"
	trans "cryptokiddies-server/crypt/transposition"
	"errors"
	"fmt"
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
		text, err := csr.CipherCaesar([]rune(text), key)
		if err != nil {
			fmt.Println("ERROR!", err)
			return nil, err
		}
		return &text, nil
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
		text, err := csr.CipherCaesar([]rune(text), key)
		if err != nil {
			return nil, err
		}
		return &text, nil
	default:
		return nil, errors.New("текст не найден")
	}

}
