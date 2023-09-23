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

// CreateComingTableProduct godoc
// @Router       /coming_table_products [post]
// @Summary      Create a new coming_table_product
// @Description  Create a new coming_table_product with the provided details
// @Tags         coming_table_products
// @Accept       json
// @Produce      json
// @Param        coming_table_product     body  models.CreateComingTableProduct  true  "data of the coming table product"
// @Success      201  {string}  string
// @Failure      400  {object}  http.ErrorResp
// @Failure      404  {object}  http.ErrorResp
// @Failure      500  {object}  http.ErrorResp
func (h *Handler) CreateComingTableProduct(c *gin.Context) {
	var newComingTableProduct models.CreateComingTableProduct
	err := c.ShouldBindJSON(&newComingTableProduct)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	resp, err := h.strg.ComingTableProduct().Create(newComingTableProduct)
	if err != nil {
		h.log.Error("error coming table product create:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	h.log.Info("coming table product created")
	c.JSON(http.StatusCreated, resp)
}

// GetComingTableProduct godoc
// @Router       /coming_table_products/{id} [get]
// @Summary      Get a oming_table_product by ID
// @Description  Retrieve a coming_table_product by its unique identifier
// @Tags         coming_table_products
// @Accept       json
// @Produce      json
// @Param        id   path    string     true    "Coming table product ID to retrieve"
// @Success      200  {object}  models.ComingTableProductResp
// @Failure      400  {object}  http.ErrorResp
// @Failure      404  {object}  http.ErrorResp
// @Failure      500  {object}  http.ErrorResp
func (h *Handler) GetComingTableProduct(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.strg.ComingTableProduct().Get(models.ComingTableProductIdReq{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		h.log.Error("error get coming table product:", logger.Error(err))
		return
	}

	h.log.Warn("Response to GetComingTableProduct", logger.Any("id:", id))
	c.JSON(http.StatusOK, resp)
}

// GetListComingTableProduct godoc
// @Router       /coming_table_products [get]
// @Summary      List coming_table_products
// @Description  get coming_table_products
// @Tags         coming_table_products
// @Accept       json
// @Produce      json
// @Param        limit    query     int  false  "limit for response"
// @Param		 page     query     int  false  "page for response"
// @Param		 category_id     query     string  false  "category_id for response"
// @Param		 barcode     query     string  false  "barcode for response"
// @Success      200  {array}   models.ComingTableProductResp
// @Failure      400  {object}  http.ErrorResp
// @Failure      404  {object}  http.ErrorResp
// @Failure      500  {object}  http.ErrorResp
func (h *Handler) GetListComingTableProduct(c *gin.Context) {
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

	resp, err := h.strg.ComingTableProduct().GetList(models.GetListComingTableProductReq{
		Page:       page,
		Limit:      limit,
		CategoryId: c.Query("category_id"),
		Barcode:    c.Query("barcode"),
	})

	if err != nil {
		h.log.Error("error ComingTableProduct GetList:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	h.log.Warn("response to GetListComingTableProduct")
	c.JSON(http.StatusOK, resp)
}

// UpdateComingTableProduct godoc
// @Router       /coming_table_products/{id} [put]
// @Summary      Update an existing coming_table_product
// @Description  Update an existing coming_table_product with the provided details
// @Tags         coming_table_products
// @Accept       json
// @Produce      json
// @Param        id       path    string     true    "Coming table product ID to update"
// @Param        coming_table   body    models.UpdateComingTableProduct  true    "Update data for the coming table product"
// @Success      200  {string}  string
// @Failure      400  {object}  http.ErrorResp
// @Failure      404  {object}  http.ErrorResp
// @Failure      500  {object}  http.ErrorResp
func (h *Handler) UpdateComingTableProduct(c *gin.Context) {
	var comingTableProduct models.ComingTableProduct
	comingTableProduct.Id = c.Param("id")
	err := c.ShouldBindJSON(&comingTableProduct)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	resp, err := h.strg.ComingTableProduct().Update(comingTableProduct)
	if err != nil {
		h.log.Error("error update coming table product:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	h.log.Warn("response to UpdateComingTableProduct")
	c.JSON(http.StatusOK, resp)
}

// DeleteComingTableProduct godoc
// @Router       /coming_table_products/{id} [delete]
// @Summary      Delete a coming_table_product
// @Description  delete a coming_table_product by its unique identifier
// @Tags         coming_table_products
// @Accept       json
// @Produce      json
// @Param        id   path    string     true    "Coming table product ID to retrieve"
// @Success      200  {object}  string
// @Failure      400  {object}  http.ErrorResp
// @Failure      404  {object}  http.ErrorResp
// @Failure      500  {object}  http.ErrorResp
func (h *Handler) DeleteComingTableProduct(c *gin.Context) {
	id := c.Param("id")
	if !helper.IsValidUUID(id) {
		h.log.Error("error ComingTableProduct id:", logger.Error(errors.New("invalid id")))
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.ComingTableProduct().Delete(models.ComingTableProductIdReq{Id: id})
	if err != nil {
		h.log.Error("error ComingTableProduct delete:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	h.log.Warn("response to DeleteComingTableProduct")
	c.JSON(http.StatusOK, resp)
}
