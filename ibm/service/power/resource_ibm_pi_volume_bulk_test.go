package power_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMPIVolumeBulkbasic(t *testing.T) {
	name := fmt.Sprintf("tf-pi-volume-%d", acctest.RandIntRange(10, 100))
	volumeRes := "ibm_pi_volume_bulk.power_volume"
	userTagsString := `["env:dev","test_tag"]`
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIVolumeBulkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIVolumeBulkConfig(name, userTagsString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIVolumeBulkExists(volumeRes),
					resource.TestCheckResourceAttr(volumeRes, "pi_count", "5"),
					resource.TestCheckResourceAttr(volumeRes, "pi_volume_name", name),
					resource.TestCheckResourceAttr(volumeRes, "pi_user_tags.#", "2"),
					resource.TestCheckTypeSetElemAttr(volumeRes, "pi_user_tags.*", "env:dev"),
					resource.TestCheckTypeSetElemAttr(volumeRes, "pi_user_tags.*", "test_tag"),
				),
			},
		},
	})
}

func testAccCheckIBMPIVolumeBulkConfig(name string, userTagsString string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_volume_bulk" "power_volume" {
			pi_cloud_instance_id    = "%[2]s"
			pi_count                = 5
			pi_user_tags            = %[3]s
			pi_volume_name          = "%[1]s"
			pi_volume_shareable     = true
			pi_volume_size          = 1
			pi_volume_type          = "tier3"
		}`, name, acc.Pi_cloud_instance_id, userTagsString)
}

func testAccCheckIBMPIVolumeBulkExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
		if err != nil {
			return err
		}

		idArr, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		cloudInstanceID := idArr[0]
		for _, volumeID := range idArr[1:] {
			client := instance.NewIBMPIVolumeClient(context.Background(), sess, cloudInstanceID)
			_, err := client.Get(volumeID)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func testAccCheckIBMPIVolumeBulkDestroy(s *terraform.State) error {
	sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_volume_bulk" {
			continue
		}

		idArr, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		cloudInstanceID := idArr[0]
		for _, volumeID := range idArr[1:] {
			client := instance.NewIBMPIVolumeClient(context.Background(), sess, cloudInstanceID)
			volume, err := client.Get(volumeID)
			if err == nil {
				log.Println("volume*****", volume.State)
				return fmt.Errorf("PI Volume still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}
