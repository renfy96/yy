package repository

import (
	"github.com/gin-gonic/gin"
	"{{ .ProjectName }}/internal/bc/{{ .FileNameTitleLower }}/domain/agg"
)

type {{ .FileName }}Repository interface {
	Save(ctx *gin.Context, {{ .FileNameTitleLower }} agg.{{ .FileName }})
}
