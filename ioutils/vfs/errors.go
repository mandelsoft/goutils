package vfs

import (
	"errors"
	"os"
	"reflect"
)

type ErrorMatcher func(err error) bool

func MatchErr(err error, match ErrorMatcher, base error) bool {
	for err != nil {
		if base == err || (match != nil && match(err)) {
			return true
		}
		switch nerr := err.(type) {
		case interface{ Unwrap() error }:
			err = nerr.Unwrap()
		default:
			err = nil
			v := reflect.ValueOf(nerr)
			if v.Kind() == reflect.Struct {
				f := v.FieldByName("Err")
				if f.IsValid() {
					err, _ = f.Interface().(error)
				}
			}
		}
	}
	return false
}

func IsErrPermission(err error) bool {
	if os.IsPermission(err) {
		return true
	}
	return MatchErr(err, os.IsPermission, ErrPermission)
}

func IsErrNotDir(err error) bool {
	return MatchErr(err, isUnderlyingErrNotDir, ErrNotDir)
}

func IsErrNotExist(err error) bool {
	if os.IsNotExist(err) {
		return true
	}
	return MatchErr(err, os.IsNotExist, ErrNotExist)
}

func IsErrExist(err error) bool {
	if os.IsExist(err) {
		return true
	}
	return MatchErr(err, os.IsExist, ErrExist)
}

func IsErrReadOnly(err error) bool {
	return MatchErr(err, nil, ErrReadOnly)
}

func NewPathError(op string, path string, err error) error {
	return &os.PathError{Op: op, Path: path, Err: err}
}

var ErrNotDir = errors.New("is no directory")
var ErrNotExist = os.ErrNotExist
var ErrPermission = os.ErrPermission
var ErrExist = os.ErrExist

var ErrReadOnly = errors.New("filehandle is not writable")
var ErrNotEmpty = errors.New("dir not empty")
