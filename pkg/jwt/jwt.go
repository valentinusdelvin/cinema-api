package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"log"
	"os"
	"strconv"
	"time"
)

type JwtService interface {
	GenerateToken(UserID uuid.UUID, IsAdmin bool, IsSubscribed bool) (string, error)
	ValidateToken(tokenString string) (uuid.UUID, bool, bool, error)
}
type jwtService struct {
	JwtKey string
	JwtExp time.Duration
}
type Claims struct {
	UserID       uuid.UUID
	IsAdmin      bool
	IsSubscribed bool
	jwt.RegisteredClaims
}

func JWTInit() JwtService {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	JwtExp, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION_TIME"))
	if err != nil {
		log.Fatalf("Error parsing JWT_EXPIRATION_TIME: %v", err)
	}
	return &jwtService{
		JwtKey: secretKey,
		JwtExp: time.Duration(JwtExp) * time.Minute,
	}
}

func (j *jwtService) GenerateToken(UserID uuid.UUID, IsAdmin bool, IsSubscribed bool) (string, error) {
	claims := Claims{
		UserID:       UserID,
		IsAdmin:      IsAdmin,
		IsSubscribed: IsSubscribed,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.JwtExp)),
			Issuer:    "Cinema API",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.JwtKey))
}

func (j *jwtService) ValidateToken(tokenString string) (uuid.UUID, bool, bool, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.JwtKey), nil
	})

	if err != nil {
		return claims.UserID, false, false, err
	}

	if !token.Valid {
		return claims.UserID, false, false, err
	}

	return claims.UserID, claims.IsAdmin, claims.IsSubscribed, nil
}
