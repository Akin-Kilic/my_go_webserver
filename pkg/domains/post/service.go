package post

import "my_go_webserver/pkg/models"

type service struct {
	repository Repository
}

type Service interface {
	GetAllPosts() ([]models.Post, error)
	GetPost(id string) (models.Post, error)
	CreatePost(post models.Post) error
	UpdatePost(id string, post models.Post) error
	DeletePost(id string) error
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAllPosts() ([]models.Post, error) {
	posts, err := s.repository.GetAllPosts()
	if err != nil {
		return posts, err
	}
	return posts, nil
}

func (s *service) GetPost(id string) (models.Post, error) {
	post, err := s.repository.GetPost(id)
	if err != nil {
		return post, err
	}
	return post, nil
}

func (s *service) CreatePost(post models.Post) error {
	err := s.repository.CreatePost(post)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) UpdatePost(id string, post models.Post) error {
	err := s.repository.UpdatePost(id, post)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) DeletePost(id string) error {
	err := s.repository.DeletePost(id)
	if err != nil {
		return err
	}
	return nil
}
