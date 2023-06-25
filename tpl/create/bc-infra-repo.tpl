package repository

import (
	"github.com/gin-gonic/gin"
	"{{ .ProjectName }}/internal/bc/{{ .FileNameTitleLower }}/domain/agg"
	"{{ .ProjectName }}/internal/bc/{{ .FileNameTitleLower }}/domain/repository"
	"{{ .ProjectName }}/internal/repository/query"
	"gorm.io/gorm"
)

type repositoryImpl struct {
	q *query.Query
}

func New{{ .FileName }}Repository(db *gorm.DB) repository.{{ .FileName }}Repository {
	return repositoryImpl{
    		q: query.Use(db),
    }
}