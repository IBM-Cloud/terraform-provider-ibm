// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	v1 "github.com/Mavrickk3/bluemix-go/api/container/containerv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMContainerWorkerPoolBasic(t *testing.T) {

	workerPoolName := fmt.Sprintf("tf-cluster-worker-%d", acctest.RandIntRange(10, 100))
	clusterName := fmt.Sprintf("tf-cluster-worker-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMContainerWorkerPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerWorkerPoolBasic(clusterName, workerPoolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool.test_pool", "worker_pool_name", workerPoolName),
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool.test_pool", "size_per_zone", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool.test_pool", "labels.%", "2"),
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool.test_pool", "state", "active"),
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool.test_pool", "disk_encryption", "true"),
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool.test_pool", "hardware", "shared"),
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool.default_pool", "size_per_zone", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool.default_pool", "labels.%", "0"),
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool.default_pool", "hardware", "shared"),
				),
			},
			{
				Config: testAccCheckIBMContainerWorkerPoolUpdate(clusterName, workerPoolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool.test_pool", "worker_pool_name", workerPoolName),
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool.test_pool", "size_per_zone", "2"),
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool.test_pool", "labels.%", "2"),
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool.test_pool", "state", "active"),
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool.test_pool", "disk_encryption", "true"),
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool.test_pool", "hardware", "shared"),
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool.default_pool", "size_per_zone", "2"),
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool.default_pool", "labels.%", "2"),
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool.default_pool", "hardware", "shared"),
				),
			},
			{
				ResourceName:      "ibm_container_worker_pool.test_pool",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIBMContainerWorkerPoolZeroSize(t *testing.T) {

	workerPoolName := fmt.Sprintf("tf-cluster-worker-%d", acctest.RandIntRange(10, 100))
	clusterName := fmt.Sprintf("tf-cluster-worker-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMContainerWorkerPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerWorkerPoolZeroSize(clusterName, workerPoolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool.test_pool", "worker_pool_name", workerPoolName),
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool.test_pool", "size_per_zone", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool.test_pool", "labels.%", "2"),
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool.test_pool", "state", "active"),
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool.test_pool", "disk_encryption", "true"),
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool.test_pool", "hardware", "shared"),
				),
			},
			{
				Config: testAccCheckIBMContainerWorkerPoolZeroSizeUpdate(clusterName, workerPoolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool.test_pool", "worker_pool_name", workerPoolName),
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool.test_pool", "size_per_zone", "0"),
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool.test_pool", "labels.%", "2"),
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool.test_pool", "state", "active"),
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool.test_pool", "disk_encryption", "true"),
					resource.TestCheckResourceAttr(
						"ibm_container_worker_pool.test_pool", "hardware", "shared"),
				),
			},
			{
				ResourceName:      "ibm_container_worker_pool.test_pool",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIBMContainerWorkerPoolInvalidSizePerZone(t *testing.T) {
	workerPoolName := fmt.Sprintf("tf-cluster-worker-%d", acctest.RandIntRange(10, 100))
	clusterName := fmt.Sprintf("tf-cluster-worker-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMContainerWorkerPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMContainerWorkerPoolInvalidSizePerZone(clusterName, workerPoolName),
				ExpectError: regexp.MustCompile("must be non-negative"),
			},
		},
	})
}

func testAccCheckIBMContainerWorkerPoolDestroy(s *terraform.State) error {

	csClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ContainerAPI()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_container_worker_pool" {
			continue
		}

		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		cluster := parts[0]
		workerPoolID := parts[1]

		target := v1.ClusterTargetHeader{
			Region: acc.CsRegion,
		}

		// Try to find the key
		_, err = csClient.WorkerPools().GetWorkerPool(cluster, workerPoolID, target)

		if err == nil {
			return fmt.Errorf("Worker pool still exists: %s", rs.Primary.ID)
		} else if !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("[ERROR] Error waiting for worker pool (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMContainerWorkerPoolBasic(clusterName, workerPoolName string) string {
	return fmt.Sprintf(`

resource "ibm_container_cluster" "testacc_cluster" {
  name             = "%[1]s"
  datacenter       = "%[2]s"
  machine_type     = "%[3]s"
  hardware         = "shared"
  public_vlan_id   = "%[4]s"
  private_vlan_id  = "%[5]s"
  kube_version     = "%[6]s"
  wait_till        = "OneWorkerNodeReady"
  operating_system = "UBUNTU_20_64"
  taints {
	key    = "key1"
	value  = "value1"
	effect = "NoSchedule"
  }
}

resource "ibm_container_worker_pool" "default_pool" {
  worker_pool_name = "default"
  machine_type     = "%[3]s"
  cluster          = ibm_container_cluster.testacc_cluster.id
  size_per_zone    = 1
  import_on_create = "true"
  taints {
	key    = "key1"
	value  = "value1"
	effect = "NoSchedule"
  }
}

resource "ibm_container_worker_pool" "test_pool" {
  worker_pool_name = "%[7]s"
  machine_type     = "%[8]s"
  cluster          = ibm_container_cluster.testacc_cluster.id
  size_per_zone    = 1
  hardware         = "shared"
  disk_encryption  = true
  labels = {
    "test"  = "test-pool"
    "test1" = "test-pool1"
  }
  taints {
	key    = "key1"
	value  = "value1"
	effect = "NoSchedule"
  }
}`, clusterName, acc.Datacenter, acc.MachineType, acc.PublicVlanID, acc.PrivateVlanID, acc.KubeVersion, workerPoolName, acc.MachineType)
}

