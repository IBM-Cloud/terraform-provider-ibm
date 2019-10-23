package containerv2_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestK8sclusterv2(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "K8sclusterv2 Suite")
}
