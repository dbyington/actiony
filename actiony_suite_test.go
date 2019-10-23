package actiony

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestActiony(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Actiony Suite")
}
