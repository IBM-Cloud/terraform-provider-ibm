// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMContainerCluster_basic(t *testing.T) {
	clusterName := fmt.Sprintf("tf-cluster-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMContainerClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerClusterBasic(clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "name", clusterName),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "default_pool_size", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "kube_version", kubeVersion),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "hardware", "shared"),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "worker_pools.#", "1"),
					resource.TestCheckResourceAttrSet(
						"ibm_container_cluster.testacc_cluster", "resource_group_id"),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "labels.%", "2"),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "tags.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "workers_info.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMContainerClusterUpdate(clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "name", clusterName),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "default_pool_size", "2"),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "kube_version", kubeUpdateVersion),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "workers_info.0.version", kubeUpdateVersion),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "hardware", "shared"),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "worker_pools.#", "1"),
					resource.TestCheckResourceAttrSet(
						"ibm_container_cluster.testacc_cluster", "resource_group_id"),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "labels.%", "3"),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "tags.#", "2"),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "workers_info.#", "4"),
				),
			},
		},
	})
}

// Removed TestAccIBMContainerCluster_trusted  testcase as is_trusted parameter is deprecated
// Removed TestAccIBMContainerClusterWaitTill testcase as many other testcase config uses wait_till parameter
// Removed TestAccIBMContainerCluster_Tag testcase added tags to basic testcase
// Removed TestAccIBMContainerClusterOptionalOrgSpace_basic as no org and space required
// Removed TestAccIBMContainerCluster_worker_count as worker_num attribute is deprecated and check covered in basic testcase
// Removed TestAccIBMContainerCluster_nosubnet_false as by default no_subnet is false

func TestAccIBMContainerClusterKmsEnable(t *testing.T) {
	clusterName := fmt.Sprintf("tf-cluster-%d", acctest.RandIntRange(10, 100))
	kmsInstanceName := fmt.Sprintf("tf-cluster-%d", acctest.RandIntRange(10, 100))
	rootKeyName := fmt.Sprintf("rootKey-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMContainerClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerClusterKmsEnable(clusterName, kmsInstanceName, rootKeyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "name", clusterName),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "kms_config.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMContainerClusterWithWorkerNumZero(t *testing.T) {
	clusterName := fmt.Sprintf("tf-cluster-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMContainerClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMContainerClusterWithWorkerNumZero(clusterName),
				ExpectError: regexp.MustCompile("must be greater than 0"),
			},
		},
	})
}

func TestAccIBMContainerClusterDiskEnc(t *testing.T) {
	clusterName := fmt.Sprintf("tf-cluster-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMContainerClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerClusterDiskEnc(clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "name", clusterName),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "default_pool_size", "1"),
				),
			},
		},
	})
}

func TestAccIBMContainerClusterPrivateSubnet(t *testing.T) {
	t.Skip()
	clusterName := fmt.Sprintf("tf-cluster-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMContainerClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerClusterPrivateSubnet(clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "name", clusterName),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "default_pool_size", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "ingress_hostname", ""),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "ingress_secret", ""),
				),
			},
		},
	})
}

