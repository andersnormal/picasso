package templr

import (
	"runtime"
	"text/template"
)

var (
	templrFuncs template.FuncMap
)

func init() {
	templrFuncs = template.FuncMap{
		"OS":   func() string { return runtime.GOOS },
		"ARCH": func() string { return runtime.GOARCH },
	}
}
