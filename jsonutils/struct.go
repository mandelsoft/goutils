package jsonutils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"slices"
	"strings"

	"github.com/modern-go/reflect2"
)

func MarshalStruct(s any) ([]byte, error) {
	v := reflect.ValueOf(s)
	for v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected struct")
	}

	values := map[string]interface{}{}

	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		ft := t.Field(i)
		if !ft.IsExported() {
			continue
		}

		omit := false
		inline := false
		name := ft.Name
		tag, ok := ft.Tag.Lookup("json")
		if ok {
			keys := strings.Split(tag, ",")
			if keys[0] != "" {
				name = keys[0]
			}
			keys = keys[1:]
			inline = slices.Contains(keys, "inline")
			omit = slices.Contains(keys, "omitempty")
		}

		f := v.Field(i)
		value := f.Interface()
		if inline {
			data, err := json.Marshal(value)
			if err != nil {
				return nil, err
			}
			value = nil
			err = json.Unmarshal(data, &value)
			if err != nil {
				return nil, err
			}
			if reflect2.IsNil(value) {
				continue
			}
			if m, ok := value.(map[string]interface{}); ok {
				for k, v := range m {
					if _, ok := values[k]; ok {
						return nil, fmt.Errorf("duplicate field %q", k)
					}
					values[k] = v
				}
				continue
			}
			return nil, fmt.Errorf("inline field must be struct")
		}

		if omit {
			v := reflect.ValueOf(value)
			if v.IsZero() {
				continue
			}
		}
		values[name] = value
	}
	return json.Marshal(values)
}

type structMap map[string]json.RawMessage

func UnmarshalStruct(data []byte, s any) error {
	v := reflect.ValueOf(s)

	if v.Kind() != reflect.Pointer || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("expected struct pointer")
	}

	var values structMap
	err := json.Unmarshal(data, &values)
	if err != nil {
		return err
	}

	v = v.Elem()
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		ft := t.Field(i)
		if !ft.IsExported() {
			continue
		}

		name := ft.Name
		inline := false
		tag, ok := ft.Tag.Lookup("json")
		if ok {
			keys := strings.Split(tag, ",")
			if keys[0] != "" {
				name = keys[0]
			}
			keys = keys[1:]
			inline = slices.Contains(keys, "inline")
		}

		rt := ft.Type
		cnt := 0
		for rt.Kind() == reflect.Pointer {
			rt = rt.Elem()
			cnt++
		}
		in := data
		if inline {
			if rt.Kind() != reflect.Struct {
				return fmt.Errorf("expected struct for inline type")
			}
		} else {
			in = values[name]
		}
		if len(in) <= 2 || string(in) == "null" {
			continue
		}
		e := reflect.New(rt)
		err := json.Unmarshal(in, e.Interface())
		if err != nil {
			return err
		}

		switch cnt {
		case 0:
			e = e.Elem()
		case 1:
		default:
			for i := 1; i < cnt; i++ {
				n := reflect.New(reflect.PointerTo(e.Type()))
				e.Set(e)
				e = n
			}
		}
		f := v.Field(i)
		f.Set(e)
	}
	return nil
}
