package repository

import (
	blog_model "hms-go/domain/blog/models"

	"gorm.io/gorm"
)

type TagRepository interface {
	GetAll() ([]blog_model.Tag, error)
	GetOne(id uint) (*blog_model.Tag, error)
	Create(tag *blog_model.Tag) (*blog_model.Tag, error)
	Update(id uint, tag *blog_model.Tag) (*blog_model.Tag, error)
	Remove(id uint) (*blog_model.Tag, error)
}

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return &tagRepository{db: db}
}

func (r *tagRepository) GetAll() ([]blog_model.Tag, error) {
	var tags []blog_model.Tag

	if err := r.db.Find(tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func (r *tagRepository) GetOne(id uint) (*blog_model.Tag, error) {
	var tag blog_model.Tag

	if err := r.db.First(&tag, id).Error; err != nil {
		return nil, err
	}

	return &tag, nil

}

func (r *tagRepository) Create(tag *blog_model.Tag) (*blog_model.Tag, error) {

	if err := r.db.Create(tag).Error; err != nil {
		return nil, err
	}

	return tag, nil

}

func (r *tagRepository) Update(id uint, tag *blog_model.Tag) (*blog_model.Tag, error) {

	tag.ID = id

	if err := r.db.Save(tag).Error; err != nil {
		return nil, err
	}

	return tag, nil
}

func (r *tagRepository) Remove(id uint) (*blog_model.Tag, error) {

	var tag blog_model.Tag

	if err := r.db.First(&tag, id).Error; err != nil {
		return nil, err
	}

	if err := r.db.Delete(&tag).Error; err != nil {
		return nil, err
	}

	return &tag, nil

}
