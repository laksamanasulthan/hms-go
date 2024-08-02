package injector

import (
	"hms-go/config"
	blog_injector "hms-go/domain/blog/injector"
	product_injector "hms-go/domain/product/injector"
)

type Injector struct {
	ProductInjector *product_injector.ProductInjector
	BlogInjector    *blog_injector.BlogInjector
	//Other domain Injector
}

func NewInjector() (*Injector, error) {

	db, err := config.DatabaseConfig()

	if err != nil {
		return nil, err
	}

	_, err = config.MigrationConfig()

	if err != nil {
		return nil, err
	}

	productInjector := product_injector.NewProductInjector(db)
	blogInjector := blog_injector.NewBlogInjector(db)
	//Other Domain Injector

	return &Injector{
		ProductInjector: productInjector,
		BlogInjector:    blogInjector,
		//Other Domain Injector
	}, nil
}
