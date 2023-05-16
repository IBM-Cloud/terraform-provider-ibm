// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	v2 "github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
)

func TestAccIBMContainerVpcClusterWorkerPoolBasic(t *testing.T) {

	name := fmt.Sprintf("tf-vpc-worker-%d", acctest.RandIntRange(10, 100))
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
				),
			},
			{
				Config: testAccCheckIBMVpcContainerWorkerPoolUpdate(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "flavor", "cx2.2x4"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "zones.#", "2"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "labels.%", "3"),
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

func TestAccIBMContainerVpcClusterWorkerPoolDedicatedHost(t *testing.T) {

	name := fmt.Sprintf("tf-vpc-worker-%d", acctest.RandIntRange(10, 100))
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

func testAccCheckIBMVpcContainerWorkerPoolBasic(name string) string {
	return fmt.Sprintf(`
	provider "ibm" {
		region="us-south"
	}
	data "ibm_resource_group" "resource_group" {
		is_default=true
	}
	data "ibm_is_vpc" "vpc" {
	  name = "cluster-squad-dallas-test"
	}
	
	data "ibm_is_subnet" "subnet1" {
	  name                     = "cluster-squad-dallas-test-01"
	}
	
	data "ibm_is_subnet" "subnet2" {
	  name                     = "cluster-squad-dallas-test-02"
	}
	
	resource "ibm_container_vpc_cluster" "cluster" {
	  name              = "%[1]s"
	  vpc_id            = data.ibm_is_vpc.vpc.id
	  flavor            = "cx2.2x4"
	  worker_count      = 1
	  resource_group_id = data.ibm_resource_group.resource_group.id
	  wait_till         = "MasterNodeReady"
	  zones {
		subnet_id = data.ibm_is_subnet.subnet1.id
		name      = "us-south-1"
	  }
	}
	resource "ibm_container_vpc_worker_pool" "test_pool" {
	  cluster           = ibm_container_vpc_cluster.cluster.id
	  worker_pool_name  = "%[1]s"
	  flavor            = "cx2.2x4"
	  vpc_id            = data.ibm_is_vpc.vpc.id
	  worker_count      = 1
	  resource_group_id = data.ibm_resource_group.resource_group.id
	  zones {
		name      = "us-south-2"
		subnet_id = data.ibm_is_subnet.subnet2.id
	  }
	  labels = {
		"test"  = "test-pool"
		"test1" = "test-pool1"
	  }
	}
		`, name)
}

func testAccCheckIBMVpcContainerWorkerPoolUpdate(name string) string {
	return fmt.Sprintf(`
	provider "ibm" {
		region="eu-de"
	}
	data "ibm_resource_group" "resource_group" {
		is_default=true
	}
	resource "ibm_is_vpc" "vpc" {
	  name = "%[1]s"
	}
	resource "ibm_is_subnet" "subnet1" {
	  name                     = "%[1]s-1"
	  vpc                      = ibm_is_vpc.vpc.id
	  zone                     = "eu-de-1"
	  total_ipv4_address_count = 256
	}
	resource "ibm_is_subnet" "subnet2" {
	  name                     = "%[1]s-2"
	  vpc                      = ibm_is_vpc.vpc.id
	  zone                     = "eu-de-2"
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
		name      = "eu-de-1"
	  }
	}
	resource "ibm_container_vpc_worker_pool" "test_pool" {
	  cluster           = ibm_container_vpc_cluster.cluster.id
	  worker_pool_name  = "%[1]s"
	  flavor            = "cx2.2x4"
	  vpc_id            = ibm_is_vpc.vpc.id
	  worker_count      = 1
	  resource_group_id = data.ibm_resource_group.resource_group.id
	  zones {
		name      = "eu-de-2"
		subnet_id = ibm_is_subnet.subnet2.id
	  }
	  zones {
		subnet_id = ibm_is_subnet.subnet1.id
		name      = "eu-de-1"
	  }
	  labels = {
		"test"  = "test-pool"
		"test1" = "test-pool1"
		"test2" = "test-pool2"
	  }
	}
		`, name)
}

func TestAccIBMContainerVpcClusterWorkerPoolEnvvar(t *testing.T) {

	name := fmt.Sprintf("tf-vpc-worker-%d", acctest.RandIntRange(10, 100))
	testChecks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttr(
			"ibm_container_vpc_worker_pool.test_pool", "flavor", "bx2.4x16"),
		resource.TestCheckResourceAttr(
			"ibm_container_vpc_worker_pool.test_pool", "zones.#", "1"),
		resource.TestCheckResourceAttr(
			"ibm_container_vpc_worker_pool.test_pool", "taints.#", "1"),
	}
	if acc.CrkID != "" {
		testChecks = append(testChecks,
			resource.TestCheckResourceAttr(
				"ibm_container_vpc_worker_pool.test_pool", "kms_instance_id", acc.KmsInstanceID),
			resource.TestCheckResourceAttr(
				"ibm_container_vpc_worker_pool.test_pool", "crk", acc.CrkID),
		)
	}
	if acc.WorkerPoolSecondaryStorage != "" {
		testChecks = append(testChecks, resource.TestCheckResourceAttr(
			"ibm_container_vpc_worker_pool.test_pool", "secondary_storage", acc.WorkerPoolSecondaryStorage),
		)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMVpcContainerWorkerPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMVpcContainerWorkerPoolEnvvar(name),
				Check:  resource.ComposeTestCheckFunc(testChecks...),
			},
			{
				Config: testAccCheckIBMVpcContainerWorkerPoolEnvvarUpdate(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "taints.#", "0"),
				),
			},
			{
				ResourceName:      "ibm_container_vpc_worker_pool.test_pool",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"kms_instance_id", "crk"},
			},
		},
	})
}

