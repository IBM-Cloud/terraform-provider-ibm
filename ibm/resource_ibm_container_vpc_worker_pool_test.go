package ibm

import (
	"fmt"
	"strings"
	"testing"

	v2 "github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccIBMVpcContainerWorkerPool_basic(t *testing.T) {

	workerPoolName := "kkkk"
	clusterName := "bmmnpu3d089sts43d4lg"
	flavor := "c2.2x4"
	vpc_id := "6015365a-9d93-4bb4-8248-79ae0db2dc26"
	worker_count := 1
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMVpcContainerWorkerPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMVpcContainerWorkerPool_basic(clusterName, workerPoolName, flavor, vpc_id, worker_count),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "worker_pool_name", workerPoolName),
				),
			},
		},
	})
}

func TestAccIBMVpcContainerWorkerPool_importBasic(t *testing.T) {
	workerPoolName := "kkkk"
	clusterName := "bmmnpu3d089sts43d4lg"
	flavor := "c2.2x4"
	vpc_id := "6015365a-9d93-4bb4-8248-79ae0db2dc26"
	worker_count := 1

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMVpcContainerWorkerPoolDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMVpcContainerWorkerPool_basic(clusterName, workerPoolName, flavor, vpc_id, worker_count),
			},

			resource.TestStep{
				ResourceName:      "ibm_container_vpc_worker_pool.test_pool",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMVpcContainerWorkerPoolDestroy(s *terraform.State) error {

	wpClient, err := testAccProvider.Meta().(ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_container_vpc_worker_pool" {
			continue
		}

		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		cluster := parts[0]
		workerPoolID := parts[1]

		target := v2.ClusterTargetHeader{}

		// Try to find the key
		_, err = wpClient.WorkerPools().GetWorkerPool(cluster, workerPoolID, target)

		if err != nil && !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("Error waiting for worker pool (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMVpcContainerWorkerPool_basic(clusterName, workerPoolName, flavor, vpc_id string, worker_count int) string {
	return fmt.Sprintf(`
		resource "ibm_container_vpc_worker_pool" "test_pool" {
			cluster          = "%s"
			worker_pool_name = "%s"
			flavor     = "%s"
			vpc_id = "%s"
			worker_count = %d
			zones = [
				{
					name= "us-south-1"
					subnet_id = "015ffb8b-efb1-4c03-8757-29335a07493b" 
				}
			]
		}
		`, clusterName, workerPoolName, flavor, vpc_id, worker_count)
}
