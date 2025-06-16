package middleware

import (
	"api-gateway/internal/util"
	"fmt"
	"net/http"
	"strings"

	"github.com/akanshgupta98/go-logger/v2"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var JWTsecret string

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Infof("API: %s Method: %s Host: %s", c.Request.URL, c.Request.Method, c.Request.Host)
		c.Next()
	}
}

func AuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract bearer token.
		JWTsecret = secret
		token, err := extractBearerToken(c.Request.Header.Get("Authorization"))
		if err != nil {
			util.ErrorJSON(c, fmt.Errorf("no token found"), nil, http.StatusUnauthorized)
			return
		}

		valid, err := validateToken(token)
		if err != nil {
			util.ErrorJSON(c, err, nil, http.StatusUnauthorized)
			return
		}
		if !valid {
			util.ErrorJSON(c, fmt.Errorf("invalid token"), nil, http.StatusUnauthorized)
			return
		} else {
			logger.Debugf("authenticated successfully!! Processing request")
		}

		c.Next()
	}
}

func extractBearerToken(header string) (string, error) {
	logger.Debugf("authorization header is: %s", header)
	pos := strings.Index(header, " ")
	logger.Debugf("pos of bearer is: %d", pos)
	if pos == -1 {
		return "", fmt.Errorf("empty bearer Auth Value")

	} else {
		logger.Debugf("returned token is: %s", header[pos+1:])
		return header[pos+1:], nil
	}

}

func validateToken(token string) (bool, error) {

	validatedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		return []byte(JWTsecret), nil
	})
	if err != nil {
		return false, err
	}
	if validatedToken.Valid {
		return true, nil
	}
	return false, nil

}
