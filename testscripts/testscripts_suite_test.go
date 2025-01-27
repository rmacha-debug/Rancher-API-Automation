package testscripts

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTestscripts(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testscripts Suite")
}
