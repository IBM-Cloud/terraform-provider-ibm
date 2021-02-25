/* IBM Confidential
*  Object Code Only Source Materials
*  5747-SM3
*  (c) Copyright IBM Corp. 2017,2021
*
*  The source code for this program is not published or otherwise divested
*  of its trade secrets, irrespective of what has been deposited with the
*  U.S. Copyright Office. */

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMContainerWorkerPoolDataSource_basic(t *testing.T) {
	workerPoolName := fmt.Sprintf("tf-cluster-worker-%d", acctest.RandIntRange(10, 100))
	clusterName := fmt.Sprintf("tf-cluster-worker-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerWorkerPoolDataSourceConfig(clusterName, workerPoolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.ibm_container_worker_pool.testacc_ds_worker_pool", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerWorkerPoolDataSourceConfig(clusterName, workerPoolName string) string {
	return fmt.Sprintf(`

resource "ibm_container_cluster" "testacc_cluster" {
  name            = "%s"
  datacenter      = "%s"
  machine_type    = "%s"
  hardware        = "shared"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
  kube_version    = "%s"
}

resource "ibm_container_worker_pool" "test_pool" {
  worker_pool_name = "%s"
  machine_type     = "%s"
  cluster          = ibm_container_cluster.testacc_cluster.id
  size_per_zone    = 1
  hardware         = "shared"
  disk_encryption  = true
  labels = {
    "test"  = "test-pool"
    "test1" = "test-pool1"
  }
}
data "ibm_container_worker_pool" "testacc_ds_worker_pool"{
  worker_pool_name = ibm_container_worker_pool.test_pool.worker_pool_name
  cluster          = ibm_container_cluster.testacc_cluster.id
}`, clusterName, datacenter, machineType, publicVlanID, privateVlanID, kubeVersion, workerPoolName, machineType)
}
