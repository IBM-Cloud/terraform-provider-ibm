// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCisLogpushJobsDataSource_Basic(t *testing.T) {
	Name := "data.ibm_cis_logpush_jobs.test"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisLogpushJobsDataSource_basic("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(Name, "id"),
				),
			},
		},
	})
}
func testAccCheckCisLogpushJobsDataSource_basic(id, CisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	data "ibm_cis_logpush_jobs" "%[1]s" {
		cis_id          = data.ibm_cis.cis.id
		domain_id       = data.ibm_cis_domain.cis_domain.domain_id
	  }
`, id, acc.CisDomainStatic)
}
