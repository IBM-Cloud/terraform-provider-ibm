package management_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestManagement(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Management Suite")
}
