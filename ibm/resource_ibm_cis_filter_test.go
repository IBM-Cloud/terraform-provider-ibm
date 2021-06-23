// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCisFilter_Basic(t *testing.T) {
	name := "ibm_cis_filter." + "test"
	filterexp := "(http.request.uri eq \"/test-update?number=5\")"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCis(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisFilter_basic("test", cisDomainStatic, "true", "Filter-creation", filterexp),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "description", "Filter-creation"),
					resource.TestCheckResourceAttr(name, "expression", filterexp),
					resource.TestCheckResourceAttr(name, "paused", "true"),
				),
			},
		},
	})
}

func TestAccIBMCisFilter_Import(t *testing.T) {
	name := "ibm_cis_filter." + "test"
	filterexp := "(http.request.uri eq \"/test-update?number=5\")"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisFilter_basic("test", cisDomainStatic, "true", "Filter-creation", filterexp),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "expression", filterexp),
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
func testAccCheckCisFilter_basic(id, cisDomainStatic, paused, description, expression string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_filter" "%[1]s" {
		cis_id          = data.ibm_cis.cis.id
		domain_id       = data.ibm_cis_domain.cis_domain.domain_id
		paused			= "true"
		description		= "Filter-creation"
		expression		= "(http.request.uri eq \"/test-update?number=5\")"
	  }
`, id, cisDomainStatic, paused, description, expression)
}
