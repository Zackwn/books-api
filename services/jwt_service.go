package services

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWT struct {
	secret        []byte
	issuer        string
	tokenLifetime time.Duration
}

func NewJWT() JWT {
	secret := os.Getenv("JWT_SECRET")
	issuer := os.Getenv("JWT_ISSUER")
	if secret == "" || issuer == "" {
		log.Fatal("missing JWT env")
	}

	return JWT{
		secret:        []byte(secret),
		issuer:        issuer,
		tokenLifetime: time.Hour * 2,
	}
}

type Claim struct {
	Sub uint `json:"sub"`
	jwt.StandardClaims
}

func (s JWT) Sign(id uint) (string, error) {
	claim := Claim{
		Sub: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.tokenLifetime).Unix(),
			Issuer:    s.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	stoken, err := token.SignedString(s.secret)
	if err != nil {
		return "", err
	}
	return stoken, nil
}

func (s JWT) Verify(token string) (uint, error) {
	var claim Claim
	_, err := jwt.ParseWithClaims(token, &claim, func(t *jwt.Token) (interface{}, error) {
		return s.secret, nil
	})
	if err != nil {
		return 0, err
	}
	return claim.Sub, nil
}
