package repository

import (
	"github.com/gin-gonic/gin"
	"{{ .ProjectName }}/internal/bc/{{ .FileNameTitleLower }}/domain/agg"
    "{{ .ProjectName }}/internal/bc/{{ .FileNameTitleLower }}/domain/agg/impl"
	"{{ .ProjectName }}/internal/bc/{{ .FileNameTitleLower }}/domain/agg"
)

type {{ .FileName }}Repository interface {
	GetById(ctx *gin.Context, id uint) agg.{{ .FileName }})
}





