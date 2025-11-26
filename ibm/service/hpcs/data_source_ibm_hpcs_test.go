// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package hpcs_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMHPCSDatasourceBasic(t *testing.T) {
	instanceName := acc.HpcsInstanceName
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:  testAccCheckIBMHPCSDatasourceConfig(instanceName),
				Destroy: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_hpcs.hpcs", "name", instanceName),
					resource.TestCheckResourceAttr("data.ibm_hpcs.hpcs", "service", "hs-crypto"),
				),
			},
		},
	})
}

func testAccCheckIBMHPCSDatasourceConfig(instanceName string) string {
	return fmt.Sprintf(`
	data "ibm_hpcs" "hpcs" {
		name              = "%s"
	}
	`, instanceName)

}
