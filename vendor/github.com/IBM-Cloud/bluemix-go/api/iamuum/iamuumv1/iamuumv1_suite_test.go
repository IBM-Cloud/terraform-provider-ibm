package iamuumv1_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIamuumv1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Iamuumv1 Suite")
}
