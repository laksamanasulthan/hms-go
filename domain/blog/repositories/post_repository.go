package repository

import (
	blog_model "hms-go/domain/blog/models"

	"gorm.io/gorm"
)

type PostRepository interface {
	GetAll() ([]blog_model.Post, error)
	GetOne(id uint) (*blog_model.Post, error)
	Create(post *blog_model.Post) (*blog_model.Post, error)
	Update(id uint, post *blog_model.Post) (*blog_model.Post, error)
	Remove(id uint) (*blog_model.Post, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) GetAll() ([]blog_model.Post, error) {
	var posts []blog_model.Post

	if err := r.db.Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *postRepository) GetOne(id uint) (*blog_model.Post, error) {

	var post blog_model.Post

	if err := r.db.First(&post, id).Error; err != nil {
		return nil, err
	}

	return &post, nil

}

func (r *postRepository) Create(post *blog_model.Post) (*blog_model.Post, error) {

	if err := r.db.Create(post).Error; err != nil {
		return nil, err
	}

	return post, nil
}

func (r *postRepository) Update(id uint, post *blog_model.Post) (*blog_model.Post, error) {

	post.ID = id

	if err := r.db.Save(post).Error; err != nil {
		return nil, err
	}

	return post, nil
}

func (r *postRepository) Remove(id uint) (*blog_model.Post, error) {

	var post blog_model.Post

	if err := r.db.First(&post, id).Error; err != nil {
		return nil, err
	}

	if err := r.db.Delete(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}
