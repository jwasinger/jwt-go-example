package main

import (
	"fmt"
	"encoding/hex"
	"time"
	"github.com/golang-jwt/jwt/v4"
)

func main() {
	var claims jwt.RegisteredClaims
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": &jwt.NumericDate{Time: time.Now()},
	})

	// Sign and get the complete encoded token as a string using the secret
	secretHex := "5e6f246df54a2b48ca21cd56c7aeba4da5630de195797df9806ddba3e0b02ee9"
	secret, err := hex.DecodeString(secretHex)
	if err != nil {
		panic(err)
	}

	tokenString, err := token.SignedString(secret)
	var keyFunc = func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	}

        _, err = jwt.ParseWithClaims(tokenString, &claims, keyFunc,
                jwt.WithValidMethods([]string{"HS256"}),
                jwt.WithoutClaimsValidation())

	if err != nil {
		panic(err)
	}
	fmt.Println(tokenString)
}
