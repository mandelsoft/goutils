package stringutils

import (
	"strings"
)

func IndentLines(orig string, gap string, skipfirst ...bool) string {
	return JoinIndentLines(strings.Split(strings.TrimPrefix(orig, "\n"), "\n"), gap, skipfirst...)
}

func JoinIndentLines(orig []string, gap string, skipfirst ...bool) string {
	if len(orig) == 0 {
		return ""
	}
	skip := false
	for _, b := range skipfirst {
		skip = skip || b
	}

	s := ""
	if !skip {
		s = gap
	}
	return s + strings.Join(orig, "\n"+gap)
}
