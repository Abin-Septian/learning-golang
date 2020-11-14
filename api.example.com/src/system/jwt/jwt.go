package jwt

import (
	"crypto/rsa"
	"errors"
	"io/ioutil"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	privateKeyPath = "./keys/app.rsa"
	pubKeyPath     = "./keys/app.rsa.pub"
	// HoursInday const
	HoursInday = 24
	// DaysInWeek const
	DaysInWeek = 7
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

func init() {
	signBytes, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		panic(err)
	}
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		panic(err)
	}
	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		panic(err)
	}
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		panic(err)
	}
}

// GetToken function
func GetToken(id int64) string {
	token := jwt.New(jwt.SigningMethodRS512)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * HoursInday * DaysInWeek).Unix()
	claims["iat"] = time.Now().Unix()
	claims["id"] = id
	token.Claims = claims

	tokenString, _ := token.SignedString(signKey)

	return tokenString
}

// IsTokenValid function
func IsTokenValid(val string) (int64, error) {
	token, err := jwt.Parse(val, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})

	switch err.(type) {
	case nil:
		if !token.Valid {
			return 0, errors.New("Token is invalid")
		}

		var userID int64

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return 0, errors.New("Token is invalid")
		}

		userID = int64(claims["id"].(float64))

		return userID, nil
	case *jwt.ValidationError:
		vErr := err.(*jwt.ValidationError)

		switch vErr.Errors {
		case jwt.ValidationErrorExpired:
			return 0, errors.New("Token Expired, get a new one")
		default:
			log.Println(vErr)
			return 0, errors.New("Error while Parsing Token")
		}
	default:
		return 0, errors.New("Unable to parse token")
	}
}
