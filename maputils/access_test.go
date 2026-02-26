package maputils_test

import (
	. "github.com/mandelsoft/goutils/maputils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type L1 map[bool]string
type L2 map[int]L1
type L3 map[string]L2

var _ = Describe("nested access", func() {
	var m = L3{
		"a": {1: {true: "true"}},
	}

	Context("Get", func() {
		It("get leaf", func() {
			Expect(AndGet(AndGet(AndGet(m, "a"), 1), true)).To(Equal("true"))
		})
		It("missed leaf", func() {
			Expect(AndGet(AndGet(AndGet(m, "a"), 1), false)).To(Equal(""))
			Expect(AndGet(AndGet(AndGet(m, "a"), 2), false)).To(Equal(""))
			Expect(AndGet(AndGet(AndGet(m, "b"), 2), false)).To(Equal(""))
		})
	})

	Context("Set", func() {
		It("set leaf", func() {
			var m L1
			m = Set(m)(true, Value("true"))
			Expect(AndGet(m, true)).To(Equal("true"))
		})

		It("set leaf", func() {
			var m L2
			// bad type inference
			m = Set(m)(5, AndSet[L1](true, Value("true")))
			Expect(AndGet(AndGet(m, 5), true)).To(Equal("true"))
		})

		It("set leaf", func() {
			var m L3
			// bad type inference
			m = Set(m)("a", AndSet[L2](5, AndSet[L1](true, Value("true"))))
			m = AndSet[L3]("a", AndSet[L2](6, AndSet[L1](false, Value("false"))))(m)
			Expect(AndGet(AndGet(AndGet(m, "a"), 5), true)).To(Equal("true"))
			Expect(AndGet(AndGet(AndGet(m, "a"), 6), false)).To(Equal("false"))
		})
	})
})
