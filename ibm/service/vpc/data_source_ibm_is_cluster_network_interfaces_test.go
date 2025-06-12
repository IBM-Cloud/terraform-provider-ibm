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

func TestAccIBMIsClusterNetworkInterfacesDataSourceBasic(t *testing.T) {
	clusterNetworkInterfaceClusterNetworkID := fmt.Sprintf("tf_cluster_network_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsClusterNetworkInterfacesDataSourceConfigBasic(clusterNetworkInterfaceClusterNetworkID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interfaces.is_cluster_network_interfaces_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interfaces.is_cluster_network_interfaces_instance", "cluster_network_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interfaces.is_cluster_network_interfaces_instance", "interfaces.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interfaces.is_cluster_network_interfaces_instance", "interfaces.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interfaces.is_cluster_network_interfaces_instance", "interfaces.0.allow_ip_spoofing"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interfaces.is_cluster_network_interfaces_instance", "interfaces.0.auto_delete"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interfaces.is_cluster_network_interfaces_instance", "interfaces.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interfaces.is_cluster_network_interfaces_instance", "interfaces.0.enable_infrastructure_nat"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interfaces.is_cluster_network_interfaces_instance", "interfaces.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interfaces.is_cluster_network_interfaces_instance", "interfaces.0.lifecycle_reasons.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interfaces.is_cluster_network_interfaces_instance", "interfaces.0.lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interfaces.is_cluster_network_interfaces_instance", "interfaces.0.mac_address"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interfaces.is_cluster_network_interfaces_instance", "interfaces.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interfaces.is_cluster_network_interfaces_instance", "interfaces.0.primary_ip.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interfaces.is_cluster_network_interfaces_instance", "interfaces.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interfaces.is_cluster_network_interfaces_instance", "interfaces.0.subnet.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interfaces.is_cluster_network_interfaces_instance", "interfaces.0.target.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interfaces.is_cluster_network_interfaces_instance", "interfaces.0.vpc.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interfaces.is_cluster_network_interfaces_instance", "interfaces.0.zone.#"),
				),
			},
		},
	})
}

func TestAccIBMIsClusterNetworkInterfacesDataSourceAllArgs(t *testing.T) {
	clusterNetworkInterfaceClusterNetworkID := fmt.Sprintf("tf_cluster_network_id_%d", acctest.RandIntRange(10, 100))
	clusterNetworkInterfaceName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsClusterNetworkInterfacesDataSourceConfig(clusterNetworkInterfaceClusterNetworkID, clusterNetworkInterfaceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interfaces.is_cluster_network_interfaces_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interfaces.is_cluster_network_interfaces_instance", "cluster_network_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interfaces.is_cluster_network_interfaces_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interfaces.is_cluster_network_interfaces_instance", "sort"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interfaces.is_cluster_network_interfaces_instance", "interfaces.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interfaces.is_cluster_network_interfaces_instance", "interfaces.0.allow_ip_spoofing"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interfaces.is_cluster_network_interfaces_instance", "interfaces.0.auto_delete"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interfaces.is_cluster_network_interfaces_instance", "interfaces.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interfaces.is_cluster_network_interfaces_instance", "interfaces.0.enable_infrastructure_nat"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interfaces.is_cluster_network_interfaces_instance", "interfaces.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interfaces.is_cluster_network_interfaces_instance", "interfaces.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interfaces.is_cluster_network_interfaces_instance", "interfaces.0.lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interfaces.is_cluster_network_interfaces_instance", "interfaces.0.mac_address"),
					resource.TestCheckResourceAttr("data.ibm_is_cluster_network_interfaces.is_cluster_network_interfaces_instance", "interfaces.0.name", clusterNetworkInterfaceName),
					// resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interfaces.is_cluster_network_interfaces_instance", "interfaces.0.protocol_state_filtering_mode"),
					resource.TestCheckResourceAttrSet("data.ibm_is_cluster_network_interfaces.is_cluster_network_interfaces_instance", "interfaces.0.resource_type"),
				),
			},
		},
	})
}

