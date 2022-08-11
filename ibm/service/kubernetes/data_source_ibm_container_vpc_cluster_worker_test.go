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

func TestAccIBMContainerVPCClusterWorkerDataSource_basic(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-cluster-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVPCClusterWorkerDataSourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_container_vpc_cluster_worker.testacc_ds_worker", "id"),
				),
			},
		},
	})
}

func TestAccIBMContainerVPCClusterWorkerDataSource_dedicatedhost(t *testing.T) {
	clusterName := acc.ClusterName
	hostpoolID := acc.HostPoolID
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVPCClusterWorkerDataSourceDedicatedHostConfig(clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_container_vpc_cluster_worker.dhost_vpc_worker", "host_pool_id", hostpoolID),
				),
			},
		},
	})
}

func testAccCheckIBMContainerVPCClusterWorkerDataSourceConfig(name string) string {
	return testAccCheckIBMVpcContainerWorkerPoolBasic(name) + `
	data "ibm_container_vpc_cluster" "testacc_ds_cluster" {
		cluster_name_id = ibm_container_vpc_cluster.cluster.id
	}
	data "ibm_container_vpc_cluster_worker" "testacc_ds_worker" {
	    cluster_name_id = ibm_container_vpc_cluster.cluster.id
	    worker_id = data.ibm_container_vpc_cluster.testacc_ds_cluster.workers[0]
	}
`
}

func testAccCheckIBMContainerVPCClusterWorkerDataSourceDedicatedHostConfig(clusterName string) string {
	return fmt.Sprintf(`
	data "ibm_container_vpc_cluster" "dhost_vpc_cluster" {
		cluster_name_id = "%s"
	}
	data "ibm_container_vpc_cluster_worker" "dhost_vpc_worker" {
	    cluster_name_id = "%s"
	    worker_id = data.ibm_container_vpc_cluster.dhost_vpc_cluster.workers[0]
	}
	`, clusterName, clusterName)
}
