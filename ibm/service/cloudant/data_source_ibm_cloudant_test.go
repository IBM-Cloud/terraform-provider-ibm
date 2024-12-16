// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cloudant_test

import (
	"fmt"
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMCloudantDataSource_basic(t *testing.T) {
	dataSourceName := "data.ibm_cloudant.instance"
	serviceName := fmt.Sprintf("terraform-test-%s", acctest.RandString(8))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCloudantDataSourceConfig(serviceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "name", serviceName),
					resource.TestCheckResourceAttr(dataSourceName, "service", "cloudantnosqldb"),
					resource.TestMatchResourceAttr(dataSourceName, flex.ResourceControllerURL, regexp.MustCompile("services/cloudantnosqldb/crn%3A.+")),
					resource.TestCheckResourceAttr(dataSourceName, "include_data_events", "false"),
					resource.TestCheckResourceAttr(dataSourceName, "capacity", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "throughput.read", "20"),
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

func testAccCheckIBMCloudantDataSourceConfig(serviceName string) string {
	return fmt.Sprintf(`

	resource "ibm_resource_instance" "cloudant" {
	  name     = "%s"
	  service  = "cloudantnosqldb"
	  plan     = "lite"
	  location = "us-south"
	}

	data "ibm_cloudant" "instance" {
	  name     = ibm_resource_instance.cloudant.name
	}

	`, serviceName)
}
