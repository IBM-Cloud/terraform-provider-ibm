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

func TestAccIBMIsClusterNetworkSubnetReservedIpsDataSourceBasic(t *testing.T) {
	clusterNetworkSubnetReservedIPClusterNetworkID := fmt.Sprintf("tf_cluster_network_id_%d", acctest.RandIntRange(10, 100))
	clusterNetworkSubnetReservedIPClusterNetworkSubnetID := fmt.Sprintf("tf_cluster_network_subnet_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsClusterNetworkSubnetReservedIpsDataSourceConfigBasic(clusterNetworkSubnetReservedIPClusterNetworkID, clusterNetworkSubnetReservedIPClusterNetworkSubnetID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet_reserved_ips.is_cluster_network_subnet_reserved_ips_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet_reserved_ips.is_cluster_network_subnet_reserved_ips_instance", "cluster_network_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet_reserved_ips.is_cluster_network_subnet_reserved_ips_instance", "cluster_network_subnet_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet_reserved_ips.is_cluster_network_subnet_reserved_ips_instance", "reserved_ips.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet_reserved_ips.is_cluster_network_subnet_reserved_ips_instance", "reserved_ips.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet_reserved_ips.is_cluster_network_subnet_reserved_ips_instance", "reserved_ips.0.address"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet_reserved_ips.is_cluster_network_subnet_reserved_ips_instance", "reserved_ips.0.auto_delete"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet_reserved_ips.is_cluster_network_subnet_reserved_ips_instance", "reserved_ips.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet_reserved_ips.is_cluster_network_subnet_reserved_ips_instance", "reserved_ips.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet_reserved_ips.is_cluster_network_subnet_reserved_ips_instance", "reserved_ips.0.lifecycle_reasons.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet_reserved_ips.is_cluster_network_subnet_reserved_ips_instance", "reserved_ips.0.lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet_reserved_ips.is_cluster_network_subnet_reserved_ips_instance", "reserved_ips.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet_reserved_ips.is_cluster_network_subnet_reserved_ips_instance", "reserved_ips.0.owner"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet_reserved_ips.is_cluster_network_subnet_reserved_ips_instance", "reserved_ips.0.resource_type"),
				),
			},
		},
	})
}

func TestAccIBMIsClusterNetworkSubnetReservedIpsDataSourceAllArgs(t *testing.T) {
	clusterNetworkSubnetReservedIPClusterNetworkID := fmt.Sprintf("tf_cluster_network_id_%d", acctest.RandIntRange(10, 100))
	clusterNetworkSubnetReservedIPClusterNetworkSubnetID := fmt.Sprintf("tf_cluster_network_subnet_id_%d", acctest.RandIntRange(10, 100))
	clusterNetworkSubnetReservedIPAddress := fmt.Sprintf("tf_address_%d", acctest.RandIntRange(10, 100))
	clusterNetworkSubnetReservedIPName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsClusterNetworkSubnetReservedIpsDataSourceConfig(clusterNetworkSubnetReservedIPClusterNetworkID, clusterNetworkSubnetReservedIPClusterNetworkSubnetID, clusterNetworkSubnetReservedIPAddress, clusterNetworkSubnetReservedIPName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet_reserved_ips.is_cluster_network_subnet_reserved_ips_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet_reserved_ips.is_cluster_network_subnet_reserved_ips_instance", "cluster_network_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet_reserved_ips.is_cluster_network_subnet_reserved_ips_instance", "cluster_network_subnet_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet_reserved_ips.is_cluster_network_subnet_reserved_ips_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet_reserved_ips.is_cluster_network_subnet_reserved_ips_instance", "sort"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet_reserved_ips.is_cluster_network_subnet_reserved_ips_instance", "reserved_ips.#"),
					resource.TestCheckResourceAttr("data.ibm_is_cluster_network_subnet_reserved_ips.is_cluster_network_subnet_reserved_ips_instance", "reserved_ips.0.address", clusterNetworkSubnetReservedIPAddress),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet_reserved_ips.is_cluster_network_subnet_reserved_ips_instance", "reserved_ips.0.auto_delete"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet_reserved_ips.is_cluster_network_subnet_reserved_ips_instance", "reserved_ips.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet_reserved_ips.is_cluster_network_subnet_reserved_ips_instance", "reserved_ips.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet_reserved_ips.is_cluster_network_subnet_reserved_ips_instance", "reserved_ips.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet_reserved_ips.is_cluster_network_subnet_reserved_ips_instance", "reserved_ips.0.lifecycle_state"),
					resource.TestCheckResourceAttr("data.ibm_is_cluster_network_subnet_reserved_ips.is_cluster_network_subnet_reserved_ips_instance", "reserved_ips.0.name", clusterNetworkSubnetReservedIPName),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet_reserved_ips.is_cluster_network_subnet_reserved_ips_instance", "reserved_ips.0.owner"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_subnet_reserved_ips.is_cluster_network_subnet_reserved_ips_instance", "reserved_ips.0.resource_type"),
				),
			},
		},
	})
}

