// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/vpc"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMIsClusterNetworksDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsClusterNetworksDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "cluster_networks.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "cluster_networks.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "cluster_networks.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "cluster_networks.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "cluster_networks.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "cluster_networks.0.lifecycle_reasons.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "cluster_networks.0.lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "cluster_networks.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "cluster_networks.0.profile.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "cluster_networks.0.profile.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "cluster_networks.0.profile.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "cluster_networks.0.profile.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "cluster_networks.0.resource_group.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "cluster_networks.0.resource_group.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "cluster_networks.0.resource_group.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "cluster_networks.0.resource_group.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "cluster_networks.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "cluster_networks.0.subnet_prefixes.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "cluster_networks.0.subnet_prefixes.0.allocation_policy"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "cluster_networks.0.subnet_prefixes.0.cidr"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "cluster_networks.0.vpc.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "cluster_networks.0.vpc.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "cluster_networks.0.vpc.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "cluster_networks.0.vpc.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "cluster_networks.0.vpc.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "cluster_networks.0.zone.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "cluster_networks.0.zone.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "cluster_networks.0.zone.0.name"),
				),
			},
		},
	})
}

func TestAccIBMIsClusterNetworksDataSourceAllArgs(t *testing.T) {
	clusterNetworkName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsClusterNetworksDataSourceConfig(clusterNetworkName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "sort"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "vpc_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "vpc_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "vpc_name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "cluster_networks.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "cluster_networks.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "cluster_networks.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "cluster_networks.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "cluster_networks.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "cluster_networks.0.lifecycle_state"),
					resource.TestCheckResourceAttr("data.ibm_is_cluster_networks.is_cluster_networks_instance", "cluster_networks.0.name", clusterNetworkName),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_networks.is_cluster_networks_instance", "cluster_networks.0.resource_type"),
				),
			},
		},
	})
}

func testAccCheckIBMIsClusterNetworksDataSourceConfigBasic() string {
	return fmt.Sprintf(`

		data "ibm_is_cluster_networks" "is_cluster_networks_instance" {
		}
	`)
}

