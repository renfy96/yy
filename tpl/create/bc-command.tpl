package application

import "{{ .ProjectName }}/internal/bc/{{ .FileNameTitleLower }}/domain/repository"

type CommandService struct {
	{{ .FileNameTitleLower }}Repo repository.{{ .FileName }}Repository
}

func NewCommandService(repo repository.{{ .FileName }}Repository) *CommandService {
	return &CommandService{ {{ .FileNameTitleLower }}Repo: repo}
}

