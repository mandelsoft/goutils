package vfs

import (
	"io"
	"io/fs"
	"os"
	"path"

	"github.com/mandelsoft/goutils/errors"
)

func Exists_(err error) bool {
	return err == nil || !IsErrNotExist(err)
}

// Exists checks if a file or directory exists.
func Exists(_fs fs.StatFS, path string) (bool, error) {
	_, err := _fs.Stat(path)
	if err == nil {
		return true, nil
	}
	if IsErrNotExist(err) {
		return false, nil
	}
	return false, err
}

// DirExists checks if a path exists and is a directory.
func DirExists(_fs fs.StatFS, path string) (bool, error) {
	fi, err := _fs.Stat(path)
	if err == nil && fi.IsDir() {
		return true, nil
	}
	if IsErrNotExist(err) {
		return false, nil
	}
	return false, err
}

// FileExists checks if a path exists and is a regular file.
func FileExists(_fs fs.StatFS, path string) (bool, error) {
	fi, err := _fs.Stat(path)
	if err == nil && fi.Mode()&os.ModeType == 0 {
		return true, nil
	}
	if IsErrNotExist(err) {
		return false, nil
	}
	return false, err
}

// IsDir checks if a given path is a directory.
func IsDir(_fs fs.StatFS, path string) (bool, error) {
	fi, err := _fs.Stat(path)
	if err != nil {
		return false, err
	}
	return fi.IsDir(), nil
}

// IsFile checks if a given path is a file.
func IsFile(_fs fs.StatFS, path string) (bool, error) {
	fi, err := _fs.Stat(path)
	if err != nil {
		return false, err
	}
	return fi.Mode()&os.ModeType == 0, nil
}

func CopyFile[F1 File, F2 File](srcfs FileFS[F1], src string, dstfs FileFS[F2], dst string) error {
	fi, err := srcfs.Lstat(src)
	if err != nil {
		return err
	}
	if !fi.Mode().IsRegular() {
		return errors.New("no regular file")
	}
	s, err := srcfs.Open(src)
	if err != nil {
		return err
	}
	defer s.Close()

	d, err := dstfs.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, fi.Mode()&os.ModePerm)
	if err != nil {
		return err
	}
	defer d.Close()

	_, err = io.Copy(d, s)
	if err != nil {
		return err
	}
	return dstfs.Chmod(dst, fi.Mode())
}

// CopyDir recursively copies a directory tree, attempting to preserve permissions.
// Source directory must exist, destination directory may exist.
// Symlinks are ignored and skipped.
func CopyDir[F1 File, F2 File](srcfs FileFS[F2], src string, dstfs FileFS[F2], dst string) error {
	si, err := srcfs.Stat(src)
	if err != nil {
		return err
	}
	if !si.IsDir() {
		return NewPathError("CopyDir", src, ErrNotDir)
	}

	di, err := dstfs.Stat(dst)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	if err == nil && !di.IsDir() {
		return NewPathError("CopyDir", dst, ErrNotDir)
	}

	err = dstfs.MkdirAll(dst, si.Mode())
	if err != nil {
		return err
	}

	entries, err := srcfs.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := path.Join(src, entry.Name())
		dstPath := path.Join(dst, entry.Name())

		if entry.IsDir() {
			err = CopyDir[F1, F2](srcfs, srcPath, dstfs, dstPath)
		} else {
			// Skip symlinks.
			if entry.Mode()&os.ModeSymlink != 0 {
				var old string
				old, err = srcfs.Readlink(srcPath)
				if err == nil {
					err = dstfs.Symlink(old, dstPath)
				}
				if err == nil {
					err = os.Chmod(dst, entry.Mode())
				}
			} else {
				err = CopyFile(srcfs, srcPath, dstfs, dstPath)
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}
