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
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMIsClusterNetworkSubnetDataSourceBasic(t *testing.T) {
	clusterNetworkSubnetClusterNetworkID := fmt.Sprintf("tf_cluster_network_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsClusterNetworkSubnetDataSourceConfigBasic(clusterNetworkSubnetClusterNetworkID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "cluster_network_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "cluster_network_subnet_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "available_ipv4_address_count"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "ip_version"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "ipv4_cidr_block"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "lifecycle_reasons.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "total_ipv4_address_count"),
				),
			},
		},
	})
}

func TestAccIBMIsClusterNetworkSubnetDataSourceAllArgs(t *testing.T) {
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
				Config: testAccCheckIBMIsClusterNetworkSubnetDataSourceConfig(clusterNetworkSubnetClusterNetworkID, clusterNetworkSubnetIPVersion, clusterNetworkSubnetIpv4CIDRBlock, clusterNetworkSubnetName, clusterNetworkSubnetTotalIpv4AddressCount),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "cluster_network_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "cluster_network_subnet_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "available_ipv4_address_count"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "ip_version"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "ipv4_cidr_block"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "lifecycle_reasons.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "lifecycle_reasons.0.code"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "lifecycle_reasons.0.message"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "lifecycle_reasons.0.more_info"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "total_ipv4_address_count"),
				),
			},
		},
	})
}

func testAccCheckIBMIsClusterNetworkSubnetDataSourceConfigBasic(clusterNetworkSubnetClusterNetworkID string) string {
	return fmt.Sprintf(`

		data "ibm_is_cluster_network_subnet" "is_cluster_network_subnet_instance" {
			cluster_network_id = "02c7-10274052-f495-4920-a67f-870eb3b87003"
			cluster_network_subnet_id = "02c7-27ca8206-f73b-4d8f-b996-913711b98ff0"
		}
	`)
}

func testAccCheckIBMIsClusterNetworkSubnetDataSourceConfig(clusterNetworkSubnetClusterNetworkID string, clusterNetworkSubnetIPVersion string, clusterNetworkSubnetIpv4CIDRBlock string, clusterNetworkSubnetName string, clusterNetworkSubnetTotalIpv4AddressCount string) string {
	return fmt.Sprintf(`
		resource "ibm_is_cluster_network_subnet" "is_cluster_network_subnet_instance" {
			cluster_network_id = "%s"
			ip_version = "%s"
			ipv4_cidr_block = "%s"
			name = "%s"
			total_ipv4_address_count = %s
		}

		data "ibm_is_cluster_network_subnet" "is_cluster_network_subnet_instance" {
			cluster_network_id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_id
			cluster_network_subnet_id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
		}
	`, clusterNetworkSubnetClusterNetworkID, clusterNetworkSubnetIPVersion, clusterNetworkSubnetIpv4CIDRBlock, clusterNetworkSubnetName, clusterNetworkSubnetTotalIpv4AddressCount)
}

func TestDataSourceIBMIsClusterNetworkSubnetClusterNetworkSubnetLifecycleReasonToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsClusterNetworkSubnetClusterNetworkSubnetLifecycleReasonToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
