package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	`github.com/mitchellh/mapstructure`
	"net/http"
)

// returns the user id that are logged in
func CheckAuth(c *gin.Context) int {
	tokenString := c.Request.Header.Get(`Authorization`)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod(`HS256`) != token.Method {
			return nil, fmt.Errorf(`unexpected signing method: %v`, token.Header[`alg`])
		}
		return []byte(`secret`), nil
	})
	if token != nil && err == nil {
		fmt.Println(`token verified`)
		claims := token.Claims.(jwt.MapClaims)
		fmt.Println(claims)
		var idAccount int
		err := mapstructure.Decode(claims["account_number"], &idAccount)
		if err != nil {
			result := gin.H{
				"message": err.Error(),
			}
			c.JSON(http.StatusUnauthorized, result)
			c.Abort()
		}
		return idAccount
	} 
	result := gin.H{
		"message": "token tidak valid",
		"error":   err.Error(),
	}
	c.JSON(http.StatusUnauthorized, result)
	c.Abort()
	return -1
}

