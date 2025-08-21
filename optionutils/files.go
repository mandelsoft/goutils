package optionutils

import (
	"encoding/base64"
	"io/fs"
	"strings"

	"github.com/mandelsoft/goutils/errors"
	"github.com/mandelsoft/goutils/general"
	"github.com/mandelsoft/goutils/ioutils"
	"github.com/mandelsoft/goutils/ioutils/vfs"
)

// ResolveData maps a data specification (typically given via command line)
// to data.
// It uses various prefixes to specify different content sources.
//   - = direct data passed as string
//   - ! base64 encoded binary data
//   - @ taken from a file
//
// The default (no such prefix given) is using content as direct data.
func ResolveData(in string, fss ...fs.FS) ([]byte, error) {
	return handlePrefix(func(in string, fs fs.FS) ([]byte, error) { return []byte(in), nil }, in, fss...)
}

// ReadFile maps a data specification (typically given via command line)
// to data like ResolveData, but uses reading from a file as default.
func ReadFile(in string, fss ...fs.FS) ([]byte, error) {
	return handlePrefix(readFile, in, fss...)
}

func handlePrefix(def func(string, fs.FS) ([]byte, error), in string, fss ...fs.FS) ([]byte, error) {
	if strings.HasPrefix(in, "=") {
		return []byte(in[1:]), nil
	}
	if strings.HasPrefix(in, "!") {
		return base64.StdEncoding.DecodeString(in[1:])
	}

	_fs := general.OptionalDefaulted[fs.FS](vfs.OSFS, fss...)
	if strings.HasPrefix(in, "@") {
		return readFile(in[1:], _fs)
	}
	return def(in, _fs)
}

func readFile(path string, _fs fs.FS) ([]byte, error) {
	path, err := ioutils.ResolvePath(path)
	if err != nil {
		return nil, err
	}
	data, err := fs.ReadFile(_fs, path)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read file %q", path)
	}
	return data, nil
}
