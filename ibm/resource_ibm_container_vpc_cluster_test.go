package ibm

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	v2 "github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
	"github.com/IBM-Cloud/bluemix-go/session"
)

func TestAccIBMContainerVpcCluster_basic(t *testing.T) {
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	clusterNamegen2 := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	randint := acctest.RandIntRange(10, 100)
	vpc := fmt.Sprintf("terraform_vpc-%d", randint)
	subnet := fmt.Sprintf("terraform_subnet-%d", randint)
	flavor := "c2.2x4"
	zone := "us-south"
	workerCount := "1"
	flavorGen2 := "bx2.2x8"
	openshiftFlavour := "bx2.16x64"
	openShiftworkerCount := "2"
	var conf *v2.ClusterInfo

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMContainerVpcClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVpcCluster_basic(zone, vpc, subnet, clusterName, flavor, workerCount),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMContainerVpcExists("ibm_container_vpc_cluster.cluster", conf),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "name", clusterName),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "worker_count", workerCount),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "flavor", flavor),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "worker_labels.%", "2"),
				),
			},
			{
				Config: testAccCheckIBMContainerVpcCluster_update(zone, vpc, subnet, clusterName, flavor, workerCount),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMContainerVpcExists("ibm_container_vpc_cluster.cluster", conf),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "name", clusterName),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "worker_count", workerCount),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "flavor", flavor),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "worker_labels.%", "3"),
				),
			},
			{
				Config: testAccCheckIBMContainerVpcClusterGen2basic(zone, vpc, subnet, clusterNamegen2, flavorGen2, workerCount),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMContainerVpcExists("ibm_container_vpc_cluster.clustergen2", conf),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.clustergen2", "name", clusterNamegen2),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.clustergen2", "worker_count", workerCount),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.clustergen2", "flavor", flavorGen2),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "worker_labels.%", "2"),
				),
			},
			{
				Config: testAccCheckIBMContainerVpcOcpClusterGen2basic(zone, vpc, subnet, clusterNamegen2, openshiftFlavour, openShiftworkerCount),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMContainerVpcExists("ibm_container_vpc_cluster.clustergen2", conf),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.clustergen2", "name", clusterNamegen2),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.clustergen2", "worker_count", openShiftworkerCount),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.clustergen2", "flavor", openshiftFlavour),
				),
			},
		},
	})
}

func TestAccIBMContainerVpcCluster_importBasic(t *testing.T) {
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	randint := acctest.RandIntRange(10, 100)
	vpc := fmt.Sprintf("vpc-%d", randint)
	subnet := fmt.Sprintf("subnet-%d", randint)
	subnet1 := fmt.Sprintf("subnet1-%d", randint)
	flavor := "c2.2x4"
	zone := "us-south"
	zone1 := "us-south"
	workerCount := "1"
	var conf *v2.ClusterInfo
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMContainerVpcClusterDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMContainerVpcCluster_basic(zone, vpc, subnet, clusterName, flavor, workerCount),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMContainerVpcExists("ibm_container_vpc_cluster.cluster", conf),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "name", clusterName),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "worker_count", workerCount),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "flavor", flavor),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "zones.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "worker_labels.%", "2"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMContainerVpcClusterZone_update(zone, zone1, vpc, subnet, subnet1, clusterName, flavor, workerCount),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMContainerVpcExists("ibm_container_vpc_cluster.cluster", conf),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "name", clusterName),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "worker_count", workerCount),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "flavor", flavor),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "zones.#", "2"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "worker_labels.%", "3"),
				),
			},

			resource.TestStep{
				ResourceName:      "ibm_container_vpc_cluster.cluster",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"wait_till", "update_all_workers"},
			},
		},
	})
}

