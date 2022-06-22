// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes_test

import (
	"fmt"
	"log"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	v2 "github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
	"github.com/IBM-Cloud/bluemix-go/session"
)

func TestAccIBMContainerVpcClusterBasic(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-cluster-%d", acctest.RandIntRange(10, 100))
	var conf *v2.ClusterInfo

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMContainerVpcClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVpcClusterBasic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMContainerVpcExists("ibm_container_vpc_cluster.cluster", conf),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "worker_count", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "flavor", "cx2.2x4"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "zones.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "worker_labels.%", "3"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "kms_config.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMContainerVpcClusterUpdate(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMContainerVpcExists("ibm_container_vpc_cluster.cluster", conf),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "worker_count", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "flavor", "cx2.2x4"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "zones.#", "2"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "worker_labels.%", "2"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "kms_config.#", "1"),
				),
			},
			{
				ResourceName:      "ibm_container_vpc_cluster.cluster",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"wait_till", "update_all_workers", "kms_config", "force_delete_storage", "wait_for_worker_update"},
			},
		},
	})
}
func TestAccIBMContainerOpenshiftClusterBasic(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-cluster-%d", acctest.RandIntRange(10, 100))
	openshiftFlavour := "bx2.16x64"
	openShiftworkerCount := "2"
	var conf *v2.ClusterInfo

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMContainerVpcClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerOcpClusterBasic(name, openshiftFlavour, openShiftworkerCount),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMContainerVpcExists("ibm_container_vpc_cluster.cluster", conf),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "worker_count", openShiftworkerCount),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "flavor", openshiftFlavour),
				),
			},
		},
	})
}

func TestAccIBMContainerVpcClusterImageSecuritySetting(t *testing.T) {
	clusterName := fmt.Sprintf("tf-vpc-cluster-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMContainerVpcClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVpcClusterImageSecuritySetting(clusterName, "true"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.testacc_vpc_cluster", "name", clusterName),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.testacc_vpc_cluster", "image_security_enforcement", "true"),
				),
			},
		},
	})
}

func TestAccIBMContainerVpcClusterDedicatedHost(t *testing.T) {
	clusterName := fmt.Sprintf("tf-vpc-cluster-dhost-%d", acctest.RandIntRange(10, 100))
	hostPoolID := acc.HostPoolID
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMContainerVpcClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVpcClusterDedicatedHostSetting(
					clusterName,
					acc.IksClusterVpcID,
					"bx2d.4x16",
					acc.IksClusterSubnetID,
					acc.IksClusterResourceGroupID,
					hostPoolID,
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.testacc_dhost_vpc_cluster", "host_pool_id", hostPoolID),
				),
			},
		},
	},
	)
}

