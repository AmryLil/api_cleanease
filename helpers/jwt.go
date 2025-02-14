package helpers

import (
	"api_cleanease/config"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

type JWT struct {
	signKey    string
	refreshKey string
}

func NewJWT(config config.ProgramConfig) JWTInterface {
	return &JWT{
		signKey:    config.SECRET,
		refreshKey: config.REFSECRET,
	}
}

func (j *JWT) GenerateJWT(userID string, userType string) map[string]any {
	var result = map[string]any{}
	var accessToken = j.GenerateToken(userID, userType)
	var refeshToken = j.generateRefreshToken(userID, userType)
	if accessToken == "" {
		return nil
	}
	result["access_token"] = accessToken
	result["refresh_token"] = refeshToken
	return result
}

func (j *JWT) GenerateToken(userID string, userType string) string {
	var claims = jwt.MapClaims{}
	claims["user_id"] = userID
	claims["user_type"] = userType
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 168).Unix()

	var sign = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	validToken, err := sign.SignedString([]byte(j.signKey))

	if err != nil {
		return ""
	}

	return validToken
}

func (j JWT) RefereshJWT(refreshToken *jwt.Token) map[string]any {
	var result = map[string]any{}
	expTime, err := refreshToken.Claims.GetExpirationTime()
	if err != nil {
		logrus.Error("get token expiration error", err.Error())
		return nil
	}
	if refreshToken.Valid && expTime.Time.Compare(time.Now()) > 0 {
		var newAccessClaim = refreshToken.Claims.(jwt.MapClaims)
		newAccessClaim["iat"] = time.Now().Unix()
		newAccessClaim["exp"] = time.Now().Add(time.Hour * 168).Unix()

		var newAccessToken = jwt.NewWithClaims(refreshToken.Method, newAccessClaim)
		newSignedAccessToken, _ := newAccessToken.SignedString(refreshToken.Signature)

		var newRefreshClaim = refreshToken.Claims.(jwt.MapClaims)
		newRefreshClaim["exp"] = time.Now().Add(time.Hour * 24).Unix()

		var newRefreshToken = jwt.NewWithClaims(refreshToken.Method, newRefreshClaim)
		newSignedRefreshToken, _ := newRefreshToken.SignedString(refreshToken.Signature)

		result["access_token"] = newSignedAccessToken
		result["refresh_token"] = newSignedRefreshToken
		return result
	}

	return nil
}

func (j *JWT) generateRefreshToken(userID string, userType string) string {
	var claims = jwt.MapClaims{}
	claims["user_id"] = userID
	claims["user_type"] = userType
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	var sign = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err := sign.SignedString([]byte(j.refreshKey))

	if err != nil {
		return ""
	}
	return refreshToken
}

func (j JWT) ExtractToken(token *jwt.Token) any {
	if token.Valid {
		var claims = token.Claims
		expTime, _ := claims.GetExpirationTime()
		fmt.Println(expTime.Time.Compare(time.Now()))
		if expTime.Time.Compare(time.Now()) > 0 {
			var mapClaim = claims.(jwt.MapClaims)
			return mapClaim["id"]
		}

		logrus.Error("Token expired")
		return nil
	}
	return nil
}

func (j JWT) ValidateToken(token string, secret string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	return parsedToken, nil
}

func (j *JWT) GetID(c *gin.Context) (int, error) {
	authHeader := c.Request.Header.Get("Authorization")

	token, err := j.ValidateToken(authHeader, os.Getenv("REFSECRET"))
	if err != nil {
		logrus.Info(err)
		return 0, err
	}

	mapClaim := token.Claims.(jwt.MapClaims)
	idFloat, ok := mapClaim["id"].(float64)
	if !ok {
		return 0, fmt.Errorf("ID not found or not a valid number")
	}
	return int(idFloat), nil
}

func (j *JWT) CheckID(c *gin.Context) any {
	authHeader := c.Request.Header.Get("Authorization")

	token, err := j.ValidateToken(authHeader, os.Getenv("REFSECRET"))
	if err != nil {
		logrus.Info(err)
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Token is not valid"})
		return nil
	}

	mapClaim := token.Claims.(jwt.MapClaims)
	id := mapClaim["id"]

	return id
}
