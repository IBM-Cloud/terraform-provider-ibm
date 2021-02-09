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

func TestAccIBMContainerVPCClusterWorkerPoolDataSource_basic(t *testing.T) {
	flavor := "c2.2x4"
	worker_count := 1
	name1 := acctest.RandIntRange(10, 100)
	name2 := acctest.RandIntRange(10, 100)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMContainerVPCClusterWorkerPoolDataSourceConfig(flavor, worker_count, name1, name2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_container_vpc_cluster_worker_pool.testacc_ds_worker_pool", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerVPCClusterWorkerPoolDataSourceConfig(flavor string, worker_count, name1, name2 int) string {
	return testAccCheckIBMVpcContainerWorkerPool_basic(flavor, worker_count, name1, name2) + fmt.Sprintf(`
	data "ibm_container_vpc_cluster_worker_pool" "testacc_ds_worker_pool" {
	    cluster = "${ibm_container_vpc_cluster.cluster.id}"
	    worker_pool_name = "${ibm_container_vpc_worker_pool.test_pool.worker_pool_name}"
	}
`)
}
