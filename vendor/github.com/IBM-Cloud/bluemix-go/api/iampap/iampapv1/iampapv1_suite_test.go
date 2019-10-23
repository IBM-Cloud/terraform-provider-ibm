package iampapv1_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestIampapv1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Iampapv1 Suite")
}
