package jsonutils_test

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mandelsoft/goutils/jsonutils"
	. "github.com/mandelsoft/goutils/testutils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type ComplexKey struct {
	Name    string
	Surname string
}

var (
	_ jsonutils.MapKeyMarshaler   = ComplexKey{}
	_ jsonutils.MapKeyUnmarshaler = (*ComplexKey)(nil)
)

func (k ComplexKey) String() string {
	return k.Name + ":" + k.Surname
}

func (k ComplexKey) MarshalMapKey() (string, error) {
	return k.String(), nil
}

func (k *ComplexKey) UnmarshalMapKey(s string) error {
	comps := strings.Split(s, ":")
	if len(comps) != 2 {
		return fmt.Errorf("invalid complex key %q", s)
	}
	k.Name = comps[0]
	k.Surname = comps[1]
	return nil
}

type ComplexMap = jsonutils.MarshalableMap[ComplexKey, int, *ComplexKey]

type EmbeddedMap struct {
	ComplexMap
}

/*
func (m EmbeddedMap) MarshalJSON() ([]byte, error) {
	return m.ComplexMap.MarshalJSON()
}

func (m *EmbeddedMap) UnmarshalJSON(data []byte) error {
	return (&m.ComplexMap).UnmarshalJSON(data)
}
*/

func (m EmbeddedMap) Get(k ComplexKey) int {
	return m.ComplexMap[k]
}

var _ = Describe("Marshalable map keys", func() {
	It("marshal/unmarshal map", func() {
		m := ComplexMap{
			ComplexKey{"alice", "jones"}: 25,
			ComplexKey{"bob", "jones"}:   26,
		}

		data := Must(json.Marshal(m))
		Expect(data).To(YAMLEqual(`
alice:jones: 25
bob:jones: 26
`))
		var u ComplexMap
		MustBeSuccessful(json.Unmarshal(data, &u))
		Expect(u).To(DeepEqual(m))
	})

	It("marshal/unmarshal embedded map", func() {
		m := EmbeddedMap{
			ComplexMap{
				ComplexKey{"alice", "jones"}: 25,
				ComplexKey{"bob", "jones"}:   26,
			},
		}

		data := Must(json.Marshal(m))
		Expect(data).To(YAMLEqual(`
alice:jones: 25
bob:jones: 26
`))
		var u EmbeddedMap
		MustBeSuccessful(json.Unmarshal(data, &u))
		Expect(u).To(DeepEqual(m))
	})

})
