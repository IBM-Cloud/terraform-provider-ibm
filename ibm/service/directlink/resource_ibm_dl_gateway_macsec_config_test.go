package directlink_test

import (
	"fmt"
	"testing"

	"github.com/IBM/networking-go-sdk/directlinkv1"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMDLGatewayMacsecConfig_Basic(t *testing.T) {
	resourceName := "ibm_dl_gateway_macsec_config.test"
	gatewayName := fmt.Sprintf("TF_MACSEC-TEST-%d", acctest.RandIntRange(00, 99))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDLGatewayMacsecConfigDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMDLGatewayMacsecConfig(gatewayName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "active", "true"),
				),
			},
		},
	})
}

func testAccIBMDLGatewayMacsecConfig(gatewayName string) string {
	return fmt.Sprintf(`
	resource ibm_dl_gateway test_dl_gateway {
  bgp_asn =  64999
  global = true 
  metered = false
  name = "%s"
  speed_mbps = 1000 
  type =  "dedicated" 
  cross_connect_router = "LAB-xcr02.dal09"
  location_name = "dal09"
  customer_name = "TEST_CUSTOMER" 
  carrier_name = "TEST_CARRIER"
  vlan=3965
} 

resource "ibm_dl_gateway_macsec_config" "test" {
    gateway = ibm_dl_gateway.test_dl_gateway.id
    active = true
    security_policy = "must_secure"
    sak_rekey {
        interval = 76
        mode = "timer"
    }
    caks {
        key {
            crn = "crn:v1:staging:public:hs-crypto:us-south:a/3f455c4c574447adbc14bda52f80e62f:b2044455-b89e-4c57-96ae-3f17c092dd31:key:ebc0fbe6-fd7c-4971-b127-71a385c8f602"
        }
        name = "AA01"
        session = "primary"
    }
    window_size = 552
}
	`, gatewayName)
}

func testAccCheckIBMDLGatewayMacsecConfigDestroy(s *terraform.State) error {
	directLink, err := directlinkClient(acc.TestAccProvider.Meta())
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_dl_gateway_macsec_cak" {
			continue
		}

		gatewayId := rs.Primary.ID

		delOptions := &directlinkv1.UnsetGatewayMacsecOptions{
			ID: &gatewayId,
		}

		_, err := directLink.UnsetGatewayMacsec(delOptions)

		if err != nil {
			return fmt.Errorf("Macsec CAK still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}
