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

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMDLPortDataSource_basic(t *testing.T) {
	name := "dl_port"
	resName := "data.ibm_dl_port.test_dl_port"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDLPortDataSourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "port_id"),
				),
			},
		},
	})
}

func testAccCheckIBMDLPortDataSourceConfig(name string) string {
	return fmt.Sprintf(`
	   data "ibm_dl_ports" "test_%s" {
	   }
	   data "ibm_dl_port" "test_dl_port" {
		   port_id = data.ibm_dl_ports.test_%s.ports[0].port_id
	   }
	  `, name, name)
}
