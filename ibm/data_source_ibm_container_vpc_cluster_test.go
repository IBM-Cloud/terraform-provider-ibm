// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMContainerVPCClusterDataSource_basic(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-cluster-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
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

func testAccCheckIBMContainerVPCClusterDataSource(name string) string {
	return testAccCheckIBMContainerVpcClusterBasic(name) + fmt.Sprintf(`
data "ibm_container_vpc_cluster" "testacc_ds_cluster" {
    cluster_name_id = ibm_container_vpc_cluster.cluster.id
}
data "ibm_container_cluster_config" "testacc_ds_cluster" {
	cluster_name_id = ibm_container_vpc_cluster.cluster.id
  }
`)
}
