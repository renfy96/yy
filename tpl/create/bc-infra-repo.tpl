package repository

import (
	"github.com/gin-gonic/gin"
	"{{ .ProjectName }}/internal/bc/{{ .FileNameTitleLower }}/domain/agg"
	"{{ .ProjectName }}/internal/bc/{{ .FileNameTitleLower }}/domain/repository"
	"gorm.io/gorm"
)

type repositoryImpl struct {
	db *gorm.DB
}

func New{{ .FileName }}Repository(db *gorm.DB) repository.{{ .FileName }}Repository {
	return repositoryImpl{
		db: db,
	}
}