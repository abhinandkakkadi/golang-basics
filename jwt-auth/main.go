package main

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	jwt.StandardClaims
	UserID string
}

func main() {

	// set custom claims
	u := UserClaims{
		UserID: "abhinand",
	}
	token,err := createToken(&u)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("generated token = ",token)
}




// generating token
func createToken(c *UserClaims) (string,error) {

	// this will create a base token by passing signing method and custom claims which we created
	token := jwt.NewWithClaims(jwt.SigningMethodHS512,c)

	// then the base token is signed using the secret key which should be of type slice of bytes
	// this returns a signed token which can be sent to the user side, so that they can use that token for auhtorization.
	tokenString,err := token.SignedString([]byte("abhinand"))
	if err != nil {
		return "",err
	}

	return tokenString,nil
} 