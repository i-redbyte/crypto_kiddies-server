package storage

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/ilya-sokolov/crypto_kiddies-server/app/rest/errors"
	. "github.com/ilya-sokolov/crypto_kiddies-server/database"
	"golang.org/x/crypto/bcrypt"
	"os"
)

type Account struct {
	BaseModel
	NickName string `json:"nickName"`
	Password []byte `json:"-"`
}

func (account *Account) Token() string {
	tk := &Token{Id: account.Id}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	return tokenString
}

type Token struct {
	Id int64
	jwt.StandardClaims
}

func ParseToken(t string, secret []byte) (*Token, error) {
	tk := &Token{}
	token, err := jwt.ParseWithClaims(t, tk, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if token == nil || !token.Valid {
		return nil, errors.InvalidToken
	}
	return tk, err
}

func CreateAccount(nickName string, pass string) (*Account, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	acc := &Account{NickName: nickName, Password: passwordHash}
	err = DB.Create(acc).Error
	if err != nil {
		return nil, err
	}
	return acc, nil
}

func GetAccount(nickName string, pass string) (*Account, error) {
	acc := &Account{}
	if err := DB.Model(&Account{}).Where("nick_name = ?", nickName).First(acc).Error; err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword(acc.Password, []byte(pass)); err != nil {
		return nil, errors.PasswordNotMatch
	}
	return acc, nil
}

func GetAccountById(id int64) (*Account, error) {
	acc := &Account{}
	if err := DB.Model(acc).Where("id = ?", id).First(acc).Error; err != nil {
		return nil, err
	}
	return acc, nil
}
