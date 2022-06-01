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
