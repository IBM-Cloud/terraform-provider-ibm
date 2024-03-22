// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMContainerVPCClusterDataSource_basic(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-cluster-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVPCClusterDataSource(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_container_vpc_cluster.testacc_ds_cluster", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_container_cluster_config.testacc_ds_cluster", "id"),
				),
			},
		},
	})
}

func TestAccIBMContainerVPCClusterDataSource_DedicatedHost(t *testing.T) {
	clusterName := fmt.Sprintf("tf-vpc-cluster-%d", acctest.RandIntRange(10, 100))
	hostPoolID := acc.HostPoolID
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVPCClusterDataSourceDedicatedHost(
					clusterName,
					acc.IksClusterVpcID,
					"bx2d.4x16",
					acc.IksClusterSubnetID,
					acc.IksClusterResourceGroupID,
					hostPoolID,
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_container_vpc_cluster.testacc_cluster_dedicatedhost", "worker_pools.0.host_pool_id", hostPoolID),
				),
			},
		},
	})
}

func testAccCheckIBMContainerVPCClusterDataSource(name string) string {
	return testAccCheckIBMContainerVpcClusterBasic(name) + `
data "ibm_container_vpc_cluster" "testacc_ds_cluster" {
    cluster_name_id = ibm_container_vpc_cluster.cluster.id
}
data "ibm_container_cluster_config" "testacc_ds_cluster" {
	cluster_name_id = ibm_container_vpc_cluster.cluster.id
  }
`
}

func testAccCheckIBMContainerVPCClusterDataSourceDedicatedHost(name, vpcID, flavor, subnetID, rgroupID, hostpoolID string) string {
	return testAccCheckIBMContainerVpcClusterDedicatedHostSetting(name, vpcID, flavor, subnetID, rgroupID, hostpoolID) + `
data "ibm_container_vpc_cluster" "testacc_cluster_dedicatedhost" {
    name = ibm_container_vpc_cluster.testacc_dhost_vpc_cluster.name
}
`
}

func testAccCheckIBMContainerVPCClusterDatasourceEnvvar(name string) string {
	return testAccCheckIBMContainerVpcClusterEnvvar(name) + `
data "ibm_container_vpc_cluster" "testacc_ds_cluster" {
     cluster_name_id = ibm_container_vpc_cluster.cluster.id
}
`
}

func TestAccIBMContainerVPCClusterDataSourceEnvvar(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-cluster-%d", acctest.RandIntRange(10, 100))
	testChecks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttrSet(
			"data.ibm_container_vpc_cluster.testacc_ds_cluster", "id"),
		resource.TestCheckResourceAttr(
			"data.ibm_container_vpc_cluster.testacc_ds_cluster", "worker_pools.#", "1"),
	}
	if acc.WorkerPoolSecondaryStorage != "" {
		testChecks = append(testChecks, resource.TestCheckResourceAttr(
			"data.ibm_container_vpc_cluster.testacc_ds_cluster", "worker_pools.0.secondary_storage.0.name", acc.WorkerPoolSecondaryStorage))
	}
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVPCClusterDatasourceEnvvar(name),
				Check:  resource.ComposeTestCheckFunc(testChecks...),
			},
		},
	})
}

// You need to set up env vars:
// export IBM_CLUSTER_VPC_ID
// export IBM_CLUSTER_VPC_RESOURCE_GROUP_ID
// export IBM_CLUSTER_VPC_SUBNET_ID
func TestAccIBMContainerVPCClusterDataSourceIngressConfig(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-cluster-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVPCClusterDatasourceIngressConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.ibm_container_vpc_cluster.cluster", "ingress_config.#", "1"),
					resource.TestCheckResourceAttr(
						"data.ibm_container_vpc_cluster.cluster", "ingress_config.0.ingress_status_report.#", "1"),
					resource.TestCheckResourceAttr(
						"data.ibm_container_vpc_cluster.cluster", "ingress_config.0.ingress_status_report.0.enabled", "true"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_container_vpc_cluster.cluster", "ingress_config.0.ingress_status_report.0.ingress_status"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_container_vpc_cluster.cluster", "ingress_config.0.ingress_status_report.0.message"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_container_vpc_cluster.cluster", "ingress_config.0.ingress_status_report.0.general_ingress_component_status.#"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_container_vpc_cluster.cluster", "ingress_config.0.ingress_status_report.0.alb_status.#"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_container_vpc_cluster.cluster", "ingress_config.0.ingress_status_report.0.secret_status.#"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_container_vpc_cluster.cluster", "ingress_config.0.ingress_status_report.0.subdomain_status.#"),
					resource.TestCheckResourceAttr(
						"data.ibm_container_vpc_cluster.cluster", "ingress_config.0.ingress_health_checker_enabled", "true"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerVPCClusterDatasourceIngressConfig(name string) string {
	return testAccConfigCheckIBMContainerVpcClusterIngressConfig(name) + `
	data "ibm_container_vpc_cluster" "cluster" {
	     cluster_name_id = ibm_container_vpc_cluster.cluster.id
	}
	`
}
