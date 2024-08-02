package product_route

import (
	product_injector "hms-go/domain/product/injector"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(router *gin.Engine, controllers product_injector.ProductControllers) {

	productRoutes := router.Group("/product")
	{
		productRoutes.GET("", controllers.ProductController.Index)
		productRoutes.GET("test-json", controllers.ProductController.TestJson)
	}

	// ptherDomainRoutes := router.Group("/other-domain-routes")
	// {
	// 	otherDomainRoutes.GET("", controllers.OtherController.Index)
	// }

}
