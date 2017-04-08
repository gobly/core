package core

import "fmt"

type Module struct {
	Name  string
	Units map[string]*Unit

	prefix string // URL prefix that will be prepended to all URL-s for this module
}

// Functional unit that will be exposed in main menu
type Unit struct {
	Name     string
	Children map[string]*Unit

	prefix string // URL prefix that will be prepended to all URL-s for this unit
}

func (app *Application) RegisterModule(name string, prefix string, url string) *Module {
	app.Modules[prefix+url] = &Module{name, make(map[string]*Unit), prefix}
	return app.Modules[prefix+url]
}

func (module *Module) CreateUnit(name string, url string) *Unit {
	module.Units[module.prefix+url] = &Unit{name, make(map[string]*Unit), module.prefix + url}
	fmt.Println(module.Units)
	return module.Units[module.prefix+url]
}

func (unit *Unit) AddChild(name string, url string) *Unit {
	unit.Children[unit.prefix+url] = &Unit{name, make(map[string]*Unit), unit.prefix + url}
	return unit.Children[unit.prefix+url]
}
