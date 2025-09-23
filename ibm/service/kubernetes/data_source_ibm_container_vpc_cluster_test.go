// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMContainerVPCClusterDataSource_basic(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-cluster-%d", acctest.RandIntRange(10, 100))
	masterNodeReadyClusterScript := testAccCheckIBMContainerVpcClusterBasic(name, "MasterNodeReady") + `
	data "ibm_container_vpc_cluster" "testacc_ds_cluster" {
		name = ibm_container_vpc_cluster.cluster.id
	}
	`
	normalClusterScriptWithConfig := testAccCheckIBMContainerVpcClusterBasic(name, "MasterNodeReady") + `
	data "ibm_container_vpc_cluster" "testacc_ds_cluster" {
		name      = ibm_container_vpc_cluster.cluster.id
		wait_till = "normal"
	}
	data "ibm_container_cluster_config" "testacc_ds_cluster" {
		cluster_name_id = data.ibm_container_vpc_cluster.testacc_ds_cluster.name
	}
	`

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: masterNodeReadyClusterScript,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_container_vpc_cluster.testacc_ds_cluster", "id"),
					resource.TestCheckResourceAttrWith("data.ibm_container_vpc_cluster.testacc_ds_cluster", "state", func(value string) error {
						switch value {
						case "deploying", "deployed":
							return nil
						}
						return fmt.Errorf("state is not deploying, it was %s", value)
					}),
				),
			},
			{
				Config: normalClusterScriptWithConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_container_vpc_cluster.testacc_ds_cluster", "id"),
					resource.TestCheckResourceAttr("data.ibm_container_vpc_cluster.testacc_ds_cluster", "state", "normal"),
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
