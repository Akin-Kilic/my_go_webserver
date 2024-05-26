package post

import (
	"my_go_webserver/pkg/models"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

type Repository interface {
	GetAllPosts() ([]models.Post, error)
	GetPost(id string) (models.Post, error)
	CreatePost(post models.Post) error
	UpdatePost(id string, post models.Post) error
	DeletePost(id string) error
}

func (r *repository) GetAllPosts() ([]models.Post, error) {
	var posts []models.Post
	err := r.db.Model(&models.Post{}).Find(&posts).Error
	if err != nil {
		return posts, err
	}
	return posts, nil
}

func (r *repository) GetPost(id string) (models.Post, error) {
	var post models.Post
	err := r.db.Model(&models.Post{}).Where("id = ?", id).First(&post).Error
	if err != nil {
		return post, err
	}
	return post, nil
}

func (r *repository) CreatePost(post models.Post) error {
	err := r.db.Model(&models.Post{}).Create(&post).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) UpdatePost(id string, post models.Post) error {
	err := r.db.Model(&models.Post{}).Where("id = ?", id).Updates(post).Error
    if err != nil {
        return err
    }
    return nil
}

func (r *repository) DeletePost(id string) error {
	err := r.db.Model(&models.Post{}).Delete(&models.Post{}, "id = ?", id).Error
	if err != nil {
		return err
	}
	return nil
}
