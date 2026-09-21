// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes_test

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	v2 "github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
	"github.com/IBM-Cloud/bluemix-go/bmxerror"
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
				Config: testAccCheckIBMContainerVpcClusterBasic(name, "OneWorkerNodeReady"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMContainerVpcClusterExists("ibm_container_vpc_cluster.cluster", conf),
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
					testAccCheckIBMContainerVpcClusterExists("ibm_container_vpc_cluster.cluster", conf),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "worker_count", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "flavor", "cx2.2x4"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "kms_config.#", "1"),
				),
			},
			{
				ResourceName:      "ibm_container_vpc_cluster.cluster",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"wait_till", "update_all_workers", "kms_config", "force_delete_storage", "wait_for_worker_update",
					"disable_outbound_traffic_protection", "flavor", "worker_count", "worker_labels", "zones",
				},
			},
		},
	})
}

func TestAccIBMContainerOpenshiftClusterBasic(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-cluster-%d", acctest.RandIntRange(10, 100))
	openshiftFlavour := "bx2.16x64"
	openShiftworkerCount := "2"
	operatingSystem := "REDHAT_8_64"
	var conf *v2.ClusterInfo

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMContainerVpcClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerOcpClusterBasic(name, openshiftFlavour, openShiftworkerCount, operatingSystem),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMContainerVpcClusterExists("ibm_container_vpc_cluster.cluster", conf),
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

func TestAccIBMContainerVpcClusterOfferingNotConfiguredIKS(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-cluster-%d", acctest.RandIntRange(10, 100))
	kubeVersion := "1.34.8"
	flavor := "bx2.2x8"
	workerCount := "1"
	var conf *v2.ClusterInfo

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMContainerVpcClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVpcClusterOffering(name, "", kubeVersion, flavor, workerCount),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMContainerVpcClusterExists("ibm_container_vpc_cluster.cluster", conf),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "offering", "kubernetes"),
				),
			},
		},
	})
}

func TestAccIBMContainerVpcClusterOfferingNotConfiguredROKS(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-cluster-%d", acctest.RandIntRange(10, 100))
	kubeVersion := "4.21_openshift"
	flavor := "bx3d.metal.64x256"
	workerCount := "2"
	var conf *v2.ClusterInfo

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMContainerVpcClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVpcClusterOffering(name, "", kubeVersion, flavor, workerCount),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMContainerVpcClusterExists("ibm_container_vpc_cluster.cluster", conf),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "offering", "openshift"),
				),
			},
		},
	})
}

func TestAccIBMContainerVpcClusterOfferingConfiguredROKS(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-cluster-%d", acctest.RandIntRange(10, 100))
	offering := "openshift"
	kubeVersion := "4.21_openshift"
	flavor := "bx3d.metal.64x256"
	workerCount := "2"
	var conf *v2.ClusterInfo

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMContainerVpcClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVpcClusterOffering(name, offering, kubeVersion, flavor, workerCount),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMContainerVpcClusterExists("ibm_container_vpc_cluster.cluster", conf),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "offering", offering),
				),
			},
		},
	})
}

func TestAccIBMContainerVpcClusterOfferingConfiguredROVS(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-cluster-%d", acctest.RandIntRange(10, 100))
	offering := "openshift-vs"
	kubeVersion := "4.21_openshift"
	flavor := "bx3d.metal.64x256"
	workerCount := "2"
	var conf *v2.ClusterInfo

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMContainerVpcClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVpcClusterOffering(name, offering, kubeVersion, flavor, workerCount),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMContainerVpcClusterExists("ibm_container_vpc_cluster.cluster", conf),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "offering", offering),
				),
			},
			{
				Config:      testAccCheckIBMContainerVpcClusterOffering(name, "openshift", kubeVersion, flavor, workerCount),
				ExpectError: regexp.MustCompile(`\[ERROR\] Modifying the 'offering' field after cluster creation is currently unsupported\.`),
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

func TestAccIBMContainerVpcClusterSecurityGroups(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-cluster-%d", acctest.RandIntRange(10, 100))
	var conf *v2.ClusterInfo

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVpcClusterSecurityGroups(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMContainerVpcClusterExists("ibm_container_vpc_cluster.cluster", conf),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "name", name),
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

func TestAccIBMContainerVPCClusterDisableOutboundTrafficProtection(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-cluster-%d", acctest.RandIntRange(10, 100))
	var conf *v2.ClusterInfo

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMContainerVpcClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVpcClusterDisableOutboundTrafficProtection(name, "1.29", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMContainerVpcClusterExists("ibm_container_vpc_cluster.cluster", conf),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "disable_outbound_traffic_protection", "false"),
				),
			},
		},
	})
}