func TestAccIBMContainerVpcClusterWorkerPoolKmsAccountEnvvar(t *testing.T) {

	name := fmt.Sprintf("tf-vpc-worker-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMVpcContainerWorkerPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMVpcContainerWorkerPoolKmsAccountEnvvar(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "flavor", "bx2.4x16"),
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
				ResourceName:      "ibm_container_vpc_worker_pool.test_pool",
				ImportState:       true,
				ImportStateVerify: true,
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

func testAccCheckIBMVpcContainerWorkerPoolEnvvar(name string) string {
	return fmt.Sprintf(testAccCheckIBMContainerVpcClusterEnvvar(name)+`
	resource "ibm_container_vpc_worker_pool" "test_pool" {
	  cluster           = ibm_container_vpc_cluster.cluster.id
	  worker_pool_name  = "%[1]s"
	  flavor            = "bx2.4x16"
	  vpc_id            = "%[2]s"
	  worker_count      = 1
	  zones {
		subnet_id = "%[3]s"
		name      = "us-south-1"
	  }
	  kms_instance_id = "%[4]s"
	  crk = "%[5]s"
	  secondary_storage = "%[6]s"
	  taints {
		key    = "key1"
		value  = "value1"
		effect = "NoSchedule"
	  }
	}
		`, name, acc.IksClusterVpcID, acc.IksClusterSubnetID, acc.KmsInstanceID, acc.CrkID, acc.WorkerPoolSecondaryStorage)
}

func testAccCheckIBMVpcContainerWorkerPoolEnvvarUpdate(name string) string {
	return fmt.Sprintf(testAccCheckIBMContainerVpcClusterEnvvar(name)+`
	resource "ibm_container_vpc_worker_pool" "test_pool" {
	  cluster           = ibm_container_vpc_cluster.cluster.id
	  worker_pool_name  = "%[1]s"
	  flavor            = "bx2.4x16"
	  vpc_id            = "%[2]s"
	  worker_count      = 1
	  zones {
		subnet_id = "%[3]s"
		name      = "us-south-1"
	  }
	  kms_instance_id = "%[4]s"
	  crk = "%[5]s"
	  secondary_storage = "%[6]s"
	}
		`, name, acc.IksClusterVpcID, acc.IksClusterSubnetID, acc.KmsInstanceID, acc.CrkID, acc.WorkerPoolSecondaryStorage)
}

func testAccCheckIBMVpcContainerWorkerPoolKmsAccountEnvvar(name string) string {
	return fmt.Sprintf(`
	resource "ibm_container_vpc_worker_pool" "test_pool" {
	  cluster           = "%[2]s"
	  worker_pool_name  = "%[1]s"
	  flavor            = "bx2.4x16"
	  vpc_id            = "%[3]s"
	  worker_count      = 1
	  zones {
		subnet_id = "%[4]s"
		name      = "us-south-1"
	  }
	  kms_instance_id = "%[5]s"
	  crk = "%[6]s"
	  kms_account_id = "%[7]s"
	}
		`, name, acc.IksClusterID, acc.IksClusterVpcID, acc.IksClusterSubnetID, acc.KmsInstanceID, acc.CrkID, acc.KmsAccountID)
}

func TestAccIBMContainerVpcOpenshiftClusterWorkerPoolBasic(t *testing.T) {

	name := fmt.Sprintf("tf-vpc-cluster-%d", acctest.RandIntRange(10, 100))
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

func TestAccIBMContainerVpcClusterWorkerPoolImportOnCreateEnvvar(t *testing.T) {

	name := fmt.Sprintf("tf-vpc-worker-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMVpcContainerWorkerPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMOpcContainerWorkerPoolImportOnCreate(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "worker_pool_name", "default"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "labels.%", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "worker_count", "1"),
				),
			},
			{
				Config: testAccCheckIBMOpcContainerWorkerPoolImportOnCreateClusterUpdate(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "worker_pool_name", "default"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "labels.%", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "worker_count", "1"),
				),
			},
			{
				Config: testAccCheckIBMOpcContainerWorkerPoolImportOnCreateWPUpdate(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "worker_pool_name", "default"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "labels.%", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "worker_count", "3"),
				),
			},
		},
	})
}
func testAccCheckIBMOpcContainerWorkerPoolImportOnCreate(name string) string {
	return fmt.Sprintf(`
	resource "ibm_container_vpc_cluster" "cluster" {
		name              = "%[1]s"
		vpc_id            = "%[2]s"
		flavor            = "bx2.4x16"
		worker_count      = 1
		resource_group_id = "%[3]s"
		zones {
			subnet_id = "%[4]s"
			name      = "us-south-1"
		}
		wait_till = "normal"
		worker_labels = {
			"test"  = "test-pool"
		}
	}

	resource "ibm_container_vpc_worker_pool" "test_pool" {
		cluster           = ibm_container_vpc_cluster.cluster.id
		vpc_id            = "%[2]s"
		flavor            = "bx2.4x16"
		worker_count      = 1
		worker_pool_name  = "default"
		zones {
			subnet_id = "%[4]s"
			name      = "us-south-1"
		}
		import_on_create  = "true"
	}
	`, name, acc.IksClusterVpcID, acc.IksClusterResourceGroupID, acc.IksClusterSubnetID)
}

