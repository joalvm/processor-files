package database

import (
	"github.com/joalvm/processor-medias/pkg/models"
	"gorm.io/gorm"
)

type MediasService struct {
	db     *gorm.DB
	medias *[]models.Media
}

func NewFilesService(db *gorm.DB) *MediasService {
	files := []models.Media{}

	db.Find(&files)

	return &MediasService{db: db, medias: &files}
}

func (s *MediasService) All() *[]models.Media {
	return s.medias
}

func (s *MediasService) Find(id int) *models.Media {
	for _, media := range *s.medias {
		if media.Id == id {
			return &media
		}
	}

	return nil
}

func (s *MediasService) Create(media *models.Media) {
	s.db.Create(media)
}

func (s *MediasService) CreateIfNotExists(media *models.Media) {
	if s.Find(media.Id) == nil {
		s.Create(media)
	}
}
