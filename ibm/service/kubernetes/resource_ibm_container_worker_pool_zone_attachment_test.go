// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes_test

import (
	"fmt"
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/sl"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const NOT_FOUND = "SoftLayer_Exception_Network_LBaaS_ObjectNotFound"

func TestAccIBMContainerWorkerPoolZoneAttachment_basic(t *testing.T) {

	workerPoolName := fmt.Sprintf("tf-cluster-worker-%d", acctest.RandIntRange(10, 100))
	clusterName := fmt.Sprintf("tf-cluster-worker-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerWorkerPoolZoneAttachmentBasic(clusterName, workerPoolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool_zone_attachment.test_zone", "private_vlan_id", acc.ZoneUpdatePrivateVlan),
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool_zone_attachment.test_zone", "public_vlan_id", acc.ZonePublicVlan),
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool_zone_attachment.test_zone", "worker_count", "2"),
				),
			},
			{
				Config: testAccCheckIBMContainerWorkerPoolZoneAttachmentUpdatePublicVlan(clusterName, workerPoolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool_zone_attachment.test_zone", "private_vlan_id", acc.ZoneUpdatePrivateVlan),
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool_zone_attachment.test_zone", "public_vlan_id", acc.ZoneUpdatePublicVlan),
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool_zone_attachment.test_zone", "worker_count", "1"),
				),
			},
			{
				ResourceName:      "ibm_container_worker_pool_zone_attachment.test_zone",
				ImportState:       true,
				ImportStateVerify: false,
			},
		},
	})
}

func TestAccIBMContainerWorkerPoolZoneAttachment_privateVlanOnly(t *testing.T) {
	workerPoolName := fmt.Sprintf("tf-cluster-worker-%d", acctest.RandIntRange(10, 100))
	clusterName := fmt.Sprintf("tf-cluster-worker-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerWorkerPoolZoneAttachmentPrivateVlanOnly(clusterName, workerPoolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool_zone_attachment.test_zone", "private_vlan_id", acc.ZoneUpdatePrivateVlan),
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool_zone_attachment.test_zone", "worker_count", "1"),
				),
			},
		},
	})
}

func TestAccIBMContainerWorkerPoolZoneAttachment_publicVlanOnly(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMLbaasDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMContainerWorkerPoolZoneAttachmentPublicVlanOnly(),
				ExpectError: regexp.MustCompile("must be specified if a public_vlan_id"),
			},
		},
	})
}

func testAccCheckIBMLbaasDestroy(s *terraform.State) error {
	sess := acc.TestAccProvider.Meta().(conns.ClientSession).SoftLayerSession()
	service := services.GetNetworkLBaaSLoadBalancerService(sess)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_lbaas" {
			continue
		}

		// Try to find the key
		_, err := service.GetLoadBalancer(sl.String(rs.Primary.ID))

		if err == nil {
			return fmt.Errorf("load balancer (%s) to be destroyed still exists", rs.Primary.ID)
		} else if apiErr, ok := err.(sl.Error); ok && apiErr.Exception != NOT_FOUND {
			return fmt.Errorf("[ERROR] Error waiting for load balancer (%s) to be destroyed: %s", rs.Primary.ID, err)
		}

	}

	return nil
}

func testAccCheckIBMContainerWorkerPoolZoneAttachmentBasic(clusterName, workerPoolName string) string {
	return fmt.Sprintf(`

resource "ibm_container_cluster" "testacc_cluster" {
  name            = "%s"
  datacenter      = "%s"
  machine_type    = "%s"
  hardware        = "shared"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
  kube_version    = "%s"
  wait_till         = "OneWorkerNodeReady"
}

resource "ibm_container_worker_pool" "test_pool" {
  worker_pool_name = "%s"
  machine_type     = "%s"
  cluster          = ibm_container_cluster.testacc_cluster.id
  size_per_zone    = 2
  hardware         = "shared"
  disk_encryption  = "true"
  labels = {
    "test"  = "test-pool"
    "test1" = "test-pool1"
  }
}

resource "ibm_container_worker_pool_zone_attachment" "test_zone" {
  cluster         = ibm_container_cluster.testacc_cluster.id
  zone            = "%s"
  worker_pool     = element(split("/", ibm_container_worker_pool.test_pool.id), 1)
  private_vlan_id = "%s"
  public_vlan_id  = "%s"
  wait_till_albs = false
}
		
		`, clusterName, acc.Datacenter, acc.MachineType, acc.PublicVlanID, acc.PrivateVlanID, acc.KubeUpdateVersion, workerPoolName, acc.MachineType, acc.Zone, acc.ZoneUpdatePrivateVlan, acc.ZonePublicVlan)
}