func testAccCheckIBMIsClusterNetworkSubnetReservedIpsDataSourceConfigBasic(clusterNetworkSubnetReservedIPClusterNetworkID string, clusterNetworkSubnetReservedIPClusterNetworkSubnetID string) string {
	return fmt.Sprintf(`

		data "ibm_is_cluster_network_subnet_reserved_ips" "is_cluster_network_subnet_reserved_ips_instance" {
			cluster_network_id        = "02c7-10274052-f495-4920-a67f-870eb3b87003"
			cluster_network_subnet_id = "02c7-27ca8206-f73b-4d8f-b996-913711b98ff0"
		}
	`)
}

func testAccCheckIBMIsClusterNetworkSubnetReservedIpsDataSourceConfig(clusterNetworkSubnetReservedIPClusterNetworkID string, clusterNetworkSubnetReservedIPClusterNetworkSubnetID string, clusterNetworkSubnetReservedIPAddress string, clusterNetworkSubnetReservedIPName string) string {
	return fmt.Sprintf(`
		resource "ibm_is_cluster_network_subnet_reserved_ip" "is_cluster_network_subnet_reserved_ip_instance" {
			cluster_network_id = "%s"
			cluster_network_subnet_id = "%s"
			address = "%s"
			name = "%s"
		}

		data "ibm_is_cluster_network_subnet_reserved_ips" "is_cluster_network_subnet_reserved_ips_instance" {
			cluster_network_id = ibm_is_cluster_network_subnet_reserved_ip.is_cluster_network_subnet_reserved_ip_instance.cluster_network_id
			cluster_network_subnet_id = ibm_is_cluster_network_subnet_reserved_ip.is_cluster_network_subnet_reserved_ip_instance.cluster_network_subnet_id
			name = ibm_is_cluster_network_subnet_reserved_ip.is_cluster_network_subnet_reserved_ip_instance.name
			sort = "name"
		}
	`, clusterNetworkSubnetReservedIPClusterNetworkID, clusterNetworkSubnetReservedIPClusterNetworkSubnetID, clusterNetworkSubnetReservedIPAddress, clusterNetworkSubnetReservedIPName)
}

