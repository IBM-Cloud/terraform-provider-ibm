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
	vpc := fmt.Sprintf("vpc-%d", randint)
	subnet := fmt.Sprintf("subnet-%d", randint)
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
	return fmt.Sprintf(`
	%s
	data "ibm_container_vpc_cluster" "testacc_ds_cluster" {
	    cluster_name_id = "${ibm_container_vpc_cluster.cluster.id}"
	}
	data "ibm_container_vpc_cluster_worker" "testacc_ds_worker" {
	    cluster_name_id = "${ibm_container_vpc_cluster.cluster.id}"
	    worker_id = "${data.ibm_container_vpc_cluster.testacc_ds_cluster.workers[0]}"
	}
`, testAccCheckIBMContainerVpcCluster_basic(zone, vpc, subnet, clusterName, flavor, workerCount))
}
