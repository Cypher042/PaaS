package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	c.JSON(http.StatusOK, "hi")
}

func Register(c *gin.Context) {

	var regRequest RegisterRequest

	if err := c.ShouldBind(&regRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}
