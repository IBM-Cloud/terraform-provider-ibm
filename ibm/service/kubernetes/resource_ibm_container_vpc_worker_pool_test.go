// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

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

	v2 "github.com/Mavrickk3/bluemix-go/api/container/containerv2"
)

func testAccCheckIBMVpcContainerWorkerPoolDestroy(s *terraform.State) error {

	wpClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_container_vpc_worker_pool" {
			continue
		}

		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		cluster := parts[0]
		workerPoolID := parts[1]

		target := v2.ClusterTargetHeader{}

		// Try to find the key
		wp, err := wpClient.WorkerPools().GetWorkerPool(cluster, workerPoolID, target)

		if err == nil {
			if wp.ActualState == "deleted" && wp.DesiredState == "deleted" {
				return nil
			}
			return fmt.Errorf("Worker pool still exists: %s", rs.Primary.ID)
		} else if !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("[ERROR] Error waiting for worker pool (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

// TestAccIBMContainerVpcClusterWorkerPoolResourceBasic ...
func TestAccIBMContainerVpcClusterWorkerPoolResourceBasic(t *testing.T) {

	name := fmt.Sprintf("tf-vpc-wp-basic-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMVpcContainerWorkerPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMVpcContainerWorkerPoolBasic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "flavor", "cx2.2x4"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "zones.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "labels.%", "2"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "worker_count", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.default_pool", "flavor", "cx2.2x4"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.default_pool", "zones.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.default_pool", "labels.%", "0"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.default_pool", "worker_count", "1"),
				),
			},
			{
				Config: testAccCheckIBMVpcContainerWorkerPoolUpdate(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "flavor", "cx2.2x4"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "zones.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "labels.%", "3"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "operating_system", "UBUNTU_24_64"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "worker_count", "2"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.default_pool", "flavor", "cx2.2x4"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.default_pool", "zones.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.default_pool", "labels.%", "2"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.default_pool", "worker_count", "2"),
				),
			},
			{
				ResourceName:            "ibm_container_vpc_worker_pool.test_pool",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"orphan_on_delete", "import_on_create"},
			},
			{
				Config:  testAccCheckIBMVpcContainerWorkerPoolUpdate(name),
				Destroy: true,
			},
		},
	})
}

func testAccCheckIBMVpcContainerWorkerPoolBasic(cluster_name string) string {
	workerpool_name := cluster_name + "-wp"
	return fmt.Sprintf(`
	data "ibm_resource_group" "resource_group" {
		is_default=true
	}
	
	resource "ibm_container_vpc_cluster" "cluster" {
	  name              = "%[3]s"
	  vpc_id            = "%[1]s"
	  flavor            = "cx2.2x4"
	  worker_count      = 1
	  resource_group_id = data.ibm_resource_group.resource_group.id
	  wait_till         = "MasterNodeReady"
	  zones {
		subnet_id = "%[2]s"
		name      = "us-south-1"
	  }
	}

	resource "ibm_container_vpc_worker_pool" "default_pool" {
		cluster           = ibm_container_vpc_cluster.cluster.id
		vpc_id            = "%[1]s"
		flavor            = "cx2.2x4"
		worker_count      = 1
		worker_pool_name  = "default"
		zones {
			subnet_id = "%[2]s"
			name      = "us-south-1"
		}
		import_on_create  = "true"
	}
	resource "ibm_container_vpc_worker_pool" "test_pool" {
	  cluster           = ibm_container_vpc_cluster.cluster.id
	  worker_pool_name  = "%[4]s"
	  flavor            = "cx2.2x4"
	  vpc_id            = "%[1]s"
	  worker_count      = 1
	  resource_group_id = data.ibm_resource_group.resource_group.id
	  operating_system  = "UBUNTU_20_64"
	  zones {
		name      = "us-south-1"
		subnet_id = "%[2]s"
	  }
	  labels = {
		"test"  = "test-pool"
		"test1" = "test-pool1"
	  }
	  depends_on = [
		  ibm_container_vpc_worker_pool.default_pool
	  ]
	}
		`, acc.IksClusterVpcID, acc.IksClusterSubnetID, cluster_name, workerpool_name)
}