func testAccCheckIBMContainerWorkerPoolUpdate(clusterName, workerPoolName string) string {
	return fmt.Sprintf(`

resource "ibm_container_cluster" "testacc_cluster" {
  name            = "%[1]s"
  datacenter      = "%[2]s"
  machine_type    = "%[3]s"
  hardware        = "shared"
  public_vlan_id  = "%[4]s"
  private_vlan_id = "%[5]s"
  kube_version    = "%[6]s"
  wait_till         = "OneWorkerNodeReady"
}

resource "ibm_container_worker_pool" "default_pool" {
  worker_pool_name = "default"
  machine_type     = "%[3]s"
  cluster          = ibm_container_cluster.testacc_cluster.id
  size_per_zone    = 2
  import_on_create = "true"
  taints {
	key    = "key1"
	value  = "value1"
	effect = "NoSchedule"
  }
  labels = {
    "test"  = "test-pool"
    "test1" = "test-pool1"
  }
  depends_on = [
		  ibm_container_worker_pool.test_pool
	  ]
  orphan_on_delete = "true"
}

resource "ibm_container_worker_pool" "test_pool" {
  worker_pool_name = "%[7]s"
  machine_type     = "%[8]s"
  cluster          = ibm_container_cluster.testacc_cluster.id
  size_per_zone    = 2
  hardware         = "shared"
  disk_encryption  = true
  operating_system = "UBUNTU_24_64"
  labels = {
    "test"  = "test-pool"
    "test1" = "test-pool1"
  }
}`, clusterName, acc.Datacenter, acc.MachineType, acc.PublicVlanID, acc.PrivateVlanID, acc.KubeVersion, workerPoolName, acc.MachineType)
}

func testAccCheckIBMContainerWorkerPoolZeroSize(clusterName, workerPoolName string) string {
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
  taints {
	key    = "key1"
	value  = "value1"
	effect = "NoSchedule"
  }
}

resource "ibm_container_worker_pool" "test_pool" {
  worker_pool_name = "%s"
  machine_type     = "%s"
  cluster          = ibm_container_cluster.testacc_cluster.id
  size_per_zone    = 1
  hardware         = "shared"
  disk_encryption  = true
  labels = {
    "test"  = "test-pool"
    "test1" = "test-pool1"
  }
  taints {
	key    = "key1"
	value  = "value1"
	effect = "NoSchedule"
  }
}`, clusterName, acc.Datacenter, acc.MachineType, acc.PublicVlanID, acc.PrivateVlanID, acc.KubeVersion, workerPoolName, acc.MachineType)
}

func testAccCheckIBMContainerWorkerPoolZeroSizeUpdate(clusterName, workerPoolName string) string {
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
  size_per_zone    = 0
  hardware         = "shared"
  disk_encryption  = true
  labels = {
    "test"  = "test-pool"
    "test1" = "test-pool1"
  }
}`, clusterName, acc.Datacenter, acc.MachineType, acc.PublicVlanID, acc.PrivateVlanID, acc.KubeVersion, workerPoolName, acc.MachineType)
}

func testAccCheckIBMContainerWorkerPoolInvalidSizePerZone(clusterName, workerPoolName string) string {
	return fmt.Sprintf(`
resource "ibm_container_worker_pool" "test_pool" {
  worker_pool_name = "%s"
  machine_type     = "%s"
  cluster          = "%s"
  size_per_zone    = -1
  hardware         = "shared"
  disk_encryption  = true

  labels = {
    "test"  = "test-pool"
    "test1" = "test-pool1"
  }
}`, workerPoolName, acc.MachineType, clusterName)
}
