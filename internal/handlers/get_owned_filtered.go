package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MalukiMuthusi/mintbase/internal/adapters"
	"github.com/MalukiMuthusi/mintbase/internal/models"
	"github.com/gin-gonic/gin"
)

type OwnedFilteredHandler struct{}

func (h *OwnedFilteredHandler) Handle(c *gin.Context) {
	var ownedFilteredParamete *models.OwnedFilteredParameter

	if err := c.ShouldBindUri(ownedFilteredParamete); err != nil {
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

	// var res map[string]interface{}

	json.NewEncoder(c.Writer).Encode(b)
}
