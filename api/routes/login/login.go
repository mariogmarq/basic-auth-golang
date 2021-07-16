package login

import (
	"GoLandPruebas/internal/services"
	"GoLandPruebas/pkg/encodeUtils"
	"GoLandPruebas/pkg/jwtUtils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	name := c.PostForm("user")
	password := string(encodeUtils.Encode([]byte(c.PostForm("password"))))

	user, err := services.UserService.GetUserByName(name, password)
	if err != nil {
		c.Status(400)
		return
	}

	signedString, err := jwtUtils.CreateToken(user.Model.ID)
	if err != nil {
		c.Status(500)
		return
	}

	c.SetCookie("token", signedString, 60*10, "/", "localhost", false, true)
	c.Status(http.StatusNoContent)
}

func ValidateLoginPost(c *gin.Context) {
	if c.PostForm("user") == "" || c.PostForm("password") == "" {
		c.Status(http.StatusBadRequest)
		c.Abort()
		return
	}

	c.Next()
}