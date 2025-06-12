// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/vpc"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMIsClusterNetworkSubnetsDataSourceBasic(t *testing.T) {
	clusterNetworkSubnetClusterNetworkID := fmt.Sprintf("tf_cluster_network_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsClusterNetworkSubnetsDataSourceConfigBasic(clusterNetworkSubnetClusterNetworkID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnets.is_cluster_network_subnets_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnets.is_cluster_network_subnets_instance", "cluster_network_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnets.is_cluster_network_subnets_instance", "subnets.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnets.is_cluster_network_subnets_instance", "subnets.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnets.is_cluster_network_subnets_instance", "subnets.0.available_ipv4_address_count"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnets.is_cluster_network_subnets_instance", "subnets.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnets.is_cluster_network_subnets_instance", "subnets.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnets.is_cluster_network_subnets_instance", "subnets.0.ip_version"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnets.is_cluster_network_subnets_instance", "subnets.0.ipv4_cidr_block"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnets.is_cluster_network_subnets_instance", "subnets.0.lifecycle_reasons.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnets.is_cluster_network_subnets_instance", "subnets.0.lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnets.is_cluster_network_subnets_instance", "subnets.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnets.is_cluster_network_subnets_instance", "subnets.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnets.is_cluster_network_subnets_instance", "subnets.0.total_ipv4_address_count"),
				),
			},
		},
	})
}

func TestAccIBMIsClusterNetworkSubnetsDataSourceAllArgs(t *testing.T) {
	clusterNetworkSubnetClusterNetworkID := fmt.Sprintf("tf_cluster_network_id_%d", acctest.RandIntRange(10, 100))
	clusterNetworkSubnetIPVersion := "ipv4"
	clusterNetworkSubnetIpv4CIDRBlock := fmt.Sprintf("tf_ipv4_cidr_block_%d", acctest.RandIntRange(10, 100))
	clusterNetworkSubnetName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	clusterNetworkSubnetTotalIpv4AddressCount := fmt.Sprintf("%d", acctest.RandIntRange(8, 16777216))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsClusterNetworkSubnetsDataSourceConfig(clusterNetworkSubnetClusterNetworkID, clusterNetworkSubnetIPVersion, clusterNetworkSubnetIpv4CIDRBlock, clusterNetworkSubnetName, clusterNetworkSubnetTotalIpv4AddressCount),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnets.is_cluster_network_subnets_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnets.is_cluster_network_subnets_instance", "cluster_network_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnets.is_cluster_network_subnets_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnets.is_cluster_network_subnets_instance", "sort"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnets.is_cluster_network_subnets_instance", "subnets.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnets.is_cluster_network_subnets_instance", "subnets.0.available_ipv4_address_count"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnets.is_cluster_network_subnets_instance", "subnets.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnets.is_cluster_network_subnets_instance", "subnets.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnets.is_cluster_network_subnets_instance", "subnets.0.id"),
					resource.TestCheckResourceAttr("data.ibm_is_cluster_network_subnets.is_cluster_network_subnets_instance", "subnets.0.ip_version", clusterNetworkSubnetIPVersion),
					resource.TestCheckResourceAttr("data.ibm_is_cluster_network_subnets.is_cluster_network_subnets_instance", "subnets.0.ipv4_cidr_block", clusterNetworkSubnetIpv4CIDRBlock),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnets.is_cluster_network_subnets_instance", "subnets.0.lifecycle_state"),
					resource.TestCheckResourceAttr("data.ibm_is_cluster_network_subnets.is_cluster_network_subnets_instance", "subnets.0.name", clusterNetworkSubnetName),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnets.is_cluster_network_subnets_instance", "subnets.0.resource_type"),
					resource.TestCheckResourceAttr("data.ibm_is_cluster_network_subnets.is_cluster_network_subnets_instance", "subnets.0.total_ipv4_address_count", clusterNetworkSubnetTotalIpv4AddressCount),
				),
			},
		},
	})
}

func testAccCheckIBMIsClusterNetworkSubnetsDataSourceConfigBasic(clusterNetworkSubnetClusterNetworkID string) string {
	return fmt.Sprintf(`

		data "ibm_is_cluster_network_subnets" "is_cluster_network_subnets_instance" {
			cluster_network_id =  "02c7-10274052-f495-4920-a67f-870eb3b87003"
		}
	`)
}

