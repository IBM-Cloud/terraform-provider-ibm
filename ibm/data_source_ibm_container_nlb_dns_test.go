// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMContainerNLBDNSDatasourceBasic(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-cluster-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerNLBDNSDataSourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_container_nlb_dns.dns", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerNLBDNSDataSourceConfig(name string) string {
	return testAccCheckIBMContainerVpcClusterBasic(name) + fmt.Sprintf(`
	data "ibm_container_nlb_dns" "dns" {
	    cluster = ibm_container_vpc_cluster.cluster.id
	}
`)
}
