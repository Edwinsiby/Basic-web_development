package helpers

import (
	"fmt"
	"main/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(user models.User, c *gin.Context) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"role": user.Role,
		"user": user.Name,
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("KEY")))

	if err == nil {
		fmt.Println("token created")
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorise", tokenString, 3600, "", "", false, true)

}

func ValidateCookie(c *gin.Context) bool {
	result := true
	cookie, _ := c.Cookie("Authorise")
	if cookie == "" {
		fmt.Println("cookie not found")
		result = false
	} else {
		fmt.Println("cookie", cookie)
		result = true
	}
	return result
}

func DeleteCookie(c *gin.Context) {
	c.SetCookie("Authorise", "", 0, "", "", true, true)
	fmt.Println("cookie deleted")
}

func FindRole(c *gin.Context) (string, string, error) {
	cookie, _ := c.Cookie("Authorise")
	if cookie == "" {
		fmt.Println("cookie not found")
		return "", "", fmt.Errorf("err")
	} else {
		token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("KEY")), nil
		})

		if err != nil {
			return "", "", err
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			role := claims["role"].(string)
			user := claims["user"].(string)
			return role, user, nil
		} else {
			return "", "", fmt.Errorf("invalid token")
		}

	}
}
