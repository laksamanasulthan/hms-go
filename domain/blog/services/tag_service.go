package service

import (
	blog_model "hms-go/domain/blog/models"
	blog_repo "hms-go/domain/blog/repositories"
)

type TagService interface {
	GetAll() ([]blog_model.Tag, error)
	GetOne(id uint) (*blog_model.Tag, error)
	Create(tag *blog_model.Tag) (*blog_model.Tag, error)
	Update(id uint, tag *blog_model.Tag) (*blog_model.Tag, error)
	Remove(id uint) (*blog_model.Tag, error)
}

type tagService struct {
	repo blog_repo.TagRepository
}

func NewTagService(repo blog_repo.TagRepository) TagService {
	return &tagService{repo: repo}
}

func (s *tagService) GetAll() ([]blog_model.Tag, error) {

	val, err := s.repo.GetAll()

	if err != nil {
		return nil, err
	}

	return val, nil
}

func (s *tagService) GetOne(id uint) (*blog_model.Tag, error) {

	val, err := s.repo.GetOne(id)

	if err != nil {
		return nil, err
	}

	return val, nil
}

func (s *tagService) Create(tag *blog_model.Tag) (*blog_model.Tag, error) {

	val, err := s.repo.Create(tag)

	if err != nil {
		return nil, err
	}

	return val, nil
}

func (s *tagService) Update(id uint, tag *blog_model.Tag) (*blog_model.Tag, error) {
	val, err := s.repo.Update(id, tag)
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (s *tagService) Remove(id uint) (*blog_model.Tag, error) {
	val, err := s.repo.Remove(id)
	if err != nil {
		return nil, err
	}

	return val, nil
}
