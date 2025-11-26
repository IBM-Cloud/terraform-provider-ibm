// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMContainerVPCClusterALBDataSource_basic(t *testing.T) {
	enable := true
	name := fmt.Sprintf("tf-vpc-alb-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVPCClusterALBDataSourceConfig(enable, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_container_vpc_cluster_alb.testacc_ds_alb", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerVPCClusterALBDataSourceConfig(enable bool, name string) string {
	return testAccCheckIBMVpcContainerALBBasic(enable, name) + `
	data "ibm_container_vpc_cluster_alb" "testacc_ds_alb" {
	    alb_id = ibm_container_vpc_cluster.cluster.albs.0.id
	}
`
}
