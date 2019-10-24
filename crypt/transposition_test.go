package crypt

import (
	"math/rand"
	"testing"
)

func getTexts() []string {
	return []string{
		"Lenin",
		"A slice literal is declared just like an array literal, except you leave out the element count",
		"Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.",
		"Go’s treatment of errors as values has served us well over the last decade. Although the standard library’s support for errors has been minimal—just the errors.New and fmt.Errorf functions, which produce errors that contain only a message—the built-in error interface allows Go programmers to add whatever information they desire. All it requires is a type that implements an Error method:",
		"А тут для примера русский текст",
		"Осенью 1888 года Ульянову было разрешено вернуться в Казань. Здесь он впоследствии вступил в один из марксистских кружков, организованных Н. Е. Федосеевым, где изучались и обсуждались сочинения К. Маркса, Ф. Энгельса и Г. В. Плеханова. В 1924 году Н. К. Крупская писала в «Правде»: «Плеханова Владимир Ильич любил страстно. Плеханов сыграл крупную роль в развитии Владимира Ильича, помог ему найти правильный революционный подход, и потому Плеханов был долгое время окружен для него ореолом: всякое самое незначительное расхождение с Плехановым он переживал крайне болезненно»",
	}
}

func getRandomString() string {
	rusRune := []rune(Rus)
	b := make([]rune, rand.Intn(100))
	for i := range b {
		b[i] = rusRune[rand.Intn(len(rusRune))]
	}
	return string(b)
}

func TestEncrypt(t *testing.T) {
	fn := func(text string, keyWord string) (bool, error) {
		encrypt, err := Encrypt([]rune(text), keyWord)
		return text == encrypt, err
	}
	for _, s := range getTexts() {
		if check, err := fn(s, getRandomString()); check || err != nil {
			t.Error("Строка ", s, " не зашифровалась")
		}
	}
}

func TestDecrypt(t *testing.T) {

}
