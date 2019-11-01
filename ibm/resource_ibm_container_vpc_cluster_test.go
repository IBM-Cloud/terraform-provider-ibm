package ibm

import (
	"fmt"
	"log"
	"strings"
	"testing"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	v2 "github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMContainerVpcCluster_basic(t *testing.T) {
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	vpcID := fmt.Sprintf("vpc_%d", acctest.RandInt())
	flavor := fmt.Sprintf("c.%d", acctest.RandInt())
	zoneID := fmt.Sprintf("zone.%d", acctest.RandInt())
	subnetID := fmt.Sprintf("subnet.%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMContainerVpcClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVpcCluster_basic(clusterName, vpcID, flavor, zoneID, subnetID),
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
	vpcID := fmt.Sprintf("vpc_%d", acctest.RandInt())
	flavor := fmt.Sprintf("c.%d", acctest.RandInt())
	zoneName := fmt.Sprintf("zone.%d", acctest.RandInt())
	subnetID := fmt.Sprintf("subnet.%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMContainerVpcClusterDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMContainerVpcCluster_basic(clusterName, vpcID, flavor, zoneName, subnetID),
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

func testAccCheckIBMContainerVpcCluster_basic(clusterName, vpc_id, flavor, subnet_id, name string) string {
	return fmt.Sprintf(`	
data "ibm_account" "acc" {
   org_guid = "${data.ibm_org.org.id}"
}

data "ibm_resource_group" "group" {
	is_default = "true"
}

resource "ibm_container_vpc_cluster" "cluster" {
	name              = "%s"
	vpc_id            = "%s"
	flavor            = "%s"
	resource_group_id = "${data.ibm_resource_group.group.id}"
	zones = [{
		 subnet_id = "%s"
		 name = "%s"
	  }
	]
}`, clusterName, vpc_id, flavor, subnet_id, name)
}
