package transposition

import (
	"errors"
	"sort"
	"strings"
)

/*
Реализация шифра перестановки
*/

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

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

func Encrypt(text []rune, keyWord string) (string, error) {
	key := getKey(keyWord)
	space := rune(' ')
	keyLength := len(key)
	textLength := len(text)
	if keyLength <= 0 {
		return "", errors.New("отсутствует ключ")
	}
	if textLength <= 0 {
		return "", errors.New("отсутствует текст шифрования")
	}
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
	return result, nil
}

func Decrypt(text []rune, keyWord string) (string, error) {
	key := getKey(keyWord)
	textLength := len(text)
	keyLength := len(key)
	space := rune(' ')
	n := textLength % keyLength
	for i := 0; i < keyLength-n; i++ {
		text = append(text, space)
	}
	if keyLength <= 0 {
		return "", errors.New("отсутствует ключ")
	}
	if textLength <= 0 {
		return "", errors.New("отсутствует текст шифрования")
	}
	result := ""
	for i := 0; i < textLength; i += keyLength {
		transposition := make([]rune, keyLength)
		for j := 0; j < keyLength; j++ {
			transposition[j] = text[i+key[j]-1]
		}
		result += string(transposition)
	}
	return result, nil
}
