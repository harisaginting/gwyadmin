package helper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	httpModel "github.com/harisaginting/gwyn/models/http"
	"github.com/harisaginting/gwyn/utils/jwt/generator"
	"golang.org/x/crypto/bcrypt"
)

var (
	AppName string
	JWTKey  string
)

func init() {
	AppName = os.Getenv("APP_NAME")
	JWTKey = "TESTJWTKEY"
}

func HashPassword(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	return string(hash), err
}

func ComparePasswords(hashedPwd, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	bytePlainPwd := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlainPwd)
	if err != nil {
		return false
	}
	return true
}

func GenerateToken(username, role, bd string) (signedToken string, err error) {
	expireAt := time.Now().Add(time.Hour * 72)
	tokenKey := generator.GenerateIdentifier()
	claims := httpModel.AuthClaim{
		Username: username,
		Role:     role,
		Bd:       bd,
		TokenKey: tokenKey,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireAt),
			Issuer:    AppName,
		},
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)
	signedToken, err = token.SignedString([]byte(JWTKey))
	return
}
