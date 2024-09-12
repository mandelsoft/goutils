package jsonutils

import (
	"encoding/json"

	"github.com/mandelsoft/goutils/errors"
)

////////////////////////////////////////////////////////////////////////////////

type MapKeyUnmarshaler interface {
	UnmarshalMapKey(string) error
}

type parseablePointer[P any] interface {
	*P
	MapKeyUnmarshaler
}

type MapKeyMarshaler interface {
	MarshalMapKey() (string, error)
}

type keyconstraint interface {
	comparable
	MapKeyMarshaler
}

type MarshalableMap[K keyconstraint, V any, P parseablePointer[K]] map[K]V

type exampleKey string

func (k exampleKey) MarshalMapKey() (string, error) {
	return string(k), nil
}
func (k *exampleKey) UnmarshalMapKey(s string) error {
	*k = exampleKey(s)
	return nil
}

var (
	_ json.Marshaler   = MarshalableMap[exampleKey, int, *exampleKey](nil)
	_ json.Unmarshaler = (*MarshalableMap[exampleKey, int, *exampleKey])(nil)
)

func (r MarshalableMap[K, V, P]) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	for k, v := range r {
		s, err := k.MarshalMapKey()
		if err != nil {
			return nil, errors.Wrapf(err, "map key %#v", k)
		}
		data, err := json.Marshal(v)
		if err != nil {
			return nil, errors.Wrapf(err, "map entry %q", s)
		}
		m[s] = data
	}
	return json.Marshal(m)
}

func (r *MarshalableMap[K, V, P]) UnmarshalJSON(bytes []byte) error {
	var m map[string]json.RawMessage

	err := json.Unmarshal(bytes, &m)
	if err != nil {
		return err
	}

	*r = MarshalableMap[K, V, P]{}
	for k, v := range m {
		var s V
		var e K

		err := P(&e).UnmarshalMapKey(k)
		if err != nil {
			return errors.Wrapf(err, "map key %q", k)
		}

		err = json.Unmarshal(v, &s)
		if err != nil {
			return errors.Wrapf(err, "map entry %q", k)
		}
		(*r)[e] = s
	}
	return nil
}