func testAccCheckIBMVpcContainerWorkerPoolUpdate(cluster_name string) string {
	workerpool_name := cluster_name + "-wp"
	return fmt.Sprintf(`
	data "ibm_resource_group" "resource_group" {
		is_default=true
	}
	
	resource "ibm_container_vpc_cluster" "cluster" {
	  name              = "%[3]s"
	  vpc_id            = "%[1]s"
	  flavor            = "cx2.2x4"
	  worker_count      = 1
	  resource_group_id = data.ibm_resource_group.resource_group.id
	  wait_till         = "MasterNodeReady"
	  zones {
		subnet_id = "%[2]s"
		name      = "us-south-1"
	  }
	}
	resource "ibm_container_vpc_worker_pool" "default_pool" {
		cluster           = ibm_container_vpc_cluster.cluster.id
		vpc_id            = "%[1]s"
		flavor            = "cx2.2x4"
		worker_count      = 2
		worker_pool_name  = "default"
		zones {
			subnet_id = "%[2]s"
			name      = "us-south-1"
		}
		import_on_create  = "true"
		labels = {
		"test"  = "default-pool"
		"test1" = "default-pool1"
	  }
	}
	resource "ibm_container_vpc_worker_pool" "test_pool" {
	  cluster           = ibm_container_vpc_cluster.cluster.id
	  worker_pool_name  = "%[4]s"
	  flavor            = "cx2.2x4"
	  vpc_id            = "%[1]s"
	  worker_count      = 2
	  resource_group_id = data.ibm_resource_group.resource_group.id
	  operating_system  = "UBUNTU_24_64"
	  zones {
		name      = "us-south-1"
		subnet_id = "%[2]s"
	  }
	  labels = {
		"test"  = "test-pool"
		"test1" = "test-pool1"
		"test2" = "test-pool2"
	  }
	  depends_on = [
		  ibm_container_vpc_worker_pool.default_pool
	  ]
	  orphan_on_delete = "true"
	}
		`, acc.IksClusterVpcID, acc.IksClusterSubnetID, cluster_name, workerpool_name)
}

// TestAccIBMContainerVpcClusterWorkerPoolResourceSecurityGroups ...
func TestAccIBMContainerVpcClusterWorkerPoolResourceSecurityGroups(t *testing.T) {

	name := fmt.Sprintf("tf-vpc-wp-secgroup-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMVpcContainerWorkerPoolSecurityGroups(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "flavor", "cx2.2x4"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "zones.#", "1"),
				),
			},
		},
	})
}

