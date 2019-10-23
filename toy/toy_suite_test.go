package toy

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestToy(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Toy Suite")
}
