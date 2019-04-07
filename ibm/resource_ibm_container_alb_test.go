package ibm

import (
	"fmt"
	"strings"
	"testing"

	v1 "github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccIBMContainerALB_Basic(t *testing.T) {
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMContainerALBDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMContainerALB_basic(clusterName, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_alb.alb", "enable", "true"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMContainerALB_basic(clusterName, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_alb.alb", "enable", "false"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerALBDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_container_alb" {
			continue
		}

		albID := rs.Primary.ID
		targetEnv := v1.ClusterTargetHeader{
			Region: "us-south",
		}

		csClient, err := testAccProvider.Meta().(ClientSession).ContainerAPI()
		if err != nil {
			return err
		}
		albAPI := csClient.Albs()
		_, err = albAPI.GetALB(albID, targetEnv)

		if err != nil && !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}
	return nil
}

func testAccCheckIBMContainerALB_basic(clusterName string, enable bool) string {
	return fmt.Sprintf(`
resource "ibm_container_cluster" "testacc_cluster" {
  name       = "%s"
  datacenter = "%s"

  default_pool_size = 1

  machine_type    = "%s"
  hardware       = "shared"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
}
resource ibm_container_alb alb {
  alb_id = "${ibm_container_cluster.testacc_cluster.albs.0.id}"
  enable = "%t"
}`, clusterName, datacenter, machineType, publicVlanID, privateVlanID, enable)
}