func testAccCheckIBMContainerVpcClusterDestroy(s *terraform.State) error {
	csClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_container_vpc_cluster" {
			continue
		}

		targetEnv := getVpcClusterTargetHeaderTestACC()
		// Try to find the key
		_, err := csClient.Clusters().GetCluster(rs.Primary.ID, targetEnv)

		if err == nil {
			return fmt.Errorf("Cluster still exists: %s", rs.Primary.ID)
		} else if !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("[ERROR] Error waiting for cluster (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
func testAccCheckIBMContainerVpcExists(n string, conf *v2.ClusterInfo) resource.TestCheckFunc {

	return func(s *terraform.State) error {

		csClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcContainerAPI()
		if err != nil {
			return err
		}

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_container_vpc_cluster" {
				continue
			}

			targetEnv := getVpcClusterTargetHeaderTestACC()

			cls, err := csClient.Clusters().GetCluster(rs.Primary.ID, targetEnv)

			if err != nil && !strings.Contains(err.Error(), "404") {
				return err
			}

			conf = cls

		}
		return nil
	}
}
func getVpcClusterTargetHeaderTestACC() v2.ClusterTargetHeader {
	c := new(bluemix.Config)
	sess, err := session.New(c)
	if err != nil {
		log.Fatal(err)
	}
	resourceGroup := sess.Config.ResourceGroup
	targetEnv := v2.ClusterTargetHeader{
		ResourceGroup: resourceGroup,
	}
	return targetEnv
}
func testAccCheckIBMContainerVpcClusterBasic(name string) string {
	return fmt.Sprintf(`
provider "ibm" {
	region ="eu-de"
}	
data "ibm_resource_group" "resource_group" {
	is_default = "true"
	//name = "Default"
}
resource "ibm_is_vpc" "vpc" {
	name = "%[1]s"
}
resource "ibm_is_subnet" "subnet" {
	name                     = "%[1]s"
	vpc                      = ibm_is_vpc.vpc.id
	zone                     = "eu-de-1"
	total_ipv4_address_count = 256
}
resource "ibm_resource_instance" "kms_instance" {
	name              = "%[1]s"
	service           = "kms"
	plan              = "tiered-pricing"
	location          = "eu-de"
}
  
resource "ibm_kms_key" "test" {
	instance_id = ibm_resource_instance.kms_instance.guid
	key_name = "%[1]s"
	standard_key =  false
	force_delete = true
}
resource "ibm_container_vpc_cluster" "cluster" {
	name              = "%[1]s"
	vpc_id            = ibm_is_vpc.vpc.id
	flavor            = "cx2.2x4"
	worker_count      = 1
	wait_till         = "OneWorkerNodeReady"
	resource_group_id = data.ibm_resource_group.resource_group.id
	zones {
		 subnet_id = ibm_is_subnet.subnet.id
		 name      = "eu-de-1"
	}
	kms_config {
		instance_id = ibm_resource_instance.kms_instance.guid
		crk_id = ibm_kms_key.test.key_id
		private_endpoint = false
	}
	worker_labels = {
	"test"  = "test-default-pool"
	"test1" = "test-default-pool1"
	"test2" = "test-default-pool2"
	}
	
  }`, name)
}
func testAccCheckIBMContainerVpcClusterUpdate(name string) string {
	return fmt.Sprintf(`
provider "ibm" {
	region ="eu-de"
}	
data "ibm_resource_group" "resource_group" {
	is_default = "true"
}
resource "ibm_is_vpc" "vpc" {
	name = "%[1]s"
}
resource "ibm_is_subnet" "subnet" {
	name                     = "%[1]s"
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
resource "ibm_resource_instance" "kms_instance" {
	name              = "%[1]s"
	service           = "kms"
	plan              = "tiered-pricing"
	location          = "eu-de"
}
  
resource "ibm_kms_key" "test" {
	instance_id = ibm_resource_instance.kms_instance.guid
	key_name = "%[1]s"
	standard_key =  false
	force_delete = true
}
resource "ibm_container_vpc_cluster" "cluster" {
	name              = "%[1]s"
	vpc_id            = ibm_is_vpc.vpc.id
	flavor            = "cx2.2x4"
	worker_count      = 1
	wait_till         = "OneWorkerNodeReady"
	resource_group_id = data.ibm_resource_group.resource_group.id
	zones {
		 subnet_id = ibm_is_subnet.subnet.id
		 name      = "eu-de-1"
	}
	zones {
		subnet_id = ibm_is_subnet.subnet2.id
		name      = "eu-de-2"
	}
	kms_config {
		instance_id = ibm_resource_instance.kms_instance.guid
		crk_id = ibm_kms_key.test.key_id
		private_endpoint = false
	}
	worker_labels = {
	"test"  = "test-default-pool"
	"test1" = "test-default-pool1"
	}
	
  }`, name)
}
func testAccCheckIBMContainerOcpClusterBasic(name, openshiftFlavour, openShiftworkerCount string) string {
	return fmt.Sprintf(`
provider "ibm" {
	region="eu-de"
}	
data "ibm_resource_group" "resource_group" {
	is_default = "true"
}
resource "ibm_is_vpc" "vpc1" {
	name = "%[1]s"
}
resource "ibm_is_subnet" "subnet" {
	name                     = "%[1]s"
	vpc                      = "${ibm_is_vpc.vpc1.id}"
	zone                     = "eu-de-1"
	total_ipv4_address_count = 256
}
resource "ibm_resource_instance" "cos_instance" {
	name     = "testcos_instance"
	service  = "cloud-object-storage"
	plan     = "standard"
	location = "global"
}
resource "ibm_container_vpc_cluster" "cluster" {
	name              = "%[1]s"
	vpc_id            = "${ibm_is_vpc.vpc1.id}"
	flavor            = "%s"
	worker_count      = "%s"
	kube_version 	  = "4.3.23_openshift"
	wait_till         = "IngressReady"
	entitlement       = "cloud_pak"
	cos_instance_crn  = ibm_resource_instance.cos_instance.id
	resource_group_id = data.ibm_resource_group.resource_group.id
	zones {
		 subnet_id = ibm_is_subnet.subnet.id
		 name      = "eu-de-1"
	  }
  }
  data "ibm_container_cluster_config" "testacc_ds_cluster" {
	cluster_name_id = ibm_container_vpc_cluster.cluster.id
  }
  `, name, openshiftFlavour, openShiftworkerCount)

}

