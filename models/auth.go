package models

import (
	"github.com/dgrijalva/jwt-go"
)

type MyCustomClaims struct {
	UId int64 `json:"uid"`
	jwt.StandardClaims
}
