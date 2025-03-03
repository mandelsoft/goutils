package stringutils

import (
	"github.com/mandelsoft/goutils/sliceutils"
)

// PadRight returns a new string of a specified minimal length in which the end of the current string is padded with a specified Unicode character.
func PadRight(str string, length int, pad byte) string {
	for len(str) < length {
		str += string(pad)
	}
	return str
}

// PadLeft returns a new string of a specified minimal length in which the beginning of the current string is padded with a specified Unicode character.
func PadLeft(str string, length int, pad byte) string {
	for len(str) < length {
		str = string(pad) + str
	}
	return str
}

// AlignLeft aligns a slice of strings by padding it right to the
// length of the longest entry.
func AlignLeft(in []string, pad byte) []string {
	maxlen := 0
	for _, s := range in {
		maxlen = max(maxlen, len(s))
	}
	return sliceutils.Transform(in, func(s string) string {
		return PadRight(s, maxlen, pad)
	})
}

// AlignRight aligns a slice of strings by padding it left to the
// length of the longest entry.
func AlignRight(in []string, pad byte) []string {
	maxlen := 0
	for _, s := range in {
		maxlen = max(maxlen, len(s))
	}
	return sliceutils.Transform(in, func(s string) string {
		return PadLeft(s, maxlen, pad)
	})
}
