// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

// TEMPORARY TEST FILE TO TEST WITH EXISTING CLUSTER

package ibm //TODO: remove the whole file.

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMContainerALBmanual_Create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMContainerALBCreateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerALBManualCreate(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_container_alb_create.alb", "enable_by_default", "true"),
					resource.TestCheckResourceAttr("ibm_container_alb_create.alb", "type", "private"),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}
func testAccCheckIBMContainerALBManualCreate() string {
	config := fmt.Sprintf(`
resource "ibm_container_alb_create" "alb" {
	enable_by_default = "true"
	type = "private"
	vlan_id = "1900403"
	zone = "dal10"
	cluster="terraform-manual-test"
}`)
	fmt.Println(config)
	return config
}
