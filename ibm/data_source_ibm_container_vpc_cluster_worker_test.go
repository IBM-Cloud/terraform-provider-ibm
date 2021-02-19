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

func TestAccIBMContainerVPCClusterWorkerDataSource_basic(t *testing.T) {
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	randint := acctest.RandIntRange(10, 100)
	vpc := fmt.Sprintf("terraform_vpc-%d", randint)
	subnet := fmt.Sprintf("terraform_subnet-%d", randint)
	flavor := "c2.2x4"
	zone := "us-south"
	workerCount := "1"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMContainerVPCClusterWorkerDataSourceConfig(zone, vpc, subnet, clusterName, flavor, workerCount),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_container_vpc_cluster_worker.testacc_ds_worker", "state", "normal"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerVPCClusterWorkerDataSourceConfig(zone, vpc, subnet, clusterName, flavor, workerCount string) string {
	return testAccCheckIBMContainerVPCClusterDataSource(zone, vpc, subnet, clusterName, flavor, workerCount) + fmt.Sprintf(`
	data "ibm_container_vpc_cluster_worker" "testacc_ds_worker" {
	    cluster_name_id = "${ibm_container_vpc_cluster.cluster.id}"
	    worker_id = "${data.ibm_container_vpc_cluster.testacc_ds_cluster.workers[0]}"
	}
`)
}
