package configurations

import (
	"encoding/json"
	"errors"
	"strings"
)

var ErrUnsupportedType = errors.New("unsupported type")
var ErrUnsupportedRequest = errors.New("unsupported request")

type Configuration interface {
	GetName() string
	Execute(cwd string) (exitCode int, err error)
}

type BaseConfigurationJSON struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

func (cfg *BaseConfigurationJSON) GetName() string {
	return cfg.Name
}

func ReadFromJSON(b []byte, vars Variables) (Configuration, error) {
	var base BaseConfigurationJSON

	err := json.Unmarshal(b, &base)
	if err != nil {
		return nil, err
	}

	switch base.Type {
	case "go":
		return NewGoConfigurationFromJSON(b, vars)
	default:
		return nil, ErrUnsupportedType
	}
}

type Variables struct {
	UserHome                string
	WorkspaceFolder         string
	WorkspaceFolderBaseName string
	File                    string
	FileWorkspaceFolder     string
	RelativeFile            string
	RelativeFileDirname     string
	FileBasename            string
	FileBasenameNoExtension string
	FileExtname             string
	FileDirname             string
	FileDirnameBasename     string
	Cwd                     string
	LineNumber              string
	SelectedText            string
	ExecPath                string
	DefaultBuildTask        string
	PathSeparator           string
}

func SubstituteVariables(s string, vars Variables) string {
	s = strings.ReplaceAll(s, "${userHome}", vars.UserHome)
	s = strings.ReplaceAll(s, "${workspaceFolder}", vars.WorkspaceFolder)
	s = strings.ReplaceAll(s, "${workspaceFolderBaseName}", vars.WorkspaceFolderBaseName)
	s = strings.ReplaceAll(s, "${file}", vars.File)
	s = strings.ReplaceAll(s, "${fileWorkspaceFolder}", vars.FileWorkspaceFolder)
	s = strings.ReplaceAll(s, "${relativeFile}", vars.RelativeFile)
	s = strings.ReplaceAll(s, "${relativeFileDirname}", vars.RelativeFileDirname)
	s = strings.ReplaceAll(s, "${fileBasename}", vars.FileBasename)
	s = strings.ReplaceAll(s, "${fileBasenameNoExtension}", vars.FileBasenameNoExtension)
	s = strings.ReplaceAll(s, "${fileExtname}", vars.FileExtname)
	s = strings.ReplaceAll(s, "${fileDirname}", vars.FileDirname)
	s = strings.ReplaceAll(s, "${fileDirnameBasename}", vars.FileDirnameBasename)
	s = strings.ReplaceAll(s, "${cwd}", vars.Cwd)
	s = strings.ReplaceAll(s, "${lineNumber}", vars.LineNumber)
	s = strings.ReplaceAll(s, "${selectedText}", vars.SelectedText)
	s = strings.ReplaceAll(s, "${execPath}", vars.ExecPath)
	s = strings.ReplaceAll(s, "${defaultBuildTask}", vars.DefaultBuildTask)
	s = strings.ReplaceAll(s, "${pathSeparator}", vars.PathSeparator)

	return s
}
