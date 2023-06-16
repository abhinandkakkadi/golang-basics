package main

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
)

func main() {

	claims := &jwt.MapClaims{
		"one": "abhinand",
		"two": "athira",
	}

	signedToken, err := GenerateToken(claims)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("signed token := ", signedToken)

	claimsReturned, err := ParseToken(signedToken)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Map claims : ", claimsReturned)
}

func GenerateToken(claims *jwt.MapClaims) (string, error) {

	// here jwt.SigningMethodHS512 is a pointer variable of type *SigningMethodHMAC
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	signedToken, err := token.SignedString([]byte("abhinand"))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ParseToken(signedToken string) (*jwt.MapClaims, error) {

	token, err := jwt.ParseWithClaims(signedToken, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {

		// here we have to assert value to *jwt.SigningMethodHMAC and not jwt.SigningMethodHS512, as the later one is not a type
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("signing method error")
		}

		return []byte("abhinand"), nil
	})

	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*jwt.MapClaims)
	return claims, nil
}
