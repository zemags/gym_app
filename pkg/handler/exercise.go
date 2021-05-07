package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createExerciseList(c *gin.Context) {
	id, _ := c.Get("userID")
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
