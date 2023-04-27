package helper

import (
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

func GenerateToken() (string, error) {
	secretKey := viper.GetString(`secretKey`)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":  "62teknologi-senior-backend-test-muhammad-hajid-al-akhtar",
		"role": 1,
	})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