func testAccCheckIBMVpcContainerWorkerPoolSecurityGroups(name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "resource_group" {
		is_default=true
	}
	resource "ibm_is_vpc" "vpc" {
		name = "%[1]s"
	}
	resource "ibm_is_security_group" "security_group1" {
		name = "%[1]s-security-group-1"
		vpc  = ibm_is_vpc.vpc.id
	}
	resource "ibm_is_security_group" "security_group2" {
		name = "%[1]s-security-group-2"
		vpc  = ibm_is_vpc.vpc.id
	}
	resource "ibm_is_subnet" "subnet1" {
		name                     = "%[1]s-subnet-1"
		vpc                      = ibm_is_vpc.vpc.id
		zone                     = "us-south-1"
		total_ipv4_address_count = 256
	}
	resource "ibm_is_subnet" "subnet2" {
		name                     = "%[1]s-subnet-2"
		vpc                      = ibm_is_vpc.vpc.id
		zone                     = "us-south-2"
		total_ipv4_address_count = 256
	}
	
	resource "ibm_container_vpc_cluster" "cluster" {
	  name              = "%[1]s"
	  vpc_id            = ibm_is_vpc.vpc.id
	  flavor            = "cx2.2x4"
	  worker_count      = 1
	  resource_group_id = data.ibm_resource_group.resource_group.id
	  wait_till         = "MasterNodeReady"
	  zones {
		subnet_id = ibm_is_subnet.subnet1.id
		name      = ibm_is_subnet.subnet1.zone
	  }
	  security_groups = [ 
		ibm_is_security_group.security_group1.id,
		"cluster",
	  ]
	}
	resource "ibm_container_vpc_worker_pool" "test_pool" {
	  cluster           = ibm_container_vpc_cluster.cluster.id
	  worker_pool_name  = "%[1]s"
	  flavor            = "cx2.2x4"
	  vpc_id            = ibm_is_vpc.vpc.id
	  worker_count      = 1
	  resource_group_id = data.ibm_resource_group.resource_group.id
	  zones {
		subnet_id = ibm_is_subnet.subnet2.id
		name      = ibm_is_subnet.subnet2.zone
	  }
	  security_groups = [ 
		ibm_is_security_group.security_group2.id,
	  ]

	}
		`, name)
}

// TestAccIBMContainerVpcClusterWorkerPoolResourceDedicatedHost ...
func TestAccIBMContainerVpcClusterWorkerPoolResourceDedicatedHost(t *testing.T) {
	if acc.HostPoolID == "" {
		fmt.Println("[WARN] Skipping TestAccIBMContainerVpcClusterWorkerPoolResourceDedicatedHost - IBM_CONTAINER_DEDICATEDHOST_POOL_ID is unset")
		return
	}
	name := fmt.Sprintf("tf-vpc-wp-dhost-%d", acctest.RandIntRange(10, 100))

	hostpoolID := acc.HostPoolID
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMVpcContainerWorkerPoolDedicatedHostCreate(
					acc.ClusterName,
					name,
					"bx2d.4x16",
					acc.IksClusterSubnetID,
					acc.IksClusterVpcID,
					acc.IksClusterResourceGroupID,
					hostpoolID,
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.vpc_worker_pool", "host_pool_id", hostpoolID),
				),
			},
		},
	})
}

func testAccCheckIBMVpcContainerWorkerPoolDedicatedHostCreate(clusterName, name, flavor, subnetID, vpcID, rgroupID, hostpoolID string) string {
	return fmt.Sprintf(`
	resource "ibm_container_vpc_worker_pool" "vpc_worker_pool" {
		cluster = "%s"
		flavor = "%s"
		worker_pool_name = "%s"
		zones {
		  subnet_id = "%s"
		  name      = "us-south-1"
		}
		worker_count      = 1
		vpc_id = "%s"
		resource_group_id = "%s"
		host_pool_id      = "%s"
	  }
	`, clusterName, flavor, name, subnetID, vpcID, rgroupID, hostpoolID)
}

// TestAccIBMContainerVpcClusterWorkerPoolResourceSecondaryStorage ...
func TestAccIBMContainerVpcClusterWorkerPoolResourceSecondaryStorage(t *testing.T) {
	if acc.WorkerPoolSecondaryStorage == "" {
		t.Fatal("IBM_WORKER_POOL_SECONDARY_STORAGE is unset")
		return
	}
	name := fmt.Sprintf("tf-vpc-wp-secstorage-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMVpcContainerWorkerPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMVpcContainerWorkerPoolSecStorage(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "flavor", "bx2.4x16"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "zones.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "secondary_storage", acc.WorkerPoolSecondaryStorage),
				),
			},
			{
				Config: testAccCheckIBMVpcContainerWorkerPoolSecStorageRemove(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "flavor", "bx2.4x16"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "zones.#", "1"),
					resource.TestCheckNoResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "secondary_storage"),
				),
			},

			{
				ResourceName:      "ibm_container_vpc_worker_pool.test_pool",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMVpcContainerWorkerPoolSecStorage(name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "resource_group" {
		is_default=true
	}
	
	resource "ibm_container_vpc_cluster" "cluster" {
	  name              = "%[1]s"
	  vpc_id            = "%[2]s"
	  flavor            = "bx2.4x16"
	  worker_count      = 1
	  resource_group_id = data.ibm_resource_group.resource_group.id
	  wait_till         = "MasterNodeReady"
	  zones {
		subnet_id = "%[3]s"
		name      = "us-south-1"
	  }
	}

	resource "ibm_container_vpc_worker_pool" "test_pool" {
	  cluster           = ibm_container_vpc_cluster.cluster.id
	  worker_pool_name  = "wp-sec-storage"
	  flavor            = "bx2.4x16"
	  vpc_id            = "%[2]s"
	  worker_count      = 1
	  zones {
		subnet_id = "%[3]s"
		name      = "us-south-1"
	  }
	  secondary_storage = "%[4]s"
	}
		`, name, acc.IksClusterVpcID, acc.IksClusterSubnetID, acc.WorkerPoolSecondaryStorage)
}

