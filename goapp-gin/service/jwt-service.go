package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(name string, isAdmin bool) string
	ValidateToken(stringToken string) (*jwt.Token, error)
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "shivaganesh@gmail.com",
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		secretKey = "secretKey"
	}
	return secretKey
}

func (service *jwtService) GenerateToken(username string, isAdmin bool) string {
	claims := &jwtCustomClaims{
		username,
		isAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    service.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return tokenString
}

func (service *jwtService) ValidateToken(stringToken string) (*jwt.Token, error) {
	return jwt.Parse(stringToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(service.secretKey), nil
	})
}
