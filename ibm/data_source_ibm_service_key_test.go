package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMServiceKeyDataSource_basic(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	serviceKey := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMServiceKeyDataSourceConfig(serviceName, serviceKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_service_key.testacc_ds_service_key", "name", serviceKey),
					resource.TestCheckResourceAttr("data.ibm_service_key.testacc_ds_service_key", "credentials.%", "3"),
				),
			},
		},
	})
}

func testAccCheckIBMServiceKeyDataSourceConfig(serviceName, serviceKey string) string {
	return fmt.Sprintf(`
data "ibm_space" "spacedata" {
  org   = "%s"
  space = "%s"
}

resource "ibm_service_instance" "service" {
  name       = "%s"
  space_guid = "${data.ibm_space.spacedata.id}"
  service    = "speech_to_text"
  plan       = "lite"
  tags       = ["cluster-service", "cluster-bind"]
}

resource "ibm_service_key" "servicekey" {
  name                  = "%s"
  service_instance_guid = "${ibm_service_instance.service.id}"
}

data "ibm_service_key" "testacc_ds_service_key" {
  name                  = "${ibm_service_key.servicekey.name}"
  service_instance_name = "${ibm_service_instance.service.name}"
	space_guid = "${data.ibm_space.spacedata.id}"
}`, cfOrganization, cfSpace, serviceName, serviceKey)

}
