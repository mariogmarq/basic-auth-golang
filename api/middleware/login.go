package middleware

import (
	"GoLandPruebas/pkg/jwtUtils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func LoginMiddleware(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusBadRequest, "no logged")
		c.Abort()
		return
	}

	claims := jwtUtils.CustomClaim{}

	parsed, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return jwtUtils.MyKey, nil
	})

	log.Println(claims)
	if err != nil || !parsed.Valid {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "invalid token")
		c.Abort()
		return
	}

	c.Next()
}