func TestAccIBMContainerVPCClusterUpdateDisableOutboundTrafficProtection(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-cluster-%d", acctest.RandIntRange(10, 100))
	var conf *v2.ClusterInfo

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMContainerVpcClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVpcClusterDisableOutboundTrafficProtection(name, "1.30", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMContainerVpcClusterExists("ibm_container_vpc_cluster.cluster", conf),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "disable_outbound_traffic_protection", "true"),
				),
			},
			{
				Config: testAccCheckIBMContainerVpcClusterDisableOutboundTrafficProtectionUpdate(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMContainerVpcClusterExists("ibm_container_vpc_cluster.cluster", conf),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "disable_outbound_traffic_protection", "false"),
				),
			},
		},
	})
}

func TestAccIBMContainerVPCClusterEnableSecureByDefault(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-cluster-%d", acctest.RandIntRange(10, 100))
	var conf *v2.ClusterInfo

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMContainerVpcClusterDestroy,
		Steps: []resource.TestStep{
			{
				// First create a cluster with a version where Secure by Default is not available
				Config: testAccCheckIBMContainerVpcClusterEnableSecureByDefault(name, "1.29", "null"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMContainerVpcClusterExists("ibm_container_vpc_cluster.cluster", conf),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "name", name),
				),
			},
			{
				// Then update it to a version where Secure by Default is available, but not applied by default
				Config: testAccCheckIBMContainerVpcClusterEnableSecureByDefault(name, "1.30", "null"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMContainerVpcClusterExists("ibm_container_vpc_cluster.cluster", conf),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "name", name),
				),
			},
			{
				// Then enable it
				Config: testAccCheckIBMContainerVpcClusterEnableSecureByDefault(name, "1.30", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMContainerVpcClusterExists("ibm_container_vpc_cluster.cluster", conf),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "enable_secure_by_default", "true"),
				),
			},
		},
	})
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
		if err != nil {
			if apiErr, ok := err.(bmxerror.RequestFailure); ok && apiErr.StatusCode() == 404 {
				return nil
			}
			return fmt.Errorf("[ERROR] Error waiting for cluster (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
		return fmt.Errorf("Cluster still exists: %s", rs.Primary.ID)
	}

	return nil
}

