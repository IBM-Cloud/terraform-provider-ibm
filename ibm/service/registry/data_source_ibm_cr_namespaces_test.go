// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package registry_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMCrNamespacesDataSourceBasic(t *testing.T) {
	namespaceName := fmt.Sprintf("terraform-tf-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCrNamespacesDataSourceConfig(namespaceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cr_namespaces.namespaces", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMCrNamespacesDataSourceConfig(namespaceName string) string {
	return testAccCheckIBMCrNamespaceConfigBasic(namespaceName) + `
	data "ibm_cr_namespaces" "namespaces" {}
`
}
