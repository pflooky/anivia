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

func GetInfrastructure(name string, t interface{}, resourceDir string) error {
	file := fmt.Sprintf(infraManifest, resourceDir, name)
	return functions.UnmarshalFile(file, t)
}

func GetMixin(name string, t interface{}, resourceDir string) error {
	file := fmt.Sprintf(mixinManifest, resourceDir, name)
	return functions.UnmarshalFile(file, t)
}

func GetResource(name string, t interface{}, resourceDir string) error {
	file := fmt.Sprintf(resourceManifest, resourceDir, name)
	return functions.UnmarshalFile(file, t)
}
