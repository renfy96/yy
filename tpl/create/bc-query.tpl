package application

import "{{ .ProjectName }}/internal/bc/{{ .FileNameTitleLower }}/domain/repository"

type QueryService struct {
    {{ .FileNameTitleLower }}Repo repository.{{ .FileName }}Repository
}

func NewQueryService(repo repository.{{ .FileName }}Repository) *QueryService {
    return &QueryService{ {{ .FileNameTitleLower }}Repo: repo}
}

