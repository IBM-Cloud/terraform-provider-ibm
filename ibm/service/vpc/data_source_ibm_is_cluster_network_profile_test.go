// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/vpc"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMIsClusterNetworkProfileDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsClusterNetworkProfileDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_profile.is_cluster_network_profile_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_profile.is_cluster_network_profile_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_profile.is_cluster_network_profile_instance", "family"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_profile.is_cluster_network_profile_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_profile.is_cluster_network_profile_instance", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_profile.is_cluster_network_profile_instance", "address_configuration_services.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_profile.is_cluster_network_profile_instance", "address_configuration_services.values"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_profile.is_cluster_network_profile_instance", "isolation_group_count.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_profile.is_cluster_network_profile_instance", "isolation_group_count.value"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_profile.is_cluster_network_profile_instance", "subnet_routing_supported.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_profile.is_cluster_network_profile_instance", "subnet_routing_supported.value"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_profile.is_cluster_network_profile_instance", "supported_instance_profiles.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_profile.is_cluster_network_profile_instance", "supported_instance_profiles.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_profile.is_cluster_network_profile_instance", "supported_instance_profiles.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_profile.is_cluster_network_profile_instance", "supported_instance_profiles.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_profile.is_cluster_network_profile_instance", "zones.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsClusterNetworkProfileDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_is_cluster_network_profile" "is_cluster_network_profile_instance" {
			name = "%s"
		}
	`, acc.ISClusterNetworkProfileName)
}

func TestDataSourceIBMIsClusterNetworkProfileInstanceProfileReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/instance/profiles/bx2-4x16"
		model["name"] = "bx2-4x16"
		model["resource_type"] = "instance_profile"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.InstanceProfileReference)
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/instance/profiles/bx2-4x16")
	model.Name = core.StringPtr("bx2-4x16")
	model.ResourceType = core.StringPtr("instance_profile")

	result, err := vpc.DataSourceIBMIsClusterNetworkProfileInstanceProfileReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsClusterNetworkProfileZoneReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1"
		model["name"] = "us-south-1"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.ZoneReference)
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1")
	model.Name = core.StringPtr("us-south-1")

	result, err := vpc.DataSourceIBMIsClusterNetworkProfileZoneReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
