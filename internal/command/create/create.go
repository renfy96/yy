package create

import (
	"fmt"
	"github.com/renfy96/yy/internal/pkg/helper"
	"github.com/renfy96/yy/tpl"
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
	Use:     "create [type] [name]",
	Short:   "Create a new application/bc",
	Example: "yy create bc user",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

	},
}
var CreateApplicationCmd = &cobra.Command{
	Use:     "application",
	Short:   "Create a new application",
	Example: "yy create application user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}
var CreateBCCmd = &cobra.Command{
	Use:     "bc",
	Short:   "Create a new bc",
	Example: "yy create bc user",
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
	case "application":
		c.createApplication()
	case "bc":
		c.createBc()
	default:
		log.Fatalf("Invalid handler type: %s", c.CreateType)
	}

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

func (c *Create) createApplication() {
	filePath := c.FilePath
	if filePath == "" {
		filePath = fmt.Sprintf("internal/%s/%s/", c.CreateType, strings.ToLower(c.FileName))
	}
	var (
		pPath = filePath + "params/"
		sPath = filePath + "service/"
	)

	genFile(c, pPath, "app-params", "params")

	genFile(c, sPath, "app-service", strings.ToLower(c.FileName))
}
func (c *Create) createBc() {
	filePath := c.FilePath
	if filePath == "" {
		filePath = fmt.Sprintf("internal/%s/%s/", c.CreateType, strings.ToLower(c.FileName))
	}
	var (
		aPath = filePath + "application/"
		sPath = filePath + "domain/"
	)

	genFile(c, aPath, "bc-query", "query_service")
	genFile(c, aPath, "bc-command", "command_service")

	genFile(c, sPath+"agg/", "bc-agg", strings.ToLower(c.FileName))
	genFile(c, sPath+"repository/", "bc-repository", strings.ToLower(c.FileName))
	genFile(c, sPath+"infrastructure/repository/", "bc-infra-repo", strings.ToLower(c.FileName))
}

func genFile(c *Create, dirPath, tmp, fileName string) {
	f := createFile(dirPath, fileName+".go")
	if f == nil {
		log.Printf("warn: file %s%s %s", dirPath, strings.ToLower(fileName)+".go", "already exists.")
		return
	}
	defer f.Close()
	t, err := template.ParseFS(tpl.CreateTemplateFS, fmt.Sprintf("create/%s.tpl", tmp))
	if err != nil {
		log.Fatalf("create %s error: %s", fileName, err.Error())
	}
	err = t.Execute(f, c)
	if err != nil {
		log.Fatalf("create %s error: %s", c.CreateType, err.Error())
	}
	log.Printf("Created new %s: %s", c.CreateType, dirPath+strings.ToLower(c.FileName)+".go")
}
