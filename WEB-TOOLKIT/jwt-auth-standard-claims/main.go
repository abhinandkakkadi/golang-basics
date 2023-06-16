package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func main() {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
		Subject:   "just a normal token",
	}
	signedToken, err := createToken(claims)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("signed token : ", signedToken)

	claimsReturned, err := GenerateToken(signedToken)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("claims returned : ", claimsReturned)
}

func createToken(claims *jwt.StandardClaims) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	signedToken, err := token.SignedString([]byte("abhinand"))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GenerateToken(signedToken string) (*jwt.StandardClaims, error) {

	token, err := jwt.ParseWithClaims(signedToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {

		// here we can get the identifier and in this case we can directly use the variable and not the type of the variable
		if token.Method.Alg() != jwt.SigningMethodHS512.Alg() {
			return nil, errors.New("signing method error")
		}

		return []byte("abhinand"), nil
	})

	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*jwt.StandardClaims)
	return claims, nil
}
