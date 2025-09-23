// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/service/vpc"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMIsClusterNetworkDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsClusterNetworkDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "cluster_network_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "lifecycle_reasons.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "profile.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "profile.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "profile.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "profile.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "resource_group.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "resource_group.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "resource_group.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "resource_group.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "subnet_prefixes.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "subnet_prefixes.0.allocation_policy"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "subnet_prefixes.0.cidr"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "vpc.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "vpc.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "vpc.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "vpc.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "vpc.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "zone.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "zone.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "zone.0.name"),
				),
			},
		},
	})
}

func TestAccIBMIsClusterNetworkDataSourceAllArgs(t *testing.T) {
	clusterNetworkName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsClusterNetworkDataSourceConfig(clusterNetworkName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "cluster_network_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "lifecycle_reasons.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "lifecycle_reasons.0.code"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "lifecycle_reasons.0.message"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "lifecycle_reasons.0.more_info"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "profile.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "resource_group.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "subnet_prefixes.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "subnet_prefixes.0.allocation_policy"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "subnet_prefixes.0.cidr"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "vpc.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network.is_cluster_network_instance", "zone.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsClusterNetworkDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_is_cluster_network" "is_cluster_network_instance" {
			cluster_network_id = "02c7-10274052-f495-4920-a67f-870eb3b87003"
		}
	`)
}

func testAccCheckIBMIsClusterNetworkDataSourceConfig(clusterNetworkName string) string {
	return fmt.Sprintf(`
		resource "ibm_is_cluster_network" "is_cluster_network_instance" {
			name = "%s"
			profile {
				href = "https://us-south.iaas.cloud.ibm.com/v1/cluster_network/profiles/h100"
				name = "h100"
				resource_type = "cluster_network_profile"
			}
			resource_group {
				href = "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345"
				id = "fee82deba12e4c0fb69c3b09d1f12345"
				name = "my-resource-group"
			}
			subnet_prefixes {
				allocation_policy = "auto"
				cidr = "10.0.0.0/24"
			}
			vpc {
				crn = "crn:v1:bluemix:public:is:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34::vpc:r006-4727d842-f94f-4a2d-824a-9bc9b02c523b"
				deleted {
					more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
				}
				href = "https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-4727d842-f94f-4a2d-824a-9bc9b02c523b"
				id = "r006-4727d842-f94f-4a2d-824a-9bc9b02c523b"
				name = "my-vpc"
				resource_type = "vpc"
			}
			zone {
				href = "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1"
				name = "us-south-1"
			}
		}

		data "ibm_is_cluster_network" "is_cluster_network_instance" {
			cluster_network_id = "cluster_network_id"
		}
	`, clusterNetworkName)
}

func TestDataSourceIBMIsClusterNetworkClusterNetworkLifecycleReasonToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["code"] = "resource_suspended_by_provider"
		model["message"] = "The resource has been suspended. Contact IBM support with the CRN for next steps."
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#resource-suspension"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.ClusterNetworkLifecycleReason)
	model.Code = core.StringPtr("resource_suspended_by_provider")
	model.Message = core.StringPtr("The resource has been suspended. Contact IBM support with the CRN for next steps.")
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#resource-suspension")

	result, err := vpc.DataSourceIBMIsClusterNetworkClusterNetworkLifecycleReasonToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsClusterNetworkClusterNetworkProfileReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/cluster_network/profiles/h100"
		model["name"] = "h100"
		model["resource_type"] = "cluster_network_profile"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.ClusterNetworkProfileReference)
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/cluster_network/profiles/h100")
	model.Name = core.StringPtr("h100")
	model.ResourceType = core.StringPtr("cluster_network_profile")

	result, err := vpc.DataSourceIBMIsClusterNetworkClusterNetworkProfileReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsClusterNetworkResourceGroupReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["href"] = "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345"
		model["id"] = "fee82deba12e4c0fb69c3b09d1f12345"
		model["name"] = "my-resource-group"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.ResourceGroupReference)
	model.Href = core.StringPtr("https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345")
	model.ID = core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345")
	model.Name = core.StringPtr("my-resource-group")

	result, err := vpc.DataSourceIBMIsClusterNetworkResourceGroupReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsClusterNetworkClusterNetworkSubnetPrefixToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["allocation_policy"] = "auto"
		model["cidr"] = "10.0.0.0/24"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.ClusterNetworkSubnetPrefix)
	model.AllocationPolicy = core.StringPtr("auto")
	model.CIDR = core.StringPtr("10.0.0.0/24")

	result, err := vpc.DataSourceIBMIsClusterNetworkClusterNetworkSubnetPrefixToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsClusterNetworkVPCReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		deletedModel := make(map[string]interface{})
		deletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		model := make(map[string]interface{})
		model["crn"] = "crn:v1:bluemix:public:is:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34::vpc:r006-4727d842-f94f-4a2d-824a-9bc9b02c523b"
		model["deleted"] = []map[string]interface{}{deletedModel}
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-4727d842-f94f-4a2d-824a-9bc9b02c523b"
		model["id"] = "r006-4727d842-f94f-4a2d-824a-9bc9b02c523b"
		model["name"] = "my-vpc"
		model["resource_type"] = "vpc"

		assert.Equal(t, result, model)
	}

	deletedModel := new(vpcv1.Deleted)
	deletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	model := new(vpcv1.VPCReference)
	model.CRN = core.StringPtr("crn:v1:bluemix:public:is:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34::vpc:r006-4727d842-f94f-4a2d-824a-9bc9b02c523b")
	model.Deleted = deletedModel
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-4727d842-f94f-4a2d-824a-9bc9b02c523b")
	model.ID = core.StringPtr("r006-4727d842-f94f-4a2d-824a-9bc9b02c523b")
	model.Name = core.StringPtr("my-vpc")
	model.ResourceType = core.StringPtr("vpc")

	result, err := vpc.DataSourceIBMIsClusterNetworkVPCReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsClusterNetworkDeletedToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.Deleted)
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	result, err := vpc.DataSourceIBMIsClusterNetworkDeletedToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsClusterNetworkZoneReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1"
		model["name"] = "us-south-1"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.ZoneReference)
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1")
	model.Name = core.StringPtr("us-south-1")

	result, err := vpc.DataSourceIBMIsClusterNetworkZoneReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
