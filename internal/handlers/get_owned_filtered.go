package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jilt/Vault-API-Filecoin/internal/adapters"
	"github.com/jilt/Vault-API-Filecoin/internal/models"
)

type OwnedFilteredHandler struct{}

func (h *OwnedFilteredHandler) Handle(c *gin.Context) {
	var ownedFilteredParamete models.OwnedFilteredParameter

	if err := c.ShouldBindUri(&ownedFilteredParamete); err != nil {
		e := models.BasicError{
			Code:    models.InvalidUserIdParam.String(),
			Message: "provide a valid user and store parameter",
		}

		c.JSON(http.StatusUnprocessableEntity, e)
		return
	}

	b, err := adapters.GetOwnedFiltered(ownedFilteredParamete)
	if err != nil {
		e := models.BasicError{
			Code:    models.InvalidUserIdParam.String(),
			Message: "failed to fetch data",
		}

		c.JSON(http.StatusInternalServerError, e)
		return
	}

	c.JSON(http.StatusOK, b)
}
