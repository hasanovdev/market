package handler

import (
	"market/models"
	"market/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateRemaining godoc
// @Router       /remainings [post]
// @Summary      Create a new remaining
// @Description  Create a new remaining with the provided details
// @Tags         remainings
// @Accept       json
// @Produce      json
// @Param        remaining     body  models.CreateRemaining  true  "data of the remaining product"
// @Success      201  {string}  string
// @Failure      400  {object}  http.ErrorResp
// @Failure      404  {object}  http.ErrorResp
// @Failure      500  {object}  http.ErrorResp
func (h *Handler) CreateRemaining(c *gin.Context) {
	var newRemaining models.CreateRemaining
	err := c.ShouldBindJSON(&newRemaining)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	resp, err := h.strg.Remaining().Create(newRemaining)
	if err != nil {
		h.log.Error("error remaining create:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	h.log.Info("remaining product created")
	c.JSON(http.StatusCreated, resp)
}
