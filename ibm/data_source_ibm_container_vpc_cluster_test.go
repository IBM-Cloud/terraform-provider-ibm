package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMContainerVPCClusterDataSource_basic(t *testing.T) {
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
				Config: testAccCheckIBMContainerVPCClusterDataSource(zone, vpc, subnet, clusterName, flavor, workerCount),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_container_vpc_cluster.testacc_ds_cluster", "worker_count", "1"),
					resource.TestCheckResourceAttr("data.ibm_container_vpc_cluster.testacc_ds_cluster", "worker_pools.#", "1"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerVPCClusterDataSource(zone, vpc, subnet, clusterName, flavor, workerCount string) string {
	return testAccCheckIBMContainerVpcCluster_basic(zone, vpc, subnet, clusterName, flavor, workerCount) + fmt.Sprintf(`
data "ibm_container_vpc_cluster" "testacc_ds_cluster" {
    cluster_name_id = "${ibm_container_vpc_cluster.cluster.id}"
}
`)
}
