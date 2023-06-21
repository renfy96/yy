package service

import (
	"github.com/gin-gonic/gin"
	"{{ .ProjectName }}/internal/boudary_context/{{ .FileNameTitleLower }}/application"
)

type Service struct {
	query *application.QueryService
	command *application.CommandService
}

func NewService(query *application.QueryService, command *application.CommandService) *Service {
	return &Service{query: query, command: command}
}