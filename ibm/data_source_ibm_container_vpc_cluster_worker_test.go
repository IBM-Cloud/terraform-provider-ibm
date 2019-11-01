package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMContainerVPCClusterWorkerDataSource_basic(t *testing.T) {
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	vpcID := fmt.Sprintf("vpc_%d", acctest.RandInt())
	flavor := fmt.Sprintf("c.%d", acctest.RandInt())
	zoneID := fmt.Sprintf("zone.%d", acctest.RandInt())
	subnetID := fmt.Sprintf("subnet.%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMContainerVPCClusterWorkerDataSourceConfig(clusterName, vpcID, flavor, zoneID, subnetID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_container_vpc_cluster_worker.testacc_ds_worker", "state", "active"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerVPCClusterWorkerDataSourceConfig(clusterName, vpc_id, flavor, subnet_id, name string) string {
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
	data "ibm_container_vpc_cluster_worker" "testacc_ds_worker" {
	    cluster_name_id = "${ibm_container_vpc_cluster.cluster.id}"
	    worker_id = "${data.ibm_container_vpc_cluster.testacc_ds_cluster.workers[0]}"
	}
`, clusterName, vpc_id, flavor, subnet_id, name)
}
