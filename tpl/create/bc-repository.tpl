package repository

import (
	"github.com/gin-gonic/gin"
	"{{ .ProjectName }}/internal/bc/{{ .FileNameTitleLower }}/domain/agg"
    "{{ .ProjectName }}/internal/bc/{{ .FileNameTitleLower }}/domain/agg/impl"
	"{{ .ProjectName }}/internal/bc/{{ .FileNameTitleLower }}/domain/agg"
)

type {{ .FileName }}Repository interface {
	Save(ctx *gin.Context, {{ .FileNameTitleLower }} agg.{{ .FileName }})
}


func (r repositoryImpl) GetById(ctx *gin.Context, id uint) agg.{{ .FileName }} {
	{{ .FileNameTitleLower }} := r.q.{{ .FileName }}
	{{ .FileNameTitleLower }}Mol, _ := {{ .FileNameTitleLower }}.WithContext(ctx).Where({{ .FileNameTitleLower }}.ID.Eq(id)).First()
	return poToAgg({{ .FileNameTitleLower }}Mol)
}




func poToAgg({{ .FileNameTitleLower }} *model.{{ .FileName }}) agg.{{ .FileName }} {
	return impl.New{{ .FileName }Impl({{ .FileNameTitleLower })
}