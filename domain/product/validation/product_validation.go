package product_validation

import "github.com/go-playground/validator/v10"

type StoreProductValidation struct {
	Name  string `json:"name" validate:"required"`
	Price string `json:"price" validate:"required"`
	Qty   uint   `json:"qty" validate:"required,min=1"`
}

type UpdateProductValidation struct {
	Name  string `json:"name" validate:"required"`
	Price string `json:"price" validate:"required"`
	Qty   uint   `json:"qty" validate:"required,min=1"`
}

func ValidateStoreProduct(product *StoreProductValidation) error {
	return validator.New().Struct(product)
}

func ValidateUpdateProduct(product *UpdateProductValidation) error {
	return validator.New().Struct(product)
}
