package iamv1_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIamv1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Iamv1 Suite")
}
