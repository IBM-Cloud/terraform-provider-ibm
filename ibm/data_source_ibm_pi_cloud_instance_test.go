/**
 * © Copyright IBM Corporation 2020. All Rights Reserved.
 *
 * Licensed under the Mozilla Public License, version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at https://mozilla.org/MPL/2.0/
 */

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMPICloudInstanceDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPICloudInstanceDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_cloud_instance.testacc_ds_cloud_instance", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPICloudInstanceDataSourceConfig() string {
	return fmt.Sprintf(`
	
data "ibm_pi_cloud_instance" "testacc_ds_cloud_instance" {
    pi_cloud_instance_id = "%s"
}`, pi_cloud_instance_id)

}
