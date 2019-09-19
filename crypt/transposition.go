package crypt

import (
	"sort"
	"strings"
)

/*
Реализация шифра перестановки
*/

var key []int

func SetKey(intKey []int) {
	key = make([]int, len(intKey))
	copy(key, intKey)
}

func SetStringKey(stringKey string) {
	key = make([]int, getLength(stringKey))
	key = getKey(stringKey)
}

func getKey(word string) []int {
	word = strings.ToLower(word)
	sortedWord := strings.Split(word, "")
	sort.Strings(sortedWord)
	usedLettersMap := make(map[string]int)
	wordLength := getLength(word)
	resultKey := make([]int, wordLength)
	for i := 0; i < wordLength; i++ {
		char := word[i]
		numberOfUsage := usedLettersMap[string(char)]
		resultKey[i] = getIndex(sortedWord, string(char)) + numberOfUsage + 1 //+1 -so that indexing does not start at 0
		numberOfUsage++
		usedLettersMap[string(char)] = numberOfUsage
	}
	return resultKey
}
func getIndex(wordSet []string, subString string) int {
	n := len(wordSet)
	for i := 0; i < n; i++ {
		if wordSet[i] == subString {
			return i
		}
	}
	return 0
}

func getLength(text string) int {
	r := []rune(text)
	return len(r)
}

func Encrypt(text string) string {
	keyLength := len(key)
	textLength := getLength(text)
	n := textLength % keyLength
	for i := 0; i < keyLength-n; i++ {
		text += " "
	}
	textLength = getLength(text)
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
