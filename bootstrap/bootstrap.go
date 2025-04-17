package bootstrap

import (
	"embed"
	"errors"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/manicar2093/gomancer/domain"
	"net/url"
	"os"
	"path"
	"strings"
	"text/template"
)

//go:embed templates/*
var templatesFS embed.FS

//go:embed core/*
var coreFS embed.FS

type (
	fileWithContent struct {
		name, tmplName string
	}
	dirsWithFiles struct {
		dirPath string
		files   []fileWithContent
	}
)

var (
	initialFiles = []dirsWithFiles{
		{
			"",
			[]fileWithContent{
				{
					".env",
					"env_file",
				},
				{
					"package.json",
					"package_json",
				},
				{
					".cz.toml",
					"cz.toml",
				},
				{
					"go.mod",
					"go.mod",
				},
				{
					"Taskfile.yml",
					"taskfile.yml",
				},
				{
					".air.toml",
					"air.toml",
				},
				{
					"README.md",
					"README.md",
				},
			},
		},
		{
			string(domain.PrismaSchemaPackagePath),
			[]fileWithContent{
				{
					"schema.prisma",
					"schema_prisma",
				},
			},
		},
		{
			string(domain.InternalDomainModelsPackagePath),
			[]fileWithContent{
				{
					"init.go",
					"internal_domain_init",
				},
			},
		},
		{
			string(domain.CmdApiControllersPackagePath),
			[]fileWithContent{
				{
					"init.go",
					"cmd_api_controllers_init",
				},
			},
		},
		{
			string(domain.CmdApiPackagePath),
			[]fileWithContent{
				{
					"main.go",
					"cmd_api_main",
				},
			},
		},
		{
			string(domain.PkgGeneratorsPackagePath),
			[]fileWithContent{
				{
					"generators.go",
					"pkg_generators_generators",
				},
			},
		},
		{
			string(domain.PkgConfigPackagePath),
			[]fileWithContent{
				{
					"config.go",
					"pkg_config_config.go",
				},
			},
		},
		{
			string(domain.PkgVersioningPackagePath),
			[]fileWithContent{
				{
					"version.go",
					"pkg_versioning_version.go",
				},
			},
		},
		{
			string(domain.GithubWorkflowsPackagePath),
			[]fileWithContent{
				{
					"bump_version.yml",
					"github_workflows_bump_version.yml",
				},
			},
		},
	}
)

type InitProjectInput struct {
	ModuleName string
}

func InitProject(input InitProjectInput) (string, error) {
	log.Info("Initiating new project...")

	asUrl, err := url.Parse(input.ModuleName)
	if err != nil {
		return "", errors.Join(err, fmt.Errorf("%s is not a valid module moduleName. follow convention of use urls as github.com/<user>/<repo>, gitlab.com/<user>/<repo>, bitbucket.com/<user>/<repo>, etc", input.ModuleName))
	}

	pathSlice := strings.Split(asUrl.Path, string(os.PathSeparator))
	projectDirName := pathSlice[len(pathSlice)-1]

	if err := os.Mkdir(projectDirName, os.ModePerm); err != nil {
		if !os.IsExist(err) {
			return "", err
		}
		if false {
			return "", fmt.Errorf("%s already exists, use -f to force creation", projectDirName)
		}
	}

	if err := copyCoreToProject(projectDirName, input); err != nil {
		return "", err
	}

	var (
		tpl          = template.Must(template.ParseFS(templatesFS, "templates/*"))
		templateData = struct {
			InitProjectInput
			ProjectName                     string
			DevEnvironment, TestEnvironment string
			GoVersion                       string
		}{
			InitProjectInput: input,
			ProjectName:      projectDirName,
			DevEnvironment:   "dev",
			TestEnvironment:  "test",
			GoVersion:        "1.23",
		}
	)
	for _, dir := range initialFiles {
		fullDirPath := path.Join(projectDirName, dir.dirPath)
		if err := os.MkdirAll(fullDirPath, os.ModePerm); err != nil {
			return "", err
		}
		for _, file := range dir.files {

			fullFilePath := path.Join(fullDirPath, file.name)
			f, err := os.OpenFile(fullFilePath, os.O_RDWR|os.O_CREATE, 0755)
			if err != nil {
				return "", err
			}
			if err := tpl.ExecuteTemplate(f, file.tmplName, templateData); err != nil {
				return "", err
			}

		}
	}
	return projectDirName, nil
}
