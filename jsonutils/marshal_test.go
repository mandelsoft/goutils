package jsonutils_test

import (
	"encoding/json"

	. "github.com/mandelsoft/goutils/testutils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mandelsoft/goutils/generics"
	"github.com/mandelsoft/goutils/jsonutils"
)

type F struct {
	Field string `json:"field,omitempty"`
}

type G struct {
	Field string `json:"data,omitempty"`
}

type G1 = G

type T struct {
	G           `json:"explicit"`
	G1          `json:",inline"`
	InlineField F `json:",inline"`
	Flat        *string
}

func (t T) MarshalJSON() ([]byte, error) {
	return jsonutils.MarshalStruct(&t)
}

func (t *T) UnmarshalJSON(data []byte) error {
	return jsonutils.UnmarshalStruct(data, t)
}

////////////////////////////////////////////////////////////////////////////////

type E struct {
	G           `json:"explicit,omitempty"`
	G1          `json:",inline"`
	InlineField F       `json:",inline"`
	Flat        *string `json:",omitempty"`
}

func (t E) MarshalJSON() ([]byte, error) {
	return jsonutils.MarshalStruct(&t)
}

func (t *E) UnmarshalJSON(data []byte) error {
	return jsonutils.UnmarshalStruct(data, t)
}

////////////////////////////////////////////////////////////////////////////////

type O struct {
	G           `json:"explicit,omitempty"`
	G1          `json:",inline,omitempty"`
	InlineField F       `json:",inline,omitempty"`
	Flat        *string `json:",omitempty"`
}

func (t O) MarshalJSON() ([]byte, error) {
	return jsonutils.MarshalStruct(&t)
}

func (t *O) UnmarshalJSON(data []byte) error {
	return jsonutils.UnmarshalStruct(data, t)
}

////////////////////////////////////////////////////////////////////////////////

type P struct {
	*G          `json:"explicit,omitempty"`
	*G1         `json:",omitempty,inline"`
	InlineField F       `json:",inline,omitempty"`
	Flat        *string `json:",omitempty"`
}

func (t P) MarshalJSON() ([]byte, error) {
	return jsonutils.MarshalStruct(&t)
}

func (t *P) UnmarshalJSON(data []byte) error {
	return jsonutils.UnmarshalStruct(data, t)
}

////////////////////////////////////////////////////////////////////////////////

var _ = Describe("Test Environment", func() {
	Context("non-empty", func() {
		DATA := `{"Flat":"pointer","data":"text","explicit":{"data":"more"},"field":"demo"}`
		sDATA := &T{
			InlineField: F{
				Field: "demo",
			},
			G: G{
				Field: "more",
			},
			G1: G1{
				Field: "text",
			},
			Flat: generics.PointerTo("pointer"),
		}

		EMPTY := `{"Flat":null,"explicit":{}}`
		sEMPTY := &T{}

		Context("marshal", func() {
			It("marshals inline", func() {
				data := Must(json.Marshal(sDATA))
				Expect(string(data)).To(Equal(DATA))
			})

			It("marshals empty", func() {
				data := Must(json.Marshal(sEMPTY))
				Expect(string(data)).To(Equal(EMPTY))
			})
		})

		Context("unmarshal", func() {
			It("unmarshals struct", func() {
				var s T
				MustBeSuccessful(json.Unmarshal([]byte(DATA), &s))
				Expect(&s).To(DeepEqual(sDATA))
			})

			It("unmarshals empty", func() {
				var s T
				MustBeSuccessful(json.Unmarshal([]byte(EMPTY), &s))
				Expect(&s).To(DeepEqual(sEMPTY))
			})
		})
	})

	////////////////////////////////////////////////////////////////////////////

	Context("pointer", func() {
		DATA := `{"Flat":"pointer","data":"text","explicit":{"data":"more"},"field":"demo"}`
		sDATA := &P{
			InlineField: F{
				Field: "demo",
			},
			G: &G{
				Field: "more",
			},
			G1: &G1{
				Field: "text",
			},
			Flat: generics.PointerTo("pointer"),
		}

		EMPTY := `{"Flat":null,"explicit":{}}`
		sEMPTY := &T{}

		Context("marshal", func() {
			It("marshals inline", func() {
				data := Must(json.Marshal(sDATA))
				Expect(string(data)).To(Equal(DATA))
			})

			It("marshals empty", func() {
				data := Must(json.Marshal(sEMPTY))
				Expect(string(data)).To(Equal(EMPTY))
			})
		})

		Context("unmarshal", func() {
			It("unmarshals struct", func() {
				var s P
				MustBeSuccessful(json.Unmarshal([]byte(DATA), &s))
				Expect(&s).To(DeepEqual(sDATA))
			})

			It("unmarshals empty", func() {
				var s T
				MustBeSuccessful(json.Unmarshal([]byte(EMPTY), &s))
				Expect(&s).To(DeepEqual(sEMPTY))
			})
		})
	})

	////////////////////////////////////////////////////////////////////////////

	Context("empty", func() {
		EMPTY := `{}`
		sEMPTY := &E{}

		Context("marshal", func() {
			It("marshals empty", func() {
				data := Must(json.Marshal(sEMPTY))
				Expect(string(data)).To(Equal(EMPTY))
			})
		})

		Context("unmarshal", func() {
			It("unmarshals empty", func() {
				var s E
				MustBeSuccessful(json.Unmarshal([]byte(EMPTY), &s))
				Expect(&s).To(DeepEqual(sEMPTY))
			})
		})
	})

	////////////////////////////////////////////////////////////////////////////

	Context("empty pointer", func() {
		EMPTY := `{}`
		sEMPTY := &P{}

		Context("marshal", func() {
			It("marshals empty", func() {
				data := Must(json.Marshal(sEMPTY))
				Expect(string(data)).To(Equal(EMPTY))
			})
		})

		Context("unmarshal", func() {
			It("unmarshals empty", func() {
				var s P
				MustBeSuccessful(json.Unmarshal([]byte(EMPTY), &s))
				Expect(&s).To(DeepEqual(sEMPTY))
			})
		})
	})

	////////////////////////////////////////////////////////////////////////////

	Context("empty2", func() {
		EMPTY := `{}`
		sEMPTY := &O{}

		Context("marshal", func() {
			It("marshals empty", func() {
				data := Must(json.Marshal(sEMPTY))
				Expect(string(data)).To(Equal(EMPTY))
			})
		})

		Context("unmarshal", func() {
			It("unmarshals empty", func() {
				var s O
				MustBeSuccessful(json.Unmarshal([]byte(EMPTY), &s))
				Expect(&s).To(DeepEqual(sEMPTY))
			})
		})
	})

	Context("standard", func() {
		sX := &X{
			F{"f1"},
			&F{"f2}"},
		}
		s := `{"field1":{"field":"f1"},"field2":{"field":"f2}"}}`

		Context("marshal", func() {
			It("marshals empty", func() {
				data := Must(json.Marshal(sX))
				Expect(string(data)).To(Equal(s))
			})
		})

		Context("unmarshal", func() {
			It("unmarshals", func() {
				var x X
				MustBeSuccessful(json.Unmarshal([]byte(s), &x))
				Expect(&x).To(DeepEqual(sX))
			})
		})
	})
})

type X struct {
	Field1 F  `json:"field1"`
	Field2 *F `json:"field2"`
}

func (t X) MarshalJSON() ([]byte, error) {
	return jsonutils.MarshalStruct(&t)
}

func (t *X) UnmarshalJSON(data []byte) error {
	return jsonutils.UnmarshalStruct(data, t)
}
