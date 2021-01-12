package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMContainerVpcClusterWorkerVolumeAttachmentDatasource(t *testing.T) {
	clusterName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	randint := acctest.RandIntRange(10, 100)
	vpc := fmt.Sprintf("terraformvpc-%d", randint)
	subnet := fmt.Sprintf("terraformsubnet-%d", randint)
	flavor := "bx2.16x64"
	zone := "us-south"
	workerCount := "1"
	volumeName := fmt.Sprintf("terraformvpcvol-%d", randint)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVpcClusterWorkerVolumeAttachDatasource_basic(zone, vpc, subnet, clusterName, flavor, workerCount, volumeName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.ibm_container_vpc_worker_storage.volume_attach_data", "status", "attached"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerVpcClusterWorkerVolumeAttachDatasource_basic(zone, vpc, subnet, clusterName, flavor, workerCount, volumeName string) string {

	return testAccCheckIBMContainerVpcClusterWorkerVolumeAttach_basic(zone, vpc, subnet, clusterName, flavor, workerCount, volumeName) +
		fmt.Sprintf(`

	data "ibm_container_vpc_worker_storage" "volume_attach_data"{
		volume_attachment_id = ibm_container_vpc_worker_storage.volume_attach.volume_attachment_id
		cluster = ibm_container_vpc_cluster.cluster.id
		worker = data.ibm_container_vpc_cluster.cluster.workers[0]
	}
	
`)
}
