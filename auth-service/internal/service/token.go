package service

import (
	"auth-service/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func issueToken(claims Claims) (Token, error) {
	t := Token{}

	// issue JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": claims.UserID,
		"exp":    time.Now().Add(time.Hour * 1),
	})

	tokenString, err := token.SignedString([]byte(service.secret))
	if err != nil {
		return t, err
	}
	t.JWTToken = tokenString

	// Issue refresh Token
	refreshToken := uuid.NewString()
	data := repository.Token{
		RefreshToken: refreshToken,
		Expiry:       7 * (time.Hour * 24),
		UserID:       claims.UserID,
	}
	err = service.model.Token.CreateToken(data)
	if err != nil {
		return t, err
	}

	t.RefreshToken = refreshToken
	return t, nil

}

func validateToken(token string) (bool, error) {

	validatedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		return []byte(service.secret), nil
	})
	if err != nil {
		return false, err
	}
	if validatedToken.Valid {
		return true, nil
	}
	return false, nil

}
