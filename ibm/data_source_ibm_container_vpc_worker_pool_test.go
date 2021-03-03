// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMContainerVPCClusterWorkerPoolDataSource_basic(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-worker-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVPCClusterWorkerPoolDataSourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_container_vpc_cluster_worker_pool.testacc_ds_worker_pool", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerVPCClusterWorkerPoolDataSourceConfig(name string) string {
	return testAccCheckIBMVpcContainerWorkerPoolBasic(name) + fmt.Sprintf(`
	data "ibm_container_vpc_cluster_worker_pool" "testacc_ds_worker_pool" {
	    cluster = "${ibm_container_vpc_cluster.cluster.id}"
	    worker_pool_name = "${ibm_container_vpc_worker_pool.test_pool.worker_pool_name}"
	}
`)
}
