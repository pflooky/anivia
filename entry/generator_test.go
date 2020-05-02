package entry

import (
	"fmt"
	"github.com/kube-sailmaker/template-gen/model"
	"github.com/kube-sailmaker/template-gen/test"
	"os"
	"testing"
)

func TestTemplateGeneratorThrowsErrorForMissingParams(t *testing.T) {
	appSpec := model.AppSpec{}
	_, err := TemplateGenerator(&appSpec, "", "", "")
	test.NotNull(t, err)
	test.EqualTo(t, "namespace is required", fmt.Sprintf("%v", err))
}

func TestTemplateGeneratorWithEmptyReleaseName(t *testing.T) {
	spec := MockSpec()
	_, err := TemplateGenerator(spec, "test", "test", "test")
	test.NotNull(t, err)
	test.EqualTo(t, "app to deploy cannot be empty", fmt.Sprintf("%v", err))
}

func TestTemplateGenerator(t *testing.T) {
	outputDir := "../tmp"

	os.RemoveAll(outputDir)
	appSpec := GetAppSpec()
	appDir := "../sample-manifest/user/apps"
	resourceDir := "../sample-manifest/provider"
	summary, err := TemplateGenerator(appSpec, appDir, resourceDir, outputDir)
	test.Null(t, err)

	stat, err := os.Stat("../tmp/busybox")
	test.Null(t, err)
	test.NotNull(t, summary)
	test.EqualTo(t, "apps", summary.Namespace)
	test.EqualTo(t, 1, len(summary.Items))
	test.EqualTo(t, true, stat.IsDir())
	os.RemoveAll(outputDir)
}

func MockSpec() *model.AppSpec {
	appList := make([]model.App, 0)
	return &model.AppSpec{
		Namespace:   "apps",
		ReleaseName: "",
		Environment: "test",
		Apps:        appList,
	}
}

func GetAppSpec() *model.AppSpec {
	busybox := model.App{
		Name:    "busybox",
		Version: "latest",
	}

	appList := make([]model.App, 0)
	appList = append(appList, busybox)

	return &model.AppSpec{
		Namespace:   "apps",
		ReleaseName: "Release-2",
		Environment: "test",
		Apps:        appList,
	}
}
