package model

type AppSpec struct {
	Namespace   string
	ReleaseName string
	Environment string
	App         []App
}

type App struct {
	Name    string
	Version string
}
