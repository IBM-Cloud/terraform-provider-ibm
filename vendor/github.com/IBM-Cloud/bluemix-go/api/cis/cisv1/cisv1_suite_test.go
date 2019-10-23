package cisv1_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCisv1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cisv1 Suite")
}
