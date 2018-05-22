package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMResourceKeyDataSource_basic(t *testing.T) {
	resourceName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resourceKey := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMResourceKeyDataSourceConfig(resourceName, resourceKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_resource_key.testacc_ds_resource_key", "name", resourceKey),
					resource.TestCheckResourceAttr("data.ibm_resource_key.testacc_ds_resource_key", "credentials.%", "7"),
					resource.TestCheckResourceAttr("data.ibm_resource_key.testacc_ds_resource_key", "role", "Viewer"),
					resource.TestCheckResourceAttr("data.ibm_resource_key.testacc_ds_resource_key1", "name", resourceKey),
					resource.TestCheckResourceAttr("data.ibm_resource_key.testacc_ds_resource_key1", "credentials.%", "7"),
					resource.TestCheckResourceAttr("data.ibm_resource_key.testacc_ds_resource_key1", "role", "Viewer"),
				),
			},
		},
	})
}

func TestAccIBMResourceKeyDataSource_mostrecent(t *testing.T) {
	resourceName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resourceKey := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMResourceKeyDataSourceConfigRecent(resourceName, resourceKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_resource_key.testacc_ds_resource_key", "name", resourceKey),
					resource.TestCheckResourceAttr("data.ibm_resource_key.testacc_ds_resource_key", "credentials.%", "7"),
					resource.TestCheckResourceAttr("data.ibm_resource_key.testacc_ds_resource_key1", "name", resourceKey),
					resource.TestCheckResourceAttr("data.ibm_resource_key.testacc_ds_resource_key1", "credentials.%", "7"),
				),
			},
		},
	})
}

func testAccCheckIBMResourceKeyDataSourceConfig(resourceName, resourceKey string) string {
	return fmt.Sprintf(`

resource "ibm_resource_instance" "resource" {
  name       = "%s"
  service    = "cloud-object-storage"
  plan       = "lite"
  location   = "global"
 
}

resource "ibm_resource_key" "resourcekey" {
  name                  = "%s"
  role                  = "Viewer"
  resource_instance_id  = "${ibm_resource_instance.resource.id}"
}

data "ibm_resource_key" "testacc_ds_resource_key" {
  name                  = "${ibm_resource_key.resourcekey.name}"
}

data "ibm_resource_key" "testacc_ds_resource_key1" {
	name                  = "${ibm_resource_key.resourcekey.name}"
	resource_instance_id  = "${ibm_resource_instance.resource.id}"
}`, resourceName, resourceKey)

}

func testAccCheckIBMResourceKeyDataSourceConfigRecent(resourceName, resourceKey string) string {
	return fmt.Sprintf(`

resource "ibm_resource_instance" "resource" {
  name       = "%s"
  service    = "cloud-object-storage"
  plan       = "lite"
  location   = "global"
 
}

resource "ibm_resource_key" "resourcekey" {
  name                  = "%s"
  role                  = "Reader"
  resource_instance_id  = "${ibm_resource_instance.resource.id}"
}

resource "ibm_resource_key" "resourcekey1" {
	name                  = "%s"
	role                  = "Viewer"
	resource_instance_id  = "${ibm_resource_instance.resource.id}"
}

data "ibm_resource_key" "testacc_ds_resource_key" {
  name                  = "${ibm_resource_key.resourcekey.name}"
  most_recent           = "true"
}

data "ibm_resource_key" "testacc_ds_resource_key1" {
	name                  = "${ibm_resource_key.resourcekey.name}"
	resource_instance_id  = "${ibm_resource_instance.resource.id}"
	most_recent           = "true"
}`, resourceName, resourceKey, resourceKey)

}
