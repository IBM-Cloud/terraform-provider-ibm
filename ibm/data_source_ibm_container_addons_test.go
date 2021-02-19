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

func TestAccIBMContainerAddOnsDataSource_basic(t *testing.T) {
	clusterName := fmt.Sprintf("terraform-%d", acctest.RandIntRange(10, 100))
	vpc := fmt.Sprintf("terraform-vpc-%d", acctest.RandIntRange(10, 100))
	subnet := fmt.Sprintf("terraform-subnet-%d", acctest.RandIntRange(10, 100))
	flavor := "c2.2x4"
	zone := "us-south"
	workerCount := "1"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMContainerAddOnsDataSourceConfig(zone, vpc, subnet, clusterName, flavor, workerCount),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_container_addons.addons", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerAddOnsDataSourceConfig(zone, vpc, subnet, clusterName, flavor, workerCount string) string {
	return testAccCheckIBMContainerAddOnsBasic(zone, vpc, subnet, clusterName, flavor, workerCount) + fmt.Sprintf(`
	data "ibm_container_addons" "addons" {
	    cluster= ibm_container_addons.addons.cluster
	}
`)
}
