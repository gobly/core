package core

type Module struct {
	Name string
	URL string
}

func (app *Application) RegisterModule(name string, url string) {
	app.Modules = append(app.Modules, Module{name, url})
}
