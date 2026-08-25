// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cloudant_test

import (
	"fmt"
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCloudantGen2DataSource_basic(t *testing.T) {
	dataSourceName := "data.ibm_cloudant.instance"
	serviceName := fmt.Sprintf("terraform-test-%s", acctest.RandString(8))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCloudantGen2DataSourceConfig(serviceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "name", serviceName),
					resource.TestCheckResourceAttr(dataSourceName, "service", "cloudantnosqldb"),
					resource.TestMatchResourceAttr(dataSourceName, flex.ResourceControllerURL, regexp.MustCompile("services/cloudantnosqldb/crn%3A.+")),
					resource.TestCheckResourceAttr(dataSourceName, "include_data_events", "false"),
					resource.TestCheckResourceAttr(dataSourceName, "capacity", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "enable_cors", "true"),
					resource.TestCheckResourceAttr(dataSourceName, "cors_config.0.allow_credentials", "true"),
					resource.TestCheckResourceAttrSet(dataSourceName, "version"),
					resource.TestCheckResourceAttrSet(dataSourceName, "features.0"),
					resource.TestCheckResourceAttrSet(dataSourceName, "features_flags.0"),
				),
			},
		},
	})
}

func testAccCheckIBMCloudantGen2DataSourceConfig(serviceName string) string {
	return fmt.Sprintf(`

	resource "ibm_cloudant" "instance" {
	  name     = "%s"
	  plan     = "standard-gen2"
	  location = "%s"

	  timeouts {
	    create = "15m"
	    update = "15m"
	    delete = "15m"
	  }
	}

	data "ibm_cloudant" "instance" {
	  name = ibm_cloudant.instance.name
	}

	`, serviceName, acc.Region())
}
