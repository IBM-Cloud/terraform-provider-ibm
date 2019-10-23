package csev2_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIamv1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Csev2 Suite")
}
