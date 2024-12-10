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

func TestAccIBMContainerNLBDNSDatasourceBasic(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-cluster-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
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
	return testAccCheckIBMContainerVpcClusterBasic(name, "OneWorkerNodeReady") + `
	data "ibm_container_nlb_dns" "dns" {
	    cluster = ibm_container_vpc_cluster.cluster.id
	}
`
}
