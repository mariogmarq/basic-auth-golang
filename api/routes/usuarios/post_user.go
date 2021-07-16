package usuarios

import (
	"GoLandPruebas/internal/services"
	"GoLandPruebas/pkg/encodeUtils"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	name := c.PostForm("user")
	password := string(encodeUtils.Encode([]byte(c.PostForm("password"))))

	services.UserService.CreateUser(name, password)
}
