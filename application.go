package core

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"
)

type Application struct {
	Version string
	Name    string
	Root    string
	Modules map[string]*Module
}

var App = Application{
	"0.0.1",
	"Gobly Engine",
	filepath.Clean(callerPath(0) + strings.Repeat("../", 4)),
	make(map[string]*Module),
}

func ShowWelcome(out io.Writer, router *Router) {
	fmt.Fprintf(out, "%s (v%s) is up and running!\n", App.Name, App.Version)
	fmt.Fprintln(out, "Activated routes: ")
	router.FPrint(out)
}
