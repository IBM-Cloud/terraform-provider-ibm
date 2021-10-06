package ibm

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMContainerVpcClusterWorkerVolumeAttachment_Basic(t *testing.T) {
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
						"ibm_container_storage_attachment.volume_attach", "status", "attached"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerVpcClusterWorkerVolumeAttach_basic(zone, vpc, subnet, clusterName, flavor, workerCount, volumeName string) string {
	return fmt.Sprintf(`
		provider "ibm" {
			region ="us-south"
		}	
		data "ibm_resource_group" "resource_group" {
			is_default = "true"
		}
		resource "ibm_is_vpc" "vpc" {
			name = "%s"
		}
		resource "ibm_is_subnet" "subnet" {
			name                     = "%s"
			vpc                      = ibm_is_vpc.vpc.id
			zone                     = "us-south-1"
			total_ipv4_address_count = 256
		}
		
		resource "ibm_container_vpc_cluster" "cluster" {
			name              = "%s"
			vpc_id            = ibm_is_vpc.vpc.id
			flavor            = "cx2.2x4"
			worker_count      = 1
			wait_till         = "OneWorkerNodeReady"
			resource_group_id = data.ibm_resource_group.resource_group.id
			zones {
				 subnet_id = ibm_is_subnet.subnet.id
				 name      = "us-south-1"
			}
			worker_labels = {
			"test"  = "test-default-pool"
			"test1" = "test-default-pool1"
			"test2" = "test-default-pool2"
			}
			
		  }

		  resource "ibm_is_volume" "storage"{
			name = "%s"
			profile = "10iops-tier"
			zone = "us-south-1"
			# capacity= 200
		}

		data "ibm_container_vpc_cluster" "cluster" {
			name = ibm_container_vpc_cluster.cluster.id
		}

		resource "ibm_container_storage_attachment" "volume_attach"{
			volume = ibm_is_volume.storage.id
			cluster = ibm_container_vpc_cluster.cluster.id
			worker = data.ibm_container_vpc_cluster.cluster.workers[0]
		}`, vpc, subnet, clusterName, volumeName)
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
