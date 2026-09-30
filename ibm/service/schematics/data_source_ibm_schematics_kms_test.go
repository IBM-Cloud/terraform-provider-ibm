// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package schematics_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSchematicsKMSDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSchematicsKMSDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_kms.schematics_kms_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_kms.schematics_kms_instance", "location"),
				),
			},
		},
	})
}

func TestAccIbmSchematicsKMSDataSourceWithLocation(t *testing.T) {
	kmsLocation := "us-south"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSchematicsKMSDataSourceConfig(kmsLocation),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_kms.schematics_kms_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_kms.schematics_kms_instance", "location"),
				),
			},
		},
	})
}

func TestAccIbmSchematicsKMSDataSourceAllArgs(t *testing.T) {
	kmsLocation := "us-south"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSchematicsKMSDataSourceConfig(kmsLocation),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_kms.schematics_kms_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_kms.schematics_kms_instance", "location"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_kms.schematics_kms_instance", "encryption_scheme"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_kms.schematics_kms_instance", "resource_group"),
				),
			},
		},
	})
}

func testAccCheckIbmSchematicsKMSDataSourceConfigBasic() string {
	return `
		data "ibm_schematics_kms" "schematics_kms_instance" {
		}
	`
}

func testAccCheckIbmSchematicsKMSDataSourceConfig(kmsLocation string) string {
	return fmt.Sprintf(`
		data "ibm_schematics_kms" "schematics_kms_instance" {
			location = "%s"
		}
	`, kmsLocation)
}
