// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIsClusterNetworkProfilesDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsClusterNetworkProfilesDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_profiles.is_cluster_network_profiles_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_profiles.is_cluster_network_profiles_instance", "profiles.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_profiles.is_cluster_network_profiles_instance", "profiles.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_profiles.is_cluster_network_profiles_instance", "profiles.0.family"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_profiles.is_cluster_network_profiles_instance", "profiles.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_profiles.is_cluster_network_profiles_instance", "profiles.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_profiles.is_cluster_network_profiles_instance", "profiles.0.supported_instance_profiles.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_profiles.is_cluster_network_profiles_instance", "profiles.0.supported_instance_profiles.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_profiles.is_cluster_network_profiles_instance", "profiles.0.supported_instance_profiles.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_profiles.is_cluster_network_profiles_instance", "profiles.0.supported_instance_profiles.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_profiles.is_cluster_network_profiles_instance", "profiles.0.zones.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsClusterNetworkProfilesDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_is_cluster_network_profiles" "is_cluster_network_profiles_instance" {
		}
	`)
}
