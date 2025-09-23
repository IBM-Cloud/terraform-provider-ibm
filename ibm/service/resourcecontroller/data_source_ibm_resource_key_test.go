// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package resourcecontroller_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMResourceKeyDataSource_basic(t *testing.T) {
	resourceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	resourceKey := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMResourceKeyDataSourceConfig(resourceName, resourceKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_resource_key.testacc_ds_resource_key", "name", resourceKey),
					resource.TestCheckResourceAttr("data.ibm_resource_key.testacc_ds_resource_key", "credentials.%", "8"),
					resource.TestCheckResourceAttr("data.ibm_resource_key.testacc_ds_resource_key", "role", "Writer"),
					resource.TestCheckResourceAttr("data.ibm_resource_key.testacc_ds_resource_key1", "name", resourceKey),
					resource.TestCheckResourceAttr("data.ibm_resource_key.testacc_ds_resource_key1", "credentials.%", "8"),
					resource.TestCheckResourceAttr("data.ibm_resource_key.testacc_ds_resource_key1", "role", "Writer"),
				),
			},
		},
	})
}

func TestAccIBMResourceKeyDataSource_mostrecent(t *testing.T) {
	resourceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	resourceKey := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMResourceKeyDataSourceConfigRecent(resourceName, resourceKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_resource_key.testacc_ds_resource_key", "name", resourceKey),
					resource.TestCheckResourceAttr("data.ibm_resource_key.testacc_ds_resource_key", "credentials.%", "8"),
					resource.TestCheckResourceAttr("data.ibm_resource_key.testacc_ds_resource_key1", "name", resourceKey),
					resource.TestCheckResourceAttr("data.ibm_resource_key.testacc_ds_resource_key1", "credentials.%", "8"),
				),
			},
		},
	})
}

func testAccCheckIBMResourceKeyDataSourceConfig(resourceName, resourceKey string) string {
	return fmt.Sprintf(`

resource "ibm_resource_instance" "resource" {
  name     = "%s"
  service  = "cloud-object-storage"
  plan     = "standard"
  location = "global"
}

resource "ibm_resource_key" "resourcekey" {
  name                 = "%s"
  role                 = "Writer"
  resource_instance_id = ibm_resource_instance.resource.id
}

data "ibm_resource_key" "testacc_ds_resource_key" {
  name = ibm_resource_key.resourcekey.name
}

data "ibm_resource_key" "testacc_ds_resource_key1" {
  name                 = ibm_resource_key.resourcekey.name
  resource_instance_id = ibm_resource_instance.resource.id
}
`, resourceName, resourceKey)

}

func testAccCheckIBMResourceKeyDataSourceConfigRecent(resourceName, resourceKey string) string {
	return fmt.Sprintf(`

resource "ibm_resource_instance" "resource" {
  name     = "%s"
  service  = "cloud-object-storage"
  plan     = "standard"
  location = "global"
}

resource "ibm_resource_key" "resourcekey" {
  name                 = "%s"
  role                 = "Reader"
  resource_instance_id = ibm_resource_instance.resource.id
}

resource "ibm_resource_key" "resourcekey1" {
  name                 = "%s"
  role                 = "Writer"
  resource_instance_id = ibm_resource_instance.resource.id
}

data "ibm_resource_key" "testacc_ds_resource_key" {
  name        = ibm_resource_key.resourcekey.name
  most_recent = "true"
}

data "ibm_resource_key" "testacc_ds_resource_key1" {
  name                 = ibm_resource_key.resourcekey.name
  resource_instance_id = ibm_resource_instance.resource.id
  most_recent          = "true"
}
`, resourceName, resourceKey, resourceKey)

}
