package main

import (
	"errors"
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
	signedToken,err := createToken(&u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("signed token = ",signedToken)
	claims,err := parseToken(signedToken)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("claims = ",claims)
}




// generating token
func createToken(c *UserClaims) (string,error) {

	// this will create a base token by passing signing method and custom claims which we created
	token := jwt.NewWithClaims(jwt.SigningMethodHS512,c)

	// then the base token is signed using the secret key which should be of type slice of bytes
	// this returns a signed token which can be sent to the user side, so that they can use that token for auhtorization.
	signedToken,err := token.SignedString([]byte("abhinand"))
	if err != nil {
		return "",err
	}

	return signedToken,nil
} 


// a function to parse the signed token and give back the claims
// if the signature does not match or if the token expire. It will be an error
func parseToken(signedToken string) (*UserClaims,error) {
	
	t,err := jwt.ParseWithClaims(signedToken,&UserClaims{},func (t *jwt.Token) (interface{},error){

		if t.Method.Alg() != jwt.SigningMethodHS512.Alg() {
			return nil,errors.New("signing algorithm error")
		}

		return []byte("abhinand"),nil
	})

	if err != nil {
		return nil,errors.New("error while checking signing method")
	}

	if !t.Valid {
		return nil,errors.New("error in parse token is not valid")
	}

	claims := t.Claims.(*UserClaims)
	return claims,nil
}

// if we were to use standard claims instead of custom claims. the second argument of jwt.ParseWithClaims() will be &jwt.StandardClaims{} instance and when trying to retrieve 
// claims from token, we will write token.Claims.(*jwt.StandardClaims) instead of giving token.Claims.(*jwt.customClaims)
// also we ca use map claims. in that insc
