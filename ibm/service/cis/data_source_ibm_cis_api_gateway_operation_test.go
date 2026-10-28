// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCisApiGatewayOperationDataSource_Basic(t *testing.T) {
	name := "data.ibm_cis_api_gateway_operation.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisApiGatewayOperationDataSourceBasic("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(name, "id"),
					resource.TestCheckResourceAttrSet(name, "endpoint"),
					resource.TestCheckResourceAttrSet(name, "method"),
					resource.TestCheckResourceAttrSet(name, "host"),
				),
			},
		},
	})
}

func testAccCheckCisApiGatewayOperationDataSourceBasic(id, cisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_api_gateway_operation" "op" {
		cis_id    = data.ibm_cis.cis.id
		domain_id = data.ibm_cis_domain.cis_domain.domain_id
		endpoint  = "/api/v1/chat"
		host      = "%[2]s"
		method    = "POST"
	}

	data "ibm_cis_api_gateway_operation" "%[1]s" {
		cis_id       = data.ibm_cis.cis.id
		domain_id    = data.ibm_cis_domain.cis_domain.domain_id
		operation_id = ibm_cis_api_gateway_operation.op.operation_id
	}
`, id, cisDomainStatic)
}
