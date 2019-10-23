package accountv2_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestAccountv2(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Accountv2 Suite")
}
