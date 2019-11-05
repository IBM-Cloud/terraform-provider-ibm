package ibm

import (
	"fmt"
	"log"
	"strings"
	"testing"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"

	v2 "github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccIBMContainerVpcCluster_basic(t *testing.T) {
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	flavor := fmt.Sprintf("c.%d", acctest.RandInt())
	zone := "us-south"
	workerCount := fmt.Sprintf("%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMContainerVpcClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVpcCluster_basic(zone, clusterName, flavor, workerCount),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.testacc_cluster", "name", clusterName),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.testacc_cluster", "vpc_id", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.testacc_cluster", "flavor", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.testacc_cluster.zones.0.id", "id", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.testacc_cluster.zones.0.subnet_id", "subnet_id", "1"),
					resource.TestCheckResourceAttrSet(
						"ibm_container_vpc_cluster.testacc_cluster", "resource_group_id"),
				),
			},
		},
	})
}

func TestAccIBMVpcContainerVpcCluster_importBasic(t *testing.T) {
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	flavor := fmt.Sprintf("c.%d", acctest.RandInt())
	zone := "us-south"
	workerCount := fmt.Sprintf("%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMContainerVpcClusterDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMContainerVpcCluster_basic(zone, clusterName, flavor, workerCount),
			},
			resource.TestStep{
				ResourceName:      "ibm_container_vpc_cluster.testacc_cluster",
				ImportState:       true,
				ImportStateVerify: true,
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
		if rs.Type != "ibm_container_cluster" {
			continue
		}

		targetEnv := getVpcClusterTargetHeaderTestACC()
		// Try to find the key
		_, err := csClient.Clusters().GetCluster(rs.Primary.ID, targetEnv)

		if err != nil && !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("Error waiting for cluster (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
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

func testAccCheckIBMContainerVpcCluster_basic(zone, clusterName, flavor, workerCount string) string {
	return fmt.Sprintf(`	
data "ibm_account" "acc" {
   org_guid = "${data.ibm_org.org.id}"
}

data "ibm_resource_group" "group" {
	is_default = "true"
}

resource "random_id" "name1" {
	byte_length = 2
}
  
resource "random_id" "name2" {
	byte_length = 2
}

locals {
	ZONE1 = "%s-1"
}
  
resource "ibm_is_vpc" "vpc1" {
	name = "vpc-${random_id.name1.hex}"
}
  
resource "ibm_is_subnet" "subnet1" {
	name                     = "subnet-${random_id.name1.hex}"
	vpc                      = "${ibm_is_vpc.vpc1.id}"
	zone                     = "${local.ZONE1}"
	total_ipv4_address_count = 256
}

resource "ibm_container_vpc_cluster" "cluster" {
	name              = "%s${random_id.name1.hex}"
	vpc_id            = "${ibm_is_vpc.vpc1.id}"
	flavor            = "%s"
	worker_count      = "%s"
	resource_group_id = "${data.ibm_resource_group.resource_group.id}"
  
	zones = [
	  {
		subnet_id = "${ibm_is_subnet.subnet1.id}"
		name      = "${local.ZONE1}"
	  },
	]
  }`, zone, clusterName, flavor, workerCount)
}
