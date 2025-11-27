// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package resourcecontroller_test

import (
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMResourceReclamationsDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMResourceReclamationsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					// Basic check to ensure reclamations attribute is present and is a list
					resource.TestCheckResourceAttrSet("data.ibm_resource_reclamations.all", "reclamations.#"),
				),
			},
		},
	})
}

func TestAccIBMResourceReclamationsDataSource_invalid(t *testing.T) {
	// Since reclamations list does not take filters, you might want to add a negative case e.g. empty provider config
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: `
                data "ibm_resource_reclamationss" "all" {
                  # Intentionally invalid config if any required fields added later, or just test error
                }`,
				ExpectError: regexp.MustCompile(``),
			},
		},
	})
}

func testAccCheckIBMResourceReclamationsDataSourceConfig() string {
	return `
data "ibm_resource_reclamations" "all" {
}
`
}
