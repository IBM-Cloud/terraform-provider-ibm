package ibm

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMContainerVpcClusterWorkerVolumeAttachment(t *testing.T) {
	clusterName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	randint := acctest.RandIntRange(10, 100)
	vpc := fmt.Sprintf("terraformvpc-%d", randint)
	subnet := fmt.Sprintf("terraformsubnet-%d", randint)
	flavor := "bx2.16x64"
	zone := "us-south"
	workerCount := "1"
	volumeName := fmt.Sprintf("terraformvpcvol-%d", randint)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMContainerVpcWorkerStorageDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVpcClusterWorkerVolumeAttach_basic(zone, vpc, subnet, clusterName, flavor, workerCount, volumeName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_storage.volume_attach", "status", "attached"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerVpcClusterWorkerVolumeAttach_basic(zone, vpc, subnet, clusterName, flavor, workerCount, volumeName string) string {
	return testAccCheckIBMContainerVpcCluster_basic(zone, vpc, subnet, clusterName, flavor, workerCount) +
		testAccCheckIBMISVolumeConfig(volumeName) +
		fmt.Sprintf(`

	data "ibm_container_vpc_cluster" "cluster"{
		name = ibm_container_vpc_cluster.cluster.name
	}

	resource "ibm_container_vpc_worker_storage" "volume_attach"{
		volume = ibm_is_volume.storage.id
		cluster = ibm_container_vpc_cluster.cluster.id
		worker = data.ibm_container_vpc_cluster.cluster.workers[0]
	}
	
`)
}

func testAccCheckIBMContainerVpcWorkerStorageDestroy(s *terraform.State) error {

	wpClient, err := testAccProvider.Meta().(ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}
	workersAPI := wpClient.Workers()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_container_vpc_worker_storgae" {
			continue
		}

		targetEnv := getVpcClusterTargetHeaderTestACC()
		// Try to find the key

		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		clusterNameorID := parts[0]
		workerID := parts[1]
		volumeAttachmentID := parts[2]

		_, attchmentErr := workersAPI.GetStorageAttachment(clusterNameorID, workerID, volumeAttachmentID, targetEnv)

		if attchmentErr == nil {
			return fmt.Errorf("Volume attachment still exists: %s", rs.Primary.ID)
		} else if !strings.Contains(attchmentErr.Error(), "404") {
			return fmt.Errorf("Error waiting for volume attachment (%s) to be destroyed: %s", rs.Primary.ID, attchmentErr)
		}
	}

	return nil
}
