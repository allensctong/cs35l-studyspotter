package src

import (
	"fmt"
	"os"
	"time"
	"math/rand"
//	"crypto/rsa"
//	"crypto/x509"
//	"encoding/pem"
	"io/ioutil"
//	"log"

	"github.com/golang-jwt/jwt/v5"
)



/*//Load private key
var pwd, _ = os.Getwd()
var privKeyLoc = pwd + "/jwtRS256.key"
var pubKeyLoc = pwd + "/jwtRS256.key.pub"

var secretKey, _ = LoadPrivKey(privKeyLoc, pubKeyLoc)

//TODO HANDLE ERRORS IN THIS FUNCTION
func LoadPrivKey(rsaPrivateKeyLocation, rsaPublicKeyLocation string) (*rsa.PrivateKey, error) {
	priv, err := ioutil.ReadFile(rsaPrivateKeyLocation)
	if err != nil {
		log.Print("No RSA private key found")
		return nil, err
	}

	privPem, _ := pem.Decode(priv)
	privPemBytes := privPem.Bytes
	var parsedKey interface{}
	parsedKey, _ = x509.ParsePKCS1PrivateKey(privPemBytes)
	privateKey := parsedKey.(*rsa.PrivateKey)

	pub, err := ioutil.ReadFile(rsaPublicKeyLocation)
	if err != nil {
		log.Print("No RSA public key found")
		return nil, err
	}

	pubPem, _ := pem.Decode(pub)
	pubPemBytes := pubPem.Bytes
	parsedKey, _ = x509.ParsePKIXPublicKey(pubPemBytes)
	pubKey := parsedKey.(*rsa.PublicKey)
	privateKey.PublicKey = *pubKey

	return privateKey, nil
}*/




func CreateToken(username string) (string, error) {
	var pwd, _ = os.Getwd()
	var secretKey, _ = ioutil.ReadFile(pwd + "/jwtHS256.key")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
		jwt.MapClaims{
			"username": username,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})
	
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (string, error) {
	var pwd, _ = os.Getwd()
	var secretKey, _ = ioutil.ReadFile(pwd + "/jwtHS256.key")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}
	tokenString = token.Raw
	return tokenString, nil
}

func GenerateKey () {
	var pwd, _ = os.Getwd()
	b := make([]byte, 32)
    if _, err := rand.Read(b); err != nil {
		panic(err)
        return
    }
	err := os.WriteFile(pwd + "/jwtHS256.key", b, 0333)
	if err != nil {
		panic(err)
	}
}