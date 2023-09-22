package api

import (
	"market/api/handler"

	_ "market/api/docs"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func NewServer(h *handler.Handler) *gin.Engine {
	r := gin.New()

	branch := r.Group("/branches")
	branch.POST("/", h.CreateBranch)
	branch.GET("/:id", h.GetBranch)
	branch.GET("/", h.GetListBranch)
	branch.PUT("/:id", h.UpdateBranch)
	branch.DELETE("/:id", h.DeleteBranch)

	category := r.Group("/categories")
	category.POST("/", h.CreateCategory)
	category.GET("/:id", h.GetCategory)
	category.GET("/", h.GetListCategory)
	category.PUT("/:id", h.UpdateCategory)
	category.DELETE("/:id", h.DeleteCategory)

	product := r.Group("/products")
	product.POST("/", h.CreateProduct)
	product.GET("/:id", h.GetProduct)
	product.GET("/", h.GetListProduct)
	product.PUT("/:id", h.UpdateProduct)
	product.DELETE("/:id", h.DeleteProduct)

	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
