package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jilt/Vault-API-Filecoin/internal/adapters"
	"github.com/jilt/Vault-API-Filecoin/internal/models"
)

type OwnersParasHandler struct{}

func (h *OwnersParasHandler) Handle(c *gin.Context) {

	var ownerParameter models.OwnerParameter

	if err := c.ShouldBindUri(&ownerParameter); err != nil {
		e := models.BasicError{
			Code:    models.InvalidTokenID.String(),
			Message: "provide a valid token ID",
		}

		c.JSON(http.StatusUnprocessableEntity, e)
		return
	}

	b, err := adapters.GetOwners(ownerParameter)
	if err != nil {

		var e models.BasicError
		var status int

		switch err {

		case models.ErrFailedFetchData:

			e = models.BasicError{
				Code:    models.InvalidTokenID.String(),
				Message: "failed to fetch data",
			}

			status = http.StatusInternalServerError

		default:

			e = models.BasicError{
				Code:    models.FailedToProcessRequest.String(),
				Message: "failed to fetch data",
			}

			status = http.StatusInternalServerError

		}

		c.JSON(status, e)
		return
	}

	c.JSON(http.StatusOK, b)

}