func testAccCheckIBMContainerVpcClusterExists(n string, conf *v2.ClusterInfo) resource.TestCheckFunc {

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

func testAccCheckIBMContainerVpcClusterBasic(name, wait_till string) string {
	region := acc.Region()
	return fmt.Sprintf(`
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
	zone                     = "%[2]s-1"
	total_ipv4_address_count = 256
}
resource "ibm_resource_instance" "kms_instance" {
	name              = "%[1]s"
	service           = "kms"
	plan              = "tiered-pricing"
	location          = "%[2]s"
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
	wait_till         = "%[3]s"
	resource_group_id = data.ibm_resource_group.resource_group.id
	zones {
		 subnet_id = ibm_is_subnet.subnet.id
		 name      = "%[2]s-1"
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

  }`, name, region, wait_till)
}

func testAccCheckIBMContainerVpcClusterUpdate(name string) string {
	region := acc.Region()
	return fmt.Sprintf(`
data "ibm_resource_group" "resource_group" {
	is_default = "true"
}
resource "ibm_is_vpc" "vpc" {
	name = "%[1]s"
}
resource "ibm_is_subnet" "subnet" {
	name                     = "%[1]s"
	vpc                      = ibm_is_vpc.vpc.id
	zone                     = "%[2]s-1"
	total_ipv4_address_count = 256
}
resource "ibm_is_subnet" "subnet2" {
	name                     = "%[1]s-2"
	vpc                      = ibm_is_vpc.vpc.id
	zone                     = "%[2]s-2"
	total_ipv4_address_count = 256
}
resource "ibm_resource_instance" "kms_instance" {
	name              = "%[1]s"
	service           = "kms"
	plan              = "tiered-pricing"
	location          = "%[2]s"
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
		 name      = "%[2]s-1"
	}
	zones {
		subnet_id = ibm_is_subnet.subnet2.id
		name      = "%[2]s-2"
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

  }`, name, region)
}

func testAccCheckIBMContainerVpcClusterDisableOutboundTrafficProtection(name, kubeVersion, disable_outbound_traffic_protection string) string {
	return fmt.Sprintf(`
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
	zone                     = "us-south-1"
	total_ipv4_address_count = 256
}
resource "ibm_resource_instance" "kms_instance" {
	name              = "%[1]s"
	service           = "kms"
	plan              = "tiered-pricing"
	location          = "us-south"
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
	kube_version      = "%[2]s"
	wait_till         = "OneWorkerNodeReady"
	resource_group_id = data.ibm_resource_group.resource_group.id
	zones {
			subnet_id = ibm_is_subnet.subnet.id
			name      = "us-south-1"
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
	disable_outbound_traffic_protection = "%[3]s"

}`, name, kubeVersion, disable_outbound_traffic_protection)
}

func testAccCheckIBMContainerVpcClusterEnableSecureByDefault(name, kubeVersion, enable_secure_by_default string) string {
	return fmt.Sprintf(`
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
	zone                     = "us-south-1"
	total_ipv4_address_count = 256
}
resource "ibm_resource_instance" "kms_instance" {
	name              = "%[1]s"
	service           = "kms"
	plan              = "tiered-pricing"
	location          = "us-south"
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
	kube_version      = "%[2]s"
	wait_till         = "OneWorkerNodeReady"
	resource_group_id = data.ibm_resource_group.resource_group.id
	zones {
			subnet_id = ibm_is_subnet.subnet.id
			name      = "us-south-1"
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
	enable_secure_by_default = %[3]s

}`, name, kubeVersion, enable_secure_by_default)
}

func testAccCheckIBMContainerVpcClusterDisableOutboundTrafficProtectionUpdate(name, disable_outbound_traffic_protection string) string {
	return fmt.Sprintf(`
data "ibm_resource_group" "resource_group" {
	is_default = "true"
}
resource "ibm_is_vpc" "vpc" {
	name = "%[1]s"
}
resource "ibm_is_subnet" "subnet" {
	name                     = "%[1]s"
	vpc                      = ibm_is_vpc.vpc.id
	zone                     = "us-south-1"
	total_ipv4_address_count = 256
}
resource "ibm_resource_instance" "kms_instance" {
	name              = "%[1]s"
	service           = "kms"
	plan              = "tiered-pricing"
	location          = "us-south"
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
			name      = "us-south-1"
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
	disable_outbound_traffic_protection = "%[2]s"

}`, name, disable_outbound_traffic_protection)
}

// previously you had to create securitygroups and use them instead
func testAccCheckIBMContainerVpcClusterSecurityGroups(name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "resource_group" {
		is_default = "true"
		//name = "Default"
	}
	resource "ibm_is_vpc" "vpc" {
		name = "%[1]s"
	}
	resource "ibm_is_security_group" "security_group" {
		name = "example-security-group"
		vpc  = ibm_is_vpc.vpc.id
	}
	resource "ibm_is_subnet" "subnet" {
		name                     = "%[1]s"
		vpc                      = ibm_is_vpc.vpc.id
		zone                     = "us-south-1"
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
			 name      = "us-south-1"
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

		security_groups = [
			ibm_is_security_group.security_group.id,
			"cluster",
		]
	}`, name)
}

func testAccCheckIBMContainerOcpClusterBasic(name, openshiftFlavour, openShiftworkerCount, operatingSystem string) string {
	return fmt.Sprintf(`
data "ibm_resource_instance" "cos_instance" {
	name     = "%[5]s"
}

resource "ibm_container_vpc_cluster" "cluster" {
	name              = "%[1]s"
	vpc_id            = "%[2]s"
	flavor            = "%[6]s"
	worker_count      = "%[7]s"
	kube_version      = "4.11_openshift"
 	operating_system  = "%[8]s"
	wait_till         = "OneWorkerNodeReady"
	entitlement       = "cloud_pak"
	cos_instance_crn  = data.ibm_resource_instance.cos_instance.id
	resource_group_id = "%[3]s"
	zones {
		 subnet_id = "%[4]s"
		 name      = "us-south-1"
	  }
  }
  data "ibm_container_cluster_config" "testacc_ds_cluster" {
	cluster_name_id = ibm_container_vpc_cluster.cluster.id
  }
  `, name, acc.IksClusterVpcID, acc.IksClusterResourceGroupID, acc.IksClusterSubnetID, acc.CosName, openshiftFlavour, openShiftworkerCount, operatingSystem)

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

func testAccCheckIBMContainerVpcClusterOffering(name, offering, kubeVersion, flavor, workerCount string) string {
	// offering is omitted from the config if the param is an empty string
	offeringConfig := ""
	if offering != "" {
		offeringConfig = fmt.Sprintf(`offering = "%s"`, offering)
	}

	return fmt.Sprintf(`
	resource "ibm_container_vpc_cluster" "cluster" {
		name              = "%s"
		vpc_id            = "%s"
		flavor            = "%s"
		worker_count      = "%s"
		kube_version      = "%s"
		resource_group_id = "%s"
		wait_till         = "MasterNodeReady"
		zones {
			subnet_id = "%s"
			name      = "us-south-1"
		}
		%s
	}`, name, acc.IksClusterVpcID, flavor, workerCount, kubeVersion, acc.IksClusterResourceGroupID, acc.IksClusterSubnetID, offeringConfig)
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

	testChecks := []resource.TestCheckFunc{
		testAccCheckIBMContainerVpcClusterExists("ibm_container_vpc_cluster.cluster", conf),
		resource.TestCheckResourceAttr(
			"ibm_container_vpc_cluster.cluster", "name", name),
		resource.TestCheckResourceAttr(
			"ibm_container_vpc_cluster.cluster", "worker_count", "1"),
		resource.TestCheckResourceAttr(
			"ibm_container_vpc_cluster.cluster", "taints.#", "1"),
	}
	if acc.WorkerPoolSecondaryStorage != "" {
		testChecks = append(testChecks, resource.TestCheckResourceAttr(
			"ibm_container_vpc_cluster.cluster", "secondary_storage", acc.WorkerPoolSecondaryStorage))
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMContainerVpcClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVpcClusterEnvvar(name),
				Check:  resource.ComposeTestCheckFunc(testChecks...),
			},
			{
				ResourceName:      "ibm_container_vpc_cluster.cluster",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"wait_till", "update_all_workers", "kms_config", "force_delete_storage", "wait_for_worker_update",
					"crk", "kms_account_id", "kms_instance_id",
				},
			},
		},
	})
}

// This test is here to help to focus on given resources, but requires everything else existing already
func TestAccIBMContainerVpcClusterBaseEnvvar(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-cluster-%d", acctest.RandIntRange(10, 100))
	var conf *v2.ClusterInfo

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMContainerVpcClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVpcClusterBaseEnvvar(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMContainerVpcClusterExists("ibm_container_vpc_cluster.cluster", conf),
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
					"wait_till", "update_all_workers", "kms_config", "force_delete_storage", "wait_for_worker_update", "albs", "disable_outbound_traffic_protection",
					//workerpool fields
					"zones", "worker_count", "flavor"},
			},
		},
	})
}

func TestAccIBMContainerVpcClusterKMSEnvvar(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-cluster-%d", acctest.RandIntRange(10, 100))
	var conf *v2.ClusterInfo

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMContainerVpcClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVpcClusterKMSEnvvar(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMContainerVpcClusterExists("ibm_container_vpc_cluster.cluster", conf),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "worker_count", "1"),
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

func TestAccIBMContainerVpcClusterNetworkPluginNotConfigured(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-cluster-%d", acctest.RandIntRange(10, 100))
	defaultNetworkPlugin := "Calico"
	var conf *v2.ClusterInfo

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMContainerVpcClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVpcClusterNetworkPlugin(name, ""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMContainerVpcClusterExists("ibm_container_vpc_cluster.cluster", conf),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "network_plugin", defaultNetworkPlugin),
				),
			},
		},
	})
}

func TestAccIBMContainerVpcClusterNetworkPluginConfigured(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-cluster-%d", acctest.RandIntRange(10, 100))
	networkPlugin := "OVNKubernetes"
	var conf *v2.ClusterInfo

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMContainerVpcClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVpcClusterNetworkPlugin(name, networkPlugin),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMContainerVpcClusterExists("ibm_container_vpc_cluster.cluster", conf),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "network_plugin", networkPlugin),
				),
			},
		},
	})
}

// You need to set up env vars:
// export IBM_CLUSTER_VPC_ID
// export IBM_CLUSTER_VPC_SUBNET_ID
// export IBM_CLUSTER_VPC_RESOURCE_GROUP_ID
// optionally for kms and cross account kms:
// export IBM_KMS_INSTANCE_ID
// export IBM_CRK_ID
// for cross account kms:
// export IBM_KMS_ACCOUNT_ID
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
		wait_till = "normal"
		kms_instance_id = "%[5]s"
		crk = "%[6]s"
		kms_account_id = "%[7]s"
		secondary_storage = "%[8]s"
		taints {
			key    = "key1"
			value  = "value1"
			effect = "NoSchedule"
		  }
	}
	`, name, acc.IksClusterVpcID, acc.IksClusterResourceGroupID, acc.IksClusterSubnetID, acc.KmsInstanceID, acc.CrkID, acc.KmsAccountID, acc.WorkerPoolSecondaryStorage)
	fmt.Println(config)
	return config
}