func testAccCheckIBMContainerVpcClusterImageSecuritySetting(name, setting string) string {
	return fmt.Sprintf(`
	resource "ibm_container_vpc_cluster" "testacc_vpc_cluster" {
		name              = "%s"
		vpc_id            = "%s"
		flavor            = "bx2.2x8"
		worker_count      = "1"
		resource_group_id = "%s"
		zones {
			subnet_id = "%s"
			name      = "us-south-1"
		  }
		image_security_enforcement = %s
	  }`, name, acc.IksClusterVpcID, acc.IksClusterResourceGroupID, acc.SubnetID, setting)
}

func testAccCheckIBMContainerVpcClusterDedicatedHostSetting(name, vpcID, flavor, subnetID, rgroupID, hostpoolID string) string {
	return fmt.Sprintf(`
	resource "ibm_container_vpc_cluster" "testacc_dhost_vpc_cluster" {
		name = "%s"
		vpc_id = "%s"
		flavor = "%s"
		zones {
		  subnet_id = "%s"
		  name      = "us-south-1"
		}
		resource_group_id = "%s"
		host_pool_id = "%s"
	}`, name, vpcID, flavor, subnetID, rgroupID, hostpoolID)
}

// This test is here to help to focus on given resources, but requires everything else existing already
func TestAccIBMContainerVpcClusterEnvvar(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-cluster-%d", acctest.RandIntRange(10, 100))
	var conf *v2.ClusterInfo

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMContainerVpcClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVpcClusterEnvvar(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMContainerVpcExists("ibm_container_vpc_cluster.cluster", conf),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "worker_count", "1"),
				),
			},
			{
				ResourceName:      "ibm_container_vpc_cluster.cluster",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"wait_till", "update_all_workers", "kms_config", "force_delete_storage", "wait_for_worker_update"},
			},
		},
	})
}

// You need to set up env vars:
// export IBM_CLUSTER_VPC_ID
// export IBM_CLUSTER_VPC_SUBNET_ID
// export IBM_CLUSTER_VPC_RESOURCE_GROUP_ID
// export IBM_KMS_INSTANCE_ID
// export IBM_CRK_ID
// for acc.IksClusterVpcID, acc.IksClusterResourceGroupID, acc.IksClusterSubnetID, acc.KmsInstanceID, acc.CrkID
func testAccCheckIBMContainerVpcClusterEnvvar(name string) string {
	config := fmt.Sprintf(`
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
		kms_instance_id = "%[5]s"
		crk = "%[6]s"
	}
	`, name, acc.IksClusterVpcID, acc.IksClusterResourceGroupID, acc.IksClusterSubnetID, acc.KmsInstanceID, acc.CrkID)
	fmt.Println(config)
	return config
}
