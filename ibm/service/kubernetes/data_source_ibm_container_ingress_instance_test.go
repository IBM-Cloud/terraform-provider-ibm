// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes_test

import (
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMContainerIngressInstanceDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMContainerIngressInstanceDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_container_vpc_cluster_alb.testacc_ds_alb", "id"),
				),
			},
		},
	})
}

func testAccIBMContainerIngressInstanceDataSourceConfig() string {
	return testAccCheckIBMContainerIngressInstanceBasic() + `
	data "ibm_container_vpc_cluster_alb" "testacc_ds_alb" {
	    alb_id = ibm_container_vpc_cluster.cluster.albs.0.id
	}
`
}
