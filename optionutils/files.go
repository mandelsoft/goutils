package optionutils

import (
	"encoding/base64"
	"strings"

	"github.com/mandelsoft/goutils/errors"
	"github.com/mandelsoft/goutils/general"
	"github.com/mandelsoft/goutils/ioutils"
	"github.com/mandelsoft/vfs/pkg/osfs"
	"github.com/mandelsoft/vfs/pkg/vfs"
)

// ResolveData maps a data specification (typically given via command line)
// to data.
// It uses various prefixes to specify different content sources.
//   - = direct data passed as string
//   - ! base64 encoded binary data
//   - @ taken from a file
//
// The default (no such prefix given) is using content as direct data.
func ResolveData(in string, fss ...vfs.FileSystem) ([]byte, error) {
	return handlePrefix(func(in string, fs vfs.FileSystem) ([]byte, error) { return []byte(in), nil }, in, fss...)
}

// ReadFile maps a data specification (typically given via command line)
// to data like ResolveData, but uses reading from a file as default.
func ReadFile(in string, fss ...vfs.FileSystem) ([]byte, error) {
	return handlePrefix(readFile, in, fss...)
}

func handlePrefix(def func(string, vfs.FileSystem) ([]byte, error), in string, fss ...vfs.FileSystem) ([]byte, error) {
	if strings.HasPrefix(in, "=") {
		return []byte(in[1:]), nil
	}
	if strings.HasPrefix(in, "!") {
		return base64.StdEncoding.DecodeString(in[1:])
	}
	if strings.HasPrefix(in, "@") {
		return readFile(in[1:], FileSystem(fss...))
	}
	return def(in, FileSystem(fss...))
}

func readFile(path string, fs vfs.FileSystem) ([]byte, error) {
	path, err := ioutils.ResolvePath(path)
	if err != nil {
		return nil, err
	}
	data, err := vfs.ReadFile(fs, path)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read file %q", path)
	}
	return data, nil
}

var _osfs = osfs.New()

func FileSystem(fss ...vfs.FileSystem) vfs.FileSystem {
	return DefaultedFileSystem(_osfs, fss...)
}

func DefaultedFileSystem(def vfs.FileSystem, fss ...vfs.FileSystem) vfs.FileSystem {
	return general.OptionalDefaulted(def, fss...)
}
