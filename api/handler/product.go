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

// CreateProduct godoc
// @Router       /products [post]
// @Summary      Create a new product
// @Description  Create a new product with the provided details
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product     body  models.CreateProduct  true  "data of the product"
// @Success      201  {string}  string
// @Failure      400  {object}  http.ErrorResp
// @Failure      404  {object}  http.ErrorResp
// @Failure      500  {object}  http.ErrorResp
func (h *Handler) CreateProduct(c *gin.Context) {
	var newProduct models.CreateProduct
	err := c.ShouldBindJSON(&newProduct)

	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	resp, err := h.strg.Product().Create(newProduct)
	if err != nil {
		h.log.Error("error product create:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	h.log.Info("product created")
	c.JSON(http.StatusCreated, resp)
}

// GetProduct godoc
// @Router       /products/{id} [get]
// @Summary      Get a product by ID
// @Description  Retrieve a product by its unique identifier
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path    string     true    "Product ID to retrieve"
// @Success      200  {object}  models.ProductResp
// @Failure      400  {object}  http.ErrorResp
// @Failure      404  {object}  http.ErrorResp
// @Failure      500  {object}  http.ErrorResp
func (h *Handler) GetProduct(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.strg.Product().Get(models.ProductIdReq{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		h.log.Error("error get product:", logger.Error(err))
		return
	}

	h.log.Warn("get response to product", logger.Any("id:", id))
	c.JSON(http.StatusOK, resp)
}

// GetListProduct godoc
// @Router       /products [get]
// @Summary      List products
// @Description  get products
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        limit    query     int  false  "limit for response"
// @Param		 page     query     int  false  "page for response"
// @Param        search     query     string false "search by name"
// @Param        barcode     query     string false "search by barcode"
// @Success      200  {array}   models.ProductResp
// @Failure      400  {object}  http.ErrorResp
// @Failure      404  {object}  http.ErrorResp
// @Failure      500  {object}  http.ErrorResp
func (h *Handler) GetListProduct(c *gin.Context) {
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

	resp, err := h.strg.Product().GetList(models.GetListProductReq{
		Page:    page,
		Limit:   limit,
		Name:    c.Query("search"),
		Barcode: c.Query("barcode"),
	})

	if err != nil {
		h.log.Error("error Product GetList:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	h.log.Warn("response to GetListProduct")
	c.JSON(http.StatusOK, resp)
}

// UpdateProduct godoc
// @Router       /products/{id} [put]
// @Summary      Update an existing product
// @Description  Update an existing product with the provided details
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id       path    string     true    "Product ID to update"
// @Param        product   body    models.UpdateProduct  true    "Update data for the product"
// @Success      200  {string}  string
// @Failure      400  {object}  http.ErrorResp
// @Failure      404  {object}  http.ErrorResp
// @Failure      500  {object}  http.ErrorResp
func (h *Handler) UpdateProduct(c *gin.Context) {
	var product models.Product
	product.Id = c.Param("id")

	err := c.ShouldBindJSON(&product)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	resp, err := h.strg.Product().Update(product)
	if err != nil {
		h.log.Error("error update product:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	h.log.Warn("response to UpdateProduct")
	c.JSON(http.StatusOK, resp)
}

// DeleteProduct godoc
// @Router       /products/{id} [delete]
// @Summary      Delete a product
// @Description  delete a product by its unique identifier
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path    string     true    "Product ID to retrieve"
// @Success      200  {object}  string
// @Failure      400  {object}  http.ErrorResp
// @Failure      404  {object}  http.ErrorResp
// @Failure      500  {object}  http.ErrorResp
func (h *Handler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if !helper.IsValidUUID(id) {
		h.log.Error("error Product id:", logger.Error(errors.New("invalid id")))
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Product().Delete(models.ProductIdReq{Id: id})
	if err != nil {
		h.log.Error("error Product delete:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	h.log.Warn("response to DeleteProduct")
	c.JSON(http.StatusOK, resp)
}
