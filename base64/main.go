package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

// Need
// sometimes when we need to store values in cookies some characters like double quotes won't be allowed, In those situation we can use base64 encoding (in case of json and all). So we can
// do that in two ways. either by specifying the encoding string(key), or by using standard encoding

func main() {

	s := "My name is ahbinand kakkadi and my home is in kannur"

	encodeStd := "abhadajndkajndakjsdn"
	// encoding using encoding string
	s64 := base64.NewEncoding(encodeStd).EncodeToString([]byte(s))

	// or we can do this without using encoding string
	// standard encoding
	s264 := base64.StdEncoding.EncodeToString([]byte(s))

	// inorder to decode to string

	bs,err := base64.StdEncoding.DecodeString(s64)
	if err != nil {
		log.Fatal("the error is ",err)
	}

	fmt.Println(s64,s264)
	fmt.Println(string(bs))

}