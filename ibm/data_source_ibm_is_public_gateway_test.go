/*
* IBM Confidential
* Object Code Only Source Materials
* 5747-SM3
* (c) Copyright IBM Corp. 2017,2021
*
* The source code for this program is not published or otherwise divested
* of its trade secrets, irrespective of what has been deposited with the
* U.S. Copyright Office.
 */

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMISPublicGatewayDatasource_basic(t *testing.T) {
	vpcname := fmt.Sprintf("tfpgw-vpc-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfpgw-name-%d", acctest.RandIntRange(10, 100))
	zone := "us-south-1"
	resName := "data.ibm_is_public_gateway.testacc_dspgw"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISPublicGatewayDataSourceConfig(vpcname, name1, zone),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						resName, "name", name1),
					resource.TestCheckResourceAttr(
						resName, "zone", zone),
				),
			},
		},
	})
}

func testAccCheckIBMISPublicGatewayDataSourceConfig(vpcname, name, zone string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
	  name = "%s"
	}

	resource "ibm_is_public_gateway" "testacc_public_gateway" {
	  name = "%s"
      vpc = ibm_is_vpc.testacc_vpc.id
      zone = "%s"
	}
	
	data "ibm_is_public_gateway" "testacc_dspgw"{
      name = ibm_is_public_gateway.testacc_public_gateway.name
	}`, vpcname, name, zone)
}
