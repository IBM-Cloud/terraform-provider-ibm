// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cloudfoundry_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMAppDomainPrivateDataSource_basic(t *testing.T) {
	name := fmt.Sprintf("terraform%d.com", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAppDomainPrivateDataSourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.ibm_app_domain_private.testacc_domain", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMAppDomainPrivateDataSourceConfig(name string) string {
	return fmt.Sprintf(`
	
		data "ibm_org" "orgdata" {
			org    = "%s"
		}

		resource "ibm_app_domain_private" "domain" {
			name = "%s"
			org_guid = data.ibm_org.orgdata.id
		}
	
		data "ibm_app_domain_private" "testacc_domain" {
			name = ibm_app_domain_private.domain.name
		}`, acc.CfOrganization, name)

}
