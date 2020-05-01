package templates

import (
	"errors"
	"fmt"
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

func Run(releaseTemplate *ReleaseTemplate, outputDir string) error {
	tmplArray := []string{"ServiceTemplate", "DeploymentTemplate", "ServiceAccountTemplate"}
	for _, application := range releaseTemplate.Application {
		appWorkDir := fmt.Sprintf("%s/%s/", outputDir, application.Name)
		cerr := createDirSafely(appWorkDir)
		if cerr != nil {
			return cerr
		}
		log.Println("Generating template for: ", application.Name)
		for _, tName := range tmplArray {
			tmpl, err := LoadTemplates(tName, &application)
			if err != nil {
				return errors.New(fmt.Sprintf("[app]: %s, [error]: %v", application.Name, err))
			}
			file, er := os.Create(fmt.Sprintf("%s/%s", appWorkDir, tmpl.Name()))
			if er != nil {
				return er
			}

			exErr := tmpl.Execute(file, &application)
			if exErr != nil {
				return errors.New(fmt.Sprintf("[app]: %s, [error]: %v", application.Name, err))
			}
		}
	}
	return nil

}
