package main

import (
	"io/ioutil"
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

	val, err := ioutil.ReadFile("jwt_secret")
	if err != nil {
		panic("could not read jwt secret")
	}

	// Sign and get the complete encoded token as a string using the secret
	secretHex := string(val)
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
