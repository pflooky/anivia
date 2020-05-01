package model

type Resource struct {
	ApiVersion string            `yaml:"apiVersion"`
	Kind       string            `yaml:"kind"`
	Metadata   map[string]string `yaml:"metadata"`
	Spec       ResourceSpec      `yaml:"spec"`
}

type ResourceSpec struct {
	ResourceTemplate []ResourceTemplate `yaml:"template"`
}

type ResourceTemplate struct {
	Name    string            `yaml:"name"`
	Element map[string]string `yaml:"element"`
	Infra   string            `yaml:"infrastructure"`
}
