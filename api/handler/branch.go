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

// CreateBranch godoc
// @Router       /branches [post]
// @Summary      Create a new branch
// @Description  Create a new branch with the provided details
// @Tags         branches
// @Accept       json
// @Produce      json
// @Param        branch     body  models.CreateBranch  true  "data of the branch"
// @Success      201  {string}  string
// @Failure      400  {object}  http.ErrorResp
// @Failure      404  {object}  http.ErrorResp
// @Failure      500  {object}  http.ErrorResp
func (h *Handler) CreateBranch(c *gin.Context) {
	var newBranch models.CreateBranch
	err := c.ShouldBindJSON(&newBranch)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	resp, err := h.strg.Branch().Create(newBranch)
	if err != nil {
		h.log.Error("error branch create:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	h.log.Info("branch created")
	c.JSON(http.StatusCreated, resp)
}

// GetBranch godoc
// @Router       /branches/{id} [get]
// @Summary      Get a branch by ID
// @Description  Retrieve a branch by its unique identifier
// @Tags         branches
// @Accept       json
// @Produce      json
// @Param        id   path    string     true    "Branch ID to retrieve"
// @Success      200  {object}  models.Branch
// @Failure      400  {object}  http.ErrorResp
// @Failure      404  {object}  http.ErrorResp
// @Failure      500  {object}  http.ErrorResp
func (h *Handler) GetBranch(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.strg.Branch().Get(models.BranchIdReq{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		h.log.Error("error get branch:", logger.Error(err))
		return
	}

	h.log.Warn("Response to GetBranch", logger.Any("id:", id))
	c.JSON(http.StatusOK, resp)
}

// GetListBranch godoc
// @Router       /branches [get]
// @Summary      List branches
// @Description  get branches
// @Tags         branches
// @Accept       json
// @Produce      json
// @Param        limit    query     int  false  "limit for response"
// @Param		 page     query     int  false  "page for response"
// @Param        search     query     string false "search by name"
// @Success      200  {array}   models.Branch
// @Failure      400  {object}  http.ErrorResp
// @Failure      404  {object}  http.ErrorResp
// @Failure      500  {object}  http.ErrorResp
func (h *Handler) GetListBranch(c *gin.Context) {
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

	resp, err := h.strg.Branch().GetList(models.GetListBranchReq{
		Page:  page,
		Limit: limit,
		Name:  c.Query("search"),
	})

	if err != nil {
		h.log.Error("error Branch GetList:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	h.log.Warn("response to GetListBranch")
	c.JSON(http.StatusOK, resp)
}

// UpdateBranch godoc
// @Router       /branches/{id} [put]
// @Summary      Update an existing branch
// @Description  Update an existing branch with the provided details
// @Tags         branches
// @Accept       json
// @Produce      json
// @Param        id       path    string     true    "Branch ID to update"
// @Param        branch   body    models.UpdateBranch  true    "Updated data for the branch"
// @Success      200  {string}  string
// @Failure      400  {object}  http.ErrorResp
// @Failure      404  {object}  http.ErrorResp
// @Failure      500  {object}  http.ErrorResp
func (h *Handler) UpdateBranch(c *gin.Context) {
	var branch models.Branch
	branch.Id = c.Param("id")
	err := c.ShouldBindJSON(&branch)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	resp, err := h.strg.Branch().Update(branch)
	if err != nil {
		h.log.Error("error update branch:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	h.log.Warn("response to UpdateBranch")
	c.JSON(http.StatusOK, resp)
}

// DeleteBranch godoc
// @Router       /branches/{id} [delete]
// @Summary      Delete a branch
// @Description  delete a branch by its unique identifier
// @Tags         branches
// @Accept       json
// @Produce      json
// @Param        id   path    string     true    "Branch ID to retrieve"
// @Success      200  {object}  string
// @Failure      400  {object}  http.ErrorResp
// @Failure      404  {object}  http.ErrorResp
// @Failure      500  {object}  http.ErrorResp
func (h *Handler) DeleteBranch(c *gin.Context) {
	id := c.Param("id")
	if !helper.IsValidUUID(id) {
		h.log.Error("error Branch id:", logger.Error(errors.New("invalid id")))
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Branch().Delete(models.BranchIdReq{Id: id})
	if err != nil {
		h.log.Error("error Branch delete:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	h.log.Warn("response to DeleteBranch")
	c.JSON(http.StatusOK, resp)
}
