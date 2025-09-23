// Copyright IBM Corp. 2017, 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package satellite_test

import (
	"fmt"
	"testing"

	"github.com/IBM-Cloud/container-services-go-sdk/kubernetesserviceapiv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
)

func TestAccSatelliteLocation_Basic(t *testing.T) {
	var instance string
	name := fmt.Sprintf("tf-satellitelocation-%d", acctest.RandIntRange(10, 100))
	managed_from := "wdc04"
	physical_address := "test-road 10, 111 test-place, testcountry"
	coreos_enabled := "true"
	capabilities := []kubernetesserviceapiv1.CapabilityManagedBySatellite{kubernetesserviceapiv1.OnPrem}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckSatelliteLocationDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckSatelliteLocationCreate(name, managed_from, physical_address, coreos_enabled, "", "", capabilities),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckSatelliteLocationExists("ibm_satellite_location.location", instance),
					resource.TestCheckResourceAttr("ibm_satellite_location.location", "location", name),
					resource.TestCheckResourceAttr("ibm_satellite_location.location", "managed_from", managed_from),
					resource.TestCheckResourceAttr("ibm_satellite_location.location", "physical_address", physical_address),
					resource.TestCheckResourceAttr("ibm_satellite_location.location", "coreos_enabled", coreos_enabled),
					resource.TestCheckResourceAttr("ibm_satellite_location.location", "capabilities.#", "1"),
					resource.TestCheckResourceAttr("ibm_satellite_location.location", "capabilities.0", "on-prem"),
				),
			},
		},
	})
}

func TestAccSatelliteLocation_Import(t *testing.T) {
	var instance string
	name := fmt.Sprintf("tf_location_%d", acctest.RandIntRange(10, 100))
	managed_from := "wdc04"
	physical_address := "test-road 10, 111 test-place, testcountry"
	coreos_enabled := "true"
	capabilities := []kubernetesserviceapiv1.CapabilityManagedBySatellite{kubernetesserviceapiv1.OnPrem}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckSatelliteLocationDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckSatelliteLocationCreate(name, managed_from, physical_address, coreos_enabled, "", "", capabilities),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckSatelliteLocationExists("ibm_satellite_location.location", instance),
					resource.TestCheckResourceAttr("ibm_satellite_location.location", "location", name),
					resource.TestCheckResourceAttr("ibm_satellite_location.location", "managed_from", managed_from),
					resource.TestCheckResourceAttr("ibm_satellite_location.location", "physical_address", physical_address),
					resource.TestCheckResourceAttr("ibm_satellite_location.location", "capabilities.#", "1"),
					resource.TestCheckResourceAttr("ibm_satellite_location.location", "capabilities.0", "on-prem"),
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

func TestAccSatelliteLocation_PodAndServiceSubnet(t *testing.T) {
	var instance string
	name := fmt.Sprintf("tf-satellitelocation-%d", acctest.RandIntRange(10, 100))
	managed_from := "wdc04"
	physical_address := "test-road 10, 111 test-place, testcountry"
	coreos_enabled := "true"
	pod_subnet := "10.69.0.0/16"
	service_subnet := "192.168.42.0/24"
	capabilities := []kubernetesserviceapiv1.CapabilityManagedBySatellite{kubernetesserviceapiv1.OnPrem}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckSatelliteLocationCreate(name, managed_from, physical_address, coreos_enabled, pod_subnet, service_subnet, capabilities),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckSatelliteLocationExists("ibm_satellite_location.location", instance),
					resource.TestCheckResourceAttr("ibm_satellite_location.location", "location", name),
					resource.TestCheckResourceAttr("ibm_satellite_location.location", "managed_from", managed_from),
					resource.TestCheckResourceAttr("ibm_satellite_location.location", "physical_address", physical_address),
					resource.TestCheckResourceAttr("ibm_satellite_location.location", "coreos_enabled", coreos_enabled),
					resource.TestCheckResourceAttr("ibm_satellite_location.location", "pod_subnet", pod_subnet),
					resource.TestCheckResourceAttr("ibm_satellite_location.location", "service_subnet", service_subnet),
					resource.TestCheckResourceAttr("ibm_satellite_location.location", "capabilities.#", "1"),
					resource.TestCheckResourceAttr("ibm_satellite_location.location", "capabilities.0", "on-prem"),
				),
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

func testAccCheckSatelliteLocationCreate(name, managed_from string, physical_address string, coreos_enabled string, pod_subnet, service_subnet string, capabilities []kubernetesserviceapiv1.CapabilityManagedBySatellite) string {
	return fmt.Sprintf(`

	data "ibm_resource_group" "res_group" {
		is_default = true
	} 

	resource "ibm_satellite_location" "location" {
		location = "%s"
		managed_from = "%s"
		physical_address = "%s"
		coreos_enabled = "%s"
		description = "test"
		zones = ["us-east-1", "us-east-2", "us-east-3"]
		resource_group_id = data.ibm_resource_group.res_group.id
		tags = ["env:dev"]
		pod_subnet = "%s"
		service_subnet = "%s"
		capabilities = %q
	}
	  
`, name, managed_from, physical_address, coreos_enabled, pod_subnet, service_subnet, capabilities)
}
