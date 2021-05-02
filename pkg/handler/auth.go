package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gym_app "github.com/zemags/gym_app/store"
)

func (h *Handler) signUp(c *gin.Context) {
	var input gym_app.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {

}
