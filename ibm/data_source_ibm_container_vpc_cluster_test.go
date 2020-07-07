package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMContainerVPCClusterDataSource_basic(t *testing.T) {
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	vpcID := fmt.Sprintf("vpc_%d", acctest.RandIntRange(10, 100))
	flavor := fmt.Sprintf("c.%d", acctest.RandIntRange(10, 100))
	zoneID := fmt.Sprintf("zone.%d", acctest.RandIntRange(10, 100))
	subnetID := fmt.Sprintf("subnet.%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMContainerVPCClusterDataSource(clusterName, vpcID, flavor, zoneID, subnetID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_container_vpc_cluster.testacc_ds_cluster", "worker_count", "1"),
					resource.TestCheckResourceAttr("data.ibm_container_vpc_cluster.testacc_ds_cluster", "worker_pools.#", "1"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerVPCClusterDataSource(clusterName, vpc_id, flavor, subnet_id, name string) string {
	return fmt.Sprintf(`
data "ibm_resource_group" "group" {
	is_default = "true"
}

resource "ibm_container_vpc_cluster" "cluster" {
	name              = "%s"
	vpc_id            = "%s"
	flavor            = "%s"
	resource_group_id = "${data.ibm_resource_group.group.id}"
	zones = [{
		 subnet_id = "%s"
		 name = "%s"
	  }
	]
}
data "ibm_container_vpc_cluster" "testacc_ds_cluster" {
    cluster_name_id = "${ibm_container_vpc_cluster.cluster.id}"
}
`, clusterName, vpc_id, flavor, subnet_id, name)
}