func testAccCheckIBMIsClusterNetworkSubnetsDataSourceConfig(clusterNetworkSubnetClusterNetworkID string, clusterNetworkSubnetIPVersion string, clusterNetworkSubnetIpv4CIDRBlock string, clusterNetworkSubnetName string, clusterNetworkSubnetTotalIpv4AddressCount string) string {
	return fmt.Sprintf(`
		resource "ibm_is_cluster_network_subnet" "is_cluster_network_subnet_instance" {
			cluster_network_id = "%s"
			ip_version = "%s"
			ipv4_cidr_block = "%s"
			name = "%s"
			total_ipv4_address_count = %s
		}

		data "ibm_is_cluster_network_subnets" "is_cluster_network_subnets_instance" {
			cluster_network_id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_id
			name = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.name
			sort = "name"
		}
	`, clusterNetworkSubnetClusterNetworkID, clusterNetworkSubnetIPVersion, clusterNetworkSubnetIpv4CIDRBlock, clusterNetworkSubnetName, clusterNetworkSubnetTotalIpv4AddressCount)
}

func TestDataSourceIBMIsClusterNetworkSubnetsClusterNetworkSubnetToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		clusterNetworkSubnetLifecycleReasonModel := make(map[string]interface{})
		clusterNetworkSubnetLifecycleReasonModel["code"] = "resource_suspended_by_provider"
		clusterNetworkSubnetLifecycleReasonModel["message"] = "The resource has been suspended. Contact IBM support with the CRN for next steps."
		clusterNetworkSubnetLifecycleReasonModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#resource-suspension"

		model := make(map[string]interface{})
		model["available_ipv4_address_count"] = int(15)
		model["created_at"] = "2019-01-01T12:00:00.000Z"
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/cluster_networks/0717-da0df18c-7598-4633-a648-fdaac28a5573/subnets/0717-7931845c-65c4-4b0a-80cd-7d9c1a6d7930"
		model["id"] = "0717-7931845c-65c4-4b0a-80cd-7d9c1a6d7930"
		model["ip_version"] = "ipv4"
		model["ipv4_cidr_block"] = "10.0.0.0/24"
		model["lifecycle_reasons"] = []map[string]interface{}{clusterNetworkSubnetLifecycleReasonModel}
		model["lifecycle_state"] = "stable"
		model["name"] = "my-cluster-network-subnet"
		model["resource_type"] = "cluster_network_subnet"
		model["total_ipv4_address_count"] = int(256)

		assert.Equal(t, result, model)
	}

	clusterNetworkSubnetLifecycleReasonModel := new(vpcv1.ClusterNetworkSubnetLifecycleReason)
	clusterNetworkSubnetLifecycleReasonModel.Code = core.StringPtr("resource_suspended_by_provider")
	clusterNetworkSubnetLifecycleReasonModel.Message = core.StringPtr("The resource has been suspended. Contact IBM support with the CRN for next steps.")
	clusterNetworkSubnetLifecycleReasonModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#resource-suspension")

	model := new(vpcv1.ClusterNetworkSubnet)
	model.AvailableIpv4AddressCount = core.Int64Ptr(int64(15))
	model.CreatedAt = CreateMockDateTime("2019-01-01T12:00:00.000Z")
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/cluster_networks/0717-da0df18c-7598-4633-a648-fdaac28a5573/subnets/0717-7931845c-65c4-4b0a-80cd-7d9c1a6d7930")
	model.ID = core.StringPtr("0717-7931845c-65c4-4b0a-80cd-7d9c1a6d7930")
	model.IPVersion = core.StringPtr("ipv4")
	model.Ipv4CIDRBlock = core.StringPtr("10.0.0.0/24")
	model.LifecycleReasons = []vpcv1.ClusterNetworkSubnetLifecycleReason{*clusterNetworkSubnetLifecycleReasonModel}
	model.LifecycleState = core.StringPtr("stable")
	model.Name = core.StringPtr("my-cluster-network-subnet")
	model.ResourceType = core.StringPtr("cluster_network_subnet")
	model.TotalIpv4AddressCount = core.Int64Ptr(int64(256))

	result, err := vpc.DataSourceIBMIsClusterNetworkSubnetsClusterNetworkSubnetToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsClusterNetworkSubnetsClusterNetworkSubnetLifecycleReasonToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["code"] = "resource_suspended_by_provider"
		model["message"] = "The resource has been suspended. Contact IBM support with the CRN for next steps."
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#resource-suspension"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.ClusterNetworkSubnetLifecycleReason)
	model.Code = core.StringPtr("resource_suspended_by_provider")
	model.Message = core.StringPtr("The resource has been suspended. Contact IBM support with the CRN for next steps.")
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#resource-suspension")

	result, err := vpc.DataSourceIBMIsClusterNetworkSubnetsClusterNetworkSubnetLifecycleReasonToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
