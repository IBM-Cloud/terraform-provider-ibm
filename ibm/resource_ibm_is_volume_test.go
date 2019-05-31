package ibm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.ibm.com/Bluemix/riaas-go-client/clients/storage"
	"github.ibm.com/riaas/rias-api/riaas/models"
)

func TestAccIBMISVolume_basic(t *testing.T) {
	var vol *models.Volume
	name := fmt.Sprintf("tf_create_step_name_%d", acctest.RandInt())
	name1 := fmt.Sprintf("tf_update_step_name_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISVolumeDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISVolumeConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVolumeExists("ibm_is_volume.storage", &vol),
					resource.TestCheckResourceAttr(
						"ibm_is_volume.storage", "name", name),
				),
			},

			resource.TestStep{
				Config: testAccCheckIBMISVolumeConfig(name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVolumeExists("ibm_is_volume.storage", &vol),
					resource.TestCheckResourceAttr(
						"ibm_is_volume.storage", "name", name1),
				),
			},
		},
	})
}

func testAccCheckIBMISVolumeDestroy(s *terraform.State) error {
	sess, _ := testAccProvider.Meta().(ClientSession).ISSession()

	VOL := storage.NewStorageClient(sess)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_vol" {
			continue
		}

		_, err := VOL.Get(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("Volume still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISVolumeExists(n string, vol **models.Volume) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, _ := testAccProvider.Meta().(ClientSession).ISSession()
		client := storage.NewStorageClient(sess)
		foundVol, err := client.Get(rs.Primary.ID)

		if err != nil {
			return err
		}

		*vol = foundVol
		return nil
	}
}

func testAccCheckIBMISVolumeConfig(name string) string {
	return fmt.Sprintf(
		`resource "ibm_is_volume" "storage"{
    name = "%s"
    profile = "10iops-tier"
    zone = "us-south-3"
    # capacity= 200
}`, name)

}
