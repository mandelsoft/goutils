package testutils_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	me "github.com/mandelsoft/goutils/testutils"
)

var _ = Describe("normalization", func() {

	It("compares with substitution variables", func() {
		exp := "A ${TEST}."
		res := "A testcase."
		vars := me.Substitutions{
			"TEST": "testcase",
		}
		Expect(res).To(me.StringEqualTrimmedWithContext(exp, me.Substitutions{}, vars))
		Expect(res).To(me.StringEqualTrimmedWithContext(exp, vars, me.Substitutions{}))
	})
})
