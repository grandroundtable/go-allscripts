package testingtsupport_test

import (
	. "github.com/grandroundtable/go-allscripts/Godeps/_workspace/src/github.com/onsi/gomega"

	"testing"
)

func TestTestingT(t *testing.T) {
	RegisterTestingT(t)
	Ω(true).Should(BeTrue())
}
