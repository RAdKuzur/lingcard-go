package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

func GenerateAccessToken(name string) (string, error) {

	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	var secretKey = []byte(os.Getenv("SECRET_KEY"))
	var minutes, err = strconv.Atoi(os.Getenv("ACCESS_TOKEN_TIME_EXPIRE"))
	if err != nil {
		return "", err
	}
	claims := jwt.MapClaims{
		"username": name,
		"exp":      time.Now().Add(time.Minute * time.Duration(minutes)).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func GenerateRefreshToken(name string) (string, error) {

	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	var secretKey = []byte(os.Getenv("SECRET_KEY"))
	var minutes, err = strconv.Atoi(os.Getenv("REFRESH_TOKEN_TIME_EXPIRE"))
	if err != nil {
		return "", err
	}
	claims := jwt.MapClaims{
		"username": name,
		"exp":      time.Now().Add(time.Minute * time.Duration(minutes)).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {

	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}
	secretKey := []byte(os.Getenv("SECRET_KEY"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
