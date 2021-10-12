// parse token & validate (validate issuer) http://happy-lifetime-iam-staging.azurewebsites.net
package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var hmacSampleSecret []byte

func Authorize(f func(c *gin.Context), roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Gets token from header
		token := c.GetHeader("Authorization")

		// Verifies token
		token, err := verifyToken(token)
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, err.Error())
			return
		}

		// Validates token
		parsedToken, err := validateToken(token)
		if err != nil {
			//c.IndentedJSON(http.StatusUnauthorized, err.Error())
			//return
		}

		// Verifies the roles
		if err := verifyRole(parsedToken, roles); err != nil {
			c.IndentedJSON(http.StatusUnauthorized, err.Error())
			return
		}

		// Adds the sub to the context
		c.Set("sub", parsedToken.Claims.(jwt.MapClaims)["sub"])

		// Calls the next handler in chain
		f(c)
	}
}

// Verifies that given user is in given role
func verifyRole(token *jwt.Token, roles []string) error {
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if len(roles) > 0 {
			role := claims["role"].([]interface{})
			okk := false
			for r := range role {
				if roles[0] == role[r].(string) {
					okk = true
				}
			}
			if !okk {
				return errors.New("auth/missing-role?role=" + roles[0])
			}
		}
	}
	return nil
}

// Verifies that given token is not empty and is of correct format
func verifyToken(token string) (resp string, error error) {

	// Verifies that a token has been given
	if token == "" {
		return "", errors.New("auth/missing-token")
	}

	// Verifies that the token starts with bearer
	if !strings.Contains(strings.ToLower(token), "bearer") {
		return "", errors.New("auth/wrong-token-format")
	}

	// Returns the substring containing only the token
	return token[7:], nil
}

// Validate token
func validateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return hmacSampleSecret, nil
	})
}
