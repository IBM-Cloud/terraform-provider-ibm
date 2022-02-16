// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package resourcecontroller_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMResourceInstanceDataSource_basic(t *testing.T) {
	instanceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:  setupResourceInstanceConfig(instanceName),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "service", "cloud-object-storage"),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance2", "service", "kms"),
				),
			},
			{
				Config:  testAccCheckIBMResourceInstanceDataSourceConfig(instanceName),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_resource_instance.testacc_ds_resource_instance", "name", instanceName),
					resource.TestCheckResourceAttr("data.ibm_resource_instance.testacc_ds_resource_instance", "service", "cloud-object-storage"),
					resource.TestCheckResourceAttr("data.ibm_resource_instance.testacc_ds_resource_instance", "plan", "standard"),
					resource.TestCheckResourceAttr("data.ibm_resource_instance.testacc_ds_resource_instance", "location", "global"),
				),
			},
			{
				Config: testAccCheckIBMResourceInstanceDataSourceConfigWithService(instanceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_resource_instance.testacc_ds_resource_instance2", "name", instanceName),
					resource.TestCheckResourceAttr("data.ibm_resource_instance.testacc_ds_resource_instance2", "service", "kms"),
					resource.TestCheckResourceAttr("data.ibm_resource_instance.testacc_ds_resource_instance2", "plan", "tiered-pricing"),
					resource.TestCheckResourceAttr("data.ibm_resource_instance.testacc_ds_resource_instance2", "location", "us-south"),
				),
			},
		},
	})
}

func setupResourceInstanceConfig(instanceName string) string {
	return fmt.Sprintf(`

resource "ibm_resource_instance" "instance" {
  name     = "%s"
  service  = "cloud-object-storage"
  plan     = "standard"
  location = "global"
}

resource "ibm_resource_instance" "instance2" {
  name     = "%s"
  service  = "kms"
  plan     = "tiered-pricing"
  location = "us-south"
}

`, instanceName, instanceName)

}

func testAccCheckIBMResourceInstanceDataSourceConfig(instanceName string) string {
	return fmt.Sprintf(`
data "ibm_resource_group" "group" {
  is_default=true
}

resource "ibm_resource_instance" "instance" {
  name     = "%s"
  service  = "cloud-object-storage"
  plan     = "standard"
  location = "global"
}

resource "ibm_resource_instance" "instance2" {
  name     = "%s"
  service  = "kms"
  plan     = "tiered-pricing"
  location = "us-south"
}

data "ibm_resource_instance" "testacc_ds_resource_instance" {
  name              = ibm_resource_instance.instance.name
  location          = "global"
  resource_group_id = data.ibm_resource_group.group.id
}
`, instanceName, instanceName)

}

func testAccCheckIBMResourceInstanceDataSourceConfigWithService(instanceName string) string {
	return fmt.Sprintf(`

resource "ibm_resource_instance" "instance" {
  name     = "%s"
  service  = "cloud-object-storage"
  plan     = "standard"
  location = "global"
}

resource "ibm_resource_instance" "instance2" {
  name     = "%s"
  service  = "kms"
  plan     = "tiered-pricing"
  location = "us-south"
}

data "ibm_resource_instance" "testacc_ds_resource_instance2" {
  name    = ibm_resource_instance.instance2.name
  service = "kms"
}
`, instanceName, instanceName)

}
