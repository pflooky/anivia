package main

import (
	"github.com/kube-sailmaker/template-gen/entry"
	"github.com/kube-sailmaker/template-gen/model"
	"os"
)

func main() {
	busybox := model.App{
		Name:    "busybox",
		Version: "latest",
	}
	nginx := model.App{
		Name:    "nginx",
		Version: "latest",
	}

	appList := make([]model.App, 0)
	appList = append(appList, busybox, nginx)

	appSpec := model.AppSpec{
		Namespace:   "apps",
		ReleaseName: "Release-2",
		Environment: "test",
		Apps:         appList,
	}
	path, _ := os.Getwd()
	appDir := path + "/sample-manifest/user/apps"
	resourceDir := path + "/sample-manifest/provider"
	outputDir := path + "/tmp"

	entry.TemplateGenerator(&appSpec, appDir, resourceDir, outputDir)
}
