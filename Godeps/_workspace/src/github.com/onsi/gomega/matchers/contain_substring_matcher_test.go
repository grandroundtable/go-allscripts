package matchers_test

import (
	. "github.com/grandroundtable/go-allscripts/Godeps/_workspace/src/github.com/onsi/gomega"
	. "github.com/grandroundtable/go-allscripts/Godeps/_workspace/src/github.com/onsi/gomega/matchers"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("ContainSubstringMatcher", func() {
	Context("when actual is a string", func() {
		It("should match against the string", func() {
			Ω("Marvelous").Should(ContainSubstring("rve"))
			Ω("Marvelous").ShouldNot(ContainSubstring("boo"))
		})
	})

	Context("when the matcher is called with multiple arguments", func() {
		It("should pass the string and arguments to sprintf", func() {
			Ω("Marvelous3").Should(ContainSubstring("velous%d", 3))
		})
	})

	Context("when actual is a stringer", func() {
		It("should call the stringer and match agains the returned string", func() {
			Ω(&myStringer{a: "Abc3"}).Should(ContainSubstring("bc3"))
		})
	})

	Context("when actual is neither a string nor a stringer", func() {
		It("should error", func() {
			success, err := (&ContainSubstringMatcher{Substr: "2"}).Match(2)
			Ω(success).Should(BeFalse())
			Ω(err).Should(HaveOccurred())
		})
	})
})