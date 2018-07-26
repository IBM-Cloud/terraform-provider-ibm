package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMContainerWorkerPoolZoneAttachment_basic(t *testing.T) {

	workerPoolName := fmt.Sprintf("terraform-%d", acctest.RandInt())
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerWorkerPoolZoneAttachment_basic(clusterName, workerPoolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool_zone_attachment.test_zone", "private_vlan_id", zoneUpdatePrivateVlan),
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool_zone_attachment.test_zone", "public_vlan_id", zonePublicVlan),
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool_zone_attachment.test_zone", "worker_count", "2"),
				),
			},
			{
				Config: testAccCheckIBMContainerWorkerPoolZoneAttachment_update_public_vlan(clusterName, workerPoolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool_zone_attachment.test_zone", "private_vlan_id", zoneUpdatePrivateVlan),
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool_zone_attachment.test_zone", "public_vlan_id", zoneUpdatePublicVlan),
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool_zone_attachment.test_zone", "worker_count", "1"),
				),
			},
		},
	})
}

func TestAccIBMContainerWorkerPoolZoneAttachment_importBasic(t *testing.T) {
	workerPoolName := fmt.Sprintf("terraform-%d", acctest.RandInt())
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMContainerWorkerPoolZoneAttachment_basic(clusterName, workerPoolName),
			},

			resource.TestStep{
				ResourceName:      "ibm_container_worker_pool_zone_attachment.test_zone",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMContainerWorkerPoolZoneAttachment_basic(clusterName, workerPoolName string) string {
	return fmt.Sprintf(`
data "ibm_org" "org" {
  org = "%s"
}

data "ibm_account" "acc" {
  org_guid = "${data.ibm_org.org.id}"
}

resource "ibm_container_cluster" "testacc_cluster" {
  name       = "%s"
  datacenter = "%s"
  account_guid = "${data.ibm_account.acc.id}"
  machine_type    = "%s"
  hardware       = "shared"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
  kube_version    = "%s"
  region = "%s"
}

resource "ibm_container_worker_pool" "test_pool" {
  worker_pool_name = "%s"
  machine_type     = "%s"
  cluster          = "${ibm_container_cluster.testacc_cluster.id}"
  size_per_zone    = 2
  hardware         = "shared"
  disk_encryption  = "true"
  region = "%s"
  labels = {
    "test" = "test-pool"

    "test1" = "test-pool1"
  }
}
resource "ibm_container_worker_pool_zone_attachment" "test_zone" {
  cluster      = "${ibm_container_cluster.testacc_cluster.id}"
  zone         = "%s"
  worker_pool  = "${element(split("/",ibm_container_worker_pool.test_pool.id),1)}"
  private_vlan_id = "%s"
  public_vlan_id  = "%s"
  region = "%s"
}
		
		`, cfOrganization, clusterName, datacenter, machineType, publicVlanID, privateVlanID, kubeUpdateVersion, csRegion, workerPoolName, machineType, csRegion, zone, zoneUpdatePrivateVlan, zonePublicVlan, csRegion)
}

func testAccCheckIBMContainerWorkerPoolZoneAttachment_update_public_vlan(clusterName, workerPoolName string) string {
	return fmt.Sprintf(`
data "ibm_org" "org" {
  org = "%s"
}

data "ibm_account" "acc" {
  org_guid = "${data.ibm_org.org.id}"
}

resource "ibm_container_cluster" "testacc_cluster" {
  name       = "%s"
  datacenter = "%s"
  account_guid = "${data.ibm_account.acc.id}"
  machine_type    = "%s"
  hardware       = "shared"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
  kube_version     = "%s"
  region = "%s"
}

resource "ibm_container_worker_pool" "test_pool" {
  worker_pool_name = "%s"
  machine_type     = "%s"
  cluster          = "${ibm_container_cluster.testacc_cluster.id}"
  size_per_zone    = 1
  hardware         = "shared"
  disk_encryption  = "true"
  region = "%s"
  labels = {
    "test" = "test-pool"

    "test1" = "test-pool1"
  }
}
resource "ibm_container_worker_pool_zone_attachment" "test_zone" {
  cluster      = "${ibm_container_cluster.testacc_cluster.id}"
  zone         = "%s"
  worker_pool  = "${element(split("/",ibm_container_worker_pool.test_pool.id),1)}"
  private_vlan_id = "%s"
  public_vlan_id  = "%s"
  region = "%s"
}
		
		`, cfOrganization, clusterName, datacenter, machineType, publicVlanID, privateVlanID, kubeUpdateVersion, csRegion, workerPoolName, machineType, csRegion, zone, zoneUpdatePrivateVlan, zoneUpdatePublicVlan, csRegion)
}