func testAccCheckIBMIsClusterNetworksDataSourceConfig(clusterNetworkName string) string {
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

		data "ibm_is_cluster_networks" "is_cluster_networks_instance" {
			resource_group_id = "resource_group_id"
			name = ibm_is_cluster_network.is_cluster_network_instance.name
			sort = "name"
			vpc_id = "vpc_id"
			vpc_crn = "crn:v1:bluemix:public:is:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34::vpc:r006-4727d842-f94f-4a2d-824a-9bc9b02c523b"
			vpc_name = "my-vpc"
		}
	`, clusterNetworkName)
}

func TestDataSourceIBMIsClusterNetworksClusterNetworkToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		clusterNetworkLifecycleReasonModel := make(map[string]interface{})
		clusterNetworkLifecycleReasonModel["code"] = "resource_suspended_by_provider"
		clusterNetworkLifecycleReasonModel["message"] = "The resource has been suspended. Contact IBM support with the CRN for next steps."
		clusterNetworkLifecycleReasonModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#resource-suspension"

		clusterNetworkProfileReferenceModel := make(map[string]interface{})
		clusterNetworkProfileReferenceModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/cluster_network/profiles/h100"
		clusterNetworkProfileReferenceModel["name"] = "h100"
		clusterNetworkProfileReferenceModel["resource_type"] = "cluster_network_profile"

		resourceGroupReferenceModel := make(map[string]interface{})
		resourceGroupReferenceModel["href"] = "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345"
		resourceGroupReferenceModel["id"] = "fee82deba12e4c0fb69c3b09d1f12345"
		resourceGroupReferenceModel["name"] = "Default"

		clusterNetworkSubnetPrefixModel := make(map[string]interface{})
		clusterNetworkSubnetPrefixModel["allocation_policy"] = "auto"
		clusterNetworkSubnetPrefixModel["cidr"] = "10.0.0.0/9"

		deletedModel := make(map[string]interface{})
		deletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		vpcReferenceModel := make(map[string]interface{})
		vpcReferenceModel["crn"] = "crn:v1:bluemix:public:is:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34::vpc:r006-4727d842-f94f-4a2d-824a-9bc9b02c523b"
		vpcReferenceModel["deleted"] = []map[string]interface{}{deletedModel}
		vpcReferenceModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-4727d842-f94f-4a2d-824a-9bc9b02c523b"
		vpcReferenceModel["id"] = "r006-4727d842-f94f-4a2d-824a-9bc9b02c523b"
		vpcReferenceModel["name"] = "my-vpc"
		vpcReferenceModel["resource_type"] = "vpc"

		zoneReferenceModel := make(map[string]interface{})
		zoneReferenceModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1"
		zoneReferenceModel["name"] = "us-south-1"

		model := make(map[string]interface{})
		model["created_at"] = "2019-01-01T12:00:00.000Z"
		model["crn"] = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::cluster-network:0717-da0df18c-7598-4633-a648-fdaac28a5573"
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/cluster_networks/0717-da0df18c-7598-4633-a648-fdaac28a5573"
		model["id"] = "0717-da0df18c-7598-4633-a648-fdaac28a5573"
		model["lifecycle_reasons"] = []map[string]interface{}{clusterNetworkLifecycleReasonModel}
		model["lifecycle_state"] = "stable"
		model["name"] = "my-cluster-network"
		model["profile"] = []map[string]interface{}{clusterNetworkProfileReferenceModel}
		model["resource_group"] = []map[string]interface{}{resourceGroupReferenceModel}
		model["resource_type"] = "cluster_network"
		model["subnet_prefixes"] = []map[string]interface{}{clusterNetworkSubnetPrefixModel}
		model["vpc"] = []map[string]interface{}{vpcReferenceModel}
		model["zone"] = []map[string]interface{}{zoneReferenceModel}

		assert.Equal(t, result, model)
	}

	clusterNetworkLifecycleReasonModel := new(vpcv1.ClusterNetworkLifecycleReason)
	clusterNetworkLifecycleReasonModel.Code = core.StringPtr("resource_suspended_by_provider")
	clusterNetworkLifecycleReasonModel.Message = core.StringPtr("The resource has been suspended. Contact IBM support with the CRN for next steps.")
	clusterNetworkLifecycleReasonModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#resource-suspension")

	clusterNetworkProfileReferenceModel := new(vpcv1.ClusterNetworkProfileReference)
	clusterNetworkProfileReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/cluster_network/profiles/h100")
	clusterNetworkProfileReferenceModel.Name = core.StringPtr("h100")
	clusterNetworkProfileReferenceModel.ResourceType = core.StringPtr("cluster_network_profile")

	resourceGroupReferenceModel := new(vpcv1.ResourceGroupReference)
	resourceGroupReferenceModel.Href = core.StringPtr("https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345")
	resourceGroupReferenceModel.ID = core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345")
	resourceGroupReferenceModel.Name = core.StringPtr("Default")

	clusterNetworkSubnetPrefixModel := new(vpcv1.ClusterNetworkSubnetPrefix)
	clusterNetworkSubnetPrefixModel.AllocationPolicy = core.StringPtr("auto")
	clusterNetworkSubnetPrefixModel.CIDR = core.StringPtr("10.0.0.0/9")

	deletedModel := new(vpcv1.Deleted)
	deletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	vpcReferenceModel := new(vpcv1.VPCReference)
	vpcReferenceModel.CRN = core.StringPtr("crn:v1:bluemix:public:is:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34::vpc:r006-4727d842-f94f-4a2d-824a-9bc9b02c523b")
	vpcReferenceModel.Deleted = deletedModel
	vpcReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-4727d842-f94f-4a2d-824a-9bc9b02c523b")
	vpcReferenceModel.ID = core.StringPtr("r006-4727d842-f94f-4a2d-824a-9bc9b02c523b")
	vpcReferenceModel.Name = core.StringPtr("my-vpc")
	vpcReferenceModel.ResourceType = core.StringPtr("vpc")

	zoneReferenceModel := new(vpcv1.ZoneReference)
	zoneReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1")
	zoneReferenceModel.Name = core.StringPtr("us-south-1")

	model := new(vpcv1.ClusterNetwork)
	model.CreatedAt = CreateMockDateTime("2019-01-01T12:00:00.000Z")
	model.CRN = core.StringPtr("crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::cluster-network:0717-da0df18c-7598-4633-a648-fdaac28a5573")
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/cluster_networks/0717-da0df18c-7598-4633-a648-fdaac28a5573")
	model.ID = core.StringPtr("0717-da0df18c-7598-4633-a648-fdaac28a5573")
	model.LifecycleReasons = []vpcv1.ClusterNetworkLifecycleReason{*clusterNetworkLifecycleReasonModel}
	model.LifecycleState = core.StringPtr("stable")
	model.Name = core.StringPtr("my-cluster-network")
	model.Profile = clusterNetworkProfileReferenceModel
	model.ResourceGroup = resourceGroupReferenceModel
	model.ResourceType = core.StringPtr("cluster_network")
	model.SubnetPrefixes = []vpcv1.ClusterNetworkSubnetPrefix{*clusterNetworkSubnetPrefixModel}
	model.VPC = vpcReferenceModel
	model.Zone = zoneReferenceModel

	result, err := vpc.DataSourceIBMIsClusterNetworksClusterNetworkToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsClusterNetworksClusterNetworkLifecycleReasonToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsClusterNetworksClusterNetworkLifecycleReasonToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsClusterNetworksClusterNetworkProfileReferenceToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsClusterNetworksClusterNetworkProfileReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsClusterNetworksResourceGroupReferenceToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsClusterNetworksResourceGroupReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsClusterNetworksClusterNetworkSubnetPrefixToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["allocation_policy"] = "auto"
		model["cidr"] = "10.0.0.0/24"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.ClusterNetworkSubnetPrefix)
	model.AllocationPolicy = core.StringPtr("auto")
	model.CIDR = core.StringPtr("10.0.0.0/24")

	result, err := vpc.DataSourceIBMIsClusterNetworksClusterNetworkSubnetPrefixToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsClusterNetworksVPCReferenceToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsClusterNetworksVPCReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsClusterNetworksDeletedToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.Deleted)
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	result, err := vpc.DataSourceIBMIsClusterNetworksDeletedToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsClusterNetworksZoneReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1"
		model["name"] = "us-south-1"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.ZoneReference)
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1")
	model.Name = core.StringPtr("us-south-1")

	result, err := vpc.DataSourceIBMIsClusterNetworksZoneReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
