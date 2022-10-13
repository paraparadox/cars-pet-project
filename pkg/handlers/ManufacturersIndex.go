package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) ManufacturersIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Manufacturers index",
	})
}
