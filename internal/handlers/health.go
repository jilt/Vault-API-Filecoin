package handlers

import (
	"net/http"

	"github.com/MalukiMuthusi/mintbase/internal/models"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func (h *HealthHandler) Handle(c *gin.Context) {
	healthOk := models.CheckHealthOK{
		Payload: "Ok",
	}

	c.JSON(http.StatusOK, healthOk)
}
