package entry

import (
	"github.com/kube-sailmaker/template-gen/model"
	"github.com/kube-sailmaker/template-gen/task"
	templates "github.com/kube-sailmaker/template-gen/template"
)

func TemplateGenerator(appSpec *model.AppSpec, appDir string, resourceDir string, outputDir string) {
	appTemplate := make([]templates.Application, 0)
	for _, app := range appSpec.App {
		application := task.ProcessApplication(&app, appSpec.ReleaseName, appSpec.Environment, appDir, resourceDir)
		appTemplate = append(appTemplate, *application)
	}

	releaseTemplate := templates.ReleaseTemplate{Application: appTemplate}
	templates.Run(&releaseTemplate, outputDir)
}
