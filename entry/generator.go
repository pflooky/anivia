package entry

import (
	"github.com/kube-sailmaker/template-gen/model"
	"github.com/kube-sailmaker/template-gen/task"
	templates "github.com/kube-sailmaker/template-gen/template"
)

func TemplateGenerator(appSpec *model.AppSpec, appDir string, resourceDir string, outputDir string) error {
	appTemplate := make([]templates.Application, 0)

	validationErr := appSpec.Validate()
	if validationErr != nil {
		return validationErr
	}
	appSpec.Normalise()

	for _, app := range appSpec.Apps {
		application, err := task.ProcessApplication(&app, appSpec.ReleaseName, appSpec.Environment, appDir, resourceDir)
		if err != nil {
			return err
		}
		appTemplate = append(appTemplate, *application)
	}

	releaseTemplate := templates.ReleaseTemplate{Application: appTemplate}
	return templates.Run(&releaseTemplate, outputDir)
}
