package kubernetes_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMContainerVpcClusterWorkerVolumeAttachment_Basic(t *testing.T) {
	randInt := acctest.RandIntRange(10, 100)
	clusterName := fmt.Sprintf("terraformcluster-%d", randInt)
	vpc := fmt.Sprintf("terraformvpc-%d", randInt)
	subnet := fmt.Sprintf("terraformsubnet-%d", randInt)
	flavor := "cx2.2x4"
	zone := "us-south"
	workerCount := "1"
	volumeName := fmt.Sprintf("terraformvpcvol-%d", randInt)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMContainerVpcWorkerStorageDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVpcClusterWorkerVolumeAttach_basic(zone, vpc, subnet, clusterName, flavor, workerCount, volumeName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_storage_attachment.volume_attach", "status", "attached"),
				),
			},
			{
				ResourceName:      "ibm_container_storage_attachment.volume_attach",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMContainerVpcClusterWorkerVolumeAttach_basic(zone, vpc, subnet, clusterName, flavor, workerCount, volumeName string) string {
	return fmt.Sprintf(`
		provider "ibm" {
			region ="%s"
		}	
		data "ibm_resource_group" "resource_group" {
			is_default = "true"
		}
		resource "ibm_is_vpc" "vpc" {
			name           = "%s"
			resource_group = data.ibm_resource_group.resource_group.id
		}
		resource "ibm_is_subnet" "subnet" {
			name                     = "%s"
			vpc                      = ibm_is_vpc.vpc.id
			zone                     = "%s-1"
			total_ipv4_address_count = 256
			resource_group           = data.ibm_resource_group.resource_group.id
		}
		
		resource "ibm_container_vpc_cluster" "cluster" {
			name              = "%s"
			vpc_id            = ibm_is_vpc.vpc.id
			flavor            = "%s"
			worker_count      = %s
			wait_till         = "OneWorkerNodeReady"
			resource_group_id = data.ibm_resource_group.resource_group.id
			zones {
				subnet_id = ibm_is_subnet.subnet.id
				name      = "%s-1"
			}
			worker_labels = {
				"test"  = "test-default-pool"
				"test1" = "test-default-pool1"
				"test2" = "test-default-pool2"
			}
			
		}

		resource "ibm_is_volume" "storage"{
			name           = "%s"
			profile        = "10iops-tier"
			zone           = "%s-1"
			# capacity     = 200
			resource_group = data.ibm_resource_group.resource_group.id
		}

		data "ibm_container_vpc_cluster" "cluster" {
			name = ibm_container_vpc_cluster.cluster.id
		}

		resource "ibm_container_storage_attachment" "volume_attach"{
			volume = ibm_is_volume.storage.id
			cluster = ibm_container_vpc_cluster.cluster.id
			worker = data.ibm_container_vpc_cluster.cluster.workers[0]
		}`, zone, vpc, subnet, zone, clusterName, flavor, workerCount, zone, volumeName, zone)
}

func testAccCheckIBMContainerVpcWorkerStorageDestroy(s *terraform.State) error {

	wpClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcContainerAPI()
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

		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		clusterNameorID := parts[0]
		workerID := parts[1]
		volumeAttachmentID := parts[2]

		_, attchmentErr := workersAPI.GetStorageAttachment(clusterNameorID, workerID, volumeAttachmentID, targetEnv)

		if attchmentErr == nil {
			return fmt.Errorf("[ERROR] Volume attachment still exists: %s", rs.Primary.ID)
		} else if !strings.Contains(attchmentErr.Error(), "404") {
			return fmt.Errorf("[ERROR] Error waiting for volume attachment (%s) to be destroyed: %s", rs.Primary.ID, attchmentErr)
		}
	}

	return nil
}
