package config

import (
	blog_model "hms-go/domain/blog/models"
	product_models "hms-go/domain/product/models"
	user_models "hms-go/domain/user/models"
)

func MigrationConfig() (bool, error) {
	db, err := DatabaseConfig()

	if err != nil {
		return false, err
	}

	//Auto-Migration
	err = db.AutoMigrate(
		&product_models.Product{},
		&user_models.User{},
		&blog_model.Tag{},
		&blog_model.Post{},
		&blog_model.PostTag{},
	)

	if err != nil {
		return false, err
	}

	return true, nil

}
