package models

import (
	u "cryptokiddies-server/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type Account struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
}

func (account *Account) Validate() (map[string]interface{}, bool) {
	// TODO: Red_byte change for good validation
	if !strings.Contains(account.Email, "@") {
		return u.Message(false, "Введите корректный E-mail"), false
	}
	if len(account.Password) < 6 {
		return u.Message(false, "Слишком короткий пароль."), false
	}

	temp := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Ошибка соединения. Пожалуйста повторите"), false
	}
	if temp.Email != "" {
		return u.Message(false, "Адрес электронной почты уже используется другим пользователем."), false
	}
	return u.Message(false, "Ok!"), true
}

func (account *Account) Create() map[string]interface{} {
	if resp, ok := account.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	GetDB().Create(account)

	if account.ID <= 0 {
		return u.Message(false, "Не удалось создать аккаунт, ошибка подключения.")
	}

	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString
	account.Password = ""

	response := u.Message(true, "Аккаунт создан")
	response["account"] = account
	return response
}

func Login(email, password string) map[string]interface{} {
	account := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "E-mail не найден")
		}
		return u.Message(false, "Ошибка соединения. Пожалуйста повторите")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message(false, "Неверные логин или пароль. Пожалуйста, попробуйте еще раз")
	}
	account.Password = ""

	//Create JWT
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	resp := u.Message(true, "Успешный вход в систему")
	resp["account"] = account
	return resp
}

func GetAccount(u uint) *Account {

	acc := &Account{}
	GetDB().Table("accounts").Where("id = ?", u).First(acc)
	if acc.Email == "" {
		return nil
	}

	acc.Password = ""
	return acc
}
