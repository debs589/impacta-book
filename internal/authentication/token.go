package authentication

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func CreateToken(userID int) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userID"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}
