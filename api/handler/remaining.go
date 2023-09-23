package handler

import (
	"errors"
	"market/models"
	"market/pkg/helper"
	"market/pkg/logger"
	"net/http"
	"strconv"

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

// GetRemaining godoc
// @Router       /remainings/{id} [get]
// @Summary      Get a reamining by ID
// @Description  Retrieve a reamining by its unique identifier
// @Tags         remainings
// @Accept       json
// @Produce      json
// @Param        id   path    string     true    "Remaining ID to retrieve"
// @Success      200  {object}  models.RemainingResp
// @Failure      400  {object}  http.ErrorResp
// @Failure      404  {object}  http.ErrorResp
// @Failure      500  {object}  http.ErrorResp
func (h *Handler) GetRemaining(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.strg.Remaining().Get(models.RemainingIdReq{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		h.log.Error("error get remaining product:", logger.Error(err))
		return
	}

	h.log.Warn("get response to remaining product", logger.Any("id:", id))
	c.JSON(http.StatusOK, resp)
}

// GetListRemaining godoc
// @Router       /remainings [get]
// @Summary      List remainings
// @Description  get remainings
// @Tags         remainings
// @Accept       json
// @Produce      json
// @Param        limit    query     int  false  "limit for response"
// @Param		 page     query     int  false  "page for response"
// @Param        branch_id     query     string false "search by branch_id"
// @Param        category_id     query     string false "search by category_id"
// @Param        barcode     query     string false "search by barcode"
// @Success      200  {array}   models.RemainingResp
// @Failure      400  {object}  http.ErrorResp
// @Failure      404  {object}  http.ErrorResp
// @Failure      500  {object}  http.ErrorResp
func (h *Handler) GetListRemaining(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		h.log.Error("error get page:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid page param")
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		h.log.Error("error get limit:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid page param")
		return
	}

	resp, err := h.strg.Remaining().GetList(models.GetListRemainingReq{
		Page:       page,
		Limit:      limit,
		BranchId:   c.Query("branch_id"),
		CategoryId: c.Query("category_id"),
		Barcode:    c.Query("barcode"),
	})

	if err != nil {
		h.log.Error("error Remaining product GetList:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	h.log.Warn("response to GetListRemaining")
	c.JSON(http.StatusOK, resp)
}

// UpdateRemaining godoc
// @Router       /remainings/{id} [put]
// @Summary      Update an existing remaining
// @Description  Update an existing remaining with the provided details
// @Tags         remainings
// @Accept       json
// @Produce      json
// @Param        id       path    string     true    "Remaining ID to update"
// @Param        remaining   body    models.UpdateRemaining  true    "Update data for the remaining product"
// @Success      200  {string}  string
// @Failure      400  {object}  http.ErrorResp
// @Failure      404  {object}  http.ErrorResp
// @Failure      500  {object}  http.ErrorResp
func (h *Handler) UpdateRemaining(c *gin.Context) {
	var remaining models.Remaining
	remaining.Id = c.Param("id")

	err := c.ShouldBindJSON(&remaining)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	resp, err := h.strg.Remaining().Update(remaining)
	if err != nil {
		h.log.Error("error update remaining product:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	h.log.Warn("response to UpdateRemaining")
	c.JSON(http.StatusOK, resp)
}

// DeleteRemaining godoc
// @Router       /remainings/{id} [delete]
// @Summary      Delete a remaining
// @Description  delete a remaining by its unique identifier
// @Tags         remainings
// @Accept       json
// @Produce      json
// @Param        id   path    string     true    "Remaining ID to retrieve"
// @Success      200  {object}  string
// @Failure      400  {object}  http.ErrorResp
// @Failure      404  {object}  http.ErrorResp
// @Failure      500  {object}  http.ErrorResp
func (h *Handler) DeleteRemaining(c *gin.Context) {
	id := c.Param("id")
	if !helper.IsValidUUID(id) {
		h.log.Error("error Remaining id:", logger.Error(errors.New("invalid id")))
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Remaining().Delete(models.RemainingIdReq{Id: id})
	if err != nil {
		h.log.Error("error Remaining delete:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	h.log.Warn("response to DeleteRemaining")
	c.JSON(http.StatusOK, resp)
}
