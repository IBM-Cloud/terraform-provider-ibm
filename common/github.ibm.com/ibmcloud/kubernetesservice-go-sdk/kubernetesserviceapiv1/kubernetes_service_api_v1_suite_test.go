// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetesserviceapiv1

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestKubernetesServiceApiV1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "KubernetesServiceApiV1 Suite")
}
