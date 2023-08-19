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

func (r repositoryImpl) GetById(ctx *gin.Context, id uint) agg.{{ .FileName }} {
	{{ .FileNameTitleLower }} := r.q.{{ .FileName }}
	{{ .FileNameTitleLower }}Mol, _ := {{ .FileNameTitleLower }}.WithContext(ctx).Where({{ .FileNameTitleLower }}.ID.Eq(id)).First()
	return poToAgg({{ .FileNameTitleLower }}Mol)
}

func poToAgg({{ .FileNameTitleLower }} *model.{{ .FileName }}) agg.{{ .FileName }} {
	return impl.New{{ .FileName }}Impl({{ .FileNameTitleLower }})
}