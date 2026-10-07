// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCisApiGatewaySchemasDataSource_Basic(t *testing.T) {
	name := "data.ibm_cis_api_gateway_schemas.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisApiGatewaySchemasDataSourceBasic("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(name, "id"),
				),
			},
		},
	})
}

func testAccCheckCisApiGatewaySchemasDataSourceBasic(id, cisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	data "ibm_cis_api_gateway_schemas" "%[1]s" {
		cis_id    = data.ibm_cis.cis.id
		domain_id = data.ibm_cis_domain.cis_domain.domain_id
	}
`, id)
}
