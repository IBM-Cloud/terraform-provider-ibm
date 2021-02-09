/**
 * © Copyright IBM Corporation 2020. All Rights Reserved.
 *
 * Licensed under the Mozilla Public License, version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at https://mozilla.org/MPL/2.0/
 */

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMContainerWorkerPoolDataSource_basic(t *testing.T) {
	workerPoolName := fmt.Sprintf("terraform-%d", acctest.RandIntRange(10, 100))
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMContainerWorkerPoolDataSourceConfig(clusterName, workerPoolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_container_worker_pool.testacc_ds_worker_pool", "state", "active"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerWorkerPoolDataSourceConfig(clusterName, workerPoolName string) string {
	return fmt.Sprintf(`

resource "ibm_container_cluster" "testacc_cluster" {
  name            = "%s"
  datacenter      = "%s"
  machine_type    = "%s"
  hardware        = "shared"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
  kube_version    = "%s"
  region          = "%s"
}

resource "ibm_container_worker_pool" "test_pool" {
  worker_pool_name = "%s"
  machine_type     = "%s"
  cluster          = ibm_container_cluster.testacc_cluster.id
  size_per_zone    = 1
  hardware         = "shared"
  disk_encryption  = true
  labels = {
    "test"  = "test-pool"
    "test1" = "test-pool1"
  }
}

data "ibm_container_worker_pool" "testacc_ds_worker_pool"{
  worker_pool_name = ibm_container_worker_pool.test_pool.worker_pool_name
  cluster          = ibm_container_cluster.testacc_cluster.id
}`, clusterName, datacenter, machineType, publicVlanID, privateVlanID, kubeVersion, csRegion, workerPoolName, machineType)
}
