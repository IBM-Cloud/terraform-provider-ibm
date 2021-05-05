// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceTagDataSource_basic(t *testing.T) {
	name := fmt.Sprintf("tf-satellitelocation-%d", acctest.RandIntRange(10, 100))
	managed_from := "wdc04"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckResourceTagReadDataSource(name, managed_from),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_resource_tag.read_tag", "tags.#", "2"),
				),
			},
		},
	})
}

func testAccCheckResourceTagReadDataSource(name, managed_from string) string {
	return fmt.Sprintf(`

	resource "ibm_satellite_location" "location" {
		location      = "%s"
		managed_from  = "%s"
		description	  = "satellite service"	
		zones		  = ["us-east-1", "us-east-2", "us-east-3"]
	}

	data "ibm_satellite_location" "get_location" {
		location   = ibm_satellite_location.location.location
	}

	resource "ibm_resource_tag" "tag" {
		resource_id = data.ibm_satellite_location.get_location.crn
		tags        =  ["env:dev", "cpu:4"]
	}

	data "ibm_resource_tag" "read_tag" {
		resource_id = ibm_resource_tag.tag.resource_id
	}
`, name, managed_from)
}
