package crypt

import (
	"sort"
	"strings"
)

/*
Реализация шифра перестановки
*/

var key []int

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func SetKey(intKey []int) {
	key = make([]int, len(intKey))
	copy(key, intKey)
}

func SetStringKey(stringKey string) {
	key = make([]int, GetLength(stringKey))
	key = getKey(stringKey)
}

func getKey(keyWord string) []int {
	keyWord = strings.ToLower(keyWord)
	word := []rune(keyWord)
	var sortedWord RuneSlice = make(RuneSlice, len(word))
	copy(sortedWord, word)
	sort.Sort(RuneSlice(sortedWord))
	usedLettersMap := make(map[string]int)
	wordLength := len(word)
	resultKey := make([]int, wordLength)
	for i := 0; i < wordLength; i++ {
		char := word[i]
		numberOfUsage := usedLettersMap[string(char)]
		resultKey[i] = getIndex(sortedWord, char) + numberOfUsage + 1 //+1 -so that indexing does not start at 0
		numberOfUsage++
		usedLettersMap[string(char)] = numberOfUsage
	}
	return resultKey
}

func getIndex(wordSet []rune, subString rune) int {
	n := len(wordSet)
	for i := 0; i < n; i++ {
		if wordSet[i] == subString {
			return i
		}
	}
	return 0
}

func GetLength(text string) int {
	r := []rune(text)
	return len(r)
}

func Encrypt(text []rune) string {
	space := rune(' ')
	keyLength := len(key)
	textLength := len(text)
	n := textLength % keyLength
	for i := 0; i < keyLength-n; i++ {
		text = append(text, space)
	}
	textLength = len(text)
	result := ""
	for i := 0; i < textLength; i += keyLength {
		transposition := make([]rune, keyLength)
		for j := 0; j < keyLength; j++ {
			transposition[key[j]-1] = text[i+j]
		}
		result += string(transposition)
	}
	return result
}

func Decrypt(text []rune) string {
	textLength := len(text)
	keyLength := len(key)
	result := ""
	for i := 0; i < textLength; i += keyLength {
		transposition := make([]rune, keyLength)
		for j := 0; j < keyLength; j++ {
			transposition[j] = text[i+key[j]-1]
		}
		result += string(transposition)
	}
	return result
}
