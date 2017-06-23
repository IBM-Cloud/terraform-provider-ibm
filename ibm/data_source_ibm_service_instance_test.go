package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMServiceInstanceDataSource_basic(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMServiceInstanceDataSourceConfig(serviceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_service_instance.testacc_ds_service_instance", "name", serviceName),
				),
			},
		},
	})
}

func testAccCheckIBMServiceInstanceDataSourceConfig(serviceName string) string {
	return fmt.Sprintf(`
	data "ibm_space" "spacedata" {
			org    = "%s"
			space  = "%s"
		}
		
		resource "ibm_service_instance" "service" {
			name              = "%s"
			space_guid        = "${data.ibm_space.spacedata.id}"
			service           = "cleardb"
			plan              = "cb5"
			tags               = ["cluster-service","cluster-bind"]
		}

	
		data "ibm_service_instance" "testacc_ds_service_instance" {
			name = "${ibm_service_instance.service.name}"
}`, cfOrganization, cfSpace, serviceName)

}
