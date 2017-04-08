package core

type Module struct {
	Name  string
	Units map[string]*Unit
}

// Functional unit that will be exposed in main menu
type Unit struct {
	Name     string
	Children map[string]*Unit
}

func (app *Application) RegisterModule(name string, url string) *Module {
	app.Modules[url] = &Module{name, make(map[string]*Unit)}
	return app.Modules[url]
}

func (module *Module) CreateUnit(name string, url string) *Unit {
	module.Units[url] = &Unit{name, make(map[string]*Unit)}
	return module.Units[url]
}

func (unit *Unit) AddChild(name string, url string) *Unit {
	unit.Children[url] = &Unit{name, make(map[string]*Unit)}
	return unit.Children[url]
}