func TestAccIBMContainerClusterPrivateAndPublicSubnet(t *testing.T) {
	t.Skip()
	clusterName := fmt.Sprintf("tf-cluster-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMContainerClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerClusterPrivateAndPublicSubnet(clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "name", clusterName),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "default_pool_size", "1"),
					resource.TestCheckResourceAttrSet(
						"ibm_container_cluster.testacc_cluster", "ingress_hostname"),
					resource.TestCheckResourceAttrSet(
						"ibm_container_cluster.testacc_cluster", "ingress_secret"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerClusterDestroy(s *terraform.State) error {
	csClient, err := testAccProvider.Meta().(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_container_cluster" {
			continue
		}
		targetEnv := containerv1.ClusterTargetHeader{}
		_, err := csClient.Clusters().Find(rs.Primary.ID, targetEnv)

		if err == nil {
			return fmt.Errorf("Cluster still exists: %s", rs.Primary.ID)
		} else if !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("Error waiting for cluster (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMContainerClusterBasic(clusterName string) string {
	return fmt.Sprintf(`

data "ibm_resource_group" "testacc_ds_resource_group" {
  is_default = "true"
}

resource "ibm_container_cluster" "testacc_cluster" {
  name       = "%s"
  datacenter = "%s"
  resource_group_id = data.ibm_resource_group.testacc_ds_resource_group.id
  default_pool_size = 1
  hardware        = "shared"
  kube_version    = "%s"
  machine_type    = "%s"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
  no_subnet       = true
  tags            = ["test"]
  timeouts {
    create = "720m"
	update = "720m"
  }
  
}	`, clusterName, datacenter, kubeVersion, machineType, publicVlanID, privateVlanID)
}

func testAccCheckIBMContainerClusterKmsEnable(clusterName, kmsInstanceName, rootKeyName string) string {
	return fmt.Sprintf(`
	
	data "ibm_resource_group" "testacc_ds_resource_group" {
		name = "default"
	}
	resource "ibm_resource_instance" "kms_instance" {
		name              = "%s"
		service           = "kms"
		plan              = "tiered-pricing"
		location          = "us-south"
	}
	resource "ibm_kms_key" "test" {
		instance_id = "${ibm_resource_instance.kms_instance.guid}"
		key_name = "%s"
		standard_key =  false
		force_delete = true
	}
	
	resource "ibm_container_cluster" "testacc_cluster" {
		name              = "%s"
		datacenter        = "%s"
		wait_till       = "MasterNodeReady"
		default_pool_size = 1
		hardware          = "shared"
		resource_group_id = data.ibm_resource_group.testacc_ds_resource_group.id
		machine_type      = "%s"
		public_vlan_id    = "%s"
		private_vlan_id   = "%s"
		kms_config {
			instance_id = ibm_resource_instance.kms_instance.guid
			crk_id = ibm_kms_key.test.key_id
			private_endpoint = false
		}
	}

`, kmsInstanceName, rootKeyName, clusterName, datacenter, machineType, publicVlanID, privateVlanID)
}

func testAccCheckIBMContainerClusterWithWorkerNumZero(clusterName string) string {
	return fmt.Sprintf(`
resource "ibm_container_cluster" "testacc_cluster" {
  name       = "%s"
  datacenter = "%s"
  default_pool_size = 0
  machine_type      = "%s"
  hardware          = "shared"
  public_vlan_id    = "%s"
  private_vlan_id   = "%s"
  no_subnet         = true
}	`, clusterName, datacenter, machineType, publicVlanID, privateVlanID)
}

func testAccCheckIBMContainerClusterDiskEnc(clusterName string) string {
	return fmt.Sprintf(`
resource "ibm_container_cluster" "testacc_cluster" {
  name       = "%s"
  datacenter = "%s"
  machine_type    = "%s"
  hardware        = "shared"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
  disk_encryption = false
  wait_till         = "MasterNodeReady"
}	`, clusterName, datacenter, machineType, publicVlanID, privateVlanID)
}

func testAccCheckIBMContainerClusterUpdate(clusterName string) string {
	return fmt.Sprintf(`

data "ibm_resource_group" "testacc_ds_resource_group" {
  is_default = "true"
}

resource "ibm_container_cluster" "testacc_cluster" {
  name       = "%s"
  datacenter = "%s"
  default_pool_size = 2
  hardware           = "shared"
  resource_group_id  = data.ibm_resource_group.testacc_ds_resource_group.id
  kube_version       = "%s"
  machine_type       = "%s"
  public_vlan_id     = "%s"
  private_vlan_id    = "%s"
  no_subnet          = true
  update_all_workers = true
  tags            = ["test", "once"]
  timeouts {
    create = "720m"
	update = "720m"
  }
}	`, clusterName, datacenter, kubeUpdateVersion, machineType, publicVlanID, privateVlanID)
}

func testAccCheckIBMContainerClusterPrivateAndPublicSubnet(clusterName string) string {
	return fmt.Sprintf(`


resource "ibm_container_cluster" "testacc_cluster" {
  name       = "%s"
  datacenter = "%s"
  machine_type    = "%s"
  hardware        = "shared"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
  no_subnet       = true
  subnet_id       = ["%s", "%s"]
}	`, clusterName, datacenter, machineType, publicVlanID, privateVlanID, privateSubnetID, publicSubnetID)
}

func testAccCheckIBMContainerClusterPrivateSubnet(clusterName string) string {
	return fmt.Sprintf(`

resource "ibm_container_cluster" "testacc_cluster" {
  name       = "%s"
  datacenter = "%s"
  machine_type    = "%s"
  hardware        = "shared"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
  no_subnet       = true
  subnet_id       = ["%s"]
}	`, clusterName, datacenter, machineType, publicVlanID, privateVlanID, privateSubnetID)
}
