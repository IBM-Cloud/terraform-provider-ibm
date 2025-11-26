// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMPIVolumeOnboardingbasic(t *testing.T) {
	name := fmt.Sprintf("tf-pi-volume-onboarding-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIVolumeOnboardingConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIVolumeOnboardingExists("ibm_pi_volume_onboarding.power_volume_onboarding"),
					resource.TestCheckResourceAttrSet("ibm_pi_volume_onboarding.power_volume_onboarding", "id"),
					resource.TestCheckResourceAttrSet("ibm_pi_volume_onboarding.power_volume_onboarding", "status"),
				),
			},
		},
	})
}

func testAccCheckIBMPIVolumeOnboardingExists(n string) resource.TestCheckFunc {
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

		ids, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		cloudInstanceID, onboardID := ids[0], ids[1]
		client := instance.NewIBMPIVolumeOnboardingClient(context.Background(), sess, cloudInstanceID)

		_, err = client.Get(onboardID)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckIBMPIVolumeOnboardingConfig(name string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_volume_onboarding" "power_volume_onboarding" {
		pi_cloud_instance_id        = "%[1]s"
		pi_description              = "%[2]s"
		pi_onboarding_volumes {
			pi_source_crn           = "%[3]s"
			pi_auxiliary_volumes {
				pi_auxiliary_volume_name = "%[4]s"
				pi_display_name = "%[2]s-1"
			}
		}
	}`, acc.Pi_cloud_instance_id, name, acc.Pi_volume_onboarding_source_crn, acc.Pi_auxiliary_volume_name)
}
