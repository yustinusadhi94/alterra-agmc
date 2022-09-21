package middlewares

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
	"time"
)

var secretKey string

type CustomClaims struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}

func InitJWTSecretKey() {
	secretKey = os.Getenv("SECRET_KEY")
}

func CreateToken(userId int) (string, error) {
	claims := CustomClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	//claims := jwt.StandardClaims{
	//	ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	//	Id:        userId,
	//	IssuedAt:  time.Now().Unix(),
	//}
	//claims["authorized"] = true
	//claims["userId"] = userId
	//claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func ExtractTokenUserId(e echo.Context) int {
	token := e.Get("user").(*jwt.Token)
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		userId := int(claims["user_id"].(float64))
		return userId
	}
	return 0
}

func ValidateToken() echo.MiddlewareFunc {
	return middleware.JWT([]byte(secretKey))
}
