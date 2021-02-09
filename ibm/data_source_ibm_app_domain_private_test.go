/**
 * © Copyright IBM Corporation 2020. All Rights Reserved.
 *
 * Licensed under the Mozilla Public License, version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at https://mozilla.org/MPL/2.0/
 */

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMAppDomainPrivateDataSource_basic(t *testing.T) {
	name := fmt.Sprintf("terraform%d.com", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
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
		}`, cfOrganization, name)

}
