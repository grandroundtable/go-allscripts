package assertion_test

import (
	. "github.com/grandroundtable/go-allscripts/Godeps/_workspace/src/github.com/onsi/gomega"
	. "github.com/onsi/ginkgo"

	"testing"
)

func TestAssertion(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Assertion Suite")
}
