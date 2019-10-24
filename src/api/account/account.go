package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"WebProject/src/common"
)

//User user
type User struct {
	UserName string `gorm:"column:user_name"`
}

//Login login
func Login(c *gin.Context) {
	c.JSON(http.StatusOK, common.ResponseOK(gin.H{"name": "hehe"}))
}
