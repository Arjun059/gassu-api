package utils

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

func CreateToken(user_email string, user_id uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_email": user_email,
			"user_id":    user_id,
			"exp":        time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	// Log the token details
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		log.Printf("Token is valid. Claims: %v", claims)
		return nil
	} else {
		log.Printf("Invalid token or claims: %v", token)
		return errors.New("invalid token or claim")
	}
}
