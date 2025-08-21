package vfs

import (
	"io/fs"
	"os"
)

var OSFS FileFS[File] = osfs{}

type osfs struct{}

func (o osfs) OpenFile(name string, flag int, perm fs.FileMode) (File, error) {
	return os.OpenFile(name, flag, perm)
}

func (o osfs) Lstat(name string) (fs.FileInfo, error) {
	return os.Lstat(name)
}

func (o osfs) Chmod(name string, mode fs.FileMode) error {
	return os.Chmod(name, mode)
}

func (o osfs) MkdirAll(name string, mode fs.FileMode) error {
	return os.MkdirAll(name, mode)
}

func (o osfs) ReadDir(name string) ([]fs.FileInfo, error) {
	d, err := os.ReadDir(name)
	if err != nil {
		return nil, err
	}

	r := make([]fs.FileInfo, len(d))
	for i, e := range d {
		r[i], err = e.Info()
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}

func (o osfs) Readlink(name string) (string, error) {
	// TODO implement me
	panic("implement me")
}

func (o osfs) Symlink(oldname string, newname string) error {
	// TODO implement me
	panic("implement me")
}

func (o osfs) Open(name string) (fs.File, error) {
	return os.Open(name)
}

func (o osfs) Stat(name string) (fs.FileInfo, error) {
	return os.Stat(name)
}