func testAccCheckIBMVpcContainerWorkerPoolSecStorageRemove(name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "resource_group" {
		is_default=true
	}
	
	resource "ibm_container_vpc_cluster" "cluster" {
	  name              = "%[1]s"
	  vpc_id            = "%[2]s"
	  flavor            = "bx2.4x16"
	  worker_count      = 1
	  resource_group_id = data.ibm_resource_group.resource_group.id
	  wait_till         = "MasterNodeReady"
	  zones {
		subnet_id = "%[3]s"
		name      = "us-south-1"
	  }
	}

	resource "ibm_container_vpc_worker_pool" "test_pool" {
	  cluster           = ibm_container_vpc_cluster.cluster.id
	  worker_pool_name  = "wp-sec-storage"
	  flavor            = "bx2.4x16"
	  vpc_id            = "%[2]s"
	  worker_count      = 1
	  zones {
		subnet_id = "%[3]s"
		name      = "us-south-1"
	  }
	}
		`, name, acc.IksClusterVpcID, acc.IksClusterSubnetID)
}

// TestAccIBMContainerVpcClusterWorkerPoolResourceKMS ...
func TestAccIBMContainerVpcClusterWorkerPoolResourceKMS(t *testing.T) {
	if acc.CrkID == "" {
		fmt.Println("[WARN] Skipping TestAccIBMContainerVpcClusterWorkerPoolResourceKMS - IBM_CRK_ID is unset")
		return
	}
	name := fmt.Sprintf("tf-vpc-wp-kms-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMVpcContainerWorkerPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMVpcContainerWorkerPoolKMS(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "flavor", "cx2.2x4"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "zones.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "kms_instance_id", acc.KmsInstanceID),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "crk", acc.CrkID),
				),
			},
			{
				ResourceName:      "ibm_container_vpc_worker_pool.test_pool",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMVpcContainerWorkerPoolKMS(name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "resource_group" {
		is_default=true
	}
	
	resource "ibm_container_vpc_cluster" "cluster" {
	  name              = "%[1]s"
	  vpc_id            = "%[2]s"
	  flavor            = "cx2.2x4"
	  worker_count      = 1
	  resource_group_id = data.ibm_resource_group.resource_group.id
	  wait_till         = "MasterNodeReady"
	  zones {
		subnet_id = "%[3]s"
		name      = "us-south-1"
	  }
	  kms_instance_id = "%[4]s"
	  crk = "%[5]s"
	}

	resource "ibm_container_vpc_worker_pool" "test_pool" {
	  cluster           = ibm_container_vpc_cluster.cluster.id
	  worker_pool_name  = "wp-kms"
	  flavor            = "cx2.2x4"
	  vpc_id            = "%[2]s"
	  worker_count      = 1
	  zones {
		subnet_id = "%[3]s"
		name      = "us-south-1"
	  }
	  kms_instance_id = "%[4]s"
	  crk = "%[5]s"
	}
		`, name, acc.IksClusterVpcID, acc.IksClusterSubnetID, acc.KmsInstanceID, acc.CrkID)
}

