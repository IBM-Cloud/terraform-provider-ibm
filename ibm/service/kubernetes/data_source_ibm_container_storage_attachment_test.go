package kubernetes_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
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
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVpcClusterWorkerVolumeAttachDatasource_basic(zone, vpc, subnet, clusterName, flavor, workerCount, volumeName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.ibm_container_storage_attachment.volume_attach_data", "status", "attached"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerVpcClusterWorkerVolumeAttachDatasource_basic(zone, vpc, subnet, clusterName, flavor, workerCount, volumeName string) string {

	return testAccCheckIBMContainerVpcClusterWorkerVolumeAttach_basic(zone, vpc, subnet, clusterName, flavor, workerCount, volumeName) +
		`
	data "ibm_container_storage_attachment" "volume_attach_data"{
		volume_attachment_id = ibm_container_storage_attachment.volume_attach.volume_attachment_id
		cluster = ibm_container_vpc_cluster.cluster.id
		worker = data.ibm_container_vpc_cluster.cluster.workers[0]
	}
	
`
}