func testAccCheckIBMContainerWorkerPoolZoneAttachmentPrivateVlanOnly(clusterName, workerPoolName string) string {
	return fmt.Sprintf(`

resource "ibm_container_cluster" "testacc_cluster" {
  name              = "%s"
  datacenter        = "%s"
  machine_type      = "%s"
  hardware          = "shared"
  public_vlan_id    = "%s"
  private_vlan_id   = "%s"
  kube_version      = "%s"
  wait_time_minutes = 180
  wait_till         = "MasterNodeReady"
}

resource "ibm_container_worker_pool" "test_pool" {
  worker_pool_name = "%s"
  machine_type     = "%s"
  cluster          = ibm_container_cluster.testacc_cluster.id
  size_per_zone    = 1
  hardware         = "shared"
  disk_encryption  = "true"
  labels = {
    "test"  = "test-pool"
    "test1" = "test-pool1"
  }
}

resource "ibm_container_worker_pool_zone_attachment" "test_zone" {
  cluster         = ibm_container_cluster.testacc_cluster.id
  zone            = "%s"
  worker_pool     = element(split("/", ibm_container_worker_pool.test_pool.id), 1)
  private_vlan_id = "%s"
}
		
		`, clusterName, acc.Datacenter, acc.MachineType, acc.PublicVlanID, acc.PrivateVlanID, acc.KubeUpdateVersion, workerPoolName, acc.MachineType, acc.Zone, acc.ZoneUpdatePrivateVlan)
}

func testAccCheckIBMContainerWorkerPoolZoneAttachmentPublicVlanOnly() string {
	return fmt.Sprintf(`
  resource "ibm_container_worker_pool_zone_attachment" "test_zone" {
    cluster        = "test"
    zone           = "ams03"
    worker_pool    = "testpool"
    public_vlan_id = "%s"
  }
		
		`, acc.PublicVlanID)
}

func testAccCheckIBMContainerWorkerPoolZoneAttachmentUpdatePublicVlan(clusterName, workerPoolName string) string {
	return fmt.Sprintf(`

resource "ibm_container_cluster" "testacc_cluster" {
  name            = "%s"
  datacenter      = "%s"
  machine_type    = "%s"
  hardware        = "shared"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
  kube_version    = "%s"
  wait_till         = "OneWorkerNodeReady"
}

resource "ibm_container_worker_pool" "test_pool" {
  worker_pool_name = "%s"
  machine_type     = "%s"
  cluster          = ibm_container_cluster.testacc_cluster.id
  size_per_zone    = 1
  hardware         = "shared"
  disk_encryption  = "true"
  labels = {
    "test"  = "test-pool"
    "test1" = "test-pool1"
  }
}

resource "ibm_container_worker_pool_zone_attachment" "test_zone" {
  cluster         = ibm_container_cluster.testacc_cluster.id
  zone            = "%s"
  worker_pool     = element(split("/", ibm_container_worker_pool.test_pool.id), 1)
  private_vlan_id = "%s"
  public_vlan_id  = "%s"
}
		
		`, clusterName, acc.Datacenter, acc.MachineType, acc.PublicVlanID, acc.PrivateVlanID, acc.KubeUpdateVersion, workerPoolName, acc.MachineType, acc.Zone, acc.ZoneUpdatePrivateVlan, acc.ZoneUpdatePublicVlan)
}
