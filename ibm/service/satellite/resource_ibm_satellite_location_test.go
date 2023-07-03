// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package satellite_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/IBM-Cloud/container-services-go-sdk/kubernetesserviceapiv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccSatelliteLocation_Basic(t *testing.T) {
	var instance string
	name := fmt.Sprintf("tf-satellitelocation-%d", acctest.RandIntRange(10, 100))
	managed_from := "wdc04"
	coreos_enabled := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckSatelliteLocationDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckSatelliteLocationCreate(name, managed_from, coreos_enabled),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckSatelliteLocationExists("ibm_satellite_location.location", instance),
					resource.TestCheckResourceAttr("ibm_satellite_location.location", "location", name),
					resource.TestCheckResourceAttr("ibm_satellite_location.location", "managed_from", managed_from),
					resource.TestCheckResourceAttr("ibm_satellite_location.location", "coreos_enabled", coreos_enabled),
				),
			},
		},
	})
}

func TestAccSatelliteLocation_Import(t *testing.T) {
	var instance string
	name := fmt.Sprintf("tf_location_%d", acctest.RandIntRange(10, 100))
	managed_from := "wdc04"
	coreos_enabled := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckSatelliteLocationDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckSatelliteLocationCreate(name, managed_from, coreos_enabled),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckSatelliteLocationExists("ibm_satellite_location.location", instance),
					resource.TestCheckResourceAttr("ibm_satellite_location.location", "location", name),
					resource.TestCheckResourceAttr("ibm_satellite_location.location", "managed_from", managed_from),
				),
			},
			{
				ResourceName:      "ibm_satellite_location.location",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckSatelliteLocationExists(n string, instance string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		ID := rs.Primary.ID
		satClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SatelliteClientSession()
		if err != nil {
			return err
		}

		getSatLocOptions := &kubernetesserviceapiv1.GetSatelliteLocationOptions{
			Controller: &ID,
		}
		instance1, resp, err := satClient.GetSatelliteLocation(getSatLocOptions)
		if err != nil {
			if resp != nil && resp.StatusCode == 404 {
				return nil
			}
			return fmt.Errorf("[ERROR] Error retrieving satellite location: %s\n Response code is: %+v", err, resp)
		}

		instance = *instance1.ID

		return nil
	}
}

func testAccCheckSatelliteLocationDestroy(s *terraform.State) error {
	satClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SatelliteClientSession()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_satellite_location" {
			continue
		}

		ID := rs.Primary.ID
		getSatLocOptions := &kubernetesserviceapiv1.GetSatelliteLocationOptions{
			Controller: &ID,
		}

		_, _, err = satClient.GetSatelliteLocation(getSatLocOptions)
		if err == nil {
			return fmt.Errorf("Satellite Location still exists: %s", rs.Primary.ID)
		}

	}
	return nil
}

func testAccCheckSatelliteLocationCreate(name, managed_from string, coreos_enabled string) string {
	return fmt.Sprintf(`

	data "ibm_resource_group" "res_group" {
		is_default = true
	} 

	resource "ibm_satellite_location" "location" {
		location = "%s"
		managed_from = "%s"
		coreos_enabled = "%s"
		description = "test"
		zones = ["us-east-1", "us-east-2", "us-east-3"]
		resource_group_id = data.ibm_resource_group.res_group.id
		tags = ["env:dev"]
	}
	  
`, name, managed_from, coreos_enabled)
}
