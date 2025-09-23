// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/Mavrickk3/bluemix-go/api/container/containerv1"
)

func TestAccIBMContainerCluster_basic(t *testing.T) {
	clusterName := fmt.Sprintf("tf-cluster-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMContainerClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerClusterBasic(clusterName, "masterNodeReady"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "name", clusterName),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "default_pool_size", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "hardware", "shared"),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "worker_pools.#", "1"),
					resource.TestCheckResourceAttrSet(
						"ibm_container_cluster.testacc_cluster", "resource_group_id"),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "labels.%", "2"),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "image_security_enforcement", "false"),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "workers_info.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMContainerClusterUpdate(clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "name", clusterName),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "default_pool_size", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "hardware", "shared"),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "worker_pools.#", "1"),
					resource.TestCheckResourceAttrSet(
						"ibm_container_cluster.testacc_cluster", "resource_group_id"),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "labels.%", "2"),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "image_security_enforcement", "true"),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "workers_info.#", "1"),
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
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
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
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
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
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
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
	clusterName := fmt.Sprintf("tf-cluster-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
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
	clusterName := fmt.Sprintf("tf-cluster-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
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

func TestAccIBMContainerClusterImageSecuritySetting(t *testing.T) {
	clusterName := fmt.Sprintf("tf-cluster-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMContainerClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerClusterImageSecuritySetting(clusterName, "true"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "name", clusterName),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "image_security_enforcement", "true"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerClusterDestroy(s *terraform.State) error {
	csClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ContainerAPI()
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
			return fmt.Errorf("[ERROR] Error waiting for cluster (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMContainerClusterBasic(clusterName, wait_till string) string {
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
  labels = {
    "test"  = "test-label"
    "test1" = "test-label1"
  }
  wait_till       = "%s"
  timeouts {
    create = "720m"
	update = "720m"
  }
  
}	`, clusterName, acc.Datacenter, acc.KubeVersion, acc.MachineType, acc.PublicVlanID, acc.PrivateVlanID, wait_till)
}

func testAccCheckIBMContainerClusterKmsEnable(clusterName, kmsInstanceName, rootKeyName string) string {
	return fmt.Sprintf(`
	
	data "ibm_resource_group" "testacc_ds_resource_group" {
		is_default=true
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

`, kmsInstanceName, rootKeyName, clusterName, acc.Datacenter, acc.MachineType, acc.PublicVlanID, acc.PrivateVlanID)
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
}	`, clusterName, acc.Datacenter, acc.MachineType, acc.PublicVlanID, acc.PrivateVlanID)
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
}	`, clusterName, acc.Datacenter, acc.MachineType, acc.PublicVlanID, acc.PrivateVlanID)
}

func testAccCheckIBMContainerClusterUpdate(clusterName string) string {
	return fmt.Sprintf(`

data "ibm_resource_group" "testacc_ds_resource_group" {
  is_default = "true"
}

resource "ibm_container_cluster" "testacc_cluster" {
  name       = "%s"
  datacenter = "%s"
  default_pool_size  = 2 # default_pool_size is applyonce, so should not modify anything in case of update
  hardware           = "shared"
  resource_group_id  = data.ibm_resource_group.testacc_ds_resource_group.id
  kube_version       = "%s"
  machine_type       = "%s"
  public_vlan_id     = "%s"
  private_vlan_id    = "%s"
  no_subnet          = true
  update_all_workers = true
  image_security_enforcement = true
  timeouts {
    create = "720m"
	update = "720m"
  }
}	`, clusterName, acc.Datacenter, acc.KubeUpdateVersion, acc.MachineType, acc.PublicVlanID, acc.PrivateVlanID)
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
}	`, clusterName, acc.Datacenter, acc.MachineType, acc.PublicVlanID, acc.PrivateVlanID, acc.PrivateSubnetID, acc.PublicSubnetID)
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
}	`, clusterName, acc.Datacenter, acc.MachineType, acc.PublicVlanID, acc.PrivateVlanID, acc.PrivateSubnetID)
}

func testAccCheckIBMContainerClusterImageSecuritySetting(clusterName string, imageSecuritySetting string) string {
	return fmt.Sprintf(`

resource "ibm_container_cluster" "testacc_cluster" {
  name       = "%s"
  datacenter = "%s"
  machine_type    = "%s"
  hardware        = "shared"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
  image_security_enforcement = %s
}	`, clusterName, acc.Datacenter, acc.MachineType, acc.PublicVlanID, acc.PrivateVlanID, imageSecuritySetting)
}
