package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) ManufacturersShow(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Manufacturers show",
	})
}
