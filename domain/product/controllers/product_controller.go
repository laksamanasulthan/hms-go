package controllers

import (
	"hms-go/domain/product/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service services.ProductService
}

func NewProductController(service services.ProductService) *ProductController {
	return &ProductController{service: service}
}

func (c *ProductController) Index(ctx *gin.Context) {

	products, err := c.service.GetAllProduct()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, products)

}

// func (c *ProductController) show(ctx *gin.Context) {

// }
// func (c *ProductController) store(ctx *gin.Context) {

// }
// func (c *ProductController) update(ctx *gin.Context) {

// }
// func (c *ProductController) delete(ctx *gin.Context) {

// }

func (c *ProductController) TestJson(ctx *gin.Context) {

	useService := c.service.TestService()
	resp := ", This is from Controller"
	ctx.JSON(200, gin.H{
		"message": useService + resp,
	})
}
