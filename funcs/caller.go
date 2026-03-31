package funcs

import (
	"runtime"
	"strings"

	"github.com/mandelsoft/goutils/general"
)

func CallerInfo(skip int, adjust ...int) (file string, line int, pkg, fname string) {
	pc, file, line, ok := runtime.Caller(skip + 1)
	line += general.Optional(adjust...)
	if ok {
		fn := runtime.FuncForPC(pc)
		if fn != nil {
			pkg, name := splitPath(fn.Name())
			_, file := splitPath(file)
			return file, line, pkg, name
		}
		return file, line, "", ""
	}
	return "", 0, "", ""
}

func splitPath(s string) (string, string) {
	idx := strings.LastIndex(s, "/")
	if idx < 0 {
		idx = strings.LastIndex(s, "\\")
	}
	if idx < 0 {
		return "", s
	}
	return s[:idx], s[idx+1:]
}
