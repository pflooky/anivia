package task

import (
	"fmt"
	"github.com/kube-sailmaker/template-gen/functions"
)

const (
	infraManifest    = "%s/infrastructure/%s.yaml"
	mixinManifest    = "%s/mixins/%s.yaml"
	resourceManifest = "%s/resources/%s.yaml"
)

func GetInfrastructure(name string, t interface{}, resourceDir string) {
	file := fmt.Sprintf(infraManifest, resourceDir, name)
	functions.UnmarshalFile(file, t)
}

func GetMixin(name string, t interface{}, resourceDir string) {
	file := fmt.Sprintf(mixinManifest, resourceDir, name)
	functions.UnmarshalFile(file, t)
}

func GetResource(name string, t interface{}, resourceDir string) {
	file := fmt.Sprintf(resourceManifest, resourceDir, name)
	functions.UnmarshalFile(file, t)
}
