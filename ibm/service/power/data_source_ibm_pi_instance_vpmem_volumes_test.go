// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMPIInstanceVpmemVolumesDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIInstanceVpmemVolumesDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_instance_vpmem_volumes.instance_vpmem_volumes_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pi_instance_vpmem_volumes.instance_vpmem_volumes_instance", "volumes.#"),
				),
			},
		},
	})
}

func testAccCheckIBMPIInstanceVpmemVolumesDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pi_instance_vpmem_volumes" "instance_vpmem_volumes_instance" {
			pi_cloud_instance_id = "%[1]s"
			pi_pvm_instance_id   = "%[2]s"
		}`, acc.Pi_cloud_instance_id, acc.Pi_instance_id)
}