func TestDataSourceIBMIsClusterNetworkSubnetReservedIpsClusterNetworkSubnetReservedIPToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		clusterNetworkSubnetReservedIPLifecycleReasonModel := make(map[string]interface{})
		clusterNetworkSubnetReservedIPLifecycleReasonModel["code"] = "resource_suspended_by_provider"
		clusterNetworkSubnetReservedIPLifecycleReasonModel["message"] = "The resource has been suspended. Contact IBM support with the CRN for next steps."
		clusterNetworkSubnetReservedIPLifecycleReasonModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#resource-suspension"

		deletedModel := make(map[string]interface{})
		deletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		clusterNetworkSubnetReservedIPTargetModel := make(map[string]interface{})
		clusterNetworkSubnetReservedIPTargetModel["deleted"] = []map[string]interface{}{deletedModel}
		clusterNetworkSubnetReservedIPTargetModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/cluster_networks/0717-da0df18c-7598-4633-a648-fdaac28a5573/interfaces/0717-ffc092f7-5d02-4b93-ab69-26860529b9fb"
		clusterNetworkSubnetReservedIPTargetModel["id"] = "0717-ffc092f7-5d02-4b93-ab69-26860529b9fb"
		clusterNetworkSubnetReservedIPTargetModel["name"] = "my-cluster-network-interface"
		clusterNetworkSubnetReservedIPTargetModel["resource_type"] = "cluster_network_interface"

		model := make(map[string]interface{})
		model["address"] = "10.1.0.6"
		model["auto_delete"] = false
		model["created_at"] = "2019-01-01T12:00:00.000Z"
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/cluster_networks/0717-da0df18c-7598-4633-a648-fdaac28a5573/subnets/0717-7931845c-65c4-4b0a-80cd-7d9c1a6d7930/reserved_ips/6d353a0f-aeb1-4ae1-832e-1110d10981bb"
		model["id"] = "6d353a0f-aeb1-4ae1-832e-1110d10981bb"
		model["lifecycle_reasons"] = []map[string]interface{}{clusterNetworkSubnetReservedIPLifecycleReasonModel}
		model["lifecycle_state"] = "stable"
		model["name"] = "my-cluster-network-subnet-reserved-ip"
		model["owner"] = "user"
		model["resource_type"] = "cluster_network_subnet_reserved_ip"
		model["target"] = []map[string]interface{}{clusterNetworkSubnetReservedIPTargetModel}

		assert.Equal(t, result, model)
	}

	clusterNetworkSubnetReservedIPLifecycleReasonModel := new(vpcv1.ClusterNetworkSubnetReservedIPLifecycleReason)
	clusterNetworkSubnetReservedIPLifecycleReasonModel.Code = core.StringPtr("resource_suspended_by_provider")
	clusterNetworkSubnetReservedIPLifecycleReasonModel.Message = core.StringPtr("The resource has been suspended. Contact IBM support with the CRN for next steps.")
	clusterNetworkSubnetReservedIPLifecycleReasonModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#resource-suspension")

	deletedModel := new(vpcv1.Deleted)
	deletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	clusterNetworkSubnetReservedIPTargetModel := new(vpcv1.ClusterNetworkSubnetReservedIPTargetClusterNetworkInterfaceReferenceClusterNetworkSubnetReservedIPTargetContext)
	clusterNetworkSubnetReservedIPTargetModel.Deleted = deletedModel
	clusterNetworkSubnetReservedIPTargetModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/cluster_networks/0717-da0df18c-7598-4633-a648-fdaac28a5573/interfaces/0717-ffc092f7-5d02-4b93-ab69-26860529b9fb")
	clusterNetworkSubnetReservedIPTargetModel.ID = core.StringPtr("0717-ffc092f7-5d02-4b93-ab69-26860529b9fb")
	clusterNetworkSubnetReservedIPTargetModel.Name = core.StringPtr("my-cluster-network-interface")
	clusterNetworkSubnetReservedIPTargetModel.ResourceType = core.StringPtr("cluster_network_interface")

	model := new(vpcv1.ClusterNetworkSubnetReservedIP)
	model.Address = core.StringPtr("10.1.0.6")
	model.AutoDelete = core.BoolPtr(false)
	model.CreatedAt = CreateMockDateTime("2019-01-01T12:00:00.000Z")
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/cluster_networks/0717-da0df18c-7598-4633-a648-fdaac28a5573/subnets/0717-7931845c-65c4-4b0a-80cd-7d9c1a6d7930/reserved_ips/6d353a0f-aeb1-4ae1-832e-1110d10981bb")
	model.ID = core.StringPtr("6d353a0f-aeb1-4ae1-832e-1110d10981bb")
	model.LifecycleReasons = []vpcv1.ClusterNetworkSubnetReservedIPLifecycleReason{*clusterNetworkSubnetReservedIPLifecycleReasonModel}
	model.LifecycleState = core.StringPtr("stable")
	model.Name = core.StringPtr("my-cluster-network-subnet-reserved-ip")
	model.Owner = core.StringPtr("user")
	model.ResourceType = core.StringPtr("cluster_network_subnet_reserved_ip")
	model.Target = clusterNetworkSubnetReservedIPTargetModel

	result, err := vpc.DataSourceIBMIsClusterNetworkSubnetReservedIpsClusterNetworkSubnetReservedIPToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsClusterNetworkSubnetReservedIpsClusterNetworkSubnetReservedIPLifecycleReasonToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["code"] = "resource_suspended_by_provider"
		model["message"] = "The resource has been suspended. Contact IBM support with the CRN for next steps."
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#resource-suspension"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.ClusterNetworkSubnetReservedIPLifecycleReason)
	model.Code = core.StringPtr("resource_suspended_by_provider")
	model.Message = core.StringPtr("The resource has been suspended. Contact IBM support with the CRN for next steps.")
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#resource-suspension")

	result, err := vpc.DataSourceIBMIsClusterNetworkSubnetReservedIpsClusterNetworkSubnetReservedIPLifecycleReasonToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsClusterNetworkSubnetReservedIpsClusterNetworkSubnetReservedIPTargetToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		deletedModel := make(map[string]interface{})
		deletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		model := make(map[string]interface{})
		model["deleted"] = []map[string]interface{}{deletedModel}
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/cluster_networks/0717-da0df18c-7598-4633-a648-fdaac28a5573/interfaces/0717-ffc092f7-5d02-4b93-ab69-26860529b9fb"
		model["id"] = "0717-ffc092f7-5d02-4b93-ab69-26860529b9fb"
		model["name"] = "my-cluster-network-interface"
		model["resource_type"] = "cluster_network_interface"

		assert.Equal(t, result, model)
	}

	deletedModel := new(vpcv1.Deleted)
	deletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	model := new(vpcv1.ClusterNetworkSubnetReservedIPTarget)
	model.Deleted = deletedModel
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/cluster_networks/0717-da0df18c-7598-4633-a648-fdaac28a5573/interfaces/0717-ffc092f7-5d02-4b93-ab69-26860529b9fb")
	model.ID = core.StringPtr("0717-ffc092f7-5d02-4b93-ab69-26860529b9fb")
	model.Name = core.StringPtr("my-cluster-network-interface")
	model.ResourceType = core.StringPtr("cluster_network_interface")

	result, err := vpc.DataSourceIBMIsClusterNetworkSubnetReservedIpsClusterNetworkSubnetReservedIPTargetToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsClusterNetworkSubnetReservedIpsDeletedToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.Deleted)
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	result, err := vpc.DataSourceIBMIsClusterNetworkSubnetReservedIpsDeletedToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsClusterNetworkSubnetReservedIpsClusterNetworkSubnetReservedIPTargetClusterNetworkInterfaceReferenceClusterNetworkSubnetReservedIPTargetContextToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		deletedModel := make(map[string]interface{})
		deletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		model := make(map[string]interface{})
		model["deleted"] = []map[string]interface{}{deletedModel}
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/cluster_networks/0717-da0df18c-7598-4633-a648-fdaac28a5573/interfaces/0717-ffc092f7-5d02-4b93-ab69-26860529b9fb"
		model["id"] = "0717-ffc092f7-5d02-4b93-ab69-26860529b9fb"
		model["name"] = "my-cluster-network-interface"
		model["resource_type"] = "cluster_network_interface"

		assert.Equal(t, result, model)
	}

	deletedModel := new(vpcv1.Deleted)
	deletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	model := new(vpcv1.ClusterNetworkSubnetReservedIPTargetClusterNetworkInterfaceReferenceClusterNetworkSubnetReservedIPTargetContext)
	model.Deleted = deletedModel
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/cluster_networks/0717-da0df18c-7598-4633-a648-fdaac28a5573/interfaces/0717-ffc092f7-5d02-4b93-ab69-26860529b9fb")
	model.ID = core.StringPtr("0717-ffc092f7-5d02-4b93-ab69-26860529b9fb")
	model.Name = core.StringPtr("my-cluster-network-interface")
	model.ResourceType = core.StringPtr("cluster_network_interface")

	result, err := vpc.DataSourceIBMIsClusterNetworkSubnetReservedIpsClusterNetworkSubnetReservedIPTargetClusterNetworkInterfaceReferenceClusterNetworkSubnetReservedIPTargetContextToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
