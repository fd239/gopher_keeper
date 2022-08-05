package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

//
//import (
//	"fmt"
//	"github.com/fd239/gopher_keeper/internal/models"
//	"github.com/golang-jwt/jwt"
//	"time"
//)
//
//// Manager is a JSON web token manager
//type Manager struct {
//	secret   string
//	duration time.Duration
//	claims   *UserClaims
//}
//
//// New returns a new JWT manager
//func New(secretKey string, tokenDuration time.Duration) *Manager {
//	return &Manager{secretKey, tokenDuration, &UserClaims{}}
//}
//
//// UserClaims is a custom JWT claims that contains some user's information
//type UserClaims struct {
//	jwt.StandardClaims
//	Username string `json:"username"`
//	Role     string `json:"role"`
//	Id       int    `json:"id"`
//}
//
//// Claims return current user claims
//func (m *Manager) Claims() *UserClaims {
//	return m.claims
//}
//
//// Generate generates and signs a new token for a user
//func (m *Manager) Generate(user *models.User) (string, error) {
//	claims := UserClaims{
//		StandardClaims: jwt.StandardClaims{
//			ExpiresAt: time.Now().Add(m.tokenDuration).Unix(),
//		},
//		Username: user.Name,
//		Role:     user.Role,
//		Id:       user.Id,
//	}
//
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//
//	return token.SignedString([]byte(m.secretKey))
//}
//
//// Verify verifies the access token string and return a user claim if the token is valid
//func (m *Manager) Verify(accessToken string) (*UserClaims, error) {
//	token, err := jwt.ParseWithClaims(
//		accessToken,
//		&UserClaims{},
//		func(token *jwt.Token) (interface{}, error) {
//			_, ok := token.Method.(*jwt.SigningMethodHMAC)
//			if !ok {
//				return nil, fmt.Errorf("unexpected token signing method")
//			}
//
//			return []byte(m.secretKey), nil
//		},
//	)
//
//	if err != nil {
//		return nil, fmt.Errorf("invalid token: %w", err)
//	}
//
//	claims, ok := token.Claims.(*UserClaims)
//	if !ok {
//		return nil, fmt.Errorf("invalid token claims")
//	}
//
//	m.claims = claims
//
//	return claims, nil
//}

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

type JWTPayload struct {
	jwt.StandardClaims
	UserId uint `json:"user_id"`
}

func (manager *JWTManager) GenerateJWT(userId uint) (string, error) {
	claims := &JWTPayload{
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

func (manager *JWTManager) ValidateJWT(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %v", token.Header["alg"])
		}
		return []byte(manager.secret), nil
	})
}