// You need to set up env vars:
// export IBM_CLUSTER_VPC_ID
// export IBM_CLUSTER_VPC_SUBNET_ID
// export IBM_CLUSTER_VPC_RESOURCE_GROUP_ID
// optionally for kms and cross account kms:
// export IBM_KMS_INSTANCE_ID
// export IBM_CRK_ID
// for cross account kms:
// export IBM_KMS_ACCOUNT_ID
func testAccCheckIBMContainerVpcClusterBaseEnvvar(name string) string {
	var kmsConfig string
	if acc.KmsInstanceID != "" {
		kmsConfig = fmt.Sprintf(`
		kms_config {
			instance_id = "%[1]s"
			crk_id = "%[2]s"
			account_id = "%[3]s"
			wait_for_apply = "true"
		}
	`, acc.KmsInstanceID, acc.CrkID, acc.KmsAccountID)
	}
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
		wait_till = "IngressReady"
		%[5]s
	}
	`, name, acc.IksClusterVpcID, acc.IksClusterResourceGroupID, acc.IksClusterSubnetID, kmsConfig)

	fmt.Println(config)
	return config
}

// You need to set up env vars:
// export IBM_CLUSTER_VPC_ID
// export IBM_CLUSTER_VPC_SUBNET_ID
// export IBM_CLUSTER_VPC_RESOURCE_GROUP_ID
// export IBM_KMS_INSTANCE_ID
// export IBM_CRK_ID
func testAccCheckIBMContainerVpcClusterKMSEnvvar(name string) string {
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
		wait_till = "normal"
		kms_config {
			instance_id = "%[5]s"
			crk_id = "%[6]s"
			private_endpoint = false
		}
	}
	`, name, acc.IksClusterVpcID, acc.IksClusterResourceGroupID, acc.IksClusterSubnetID, acc.KmsInstanceID, acc.CrkID)
	fmt.Println(config)
	return config
}

