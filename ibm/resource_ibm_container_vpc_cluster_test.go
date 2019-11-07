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
	randint := acctest.RandIntRange(10, 100)
	vpc := fmt.Sprintf("vpc-%d", randint)
	subnet := fmt.Sprintf("subnet-%d", randint)
	flavor := "c2.2x4"
	zone := "us-south"
	workerCount := "1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMContainerVpcClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVpcCluster_basic(zone, vpc, subnet, clusterName, flavor, workerCount),
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
	randint := acctest.RandIntRange(10, 100)
	vpc := fmt.Sprintf("vpc-%d", randint)
	subnet := fmt.Sprintf("subnet-%d", randint)
	flavor := "c2.2x4"
	zone := "us-south"
	workerCount := "1"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMContainerVpcClusterDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMContainerVpcCluster_basic(zone, vpc, subnet, clusterName, flavor, workerCount),
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

func testAccCheckIBMContainerVpcCluster_basic(zone, vpc, subnet, clusterName, flavor, workerCount string) string {
	return fmt.Sprintf(`
	
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
	resource_group_id = "${data.ibm_resource_group.resource_group.id}"
	zones = [{
		 subnet_id = "${ibm_is_subnet.subnet1.id}"
		 name      = "${local.ZONE1}"
	  }
	]
  }`, zone, vpc, subnet, clusterName, flavor, workerCount)

}
