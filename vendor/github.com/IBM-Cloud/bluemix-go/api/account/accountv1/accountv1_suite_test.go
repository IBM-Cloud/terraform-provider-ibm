package accountv1_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestAccountv1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Accountv1 Suite")
}
