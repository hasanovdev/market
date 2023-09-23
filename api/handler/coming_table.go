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

// CreateComingTable godoc
// @Router       /coming_tables [post]
// @Summary      Create a new coming_table
// @Description  Create a new coming_table with the provided details
// @Tags         coming_tables
// @Accept       json
// @Produce      json
// @Param        coming_table     body  models.CreateComingTable  true  "data of the coming table"
// @Success      201  {string}  string
// @Failure      400  {object}  http.ErrorResp
// @Failure      404  {object}  http.ErrorResp
// @Failure      500  {object}  http.ErrorResp
func (h *Handler) CreateComingTable(c *gin.Context) {
	var newComingTable models.CreateComingTable
	err := c.ShouldBindJSON(&newComingTable)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	resp, err := h.strg.ComingTable().Create(newComingTable)
	if err != nil {
		h.log.Error("error coming table create:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	h.log.Info("coming table created")
	c.JSON(http.StatusCreated, resp)
}

// GetComingTable godoc
// @Router       /coming_tables/{id} [get]
// @Summary      Get a coming_table by ID
// @Description  Retrieve a coming_table by its unique identifier
// @Tags         coming_tables
// @Accept       json
// @Produce      json
// @Param        id   path    string     true    "Coming table ID to retrieve"
// @Success      200  {object}  models.ComingTableResp
// @Failure      400  {object}  http.ErrorResp
// @Failure      404  {object}  http.ErrorResp
// @Failure      500  {object}  http.ErrorResp
func (h *Handler) GetComingTable(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.strg.ComingTable().Get(models.ComingTableIdReq{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		h.log.Error("error get coming table:", logger.Error(err))
		return
	}

	h.log.Warn("Response to GetComingTable", logger.Any("id:", id))
	c.JSON(http.StatusOK, resp)
}

// GetListComingTable godoc
// @Router       /coming_tables [get]
// @Summary      List coming_tables
// @Description  get coming_tables
// @Tags         coming_tables
// @Accept       json
// @Produce      json
// @Param        limit    query     int  false  "limit for response"
// @Param		 page     query     int  false  "page for response"
// @Param		 branch_id     query     string  false  "branch_id for response"
// @Param		 coming_id     query     string  false  "coming_id for response"
// @Success      200  {array}   models.ComingTableResp
// @Failure      400  {object}  http.ErrorResp
// @Failure      404  {object}  http.ErrorResp
// @Failure      500  {object}  http.ErrorResp
func (h *Handler) GetListComingTable(c *gin.Context) {
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

	resp, err := h.strg.ComingTable().GetList(models.GetListComingTableReq{
		Page:     page,
		Limit:    limit,
		BranchId: c.Query("branch_id"),
		ComingId: c.Query("coming_id"),
	})

	if err != nil {
		h.log.Error("error ComingTable GetList:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	h.log.Warn("response to GetListComingTable")
	c.JSON(http.StatusOK, resp)
}

// UpdateComingTable godoc
// @Router       /coming_tables/{id} [put]
// @Summary      Update an existing coming_table
// @Description  Update an existing coming_table with the provided details
// @Tags         coming_tables
// @Accept       json
// @Produce      json
// @Param        id       path    string     true    "Coming table ID to update"
// @Param        coming_table   body    models.UpdateComingTable  true    "Update data for the coming table"
// @Success      200  {string}  string
// @Failure      400  {object}  http.ErrorResp
// @Failure      404  {object}  http.ErrorResp
// @Failure      500  {object}  http.ErrorResp
func (h *Handler) UpdateComingTable(c *gin.Context) {
	var comingTable models.ComingTable
	comingTable.Id = c.Param("id")
	err := c.ShouldBindJSON(&comingTable)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	resp, err := h.strg.ComingTable().Update(comingTable)
	if err != nil {
		h.log.Error("error update coming table:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	h.log.Warn("response to UpdateComingTable")
	c.JSON(http.StatusOK, resp)
}

// DeleteComingTable godoc
// @Router       /coming_tables/{id} [delete]
// @Summary      Delete a coming_table
// @Description  delete a coming_table by its unique identifier
// @Tags         coming_tables
// @Accept       json
// @Produce      json
// @Param        id   path    string     true    "Coming table ID to retrieve"
// @Success      200  {object}  string
// @Failure      400  {object}  http.ErrorResp
// @Failure      404  {object}  http.ErrorResp
// @Failure      500  {object}  http.ErrorResp
func (h *Handler) DeleteComingTable(c *gin.Context) {
	id := c.Param("id")
	if !helper.IsValidUUID(id) {
		h.log.Error("error ComingTable id:", logger.Error(errors.New("invalid id")))
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.ComingTable().Delete(models.ComingTableIdReq{Id: id})
	if err != nil {
		h.log.Error("error ComingTable delete:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	h.log.Warn("response to DeleteComingTable")
	c.JSON(http.StatusOK, resp)
}
