// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCrNamespacesDataSourceBasic(t *testing.T) {
	namespaceName := fmt.Sprintf("terraform-tf-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
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
	return testAccCheckIBMCrNamespaceConfigBasic(namespaceName) + fmt.Sprintf(`
	data "ibm_cr_namespaces" "namespaces" {}
`)
}
