package main

import (
	"flag"
	"fmt"
	"github.com/kube-sailmaker/template-gen/entry"
	"os"
)

func main() {
	path, _ := os.Getwd()
	outputDir := path + "\\generated"
	buildDir := path + "\\sample\\build"

	generate := flag.String("generate", "vx-pipeline", "Choose to generate either vx-pipeline or jenkins")

	//192.168.99.100
	switch *generate {
	case "vx-pipeline":
		{
			err := entry.RunVxPipelineTemplates(buildDir, outputDir)
			if err != nil {
				fmt.Println(err)
			}
		}
	case "jenkins":
		{
			err := entry.TemplateGenerator(buildDir, outputDir)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

}
