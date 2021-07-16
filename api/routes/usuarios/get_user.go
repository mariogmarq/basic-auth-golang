package usuarios

import (
	"GoLandPruebas/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUsuarios(c *gin.Context) {
	id := c.Param("id")
	i, _ := strconv.Atoi(id)
	user, err := services.UserService.GetUserById(uint(i))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error happened")
		return
	}
	c.JSON(http.StatusOK, user)
}

func GetAllUsuarios(c *gin.Context) {
	users, _ := services.UserService.GetAllUsers()
	c.JSON(http.StatusOK, users)
}