package injector

import (
	blog_controller "hms-go/domain/blog/controllers"
	blog_repository "hms-go/domain/blog/repositories"
	blog_service "hms-go/domain/blog/services"

	"gorm.io/gorm"
)

type BlogControllers struct {
	TagController *blog_controller.TagController
	//Other Controlller
}
type BlogInjector struct {
	Controllers BlogControllers
}

func NewBlogInjector(db *gorm.DB) *BlogInjector {

	tagRepo := blog_repository.NewTagRepository(db)
	tagService := blog_service.NewTagService(tagRepo)
	tagController := blog_controller.NewTagController(tagService)

	//Other DI Repo
	//Other DI Service
	//Other DI Controller

	controllers := BlogControllers{
		TagController: tagController,
		//Other Controller
	}

	return &BlogInjector{
		Controllers: controllers,
	}
}
