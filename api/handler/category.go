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

// CreateCategory godoc
// @Router       /categories [post]
// @Summary      Create a new category
// @Description  Create a new category with the provided details
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        category     body  models.CreateCategory  true  "data of the category"
// @Success      201  {string}  string
// @Failure      400  {object}  http.ErrorResp
// @Failure      404  {object}  http.ErrorResp
// @Failure      500  {object}  http.ErrorResp
func (h *Handler) CreateCategory(c *gin.Context) {
	var newCategory models.CreateCategory
	err := c.ShouldBindJSON(&newCategory)

	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	resp, err := h.strg.Category().Create(newCategory)
	if err != nil {
		h.log.Error("error category create:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	h.log.Info("category created")
	c.JSON(http.StatusCreated, resp)
}

// GetCategory godoc
// @Router       /categories/{id} [get]
// @Summary      Get a category by ID
// @Description  Retrieve a category by its unique identifier
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        id   path    string     true    "Category ID to retrieve"
// @Success      200  {object}  models.CategoryResp
// @Failure      400  {object}  http.ErrorResp
// @Failure      404  {object}  http.ErrorResp
// @Failure      500  {object}  http.ErrorResp
func (h *Handler) GetCategory(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.strg.Category().Get(models.CategoryIdReq{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		h.log.Error("error get category:", logger.Error(err))
		return
	}

	h.log.Warn("get response to category", logger.Any("id:", id))
	c.JSON(http.StatusOK, resp)
}

// GetListCategory godoc
// @Router       /categories [get]
// @Summary      List categories
// @Description  get categories
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        limit    query     int  false  "limit for response"
// @Param		 page     query     int  false  "page for response"
// @Param        search     query     string false "search by name"
// @Success      200  {array}   models.CategoryResp
// @Failure      400  {object}  http.ErrorResp
// @Failure      404  {object}  http.ErrorResp
// @Failure      500  {object}  http.ErrorResp
func (h *Handler) GetListCategory(c *gin.Context) {
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

	resp, err := h.strg.Category().GetList(models.GetListCategoryReq{
		Page:  page,
		Limit: limit,
		Name:  c.Query("search"),
	})

	if err != nil {
		h.log.Error("error Category GetList:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	h.log.Warn("response to GetListCategory")
	c.JSON(http.StatusOK, resp)
}

// UpdateCategory godoc
// @Router       /categories/{id} [put]
// @Summary      Update an existing category
// @Description  Update an existing category with the provided details
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        id       path    string     true    "Category ID to update"
// @Param        category   body    models.UpdateCategory  true    "Update data for the category"
// @Success      200  {string}  string
// @Failure      400  {object}  http.ErrorResp
// @Failure      404  {object}  http.ErrorResp
// @Failure      500  {object}  http.ErrorResp
func (h *Handler) UpdateCategory(c *gin.Context) {
	var category models.Category
	category.Id = c.Param("id")

	err := c.ShouldBindJSON(&category)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	resp, err := h.strg.Category().Update(category)
	if err != nil {
		h.log.Error("error update category:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	h.log.Warn("response to UpdateCategory")
	c.JSON(http.StatusOK, resp)
}

// DeleteCategory godoc
// @Router       /categories/{id} [delete]
// @Summary      Delete a category
// @Description  delete a category by its unique identifier
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        id   path    string     true    "Category ID to retrieve"
// @Success      200  {object}  string
// @Failure      400  {object}  http.ErrorResp
// @Failure      404  {object}  http.ErrorResp
// @Failure      500  {object}  http.ErrorResp
func (h *Handler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	if !helper.IsValidUUID(id) {
		h.log.Error("error Category id:", logger.Error(errors.New("invalid id")))
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Category().Delete(models.CategoryIdReq{Id: id})
	if err != nil {
		h.log.Error("error Category delete:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	h.log.Warn("response to DeleteCategory")
	c.JSON(http.StatusOK, resp)
}
