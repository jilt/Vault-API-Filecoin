package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jilt/Vault-API-Filecoin/internal/adapters"
	"github.com/jilt/Vault-API-Filecoin/internal/models"
)

type OwnedHandler struct{}

func (h *OwnedHandler) Handle(c *gin.Context) {
	var userID models.UserIDParameter

	if err := c.ShouldBindUri(&userID); err != nil {
		e := models.BasicError{
			Code:    models.InvalidUserIdParam.String(),
			Message: "provide a valid user parameter",
		}

		c.JSON(http.StatusUnprocessableEntity, e)
		return
	}

	b, err := adapters.GetOwnedByUser(userID)
	if err != nil {
		e := models.BasicError{
			Code:    models.InvalidUserIdParam.String(),
			Message: "failed to fetch data",
		}

		c.JSON(http.StatusInternalServerError, e)
		return
	}

	// var res map[string]interface{}

	c.JSON(http.StatusOK, b)
}
