/* IBM Confidential
*  Object Code Only Source Materials
*  5747-SM3
*  (c) Copyright IBM Corp. 2017,2021
*
*  The source code for this program is not published or otherwise divested
*  of its trade secrets, irrespective of what has been deposited with the
*  U.S. Copyright Office. */

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMTransitGatewayDataSource_basic(t *testing.T) {
	gatewayname := fmt.Sprintf("gateway-name-%d", acctest.RandIntRange(10, 100))
	location := fmt.Sprintf("us-south")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMTransitGatewayDataSourceConfig(gatewayname, location),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.ibm_tg_gateway.test_tg_gateway", "name", gatewayname),
					resource.TestCheckResourceAttr(
						"data.ibm_tg_gateway.test_tg_gateway", "location", location),
				),
			},
		},
	})
}

func testAccCheckIBMTransitGatewayDataSourceConfig(gatewayname, location string) string {
	return fmt.Sprintf(`
	
	resource "ibm_tg_gateway" "test_tg_gateway" {
		name="%s"
		location="%s"
		global=true
	  }
	
	   data "ibm_tg_gateway" "test_tg_gateway" {
			name = ibm_tg_gateway.test_tg_gateway.name
		 }
	  `, gatewayname, location)
}