func testAccCheckIBMIsClusterNetworkInterfacesDataSourceConfigBasic(clusterNetworkInterfaceClusterNetworkID string) string {
	return fmt.Sprintf(`

		data "ibm_is_cluster_network_interfaces" "is_cluster_network_interfaces_instance" {
			cluster_network_id = "02c7-10274052-f495-4920-a67f-870eb3b87003"
		}
	`)
}

func testAccCheckIBMIsClusterNetworkInterfacesDataSourceConfig(clusterNetworkInterfaceClusterNetworkID string, clusterNetworkInterfaceName string) string {
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

		data "ibm_is_cluster_network_interfaces" "is_cluster_network_interfaces_instance" {
			cluster_network_id = ibm_is_cluster_network_interface.is_cluster_network_interface_instance.cluster_network_id
			name = ibm_is_cluster_network_interface.is_cluster_network_interface_instance.name
			sort = "name"
		}
	`, clusterNetworkInterfaceClusterNetworkID, clusterNetworkInterfaceName)
}

func TestDataSourceIBMIsClusterNetworkInterfacesClusterNetworkInterfaceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		clusterNetworkInterfaceLifecycleReasonModel := make(map[string]interface{})
		clusterNetworkInterfaceLifecycleReasonModel["code"] = "resource_suspended_by_provider"
		clusterNetworkInterfaceLifecycleReasonModel["message"] = "The resource has been suspended. Contact IBM support with the CRN for next steps."
		clusterNetworkInterfaceLifecycleReasonModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#resource-suspension"

		deletedModel := make(map[string]interface{})
		deletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		clusterNetworkSubnetReservedIPReferenceModel := make(map[string]interface{})
		clusterNetworkSubnetReservedIPReferenceModel["address"] = "10.1.0.6"
		clusterNetworkSubnetReservedIPReferenceModel["deleted"] = []map[string]interface{}{deletedModel}
		clusterNetworkSubnetReservedIPReferenceModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/cluster_networks/0717-da0df18c-7598-4633-a648-fdaac28a5573/subnets/0717-7931845c-65c4-4b0a-80cd-7d9c1a6d7930/reserved_ips/63341ffa-1037-4b50-be40-676e3e9ac0c7"
		clusterNetworkSubnetReservedIPReferenceModel["id"] = "63341ffa-1037-4b50-be40-676e3e9ac0c7"
		clusterNetworkSubnetReservedIPReferenceModel["name"] = "my-cluster-network-subnet-reserved-ip"
		clusterNetworkSubnetReservedIPReferenceModel["resource_type"] = "cluster_network_subnet_reserved_ip"

		clusterNetworkSubnetReferenceModel := make(map[string]interface{})
		clusterNetworkSubnetReferenceModel["deleted"] = []map[string]interface{}{deletedModel}
		clusterNetworkSubnetReferenceModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/cluster_networks/0717-da0df18c-7598-4633-a648-fdaac28a5573/subnets/0717-7931845c-65c4-4b0a-80cd-7d9c1a6d7930"
		clusterNetworkSubnetReferenceModel["id"] = "0717-7931845c-65c4-4b0a-80cd-7d9c1a6d7930"
		clusterNetworkSubnetReferenceModel["name"] = "my-cluster-network-subnet"
		clusterNetworkSubnetReferenceModel["resource_type"] = "cluster_network_subnet"

		clusterNetworkInterfaceTargetModel := make(map[string]interface{})
		clusterNetworkInterfaceTargetModel["href"] = "https://us-south.iaas.cloud.ibm.com/v1/instances/0717_e21b7391-2ca2-4ab5-84a8-b92157a633b0/cluster_network_attachments/0717-fb880975-db45-4459-8548-64e3995ac213"
		clusterNetworkInterfaceTargetModel["id"] = "0717-fb880975-db45-4459-8548-64e3995ac213"
		clusterNetworkInterfaceTargetModel["name"] = "my-instance-network-attachment"
		clusterNetworkInterfaceTargetModel["resource_type"] = "instance_cluster_network_attachment"

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
		model["allow_ip_spoofing"] = true
		model["auto_delete"] = false
		model["created_at"] = "2019-01-01T12:00:00.000Z"
		model["enable_infrastructure_nat"] = false
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/cluster_networks/0717-da0df18c-7598-4633-a648-fdaac28a5573/interfaces/0717-ffc092f7-5d02-4b93-ab69-26860529b9fb"
		model["id"] = "0717-ffc092f7-5d02-4b93-ab69-26860529b9fb"
		model["lifecycle_reasons"] = []map[string]interface{}{clusterNetworkInterfaceLifecycleReasonModel}
		model["lifecycle_state"] = "stable"
		model["mac_address"] = "02:00:4D:45:45:4D"
		model["name"] = "my-cluster-network-interface"
		model["primary_ip"] = []map[string]interface{}{clusterNetworkSubnetReservedIPReferenceModel}
		// model["protocol_state_filtering_mode"] = "enabled"
		model["resource_type"] = "cluster_network_interface"
		model["subnet"] = []map[string]interface{}{clusterNetworkSubnetReferenceModel}
		model["target"] = []map[string]interface{}{clusterNetworkInterfaceTargetModel}
		model["vpc"] = []map[string]interface{}{vpcReferenceModel}
		model["zone"] = []map[string]interface{}{zoneReferenceModel}

		assert.Equal(t, result, model)
	}

	clusterNetworkInterfaceLifecycleReasonModel := new(vpcv1.ClusterNetworkInterfaceLifecycleReason)
	clusterNetworkInterfaceLifecycleReasonModel.Code = core.StringPtr("resource_suspended_by_provider")
	clusterNetworkInterfaceLifecycleReasonModel.Message = core.StringPtr("The resource has been suspended. Contact IBM support with the CRN for next steps.")
	clusterNetworkInterfaceLifecycleReasonModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#resource-suspension")

	deletedModel := new(vpcv1.Deleted)
	deletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	clusterNetworkSubnetReservedIPReferenceModel := new(vpcv1.ClusterNetworkSubnetReservedIPReference)
	clusterNetworkSubnetReservedIPReferenceModel.Address = core.StringPtr("10.1.0.6")
	clusterNetworkSubnetReservedIPReferenceModel.Deleted = deletedModel
	clusterNetworkSubnetReservedIPReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/cluster_networks/0717-da0df18c-7598-4633-a648-fdaac28a5573/subnets/0717-7931845c-65c4-4b0a-80cd-7d9c1a6d7930/reserved_ips/63341ffa-1037-4b50-be40-676e3e9ac0c7")
	clusterNetworkSubnetReservedIPReferenceModel.ID = core.StringPtr("63341ffa-1037-4b50-be40-676e3e9ac0c7")
	clusterNetworkSubnetReservedIPReferenceModel.Name = core.StringPtr("my-cluster-network-subnet-reserved-ip")
	clusterNetworkSubnetReservedIPReferenceModel.ResourceType = core.StringPtr("cluster_network_subnet_reserved_ip")

	clusterNetworkSubnetReferenceModel := new(vpcv1.ClusterNetworkSubnetReference)
	clusterNetworkSubnetReferenceModel.Deleted = deletedModel
	clusterNetworkSubnetReferenceModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/cluster_networks/0717-da0df18c-7598-4633-a648-fdaac28a5573/subnets/0717-7931845c-65c4-4b0a-80cd-7d9c1a6d7930")
	clusterNetworkSubnetReferenceModel.ID = core.StringPtr("0717-7931845c-65c4-4b0a-80cd-7d9c1a6d7930")
	clusterNetworkSubnetReferenceModel.Name = core.StringPtr("my-cluster-network-subnet")
	clusterNetworkSubnetReferenceModel.ResourceType = core.StringPtr("cluster_network_subnet")

	clusterNetworkInterfaceTargetModel := new(vpcv1.ClusterNetworkInterfaceTargetInstanceClusterNetworkAttachmentReferenceClusterNetworkInterfaceContext)
	clusterNetworkInterfaceTargetModel.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/instances/0717_e21b7391-2ca2-4ab5-84a8-b92157a633b0/cluster_network_attachments/0717-fb880975-db45-4459-8548-64e3995ac213")
	clusterNetworkInterfaceTargetModel.ID = core.StringPtr("0717-fb880975-db45-4459-8548-64e3995ac213")
	clusterNetworkInterfaceTargetModel.Name = core.StringPtr("my-instance-network-attachment")
	clusterNetworkInterfaceTargetModel.ResourceType = core.StringPtr("instance_cluster_network_attachment")

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

	model := new(vpcv1.ClusterNetworkInterface)
	model.AllowIPSpoofing = core.BoolPtr(true)
	model.AutoDelete = core.BoolPtr(false)
	model.CreatedAt = CreateMockDateTime("2019-01-01T12:00:00.000Z")
	model.EnableInfrastructureNat = core.BoolPtr(false)
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/cluster_networks/0717-da0df18c-7598-4633-a648-fdaac28a5573/interfaces/0717-ffc092f7-5d02-4b93-ab69-26860529b9fb")
	model.ID = core.StringPtr("0717-ffc092f7-5d02-4b93-ab69-26860529b9fb")
	model.LifecycleReasons = []vpcv1.ClusterNetworkInterfaceLifecycleReason{*clusterNetworkInterfaceLifecycleReasonModel}
	model.LifecycleState = core.StringPtr("stable")
	model.MacAddress = core.StringPtr("02:00:4D:45:45:4D")
	model.Name = core.StringPtr("my-cluster-network-interface")
	model.PrimaryIP = clusterNetworkSubnetReservedIPReferenceModel
	// model.ProtocolStateFilteringMode = core.StringPtr("enabled")
	model.ResourceType = core.StringPtr("cluster_network_interface")
	model.Subnet = clusterNetworkSubnetReferenceModel
	model.Target = clusterNetworkInterfaceTargetModel
	model.VPC = vpcReferenceModel
	model.Zone = zoneReferenceModel

	result, err := vpc.DataSourceIBMIsClusterNetworkInterfacesClusterNetworkInterfaceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsClusterNetworkInterfacesClusterNetworkInterfaceLifecycleReasonToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsClusterNetworkInterfacesClusterNetworkInterfaceLifecycleReasonToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsClusterNetworkInterfacesClusterNetworkSubnetReservedIPReferenceToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsClusterNetworkInterfacesClusterNetworkSubnetReservedIPReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsClusterNetworkInterfacesDeletedToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.Deleted)
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	result, err := vpc.DataSourceIBMIsClusterNetworkInterfacesDeletedToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsClusterNetworkInterfacesClusterNetworkSubnetReferenceToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsClusterNetworkInterfacesClusterNetworkSubnetReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsClusterNetworkInterfacesClusterNetworkInterfaceTargetToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsClusterNetworkInterfacesClusterNetworkInterfaceTargetToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsClusterNetworkInterfacesClusterNetworkInterfaceTargetInstanceClusterNetworkAttachmentReferenceClusterNetworkInterfaceContextToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsClusterNetworkInterfacesClusterNetworkInterfaceTargetInstanceClusterNetworkAttachmentReferenceClusterNetworkInterfaceContextToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsClusterNetworkInterfacesVPCReferenceToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsClusterNetworkInterfacesVPCReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsClusterNetworkInterfacesZoneReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1"
		model["name"] = "us-south-1"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.ZoneReference)
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1")
	model.Name = core.StringPtr("us-south-1")

	result, err := vpc.DataSourceIBMIsClusterNetworkInterfacesZoneReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