func TestAccIBMContainerVpcCluster_KmsEnable(t *testing.T) {
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	kmsInstanceName := fmt.Sprintf("kmsInstance_%d", acctest.RandIntRange(10, 100))
	rootKeyName := fmt.Sprintf("rootKey_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMContainerVpcClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVpcCluster_KmsEnable(clusterName, kmsInstanceName, rootKeyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.testacc_cluster", "name", clusterName),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.testacc_cluster", "kms_config.#", "1"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerVpcClusterDestroy(s *terraform.State) error {
	csClient, err := testAccProvider.Meta().(ClientSession).VpcContainerAPI()
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
			return fmt.Errorf("Error waiting for cluster (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMContainerVpcExists(n string, conf *v2.ClusterInfo) resource.TestCheckFunc {

	return func(s *terraform.State) error {

		csClient, err := testAccProvider.Meta().(ClientSession).VpcContainerAPI()
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

func testAccCheckIBMContainerVpcCluster_basic(zone, vpc, subnet, clusterName, flavor, workerCount string) string {
	return fmt.Sprintf(`
provider "ibm" {
  generation = 2
}
data "ibm_resource_group" "resource_group" {
  is_default = "true"
}

locals {
  ZONE1 = "%s-1"
}

resource "ibm_is_vpc" "vpc1" {
  name = "%s"
}

resource "ibm_is_subnet" "subnet1" {
  name                     = "%s"
  vpc                      = "${ibm_is_vpc.vpc1.id}"
  zone                     = "${local.ZONE1}"
  total_ipv4_address_count = 256
}

resource "ibm_container_vpc_cluster" "cluster" {
  name              = "%s"
  vpc_id            = "${ibm_is_vpc.vpc1.id}"
  flavor            = "%s"
  worker_count      = "%s"
  wait_till         = "OneWorkerNodeReady"
  resource_group_id = "${data.ibm_resource_group.resource_group.id}"
  zones {
	subnet_id = "${ibm_is_subnet.subnet1.id}"
	name      = "${local.ZONE1}"
  }

  worker_labels = {
	"test"  = "test-default-pool"
	"test1" = "test-default-pool1"
  }
}
`, zone, vpc, subnet, clusterName, flavor, workerCount)
}

func testAccCheckIBMContainerVpcCluster_update(zone, vpc, subnet, clusterName, flavor, workerCount string) string {
	return fmt.Sprintf(`
provider "ibm" {
  generation = 1
}
data "ibm_resource_group" "resource_group" {
  is_default = "true"
}

locals {
  ZONE1 = "%s-1"
}

resource "ibm_is_vpc" "vpc1" {
  name = "%s"
}

resource "ibm_is_subnet" "subnet1" {
  name                     = "%s"
  vpc                      = "${ibm_is_vpc.vpc1.id}"
  zone                     = "${local.ZONE1}"
  total_ipv4_address_count = 256
}

resource "ibm_resource_instance" "kms_instance1" {
  name     = "test_kms"
  service  = "kms"
  plan     = "tiered-pricing"
  location = "us-south"
}

resource "ibm_kms_key" "test" {
  instance_id  = "${ibm_resource_instance.kms_instance1.id}"
  key_name     = "test_root_key"
  standard_key = false
  force_delete = true
}

resource "ibm_container_vpc_cluster" "cluster" {
  name              = "%s"
  vpc_id            = "${ibm_is_vpc.vpc1.id}"
  flavor            = "%s"
  worker_count      = "%s"
  wait_till         = "OneWorkerNodeReady"
  resource_group_id = "${data.ibm_resource_group.resource_group.id}"
  zones {
	subnet_id = "${ibm_is_subnet.subnet1.id}"
	name      = "${local.ZONE1}"
  }

  worker_labels = {
	"test"  = "test-default-pool"
	"test1" = "test-default-pool1"
	"test2" = "test-default-pool2"
  }
}
`, zone, vpc, subnet, clusterName, flavor, workerCount)
}

func testAccCheckIBMContainerVpcCluster_KmsEnable(clusterName, KmsInstance, rootKey string) string {
	return fmt.Sprintf(`
provider "ibm" {
        generation =2
}
data "ibm_resource_group" "resource_group" {
        is_default = "true"
}

locals {
        ZONE1 = "us-south-1"
}

resource "ibm_is_vpc" "vpc1" {
        name = "testvpc"
}

resource "ibm_is_subnet" "subnet1" {
        name                     = "testsubnet"
        vpc                      = "${ibm_is_vpc.vpc1.id}"
        zone                     = "us-south-1"
        total_ipv4_address_count = 256
}

resource "ibm_resource_instance" "kms_instance1" {
	name              = "%s"
	service           = "kms"
	plan              = "tiered-pricing"
	location          = "us-south"
}
  
resource "ibm_kms_key" "test" {
	instance_id = "${ibm_resource_instance.kms_instance1.guid}"
	key_name = "%s"
	standard_key =  false
	force_delete = true
}

resource "ibm_container_vpc_cluster" "testacc_cluster" {
        name              = "%s"
        vpc_id            = "${ibm_is_vpc.vpc1.id}"
        flavor            = "bx2.2x8"
        worker_count      = "1"
        wait_till         = "OneWorkerNodeReady"
        resource_group_id = "${data.ibm_resource_group.resource_group.id}"
        zones {
                 subnet_id = "${ibm_is_subnet.subnet1.id}"
                 name      = "us-south-1"
          }

		kms_config {
			instance_id = ibm_resource_instance.kms_instance1.guid
			crk_id = ibm_kms_key.test.key_id
			private_endpoint = false
		}		
}`, KmsInstance, rootKey, clusterName)

}

func testAccCheckIBMContainerVpcClusterZone_update(zone, zone1, vpc, subnet, subnet1, clusterName, flavor, workerCount string) string {
	return fmt.Sprintf(`
provider "ibm" {
  generation = 1
}
data "ibm_resource_group" "resource_group" {
  is_default = "true"
}

locals {
  ZONE1 = "%s-1"
}
locals {
  ZONE2 = "%s-2"
}
resource "ibm_is_vpc" "vpc1" {
  name = "%s"
}

resource "ibm_is_subnet" "subnet1" {
  name                     = "%s"
  vpc                      = "${ibm_is_vpc.vpc1.id}"
  zone                     = "${local.ZONE1}"
  total_ipv4_address_count = 256
}
resource "ibm_is_subnet" "subnet2" {
  name                     = "%s"
  vpc                      = "${ibm_is_vpc.vpc1.id}"
  zone                     = "${local.ZONE2}"
  total_ipv4_address_count = 256
}
resource "ibm_container_vpc_cluster" "cluster" {
  name              = "%s"
  vpc_id            = "${ibm_is_vpc.vpc1.id}"
  flavor            = "%s"
  worker_count      = "%s"
  wait_till         = "OneWorkerNodeReady"
  resource_group_id = "${data.ibm_resource_group.resource_group.id}"
  zones {
	subnet_id = "${ibm_is_subnet.subnet1.id}"
	name      = "${local.ZONE1}"
  }
  zones {
	subnet_id = "${ibm_is_subnet.subnet2.id}"
	name      = "${local.ZONE2}"
  }
  worker_labels = {
	"test"  = "test-default-pool"
	"test1" = "test-default-pool1"
	"test2" = "test-default-pool2"
  }
}
`, zone, zone1, vpc, subnet, subnet1, clusterName, flavor, workerCount)

}

func testAccCheckIBMContainerVpcClusterGen2basic(zone, vpc, subnet, clusterName, flavor, workerCount string) string {
	return fmt.Sprintf(`
provider "ibm" {
	generation =2
}	
data "ibm_resource_group" "resource_group" {
	is_default = "true"
}

locals {
	ZONE1 = "%s-1"
}
  
resource "ibm_is_vpc" "vpc1" {
	name = "%s"
}
  
resource "ibm_is_subnet" "subnet1" {
	name                     = "%s"
	vpc                      = "${ibm_is_vpc.vpc1.id}"
	zone                     = "${local.ZONE1}"
	total_ipv4_address_count = 256
}

resource "ibm_resource_instance" "kms_instance1" {
    name              = "test_kms"
    service           = "kms"
    plan              = "tiered-pricing"
    location          = "us-south"
}
  
resource "ibm_kms_key" "test" {
    instance_id = "${ibm_resource_instance.kms_instance1.guid}"
    key_name = "test_root_key"
    standard_key =  false
    force_delete = true
}


resource "ibm_container_vpc_cluster" "clustergen2" {
	name              = "%s"
	vpc_id            = "${ibm_is_vpc.vpc1.id}"
	flavor            = "%s"
	worker_count      = "%s"
	kube_version 	  = "1.17.7"
	wait_till         = "OneWorkerNodeReady"
	resource_group_id = "${data.ibm_resource_group.resource_group.id}"
	zones {
		 subnet_id = "${ibm_is_subnet.subnet1.id}"
		 name      = "${local.ZONE1}"
	  }

	worker_labels = {
	"test"  = "test-default-pool"
	"test1" = "test-default-pool1"
	"test2" = "test-default-pool2"
	}
	
  }`, zone, vpc, subnet, clusterName, flavor, workerCount)

}

func testAccCheckIBMContainerVpcOcpClusterGen2basic(zone, vpc, subnet, clusterName, flavor, workerCount string) string {
	return fmt.Sprintf(`
provider "ibm" {
	generation =2
}	
data "ibm_resource_group" "resource_group" {
	is_default = "true"
}

locals {
	ZONE1 = "%s-1"
}
  
resource "ibm_is_vpc" "vpc1" {
	name = "%s"
}

resource "ibm_resource_instance" "cos_instance" {
	name     = "testcos_instance"
	service  = "cloud-object-storage"
	plan     = "standard"
	location = "global"
  }
  
resource "ibm_is_subnet" "subnet1" {
	name                     = "%s"
	vpc                      = "${ibm_is_vpc.vpc1.id}"
	zone                     = "${local.ZONE1}"
	total_ipv4_address_count = 256
}

resource "ibm_container_vpc_cluster" "clustergen2" {
	name              = "%s"
	vpc_id            = "${ibm_is_vpc.vpc1.id}"
	flavor            = "%s"
	worker_count      = "%s"
	kube_version 	  = "4.3.23_openshift"
	wait_till         = "IngressReady"
	entitlement       = "cloud_pak"
	cos_instance_crn  = ibm_resource_instance.cos_instance.id

	resource_group_id = "${data.ibm_resource_group.resource_group.id}"
	zones {
		 subnet_id = "${ibm_is_subnet.subnet1.id}"
		 name      = "${local.ZONE1}"
	  }
  }`, zone, vpc, subnet, clusterName, flavor, workerCount)

}
