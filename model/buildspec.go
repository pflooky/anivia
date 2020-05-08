package model

import (
	"errors"
	"fmt"
	"strings"
)

type BuildSpec struct {
	Name      string            `yaml:"name"`
	Namespace string            `yaml:"namespace"`
	Folder    []string          `yaml:"folder"`
	Params    []Parameter       `yaml:"params"`
	Repo      map[string]string `yaml:"repo"`
	Build     Build             `yaml:"build"`
	Quality   Quality           `yaml:"quality"`
	Security  Security          `yaml:"security"`
}

type Parameter struct {
	Name        string `yaml:"name"`
	Default     string `yaml:"default-value"`
	Description string `yaml:"description"`
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
	if as.Namespace == "" {
		return errors.New("namespace is required")
	}
	if as.Repo == nil {
		return errors.New("repo:path is required")
	}
	if as.Build.Step == nil {
		return errors.New("build steps are required")
	}
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
}

func (as *Step) Normalise() {
	if as.Description == "" {
		as.Description = as.Type + " step"
	}
}
