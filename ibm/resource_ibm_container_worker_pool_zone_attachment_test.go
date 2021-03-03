// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMContainerWorkerPoolZoneAttachment_basic(t *testing.T) {

	workerPoolName := fmt.Sprintf("tf-cluster-worker-%d", acctest.RandIntRange(10, 100))
	clusterName := fmt.Sprintf("tf-cluster-worker-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerWorkerPoolZoneAttachmentBasic(clusterName, workerPoolName),
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
				Config: testAccCheckIBMContainerWorkerPoolZoneAttachmentUpdatePublicVlan(clusterName, workerPoolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool_zone_attachment.test_zone", "private_vlan_id", zoneUpdatePrivateVlan),
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool_zone_attachment.test_zone", "public_vlan_id", zoneUpdatePublicVlan),
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
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerWorkerPoolZoneAttachmentPrivateVlanOnly(clusterName, workerPoolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool_zone_attachment.test_zone", "private_vlan_id", zoneUpdatePrivateVlan),
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool_zone_attachment.test_zone", "worker_count", "1"),
				),
			},
		},
	})
}

func TestAccIBMContainerWorkerPoolZoneAttachment_publicVlanOnly(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMLbaasDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMContainerWorkerPoolZoneAttachmentPublicVlanOnly(),
				ExpectError: regexp.MustCompile("must be specified if a public_vlan_id"),
			},
		},
	})
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
		
		`, clusterName, datacenter, machineType, publicVlanID, privateVlanID, kubeUpdateVersion, workerPoolName, machineType, zone, zoneUpdatePrivateVlan, zonePublicVlan)
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
		
		`, clusterName, datacenter, machineType, publicVlanID, privateVlanID, kubeUpdateVersion, workerPoolName, machineType, zone, zoneUpdatePrivateVlan)
}

func testAccCheckIBMContainerWorkerPoolZoneAttachmentPublicVlanOnly() string {
	return fmt.Sprintf(`
  resource "ibm_container_worker_pool_zone_attachment" "test_zone" {
    cluster        = "test"
    zone           = "ams03"
    worker_pool    = "testpool"
    public_vlan_id = "%s"
  }
		
		`, publicVlanID)
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
		
		`, clusterName, datacenter, machineType, publicVlanID, privateVlanID, kubeUpdateVersion, workerPoolName, machineType, zone, zoneUpdatePrivateVlan, zoneUpdatePublicVlan)
}
