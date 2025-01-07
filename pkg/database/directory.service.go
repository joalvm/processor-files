package database

import (
	"github.com/joalvm/processor-medias/pkg/models"
	"gorm.io/gorm"
)

type DirectoryService struct {
	db          *gorm.DB
	directories []models.Directory
}

func NewDirectoryService(db *gorm.DB) *DirectoryService {
	directories := []models.Directory{}

	db.Find(directories)

	return &DirectoryService{db: db, directories: directories}
}

func (s *DirectoryService) All() []models.Directory {
	return s.directories
}

func (s *DirectoryService) Find(id int) *models.Directory {
	for _, directory := range s.directories {
		if directory.Id == id {
			return &directory
		}
	}

	return nil
}

func (s *DirectoryService) Create(directory models.Directory) {
	s.db.Create(&directory)
	s.directories = append(s.directories, directory)
}
