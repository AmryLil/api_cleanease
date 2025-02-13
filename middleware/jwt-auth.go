package middlewares

import (
	"api_cleanease/helpers"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func AuthorizeJWT(jwtService helpers.JWTInterface, role int, secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			if role == -1 {
				c.Next()
				return
			}
			response := helpers.BuildErrorResponse("no token found")
			c.JSON(http.StatusBadRequest, response)
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwtService.ValidateToken(tokenString, secret)
		if err != nil {
			log.Println(err)
			response := helpers.BuildErrorResponse("token is not valid - " + err.Error())
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userIDStr, ok := claims["user_id"].(string)
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id is missing or not a string"})
				c.Abort()
				return
			}
			userID, err := uuid.Parse(userIDStr)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid UUID format"})
				c.Abort()
				return
			}
			personIDStr, ok := claims["person_id"].(string)
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "person_id is missing or not a string"})
				c.Abort()
				return
			}
			personID, err := uuid.Parse(personIDStr)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid UUID format"})
				c.Abort()
				return
			}
			userTypeStr, ok := claims["user_type"].(string)
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "user_type is missing or not a string"})
				c.Abort()
				return
			}
			userType, err := strconv.Atoi(userTypeStr)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "user_type is not a valid integer"})
				c.Abort()
				return
			}

			c.Set("user_id", userID)
			c.Set("person_id", personID)
			c.Set("user_type", userType)
			log.Println("User ID set in context:", userID)
			log.Println("user type set in context:", userType)

			if role == 0 || role == -1 || userType == 3 || (userType == 1 || userType == 2) {
				c.Next()
				return
			}

			if userType != role {
				response := helpers.BuildErrorResponse("this user cannot access this endpoint")
				c.JSON(http.StatusUnauthorized, response)
				c.Abort()
				return
			}

			c.Next()
			return
		}

		response := helpers.BuildErrorResponse("invalid token claims")
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
	}
}
