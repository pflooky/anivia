package templates

import (
	"fmt"
	"log"
	"os"
)

func Run(releaseTemplate *ReleaseTemplate, outputDir string) {
	tmplArray := []string{"ServiceTemplate", "DeploymentTemplate", "ServiceAccountTemplate"}
	os.MkdirAll(outputDir, os.ModePerm)
	for _, application := range releaseTemplate.Application {
		os.Chdir(outputDir)
		os.Mkdir(application.Name, os.ModePerm)
		os.Chdir(application.Name)
		fmt.Println("Generating template for: ", application.Name)
		for _, tName := range tmplArray {
			tmpl := LoadTemplates(tName, &application)

			file, er := os.Create(tmpl.Name())
			if er != nil {
				log.Fatal("error ", er)
			}

			err := tmpl.Execute(file, &application)
			if err != nil {
				log.Fatal("error ", err)
			}
		}
	}

}
