package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) ManufacturersStore(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Manufacturers store",
	})
}
