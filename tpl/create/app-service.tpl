package service

import (
	"{{ .ProjectName }}/internal/bc/{{ .FileNameTitleLower }}/application"
)

type Service struct {
	query *application.QueryService
	command *application.CommandService
}

func NewService(query *application.QueryService, command *application.CommandService) *Service {
	return &Service{query: query, command: command}
}