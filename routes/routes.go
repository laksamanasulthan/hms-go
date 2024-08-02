package routes

import (
	blog_route "hms-go/domain/blog/route"
	product_route "hms-go/domain/product/route"
	"hms-go/injector"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, injector *injector.Injector) {
	//``
	product_route.RegisterProductRoutes(router, injector.ProductInjector.Controllers)
	blog_route.RegisterBlogRoutes(router, injector.BlogInjector.Controllers)

}
