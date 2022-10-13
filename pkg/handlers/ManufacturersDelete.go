package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) ManufacturersDelete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Manufacturers delete",
	})
}
