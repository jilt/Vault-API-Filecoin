package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jilt/Vault-API-Filecoin/internal/models"
)

type HealthHandler struct{}

func (h *HealthHandler) Handle(c *gin.Context) {
	healthOk := models.CheckHealthOK{
		Payload: "Ok",
	}

	c.JSON(http.StatusOK, healthOk)
}
