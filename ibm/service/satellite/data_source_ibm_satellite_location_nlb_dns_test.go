// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package satellite_test

import (
	"fmt"
	"testing"

	"github.com/IBM-Cloud/container-services-go-sdk/kubernetesserviceapiv1"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMSatelliteLocationNLBDNSListBasic(t *testing.T) {
	name := fmt.Sprintf("tf-satellitelocation-%d", acctest.RandIntRange(10, 100))
	managed_from := "wdc04"
	physical_address := "test location address"
	capabilities := []kubernetesserviceapiv1.CapabilityManagedBySatellite{kubernetesserviceapiv1.OnPrem}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSatelliteLocationNLBDNSListConfig(name, managed_from, physical_address, capabilities),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_satellite_location_nlb_dns.dns_list", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMSatelliteLocationNLBDNSListConfig(name, managed_from, physical_address string, capabilities []kubernetesserviceapiv1.CapabilityManagedBySatellite) string {
	return testAccCheckSatelliteLocationDataSource(name, managed_from, physical_address, capabilities) + `
	data ibm_satellite_location_nlb_dns dns_list {
		location = ibm_satellite_location.location.id
	}
`
}
