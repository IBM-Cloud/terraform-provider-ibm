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
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMIsClusterNetworkInterfaceDataSourceBasic(t *testing.T) {
	clusterNetworkInterfaceClusterNetworkID := fmt.Sprintf("tf_cluster_network_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsClusterNetworkInterfaceDataSourceConfigBasic(clusterNetworkInterfaceClusterNetworkID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "cluster_network_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "cluster_network_interface_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "allow_ip_spoofing"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "auto_delete"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "enable_infrastructure_nat"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "lifecycle_reasons.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "mac_address"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "primary_ip.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "subnet.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "vpc.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "zone.#"),
				),
			},
		},
	})
}

func TestAccIBMIsClusterNetworkInterfaceDataSourceAllArgs(t *testing.T) {
	clusterNetworkInterfaceClusterNetworkID := fmt.Sprintf("tf_cluster_network_id_%d", acctest.RandIntRange(10, 100))
	clusterNetworkInterfaceName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsClusterNetworkInterfaceDataSourceConfig(clusterNetworkInterfaceClusterNetworkID, clusterNetworkInterfaceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "cluster_network_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "cluster_network_interface_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "allow_ip_spoofing"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "auto_delete"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "enable_infrastructure_nat"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "lifecycle_reasons.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "lifecycle_reasons.0.code"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "lifecycle_reasons.0.message"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "lifecycle_reasons.0.more_info"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "mac_address"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "primary_ip.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "protocol_state_filtering_mode"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "subnet.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "target.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "vpc.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "zone.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsClusterNetworkInterfaceDataSourceConfigBasic(clusterNetworkInterfaceClusterNetworkID string) string {
	return fmt.Sprintf(`
		data "ibm_is_cluster_network_interface" "is_cluster_network_interface_instance" {
			cluster_network_id = "02c7-10274052-f495-4920-a67f-870eb3b87003"
			cluster_network_interface_id = "02c7-fcc6bdf2-56f5-40ad-abce-42950007857c"
		}
	`)
}

func testAccCheckIBMIsClusterNetworkInterfaceDataSourceConfig(clusterNetworkInterfaceClusterNetworkID string, clusterNetworkInterfaceName string) string {
	return fmt.Sprintf(`
		resource "ibm_is_cluster_network_interface" "is_cluster_network_interface_instance" {
			cluster_network_id = "%s"
			name = "%s"
			primary_ip {
				address = "10.1.0.6"
				deleted {
					more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
				}
				href = "https://us-south.iaas.cloud.ibm.com/v1/cluster_networks/0717-da0df18c-7598-4633-a648-fdaac28a5573/subnets/0717-7931845c-65c4-4b0a-80cd-7d9c1a6d7930/reserved_ips/6d353a0f-aeb1-4ae1-832e-1110d10981bb"
				id = "6d353a0f-aeb1-4ae1-832e-1110d10981bb"
				name = "my-cluster-network-subnet-reserved-ip"
				resource_type = "cluster_network_subnet_reserved_ip"
			}
			subnet {
				deleted {
					more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
				}
				href = "https://us-south.iaas.cloud.ibm.com/v1/cluster_networks/0717-da0df18c-7598-4633-a648-fdaac28a5573/subnets/0717-7931845c-65c4-4b0a-80cd-7d9c1a6d7930"
				id = "0717-7931845c-65c4-4b0a-80cd-7d9c1a6d7930"
				name = "my-cluster-network-subnet"
				resource_type = "cluster_network_subnet"
			}
		}

		data "ibm_is_cluster_network_interface" "is_cluster_network_interface_instance" {
			cluster_network_id = ibm_is_cluster_network_interface.is_cluster_network_interface_instance.cluster_network_id
			cluster_network_interface_id = ibm_is_cluster_network_interface.is_cluster_network_interface_instance.cluster_network_interface_id
		}
	`, clusterNetworkInterfaceClusterNetworkID, clusterNetworkInterfaceName)
}

func TestDataSourceIBMIsClusterNetworkInterfaceClusterNetworkInterfaceLifecycleReasonToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["code"] = "resource_suspended_by_provider"
		model["message"] = "The resource has been suspended. Contact IBM support with the CRN for next steps."
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#resource-suspension"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.ClusterNetworkInterfaceLifecycleReason)
	model.Code = core.StringPtr("resource_suspended_by_provider")
	model.Message = core.StringPtr("The resource has been suspended. Contact IBM support with the CRN for next steps.")
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#resource-suspension")

	result, err := vpc.DataSourceIBMIsClusterNetworkInterfaceClusterNetworkInterfaceLifecycleReasonToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsClusterNetworkInterfaceClusterNetworkSubnetReservedIPReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		deletedModel := make(map[string]interface{})
		deletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		model := make(map[string]interface{})
		model["address"] = "10.1.0.6"
		model["deleted"] = []map[string]interface{}{deletedModel}
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/cluster_networks/0717-da0df18c-7598-4633-a648-fdaac28a5573/subnets/0717-7931845c-65c4-4b0a-80cd-7d9c1a6d7930/reserved_ips/6d353a0f-aeb1-4ae1-832e-1110d10981bb"
		model["id"] = "6d353a0f-aeb1-4ae1-832e-1110d10981bb"
		model["name"] = "my-cluster-network-subnet-reserved-ip"
		model["resource_type"] = "cluster_network_subnet_reserved_ip"

		assert.Equal(t, result, model)
	}

	deletedModel := new(vpcv1.Deleted)
	deletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	model := new(vpcv1.ClusterNetworkSubnetReservedIPReference)
	model.Address = core.StringPtr("10.1.0.6")
	model.Deleted = deletedModel
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/cluster_networks/0717-da0df18c-7598-4633-a648-fdaac28a5573/subnets/0717-7931845c-65c4-4b0a-80cd-7d9c1a6d7930/reserved_ips/6d353a0f-aeb1-4ae1-832e-1110d10981bb")
	model.ID = core.StringPtr("6d353a0f-aeb1-4ae1-832e-1110d10981bb")
	model.Name = core.StringPtr("my-cluster-network-subnet-reserved-ip")
	model.ResourceType = core.StringPtr("cluster_network_subnet_reserved_ip")

	result, err := vpc.DataSourceIBMIsClusterNetworkInterfaceClusterNetworkSubnetReservedIPReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsClusterNetworkInterfaceDeletedToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.Deleted)
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	result, err := vpc.DataSourceIBMIsClusterNetworkInterfaceDeletedToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsClusterNetworkInterfaceClusterNetworkSubnetReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		deletedModel := make(map[string]interface{})
		deletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		model := make(map[string]interface{})
		model["deleted"] = []map[string]interface{}{deletedModel}
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/cluster_networks/0717-da0df18c-7598-4633-a648-fdaac28a5573/subnets/0717-7931845c-65c4-4b0a-80cd-7d9c1a6d7930"
		model["id"] = "0717-7931845c-65c4-4b0a-80cd-7d9c1a6d7930"
		model["name"] = "my-cluster-network-subnet"
		model["resource_type"] = "cluster_network_subnet"

		assert.Equal(t, result, model)
	}

	deletedModel := new(vpcv1.Deleted)
	deletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	model := new(vpcv1.ClusterNetworkSubnetReference)
	model.Deleted = deletedModel
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/cluster_networks/0717-da0df18c-7598-4633-a648-fdaac28a5573/subnets/0717-7931845c-65c4-4b0a-80cd-7d9c1a6d7930")
	model.ID = core.StringPtr("0717-7931845c-65c4-4b0a-80cd-7d9c1a6d7930")
	model.Name = core.StringPtr("my-cluster-network-subnet")
	model.ResourceType = core.StringPtr("cluster_network_subnet")

	result, err := vpc.DataSourceIBMIsClusterNetworkInterfaceClusterNetworkSubnetReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsClusterNetworkInterfaceClusterNetworkInterfaceTargetToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/instances/0717_e21b7391-2ca2-4ab5-84a8-b92157a633b0/cluster_network_attachments/0717-fb880975-db45-4459-8548-64e3995ac213"
		model["id"] = "0717-fb880975-db45-4459-8548-64e3995ac213"
		model["name"] = "my-instance-network-attachment"
		model["resource_type"] = "instance_cluster_network_attachment"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.ClusterNetworkInterfaceTarget)
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/instances/0717_e21b7391-2ca2-4ab5-84a8-b92157a633b0/cluster_network_attachments/0717-fb880975-db45-4459-8548-64e3995ac213")
	model.ID = core.StringPtr("0717-fb880975-db45-4459-8548-64e3995ac213")
	model.Name = core.StringPtr("my-instance-network-attachment")
	model.ResourceType = core.StringPtr("instance_cluster_network_attachment")

	result, err := vpc.DataSourceIBMIsClusterNetworkInterfaceClusterNetworkInterfaceTargetToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsClusterNetworkInterfaceClusterNetworkInterfaceTargetInstanceClusterNetworkAttachmentReferenceClusterNetworkInterfaceContextToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/instances/0717_e21b7391-2ca2-4ab5-84a8-b92157a633b0/cluster_network_attachments/0717-fb880975-db45-4459-8548-64e3995ac213"
		model["id"] = "0717-fb880975-db45-4459-8548-64e3995ac213"
		model["name"] = "my-instance-network-attachment"
		model["resource_type"] = "instance_cluster_network_attachment"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.ClusterNetworkInterfaceTargetInstanceClusterNetworkAttachmentReferenceClusterNetworkInterfaceContext)
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/instances/0717_e21b7391-2ca2-4ab5-84a8-b92157a633b0/cluster_network_attachments/0717-fb880975-db45-4459-8548-64e3995ac213")
	model.ID = core.StringPtr("0717-fb880975-db45-4459-8548-64e3995ac213")
	model.Name = core.StringPtr("my-instance-network-attachment")
	model.ResourceType = core.StringPtr("instance_cluster_network_attachment")

	result, err := vpc.DataSourceIBMIsClusterNetworkInterfaceClusterNetworkInterfaceTargetInstanceClusterNetworkAttachmentReferenceClusterNetworkInterfaceContextToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsClusterNetworkInterfaceVPCReferenceToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsClusterNetworkInterfaceVPCReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsClusterNetworkInterfaceZoneReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1"
		model["name"] = "us-south-1"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.ZoneReference)
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1")
	model.Name = core.StringPtr("us-south-1")

	result, err := vpc.DataSourceIBMIsClusterNetworkInterfaceZoneReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
