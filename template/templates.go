package templates

import (
	"errors"
	"fmt"
	"github.com/kube-sailmaker/template-gen/model"
	"strings"
	"text/template"
)

var JenkinsBuildJobTemplate = `{{ range $entry := .Folder }}folder('{{ $entry }}')
{{ end }}
job('{{ .Namespace | ToFolder }}/{{ .Name }}') {
    displayName('{{ .Name }}-build')
    description('Build for {{ .Name }}')
	parameters {
		{{ range $entry := .Params }}stringParam('{{ $entry.Name }}', '{{ $entry.Default }}', '{{ $entry.Description }}')
		{{ end }}
	}
    scm {
        git('nexus-pathway/build/app-builds', 'master')
    }
    steps {
        groovyScriptFile('generated/builds/{{ .Namespace | ToFolder }}/{{ .Name }}-build.groovy')
    }
}
`

var BuildTemplate = `pipeline {
	{{ $name := .Name }}
	{{ $lang := .Build.Language | ToUpper }}
    agent any
    stages {
        stage("Checkout from Stash") {
            steps {
                sh "git clone nexus-stash-path/{{ .Name }}"
                sh "git checkout tags/$APP_VER -b master"
            }
        }{{ range $step := .Build.Step }}
		{{ if eq $step.Type "gradlew" }}stage("Build {{ $name }}") {
            steps {
				sh "JAVA_HOME=$JAVA{{ index $step.Args "jdk" }}_HOME"
                sh "gradlew {{ index $step.Args "tasks" }}"
            }
        }{{ else if eq $step.Type "docker" }}stage("Push image") {
            steps {
                sh "docker build -t nexus-images/{{ $name }}:$APP_VER ."
                sh "docker tag nexus-images/{{ $name }}:$APP_VER nexus-images/{{ $name }}:latest"
                sh "docker push nexus-images/{{ $name }}:$APP_VER"
                sh "docker push nexus-images/{{ $name }}:latest"
            }
        }{{ else if eq $step.Type "script" }}stage("Custom script") {
			steps {
				sh 
			}
		}{{ end }}{{ end }}
    }
}
`

var QualityTemplate = `
`

var SecurityTemplate = `
`

var VxPipelineTemplate = `NEXUS:
	display_name: "Project Nexus"
	updateExistingJobs: "on"
	repositories:
		{{ range $build := . }}"{{ $build.Name }}":
			branches:
				"{{ $build.Branches }}":
					"{{ $build.View }}":
						webhooks: "on"
						filename: "{{ $build.Filename }}"
						{{ if $build.Params }}parameters:{{range $param := $build.Params }}
							- {{ $param.Type }}:
								name: "{{ $param.Name }}"
								default: "{{ $param.Default }}"
								description: "{{ $param.Description }}"
								{{ if eq $param.Type "choice" }}choices: {{ $param.Choices | ToChoices }}{{ end }}{{ end }}{{ end }}
		{{ end }}
`

func LoadVxPipelineTemplate() (*template.Template, error) {
	return getTemplate("NEXUS.yaml", VxPipelineTemplate)
}

//LoadTemplates parse static template to helm chart
func LoadTemplates(tName string, app *model.BuildSpec) (*template.Template, error) {
	switch tName {
	case "jenkins-build-job":
		return getTemplate(fmt.Sprintf("%s-jenkins-build-job.yaml", app.Name), JenkinsBuildJobTemplate)
	case "build-script":
		return getTemplate(fmt.Sprintf("%s-build.groovy", app.Name), BuildTemplate)
	case "quality-job":
		return getTemplate(fmt.Sprintf("%s-quality.groovy", app.Name), QualityTemplate)
	case "security-job":
		return getTemplate(fmt.Sprintf("%s-security.groovy", app.Name), SecurityTemplate)
	}
	return nil, nil
}

func getTemplate(name string, templateType string) (*template.Template, error) {
	funcMap := template.FuncMap{
		"ToUpper":   strings.ToUpper,
		"ToLower":   strings.ToLower,
		"ToFolder":  ToFolder,
		"ToChoices": ToChoices,
	}

	tmpl, err := template.New(name).Funcs(funcMap).Parse(templateType)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error parsing %v ", err))
	}
	return tmpl, nil
}

func ToFolder(s string) string {
	return strings.ReplaceAll(s, "::", "/")
}

func ToChoices(s []string) []string {
	quotedS := make([]string, len(s))
	for i, v := range s {
		quotedS[i] = fmt.Sprintf("\"%s\"", v)
	}
	return quotedS
}
