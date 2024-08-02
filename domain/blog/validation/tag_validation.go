package validation

import "github.com/go-playground/validator/v10"

type StoreTagValidation struct {
	Tag string `json:"tag" validate:"required"`
}

type UpdateTagValidation struct {
	Name  string `json:"name" validate:"required"`
	Price string `json:"price" validate:"required"`
	Qty   uint   `json:"qty" validate:"required,min=1"`
}

func ValidateStoreProduct(v *StoreTagValidation) error {
	return validator.New().Struct(v)
}

func ValidateUpdateProduct(v *UpdateTagValidation) error {
	return validator.New().Struct(v)
}
