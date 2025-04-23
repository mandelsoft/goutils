package datautils

import (
	"bytes"
	"io"

	"github.com/mandelsoft/goutils/errors"
)

type DataGetter interface {
	// Get returns the content as byte array
	Get() ([]byte, error)
}

type DataReader interface {
	// Reader returns a reader to incrementally access byte stream content
	Reader() (io.ReadCloser, error)
}

type GenericData = interface{}

type GenericDataGetter interface {
	Get() (GenericData, error)
}

const KIND_DATASOURCE = "data source"

// GetData provides data as byte sequence from some generic
// data sources like byte arrays, strings, DataReader and
// DataGetters. This means we can pass all BlobAccess or DataAccess
// objects.
// If no an unknown data source is passes an ErrInvalid(KIND_DATASOURCE)
// is returned.
func GetData(src GenericData) ([]byte, error) {
	switch t := src.(type) {
	case []byte:
		return t, nil
	case string:
		return []byte(t), nil
	case DataGetter:
		return t.Get()
	case DataReader:
		var buf bytes.Buffer
		r, err := t.Reader()
		if err != nil {
			return nil, err
		}
		defer r.Close()
		_, err = io.Copy(&buf, r)
		if err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	}
	return nil, errors.ErrInvalidType(KIND_DATASOURCE, src)
}

// GetGenericData evaluates some input provided by well-known
// types or interfaces and provides some data output
// by mapping the input to either a byte sequence or
// some specialized object.
// If the input type is not known an ErrInvalid(KIND_DATASOURCE)
// // is returned.
// In extension to GetData, it additionally evaluates the interface
// GenericDataGetter to map the input to some evaluated object.
func GetGenericData(src GenericData) (interface{}, error) {
	switch t := src.(type) {
	case GenericDataGetter:
		return t.Get()
	default:
		return GetData(src)
	}
}
