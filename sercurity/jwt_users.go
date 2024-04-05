package sercurity

import (
	"app/usecases/dto"
	"fmt"

	"github.com/golang-jwt/jwt"
	"time"
)

const SECRET_KEY_USERS = "learngolanglalalafdfds"

type JwtCustomClaimsUsers struct {
	UserId string
	Role   string
	jwt.StandardClaims
}

func GenTokenUsers(user dto.Admin) (string, error) {
	claims := &JwtCustomClaims{
		UserId: user.UserId,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(SECRET_KEY_USERS))
	if err != nil {
		fmt.Println("Loi tao token", err.Error())
		return "Tao Token Loi!", err
	}

	return result, nil

}
