package task

import (
	"github.com/kube-sailmaker/template-gen/functions"
	"github.com/kube-sailmaker/template-gen/model"
	"os"
	"path/filepath"
)

func GenerateBuildSpecs(buildDir string) ([]model.BuildSpec, error) {
	buildSpecs := make([]model.BuildSpec, 0)

	err := filepath.Walk(buildDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		buildSpec := &model.BuildSpec{}
		err = functions.UnmarshalFile(path, buildSpec)
		if err != nil {
			return err
		}

		validErr := buildSpec.Validate()
		if validErr != nil {
			return validErr
		}

		buildSpec.Normalise()
		buildSpecs = append(buildSpecs, *buildSpec)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return buildSpecs, nil
}
