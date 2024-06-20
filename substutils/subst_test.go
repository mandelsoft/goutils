package substutils_test

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mandelsoft/goutils/substutils"
)

type MapData struct {
	Alice string `json:"alice"`
	Bob   string `json:"bob"`
}

func CheckSubst(s Substitution, variable string, result string) {
	r, ok := s.Substitute(variable)
	ExpectWithOffset(1, ok).To(Equal(result != ""))
	ExpectWithOffset(1, r).To(Equal(result))

}

var _ = Describe("Test Environment", func() {
	Context("flavors", func() {
		It("string map", func() {
			s := SubstitutionMap{
				"alice": "25",
				"bob":   "26",
			}
			CheckSubst(s, "alice", "25")
			CheckSubst(s, "marie", "")
		})

		It("subst list", func() {
			s := SubstList("alice", "25", "bob", "26")
			CheckSubst(s, "alice", "25")
			CheckSubst(s, "marie", "")
		})

		It("subst object", func() {
			s := SubstFrom(&MapData{"25", "26"})
			CheckSubst(s, "alice", "25")
			CheckSubst(s, "marie", "")
		})

		It("env", func() {
			r, ok := EnvSubstitution().Substitute("PATH")
			Expect(ok).To(BeTrue())
			Expect(r).NotTo(BeEmpty())
		})

		It("env with prefix", func() {
			r, ok := EnvSubstitution("pre", "fix_").Substitute("prefix_PATH")
			Expect(ok).To(BeTrue())
			Expect(r).NotTo(BeEmpty())
		})

		It("missing env", func() {
			os.Unsetenv("LABER")
			r, ok := EnvSubstitution().Substitute("LABEL")
			Expect(ok).To(BeFalse())
			Expect(r).To(BeEmpty())
		})
	})

	Context("generic", func() {
		It("join", func() {
			s1 := SubstList("alice", "25", "bob", "26")
			s2 := SubstList("tom", "27", "bob", "100")

			s := Join(s1, s2)
			CheckSubst(s, "alice", "25")
			CheckSubst(s, "bob", "26")
			CheckSubst(s, "tom", "27")
			CheckSubst(s, "marie", "")
		})

		It("override", func() {
			s1 := SubstList("alice", "25", "bob", "26")
			s2 := SubstList("tom", "27", "bob", "100")

			s := Override(s1, s2)
			CheckSubst(s, "alice", "25")
			CheckSubst(s, "bob", "100")
			CheckSubst(s, "tom", "27")
			CheckSubst(s, "marie", "")
		})

		It("merge", func() {
			s1 := SubstList("alice", "25", "bob", "26")
			s2 := SubstList("tom", "27", "bob", "100")

			s := MergeMapSubstitution(s1, s2)
			CheckSubst(s, "alice", "25")
			CheckSubst(s, "bob", "100")
			CheckSubst(s, "tom", "27")
			CheckSubst(s, "marie", "")
		})
	})
})
