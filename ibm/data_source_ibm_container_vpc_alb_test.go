/*
* IBM Confidential
* Object Code Only Source Materials
* 5747-SM3
* (c) Copyright IBM Corp. 2017,2021
*
* The source code for this program is not published or otherwise divested
* of its trade secrets, irrespective of what has been deposited with the
* U.S. Copyright Office.
 */

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMContainerVPCClusterALBDataSource_basic(t *testing.T) {
	flavor := "c2.2x4"
	enable := true
	worker_count := 1
	name1 := acctest.RandIntRange(10, 100)
	name2 := acctest.RandIntRange(10, 100)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMContainerVPCClusterALBDataSourceConfig(enable, flavor, worker_count, name1, name2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_container_vpc_cluster_alb.testacc_ds_alb", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerVPCClusterALBDataSourceConfig(enable bool, flavor string, worker_count, name1, name2 int) string {
	return testAccCheckIBMVpcContainerALB_basic(enable, flavor, worker_count, name1, name2) + fmt.Sprintf(`
	data "ibm_container_vpc_cluster_alb" "testacc_ds_alb" {
	    alb_id = "${ibm_container_vpc_cluster.cluster.albs.0.id}"
	}
`)
}
