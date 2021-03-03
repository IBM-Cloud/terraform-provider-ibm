// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMContainerAddOnsDataSource_basic(t *testing.T) {
	name := fmt.Sprintf("tf-cluster-addon-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerAddOnsDataSourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_container_addons.addons", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerAddOnsDataSourceConfig(name string) string {
	return testAccCheckIBMContainerAddOnsBasic(name) + fmt.Sprintf(`
	data "ibm_container_addons" "addons" {
	    cluster= ibm_container_addons.addons.cluster
	}
`)
}