// TestAccIBMContainerVpcClusterWorkerPoolResourceKmsAccount ...
func TestAccIBMContainerVpcClusterWorkerPoolResourceKmsAccount(t *testing.T) {
	if acc.KmsAccountID == "" {
		fmt.Println("[WARN] Skipping TestAccIBMContainerVpcClusterWorkerPoolResourceKmsAccount - IBM_KMS_ACCOUNT_ID is unset")
		return
	}
	name := fmt.Sprintf("tf-vpc-wp-kmsacc-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMVpcContainerWorkerPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMVpcContainerWorkerPoolKmsAccount(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "flavor", "cx2.2x4"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "zones.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "kms_instance_id", acc.KmsInstanceID),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "kms_account_id", acc.KmsAccountID),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "crk", acc.CrkID),
				),
			},
			{
				ResourceName:            "ibm_container_vpc_worker_pool.test_pool",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"kms_account_id"},
			},
		},
	})
}

func testAccCheckIBMVpcContainerWorkerPoolKmsAccount(name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "resource_group" {
		is_default=true
	}
	
	resource "ibm_container_vpc_cluster" "cluster" {
	  name              = "%[1]s"
	  vpc_id            = "%[2]s"
	  flavor            = "cx2.2x4"
	  worker_count      = 1
	  resource_group_id = data.ibm_resource_group.resource_group.id
	  wait_till         = "MasterNodeReady"
	  zones {
		subnet_id = "%[3]s"
		name      = "us-south-1"
	  }
	  kms_instance_id = "%[4]s"
	  crk = "%[5]s"
	  kms_account_id = "%[6]s"
	}
	  
	resource "ibm_container_vpc_worker_pool" "test_pool" {
	  cluster           = ibm_container_vpc_cluster.cluster.id
	  worker_pool_name  = "wp-kms"
	  flavor            = "cx2.2x4"
	  vpc_id            = "%[2]s"
	  worker_count      = 1
	  zones {
		subnet_id = "%[3]s"
		name      = "us-south-1"
	  }
	  kms_instance_id = "%[4]s"
	  crk = "%[5]s"
	  kms_account_id = "%[6]s"
	}
		`, name, acc.IksClusterVpcID, acc.IksClusterSubnetID, acc.KmsInstanceID, acc.CrkID, acc.KmsAccountID)
}

// TestAccIBMContainerVpcOpenshiftClusterWorkerPoolBasic ...
func TestAccIBMContainerVpcOpenshiftClusterWorkerPoolResourceBasic(t *testing.T) {

	name := fmt.Sprintf("tf-vpc-oc-wp-basic-%d", acctest.RandIntRange(10, 100))
	openshiftFlavour := "bx2.16x64"
	openShiftworkerCount := "2"
	operatingSystem := "REDHAT_8_64"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMVpcContainerWorkerPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMOpcContainerWorkerPoolBasic(name, openshiftFlavour, openShiftworkerCount, operatingSystem),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "flavor", openshiftFlavour),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "zones.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "labels.%", "2"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "operating_system", operatingSystem),
				),
			},
			{
				ResourceName:            "ibm_container_vpc_worker_pool.test_pool",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"entitlement"},
			},
		},
	})
}

func testAccCheckIBMOpcContainerWorkerPoolBasic(name, openshiftFlavour, openShiftworkerCount, operatingSystem string) string {
	return testAccCheckIBMContainerOcpClusterBasic(name, openshiftFlavour, openShiftworkerCount, operatingSystem) +
		fmt.Sprintf(`

	resource "ibm_container_vpc_worker_pool" "test_pool" {
	  cluster           = ibm_container_vpc_cluster.cluster.id
	  worker_pool_name  = "%[1]s"
	  flavor            = "%[5]s"
	  worker_count      = "%[6]s"
	  vpc_id            = "%[2]s"
	  resource_group_id = "%[3]s"
 	  operating_system  = "%[7]s"
	  entitlement       = "cloud_pak"
	  zones {
		subnet_id = "%[4]s"
		name      = "us-south-1"
	  }
	  labels = {
		"test"  = "test-pool"
		"test1" = "test-pool1"
	  }
	}
		`, name, acc.IksClusterVpcID, acc.IksClusterResourceGroupID, acc.IksClusterSubnetID, openshiftFlavour, openShiftworkerCount, operatingSystem)
}
