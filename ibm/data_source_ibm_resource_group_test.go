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
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMResourceGroupDataSource_Basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMResourceGroupDataSourceConfigDefault(),
			},
			{
				Config: testAccCheckIBMResourceGroupDataSourceConfigWithName(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_resource_group.testacc_ds_resource_group_name", "name", "default"),
				),
			},
		},
	})
}

func TestAccIBMResourceGroupDataSource_Default_false(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMResourceGroupDataSourceDefaultFalse(),
				ExpectError: regexp.MustCompile(`Missing required properties. Need a resource group name, or the is_default true`),
			},
		},
	})
}

func testAccCheckIBMResourceGroupDataSourceConfigDefault() string {
	return fmt.Sprintf(`
	
data "ibm_resource_group" "testacc_ds_resource_group" {
	is_default = "true"
}`)

}

func testAccCheckIBMResourceGroupDataSourceConfigWithName() string {
	return fmt.Sprintf(`

data "ibm_resource_group" "testacc_ds_resource_group_name" {
	name = "default"
}`)

}

func testAccCheckIBMResourceGroupDataSourceDefaultFalse() string {
	return fmt.Sprintf(`
	
data "ibm_resource_group" "testacc_ds_resource_group" {
	is_default = "false"
}`)

}
