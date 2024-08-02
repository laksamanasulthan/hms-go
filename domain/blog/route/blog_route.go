package route

import (
	blog_injector "hms-go/domain/blog/injector"

	"github.com/gin-gonic/gin"
)

func RegisterBlogRoutes(router *gin.Engine, ctrl blog_injector.BlogControllers) {

	blogRoutes := router.Group("/blog")
	{
		blogRoutes.GET("", ctrl.TagController.GetAll)
		blogRoutes.GET(":id", ctrl.TagController.GetOne)
		blogRoutes.POST("", ctrl.TagController.Create)
	}

	// ptherDomainRoutes := router.Group("/other-domain-routes")
	// {
	// 	otherDomainRoutes.GET("", controllers.OtherController.Index)
	// }

}
