package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecret"

func GenerateToken(email string, userId int64) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"userId": userId,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	}) //generate new token with data attached to it.
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) error {
	// in golang -> interface{} means any type.
	// claims -> information in token.
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return secretKey, nil
	}) //parses the received token and extracts it's information.
	if err != nil {
		return errors.New("Could not parse token")
	}
	tokenIsValid := parsedToken.Valid
	if !tokenIsValid{
		return errors.New("Invalid token")
	}
	// claims, ok := parsedToken.Claims(jwt.MapClaims)
	// if !ok {
	// 	return errors.New("Invalid token claims")
	// }
	// email := claims["email"].(string)
	// userId := claims["userId"].(int64)
	return nil
}