package model

import (
	"errors"
	"fmt"
	"strings"
)

type BuildSpec struct {
	Name      string            `yaml:"name"`
	Namespace string            `yaml:"namespace"`
	Branches  string            `yaml:"branches"`
	View      string            `yaml:"view"`
	Filename  string            `yaml:"filename"`
	Folder    []string          `yaml:"folder"`
	Params    []Parameter       `yaml:"params"`
	Repo      map[string]string `yaml:"repo"`
	Build     Build             `yaml:"build"`
	Quality   Quality           `yaml:"quality"`
	Security  Security          `yaml:"security"`
}

type Parameter struct {
	Name        string   `yaml:"name"`
	Type        string   `yaml:"type"`
	Default     string   `yaml:"default-value"`
	Description string   `yaml:"description"`
	Choices     []string `yaml:"choices"`
}

type Build struct {
	Language string `yaml:"lang"`
	Step     []Step `yaml:"step"`
}

type Step struct {
	Type        string            `yaml:"type"`
	Description string            `yaml:"description"`
	Args        map[string]string `yaml:"args"`
}

type Quality struct {
	Enabled bool `yaml:"enabled"`
}

type Security struct {
	Enabled bool `yaml:"enabled"`
}

func (as *BuildSpec) Validate() error {
	if as.Name == "" {
		return errors.New("name is required")
	}

	appName := as.Name
	errorMsg := fmt.Sprintf("app-name=%s, error=", appName)

	if as.Namespace == "" {
		return errors.New(errorMsg + "namespace is required")
	}
	if as.Repo == nil {
		return errors.New(errorMsg + "repo:path is required")
	}
	if as.Params != nil {
		for _, param := range as.Params {
			if param.Name == "" {
				return errors.New(errorMsg + "parameter requires name to be defined")
			}
			if param.Type == "" {
				return errors.New(errorMsg + "parameter requires type to be defined as either string, boolean or choice")
			}
		}
	}
	//if as.Build.Step == nil {
	//	return errors.New(errorMsg + "build steps are required")
	//}
	//for _, step := range as.Build.Step {
	//	if step.Type == "" {
	//		return errors.New(errorMsg + "type is required for build step")
	//	}
	//}
	return nil
}

func (as *BuildSpec) Normalise() {
	if as.Folder == nil {
		split := strings.Split(as.Namespace, "::")
		folderStructure := make([]string, len(split))
		for i, s := range split {
			if i == 0 {
				folderStructure[i] = s
			} else {
				folderStructure[i] = fmt.Sprintf("%s/%s", folderStructure[i-1], s)
			}
		}

		as.Folder = folderStructure
	}
	if as.Branches == "" {
		as.Branches = "(master|develop|release/*|feature/*|features/*|hotfix/*)"
	}
	if as.View == "" {
		as.View = "build"
	}
	if as.Filename == "" {
		as.Filename = "jenkins/deploy/JenkinsFile"
	}
}

func (as *Step) Normalise() {
	if as.Description == "" {
		as.Description = as.Type + " step"
	}
}
