package vfs

import (
	"io"
	"io/fs"
)

////////////////////////////////////////////////////////////////////////////////

type FileFS[F File] interface {
	fs.StatFS
	OpenFile(name string, flag int, perm fs.FileMode) (F, error)
	Lstat(name string) (fs.FileInfo, error)
	Chmod(name string, mode fs.FileMode) error
	MkdirAll(name string, mode fs.FileMode) error
	ReadDir(name string) ([]fs.FileInfo, error)
	Readlink(name string) (string, error)
	Symlink(oldname string, newname string) error
}

type File interface {
	fs.File
	io.WriteCloser
}
