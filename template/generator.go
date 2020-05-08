package templates

import (
	"errors"
	"fmt"
	"github.com/kube-sailmaker/template-gen/model"
	"log"
	"os"
	"path/filepath"
)

func createDirSafely(fileName string) error {
	dirName := filepath.Dir(fileName)
	if _, serr := os.Stat(dirName); serr != nil {
		merr := os.MkdirAll(dirName, os.ModePerm)
		if merr != nil {
			return merr
		}
	}
	return nil
}

func RunTemplates(buildSpecs *[]model.BuildSpec, outputDir string) error {
	tmplArray := []string{"jenkins-build-job", "build-script"}

	for _, buildSpec := range *buildSpecs {
		for _, tName := range tmplArray {
			log.Println(fmt.Sprintf("Generating %s template for: %s", tName, buildSpec.Name))
			appWorkDir := fmt.Sprintf("%s\\build\\%s\\%s\\", outputDir, buildSpec.Name, tName)
			appDirErr := createDirSafely(appWorkDir)
			if appDirErr != nil {
				return appDirErr
			}

			tmpl, tmplErr := LoadTemplates(tName, &buildSpec)
			if tmplErr != nil {
				return errors.New(fmt.Sprintf("[app]: %s, [error]: %v", buildSpec.Name, tmplErr))
			}
			file, fileErr := os.Create(fmt.Sprintf("%s\\%s", appWorkDir, tmpl.Name()))
			if fileErr != nil {
				return fileErr
			}

			exErr := tmpl.Execute(file, &buildSpec)
			if exErr != nil {
				return errors.New(fmt.Sprintf("[app]: %s, [error]: %v", buildSpec.Name, exErr))
			}
		}
	}

	return nil
}
