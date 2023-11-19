package security

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	GenerateToken(username string, email string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type ChatCustomClaims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

func NewJWTService() JWTService {
	return &jwtServices{
		secretKey: "jdnfksdmfksd", // TODO BETTER
		issure:    "chat_app",
	}
}

func (j *jwtServices) GenerateToken(username string, email string) (signed string, err error) {
	// Creating claims
	claims := ChatCustomClaims{
		username,
		email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    j.issure,
		},
	}

	// Create and sign token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err = token.SignedString(j.secretKey)

	return
}

func (j *jwtServices) ValidateToken(signedToken string) (token *jwt.Token, err error) {
	token, err = jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(j.secretKey), nil
	})

	return
}
