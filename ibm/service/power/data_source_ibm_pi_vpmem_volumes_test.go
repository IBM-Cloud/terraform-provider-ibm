// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMPIVpmemVolumesDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIVpmemVolumesDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_vpmem_volumes.vpmem_volumes_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pi_vpmem_volumes.vpmem_volumes_instance", "volumes.#"),
				),
			},
		},
	})
}

func testAccCheckIBMPIVpmemVolumesDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pi_vpmem_volumes" "vpmem_volumes_instance" {
		   pi_cloud_instance_id = "%s"
		}`, acc.Pi_cloud_instance_id)
}
