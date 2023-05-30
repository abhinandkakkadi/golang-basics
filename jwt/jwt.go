package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	// "github.com/golang-jwt/jwt/v4"
)

// in case of HMAC same key is used for signing as well as to check whether the token is valid
// in all the other signing method provided my go, there will be a private kry used for signing and a public key which is used to check the validity of the key
// the other one that i mentioned above was ECDSA and RSA which are kind of identical

type userClaims struct {
	jwt.StandardClaims
	SessionID int64
}

func (u *userClaims) Valid() error {
	
	if !u.VerifyExpiresAt(time.Now().Unix(),true) {
		return fmt.Errorf("token has expired")
	}

	if u.SessionID == 0 {
		return fmt.Errorf("invalid session ID")
	}

	return nil
}


func CreateToken(c *userClaims) (string,error) {
	
	token := jwt.NewWithClaims(jwt.SigningMethodES256,c)

	signedToken,err := token.SignedString([]byte("abhinand"))
	if err != nil {
		return "",err
	}
	return signedToken,nil
}
 