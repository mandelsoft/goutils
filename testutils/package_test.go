package testutils_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	me "github.com/mandelsoft/goutils/testutils"
)

var _ = Describe("package tests", func() {
	It("go module name", func() {
		mod := me.Must(me.GetModuleName())
		Expect(mod).To(Equal("github.com/mandelsoft/goutils"))
	})
})
