package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMResourceInstanceDataSource_basic(t *testing.T) {
	instanceName := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:  setupResourceInstanceConfig(instanceName),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_resource_instance.instance", "service", "cloud-object-storage"),
					resource.TestCheckResourceAttr("ibm_resource_instance.instance2", "service", "kms"),
				),
			},
			resource.TestStep{
				Config:  testAccCheckIBMResourceInstanceDataSourceConfig(instanceName),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_resource_instance.testacc_ds_resource_instance", "name", instanceName),
					resource.TestCheckResourceAttr("data.ibm_resource_instance.testacc_ds_resource_instance", "service", "cloud-object-storage"),
					resource.TestCheckResourceAttr("data.ibm_resource_instance.testacc_ds_resource_instance", "plan", "lite"),
					resource.TestCheckResourceAttr("data.ibm_resource_instance.testacc_ds_resource_instance", "location", "global"),
				),
			},
			resource.TestStep{
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
  name       = "%s"
  service    = "cloud-object-storage"
  plan       = "lite"
  location   = "global"
  tags       = ["tag1", "tag2"]
}

resource "ibm_resource_instance" "instance2" {
	name       = "%s"
	service    = "kms"
	plan       = "tiered-pricing"
	location   = "us-south"
	tags       = ["tag3", "tag4"]
  }

`, instanceName, instanceName)

}

func testAccCheckIBMResourceInstanceDataSourceConfig(instanceName string) string {
	return fmt.Sprintf(`
data "ibm_resource_group" "group" {
  name = "default"
}

resource "ibm_resource_instance" "instance" {
	name       = "%s"
	service    = "cloud-object-storage"
	plan       = "lite"
	location   = "global"
	tags       = ["tag1", "tag2"]
}

resource "ibm_resource_instance" "instance2" {
	name       = "%s"
	service    = "kms"
	plan       = "tiered-pricing"
	location   = "us-south"
	tags       = ["tag3", "tag4"]
  }

data "ibm_resource_instance" "testacc_ds_resource_instance" {
  name = "${ibm_resource_instance.instance.name}"
  location = "global"
  resource_group_id = "${data.ibm_resource_group.group.id}"
}
`, instanceName, instanceName)

}

func testAccCheckIBMResourceInstanceDataSourceConfigWithService(instanceName string) string {
	return fmt.Sprintf(`

resource "ibm_resource_instance" "instance" {
	name       = "%s"
	service    = "cloud-object-storage"
	plan       = "lite"
	location   = "global"
	tags       = ["tag1", "tag2"]
}

resource "ibm_resource_instance" "instance2" {
	name       = "%s"
	service    = "kms"
	plan       = "tiered-pricing"
	location   = "us-south"
	tags       = ["tag3", "tag4"]
  }

data "ibm_resource_instance" "testacc_ds_resource_instance2" {
  name = "${ibm_resource_instance.instance2.name}"
  service = "kms"
}
`, instanceName, instanceName)

}
