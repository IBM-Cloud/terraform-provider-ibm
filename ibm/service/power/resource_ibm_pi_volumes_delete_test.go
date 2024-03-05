// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
)

func TestAccIBMPIVolumesDeleteBasic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIVolumesDeleteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIVolumesDeleteConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_pi_volumes_delete.volumes_delete", "pi_cloud_instance_id", acc.Pi_cloud_instance_id),
				),
			},
		},
	})
}

func testAccCheckIBMPIVolumesDeleteConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_pi_volumes_delete" "volumes_delete" {
			pi_cloud_instance_id = "%s"
			pi_volume_ids = ["c1dc81e9-85f0-4e32-8c58-f22cc96a0037", "61b26eb2-50be-40b4-bda7-3c9b1662a742"]
		}
	`, acc.Pi_cloud_instance_id)
}

func testAccCheckIBMPIVolumesDeleteDestroy(s *terraform.State) error {
	sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_volumes_delete" {
			continue
		}
		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		volumeC := st.NewIBMPIVolumeClient(context.Background(), sess, parts[0])
		for _, volumeID := range parts[1:] {
			volume, err := volumeC.Get(volumeID)
			if err == nil {
				log.Println("volume*****", volume.State)
				return fmt.Errorf("PI Volume still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}
