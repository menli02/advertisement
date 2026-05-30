package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	AccessTokenExpiry  = 3600 * time.Second  // 1 hour
	RefreshTokenExpiry = 86400 * time.Second // 24 hours
	TempTokenExpiry    = 300 * time.Second   // 5 minutes
)

type Claims struct {
	UserId    int64  `json:"userId"`
	UserType  string `json:"userType"`
	TokenType string `json:"tokenType"`
	jwt.RegisteredClaims
}

func NewClaims(userId int64, userType, tokenType string) *Claims {
	var expiry time.Duration
	if tokenType == "refresh" {
		expiry = RefreshTokenExpiry
	} else {
		expiry = AccessTokenExpiry
	}
	now := time.Now()
	return &Claims{
		UserId:    userId,
		UserType:  userType,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(expiry)),
		},
	}
}
