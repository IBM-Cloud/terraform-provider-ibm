package icdv4_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIcdv4(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Icdv4 Suite")
}