func testAccCheckIBMContainerVpcClusterNetworkPlugin(name, networkPlugin string) string {
	region := acc.Region()

	// networkPlugin is omitted from the config if the param is an empty string
	networkPluginConfig := ""
	if networkPlugin != "" {
		networkPluginConfig = fmt.Sprintf(`network_plugin = "%s"`, networkPlugin)
	}

	return fmt.Sprintf(`
data "ibm_resource_group" "resource_group" {
	is_default = "true"
}
resource "ibm_is_vpc" "vpc" {
	name = "%[1]s"
}
resource "ibm_is_subnet" "subnet" {
	name                     = "%[1]s"
	vpc                      = ibm_is_vpc.vpc.id
	zone                     = "%[2]s-1"
	total_ipv4_address_count = 256
}
resource "ibm_container_vpc_cluster" "cluster" {
	name              = "%[1]s"
	vpc_id            = ibm_is_vpc.vpc.id
	flavor            = "cx2.2x4"
	worker_count      = 1
	kube_version      = "4.20_openshift"
	wait_till         = "OneWorkerNodeReady"
	resource_group_id = data.ibm_resource_group.resource_group.id
	zones {
		subnet_id = ibm_is_subnet.subnet.id
		name      = "%[2]s-1"
	}
	%[3]s
}`, name, region, networkPluginConfig)
}
