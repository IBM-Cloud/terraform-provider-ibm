// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCisRateLimitDataSource_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisRateLimitDataSourceConfigBasic1(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cis_rate_limit.ratelimit", "cis_id"),
				),
			},
		},
	})
}

func testAccCheckIBMCisRateLimitDataSourceConfigBasic1() string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + `
	data "ibm_cis_rate_limit" "ratelimit" {
		cis_id    = data.ibm_cis.cis.id
		domain_id = data.ibm_cis_domain.cis_domain.id
	}
	`
}