func testAccCheckIBMOpcContainerWorkerPoolImportOnCreateClusterUpdate(name string) string {
	return fmt.Sprintf(`
	resource "ibm_container_vpc_cluster" "cluster" {
		name              = "%[1]s"
		vpc_id            = "%[2]s"
		flavor            = "bx2.4x16"
		worker_count      = 3
		resource_group_id = "%[3]s"
		zones {
			subnet_id = "%[4]s"
			name      = "us-south-1"
		}
		wait_till = "normal"
		worker_labels = {
			"test"  = "test-pool"
		}
	}

	resource "ibm_container_vpc_worker_pool" "test_pool" {
		cluster           = ibm_container_vpc_cluster.cluster.id
		vpc_id            = "%[2]s"
		flavor            = "bx2.4x16"
		worker_count      = 1
		worker_pool_name  = "default"
		zones {
			subnet_id = "%[4]s"
			name      = "us-south-1"
		}
		import_on_create  = "true"
	}
	`, name, acc.IksClusterVpcID, acc.IksClusterResourceGroupID, acc.IksClusterSubnetID)
}

func testAccCheckIBMOpcContainerWorkerPoolImportOnCreateWPUpdate(name string) string {
	return fmt.Sprintf(`
	resource "ibm_container_vpc_cluster" "cluster" {
		name              = "%[1]s"
		vpc_id            = "%[2]s"
		flavor            = "bx2.4x16"
		worker_count      = 1
		resource_group_id = "%[3]s"
		zones {
			subnet_id = "%[4]s"
			name      = "us-south-1"
		}
		wait_till = "normal"
		worker_labels = {
			"test"  = "test-pool"
		}
	}

	resource "ibm_container_vpc_worker_pool" "test_pool" {
		cluster           = ibm_container_vpc_cluster.cluster.id
		vpc_id            = "%[2]s"
		flavor            = "bx2.4x16"
		worker_count      = 3
		worker_pool_name  = "default"
		zones {
			subnet_id = "%[4]s"
			name      = "us-south-1"
		}
		import_on_create  = "true"
	}
	`, name, acc.IksClusterVpcID, acc.IksClusterResourceGroupID, acc.IksClusterSubnetID)
}
