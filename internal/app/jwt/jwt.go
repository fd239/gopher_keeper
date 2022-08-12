package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTManager struct {
	secret   string
	duration time.Duration
}

func NewJWTManager(secret string, duration time.Duration) *JWTManager {
	return &JWTManager{
		secret:   secret,
		duration: duration,
	}
}

type Payload struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

func (manager *JWTManager) GenerateJWT(userId string) (string, error) {
	claims := Payload{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.duration).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "server",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(manager.secret))
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}

	return tokenString, nil
}

//VerifyToken verifies given token
func (manager *JWTManager) VerifyToken(accessToken string) (*Payload, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&Payload{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return []byte(manager.secret), nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	if token.Valid {
		claims, ok := token.Claims.(*Payload)
		if !ok {
			return nil, fmt.Errorf("invalid token claims")
		}
		return claims, err
	}

	return nil, err
}
