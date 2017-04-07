package core

import (
	"path/filepath"
	"io"
	"fmt"
	"strings"
)

type Application struct {
	Version string
	Name string
	Root string
	Modules []Module
}

var App = Application {
	"0.0.1",
	"Gobly Engine",
	filepath.Clean(callerPath(0) + strings.Repeat("../", 4)),
	make([]Module, 0),
}

func ShowWelcome(out io.Writer, router *Router) {
	fmt.Fprintln(out, "Gobly Web Framweork is up and running!")
	fmt.Fprintln(out, "Activated routes: ")
	router.FPrint(out)
}
