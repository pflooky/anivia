package entry

import (
	"github.com/kube-sailmaker/template-gen/task"
	templates "github.com/kube-sailmaker/template-gen/template"
	"log"
)

func TemplateGenerator(buildDir string, outputDir string) error {
	log.Println("Generating build templates found in " + buildDir)
	buildSpecs, buildErr := task.GenerateBuildSpecs(buildDir)
	if buildErr != nil {
		return buildErr
	}

	err := templates.RunTemplates(&buildSpecs, outputDir)
	if err != nil {
		return err
	}

	return nil
}
