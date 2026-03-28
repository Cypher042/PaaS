package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}


func (h *Handler) Login(c *gin.Context) {

	var loginReq LoginRequest

	if err := c.ShouldBind(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user, err := h.service.Login(loginReq)

	if err != nil{

	}

	c.JSON(http.StatusOK, user)



}

func (h *Handler) Register(c *gin.Context) {

	var regRequest RegisterRequest

	if err := c.ShouldBind(&regRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user, err := h.service.Register(regRequest)

	if err != nil{

	}

	c.JSON(http.StatusOK, user)
}
