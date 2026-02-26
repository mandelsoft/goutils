package transformer_test

import (
	"fmt"
	"github.com/mandelsoft/goutils/transformer"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type Element interface {
	GetName() string
	SetName(string)
}

func ElementWorker(p Element) string {
	return "work on element " + p.GetName()
}

type Person struct {
	Name string
}

func (p *Person) GetName() string {
	return p.Name
}
func (p *Person) SetName(n string) {
	p.Name = n
}

type Worker[T Element] = func(T) string

func PersonWorker(p *Person) string {
	return "work with person " + p.GetName()
}

func Convert[T Element](f Worker[T]) Worker[Element] {
	return func(element Element) string {
		element.SetName(element.GetName() + "(converted)")
		return f(any(element).(T))
	}
}

func Transform[T Element](in Worker[T]) Worker[Element] {
	return transformer.Optimized[Worker[T], Worker[Element]](Convert[T])(in)
}

////////////////////////////////////////////////////////////////////////////////

func StringConverter(s any) string {
	return fmt.Sprintf("converted(%v)", s)
}

func TransformToString[T any](in T, conv transformer.Transformer[T, string]) string {
	return transformer.OptimizedTransform[T, string](in, conv)
}

////////////////////////////////////////////////////////////////////////////////

var _ = Describe("Mapper Test Environment", func() {
	var p *Person

	BeforeEach(func() {
		p = &Person{"alice"}
	})

	Context("simple converter", func() {
		It("transforms any object", func() {
			Expect(TransformToString(5, StringConverter)).To(Equal("converted(5)"))
		})

		It("skip transformsation forstrings", func() {
			Expect(StringConverter("test")).To(Equal("converted(test)"))
			Expect(TransformToString("test", StringConverter)).To(Equal("test"))
		})
	})

	Context("not optimized", func() {
		It("not optimizable", func() {
			f := Convert(PersonWorker)
			Expect(f(p)).To(Equal("work with person alice(converted)"))
		})

		It("not optimized, but optimizable", func() {
			f := Convert(ElementWorker)
			Expect(f(p)).To(Equal("work on element alice(converted)"))
		})
	})

	Context("optimized", func() {
		It("not optimizable", func() {
			f := transformer.Optimized[Worker[*Person], Worker[Element]](Convert[*Person])(PersonWorker)
			Expect(f(p)).To(Equal("work with person alice(converted)"))
		})
		It("not optimizable", func() {
			f := Transform(PersonWorker)
			Expect(f(p)).To(Equal("work with person alice(converted)"))
		})

		It("optimized", func() {
			f := transformer.Optimized[Worker[Element], Worker[Element]](Convert[Element])(ElementWorker)
			Expect(f(p)).To(Equal("work on element alice"))
		})
		It("optimized", func() {
			f := Transform(ElementWorker)
			Expect(f(p)).To(Equal("work on element alice"))
		})
	})
})
