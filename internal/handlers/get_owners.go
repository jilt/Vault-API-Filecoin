package handlers

import (
	"net/http"

	"github.com/jilt/Vault-API-Filecoin/tree/main/internal/adapters"
	"github.com/jilt/Vault-API-Filecoin/tree/main/internal/models"
	"github.com/gin-gonic/gin"
)

type OwnersHandler struct{}

func (h *OwnersHandler) Handle(c *gin.Context) {

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
