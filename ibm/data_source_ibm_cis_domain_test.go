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

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMCisDomainDataSource_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisDomainDataSourceConfigBasic1(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_cis_domain.cis_domain", "status", "active"),
					resource.TestCheckResourceAttr("data.ibm_cis_domain.cis_domain", "original_name_servers.#", "2"),
					resource.TestCheckResourceAttr("data.ibm_cis_domain.cis_domain", "name_servers.#", "2"),
				),
			},
		},
	})
}

func testAccCheckIBMCisDomainDataSourceConfigBasic1() string {
	name := cisDomainStatic
	instance := cisInstance
	resourceGroup := cisResourceGroup
	return fmt.Sprintf(`
	data "ibm_cis_domain" "cis_domain" {
		cis_id = data.ibm_cis.cis.id
		domain = "%[1]s"
	}
	data "ibm_resource_group" "test_acc" {
		name = "%[2]s"
	}
	data "ibm_cis" "cis" {
		resource_group_id = data.ibm_resource_group.test_acc.id
		name = "%[3]s"
	}
	`, name, resourceGroup, instance)
}
