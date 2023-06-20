package create

import (
	"fmt"
	"github.com/go-nunu/nunu/tpl"
	"github.com/renfy96/yy/internal/pkg/helper"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Create struct {
	ProjectName        string
	CreateType         string
	FilePath           string
	FileName           string
	FileNameTitleLower string
	FileNameFirstChar  string
	IsFull             bool
}

func NewCreate() *Create {
	return &Create{}
}

var CreateCmd = &cobra.Command{
	Use:     "create [type] [handler-name]",
	Short:   "Create a new handler/service/repository.tpl/model",
	Example: "yy create handler user",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

	},
}
var CreateHandlerCmd = &cobra.Command{
	Use:     "handler",
	Short:   "Create a new handler",
	Example: "yy create handler user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}
var CreateServiceCmd = &cobra.Command{
	Use:     "service",
	Short:   "Create a new service",
	Example: "yy create service user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}
var CreateRepositoryCmd = &cobra.Command{
	Use:     "repository.tpl",
	Short:   "Create a new repository.tpl",
	Example: "yy create repository.tpl user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}
var CreateModelCmd = &cobra.Command{
	Use:     "model",
	Short:   "Create a new model",
	Example: "yy create model user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}
var CreateAllCmd = &cobra.Command{
	Use:     "all",
	Short:   "Create a new handler & service & repository.tpl & model",
	Example: "yy create all user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}

func runCreate(cmd *cobra.Command, args []string) {
	c := NewCreate()
	c.ProjectName = helper.GetProjectName(".")
	c.CreateType = cmd.Use
	c.FilePath, c.FileName = filepath.Split(args[0])
	c.FileName = strings.ReplaceAll(strings.ToUpper(string(c.FileName[0]))+c.FileName[1:], ".go", "")
	c.FileNameTitleLower = strings.ToLower(string(c.FileName[0])) + c.FileName[1:]
	c.FileNameFirstChar = string(c.FileNameTitleLower[0])

	switch c.CreateType {
	case "handler", "service", "repository.tpl", "model":
		c.genFile()
	case "all":

		c.CreateType = "handler"
		c.genFile()

		c.CreateType = "service"
		c.genFile()

		c.CreateType = "repository.tpl"
		c.genFile()

		c.CreateType = "model"
		c.genFile()
	default:
		log.Fatalf("Invalid handler type: %s", c.CreateType)
	}

}
func (c *Create) genFile() {
	filePath := c.FilePath
	if filePath == "" {
		filePath = fmt.Sprintf("internal/%s/", c.CreateType)
	}
	f := createFile(filePath, strings.ToLower(c.FileName)+".go")
	if f == nil {
		log.Printf("warn: file %s%s %s", filePath, strings.ToLower(c.FileName)+".go", "already exists.")
		return
	}
	defer f.Close()

	t, err := template.ParseFS(tpl.CreateTemplateFS, fmt.Sprintf("create/%s.tpl", c.CreateType))
	if err != nil {
		log.Fatalf("create %s error: %s", c.CreateType, err.Error())
	}
	err = t.Execute(f, c)
	if err != nil {
		log.Fatalf("create %s error: %s", c.CreateType, err.Error())
	}
	log.Printf("Created new %s: %s", c.CreateType, filePath+strings.ToLower(c.FileName)+".go")

}
func createFile(dirPath string, filename string) *os.File {
	filePath := dirPath + filename
	// 创建文件夹
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create dir %s: %v", dirPath, err)
	}
	stat, _ := os.Stat(filePath)
	if stat != nil {
		return nil
	}
	// 创建文件
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", filePath, err)
	}

	return file
}
