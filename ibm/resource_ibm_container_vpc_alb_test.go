package ibm

import (
	"fmt"
	"strings"
	"testing"

	v2 "github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMVpcContainerALB_Basic(t *testing.T) {
	albID := "private-crbm64u3ed02o93vv36hb0-alb1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMVpcContainerALBDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMVpcContainerALB_basic(albID, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_alb.alb", "enable", "true"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMVpcContainerALB_basic(albID, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_alb.alb", "enable", "false"),
				),
			},
		},
	})
}

func testAccCheckIBMVpcContainerALBDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_container_vpc_alb" {
			continue
		}

		albID := rs.Primary.ID
		targetEnv := v2.ClusterTargetHeader{}

		csClient, err := testAccProvider.Meta().(ClientSession).VpcContainerAPI()
		if err != nil {
			return err
		}
		albAPI := csClient.Albs()
		_, err = albAPI.GetAlb(albID, targetEnv)

		if err != nil && !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}
	return nil
}

func testAccCheckIBMVpcContainerALB_basic(albID string, enable bool) string {
	return fmt.Sprintf(`
resource ibm_container_vpc_alb alb {
  alb_id = "%s"
  enable = "%t"
}`, albID, enable)
}
