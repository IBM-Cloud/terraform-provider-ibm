package mccpv2_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCfv2(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cfv2 Suite")
}
