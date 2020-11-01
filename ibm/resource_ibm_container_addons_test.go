package ibm

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	v1 "github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
)

func TestAccIBMContainerAddOns_Basic(t *testing.T) {
	clusterName := fmt.Sprintf("terraform-%d", acctest.RandIntRange(10, 100))
	vpc := fmt.Sprintf("terraform-vpc-%d", acctest.RandIntRange(10, 100))
	subnet := fmt.Sprintf("terraform-subnet-%d", acctest.RandIntRange(10, 100))
	flavor := "c2.2x4"
	zone := "us-south"
	workerCount := "1"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMContainerAddOnsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMContainerAddOnsBasic(clusterName, zone, vpc, subnet, flavor, workerCount),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_addons.addons", "addons.#", "3"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMContainerAddOnsUpdate(zone, vpc, subnet, clusterName, flavor, workerCount),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_addons.addons", "addons.#", "3"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerAddOnsDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_container_addons" {
			continue
		}
		targetEnv := v1.ClusterTargetHeader{
			Region: "us-south",
		}
		csClient, err := testAccProvider.Meta().(ClientSession).ContainerAPI()
		if err != nil {
			return err
		}
		cluster := rs.Primary.ID
		addOnAPI := csClient.AddOns()
		_, err = addOnAPI.GetAddons(cluster, targetEnv)
		if err == nil {
			return fmt.Errorf("AddOns still exists: %s", rs.Primary.ID)
		} else if !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("Error checking if AddOns (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}
	return nil
}

func testAccCheckIBMContainerAddOnsBasic(zone, vpc, subnet, clusterName, flavor, workerCount string) string {
	return testAccCheckIBMContainerVpcOcpClusterGen2basic(zone, vpc, subnet, clusterName, flavor, workerCount) + fmt.Sprintf(`
	resource "ibm_container_addons" "addons" {
	cluster = ibm_container_vpc_cluster.clustergen2.name
	addons {
		name    = "istio"
		version = "1.6"
	}
	addons {
		name    = "vpc-block-csi-driver"
		version = "2.0.3"
	}
	addons {
		name    = "static-route"
		version = "1.0.0"
	}
}`)
}
func testAccCheckIBMContainerAddOnsUpdate(clusterName, zone, vpc, subnet, flavor, workerCount string) string {
	return testAccCheckIBMContainerVpcOcpClusterGen2basic(zone, vpc, subnet, clusterName, flavor, workerCount) + fmt.Sprintf(`
	resource "ibm_container_addons" "addons" {
	cluster = ibm_container_vpc_cluster.clustergen2.name
	addons {
		name    = "istio"
		version = "1.7"
	}
	addons {
		name    = "vpc-block-csi-driver"
		version = "2.0.3"
	}
	addons {
		name    = "kube-terminal"
		version = "1.0.0"
	}
}`)
}
