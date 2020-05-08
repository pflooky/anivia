package main

import (
	"fmt"
	"github.com/kube-sailmaker/template-gen/entry"
	"os"
)

func main() {
	path, _ := os.Getwd()
	outputDir := path + "\\generated"
	buildDir := path + "\\sample\\build"

	//192.168.99.100
	err := entry.TemplateGenerator(buildDir, outputDir)
	if err != nil {
		fmt.Println(err)
	}
}
