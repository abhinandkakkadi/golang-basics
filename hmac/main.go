package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)


func main() {

	c := getCode("test@example.com")
	fmt.Println(c)
	c = getCode("test@exampl.com")
	fmt.Println(c)
}

func getCode(s string) string {

	h := hmac.New(sha256.New,[]byte("secret-key"))
	io.WriteString(h,s)
	return fmt.Sprintf("%x",h.Sum(nil))

}

// when we want to store something on client's computer and if we want to verify that the stored value is not being tampered with. We can use hmac