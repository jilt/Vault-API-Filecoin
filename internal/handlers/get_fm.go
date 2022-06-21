package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jilt/Vault-API-Filecoin/internal/adapters"
	"github.com/jilt/Vault-API-Filecoin/internal/models"
)

type FmHandler struct{}

func (h *FmHandler) Handle(c *gin.Context) {

	var fmParameter models.FmParameter

	if err := c.ShouldBindUri(&fmParameter); err != nil {
		e := models.BasicError{
			Code:    models.InvalidTokenID.String(),
			Message: "provide a valid token ID",
		}

		c.JSON(http.StatusUnprocessableEntity, e)
		return
	}

	b, err := adapters.GetFm(fmParameter)
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
