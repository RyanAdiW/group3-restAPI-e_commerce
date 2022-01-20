package middlewares

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte("rahasia"),
	})
}

func CreateToken(userid int, userName string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = userid
	claims["user_name"] = userName
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("rahasia"))
}

func GetUserName(e echo.Context) (string, error) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userName := claims["user_name"].(string)
		if userName == "" {
			return userName, fmt.Errorf("empty user_name")
		}
		return userName, nil
	}
	return "", fmt.Errorf("invalid user")
}

func GetId(jwtSecret string, e echo.Context) (int, error) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userid := int(claims["id"].(float64))
		if userid == 0 {
			return userid, fmt.Errorf("invalid id")
		}
		return userid, nil
	}
	return 0, fmt.Errorf("invalid user")
}
