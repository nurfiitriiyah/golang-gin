package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
	"time"
	"unicode/utf8"
)

func auth(c *gin.Context) {
	authorizationHeader := c.Request.Header.Get("Authorization")
	if authorizationHeader != "" {
		bearerToken := strings.Split(authorizationHeader, " ")
		token, err := parseBearerToken(bearerToken[0])
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
		} else {
			decoded := token.Claims
			c.JSON(http.StatusOK, decoded)
		}

	}
}

func parseBearerToken(bearerToken string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET_KEY")
	bearersToken := strings.Split(bearerToken, " ")
	return jwt.Parse(bearersToken[0], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("error")
			return nil, fmt.Errorf("There was an error")
		}
		return []byte(secret), nil
	})
}

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

func encrypt(s string) string {
	myPassword := s
	bv := []byte(myPassword)
	hash := sha256.New()
	hash.Write(bv)
	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)
	return string(mdStr)
}

func createIds(s string) string {
	dt := time.Now()
	createId := s + dt.Format("01022006150405")
	return string(createId)
}
