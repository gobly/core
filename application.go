package core

import (
	"path/filepath"
	"runtime"
	"path"
	"io"
	"fmt"
	"strings"
)

type Application struct {
	Version string
	Name string
	Root string
}

var App = Application {
	"0.0.1",
	"Gobly Engine",
	filepath.Clean(callerPath(0) + strings.Repeat("../", 4)),
}

var callerPath = func(skip int) string {
	_, filename, _, success := runtime.Caller(skip)
	if !success {
		panic("No caller information")
	}

	return path.Dir(filename)
}

var CurrentFolder = func() string {
	return moduleRoot(callerPath(2))
}

var CurrentModule = func() string {
	return moduleRoot(callerPath(3))
}

func ShowWelcome(out io.Writer, router *Router) {
	fmt.Fprintln(out, "Gobly Web Framweork is up and running!")
	fmt.Fprintln(out, "Activated routes: ")
	router.FPrint(out)
}

func moduleRoot(path string) string {
	rPath, err := filepath.Rel(App.Root, path)
	if err != nil {
		panic("Cannot determine package path from " + rPath)
	}

	split := strings.Split(filepath.ToSlash(rPath), "/")

	// GitHub packages use the format github.com/user/package by default. So use first three tokens as a package ID
	if strings.HasPrefix(rPath, "github.com") {
		return filepath.Join(App.Root, split[0], split[1], split[2])
	}

	return filepath.Join(App.Root, split[0])
}

