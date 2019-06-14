package ibm

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"testing"
)

func TestAccIBMCseInstance_basic(t *testing.T) {
	var cisInstanceOne string
	name := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	rsName := "ibm_cse_instance" + ".inst-" + name
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCseInstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCseInstance_basic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCseInstanceExists(rsName, &cisInstanceOne),
					resource.TestCheckResourceAttr(rsName, "region", "us-south"),
					resource.TestCheckResourceAttr(rsName, "service", "terraform-1"),
					resource.TestCheckResourceAttr(rsName, "customer", "customer-"+name),
				),
			},
		},
	})
}

func testAccCheckIBMCseInstanceExists(n string, tfCseId *string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		cseClient, err := testAccProvider.Meta().(ClientSession).CseAPI()
		if err != nil {
			return err
		}

		seAPI := cseClient.ServiceEndpoints()

		instanceID := rs.Primary.ID
		srvObj, err := seAPI.GetServiceEndpoint(instanceID)

		if err != nil {
			return err
		}

		*tfCseId = srvObj.Service.Srvid

		return nil
	}
}

func testAccCheckIBMCseInstanceDestroy(s *terraform.State) error {
	cseClient, err := testAccProvider.Meta().(ClientSession).CseAPI()
	if err != nil {
		return err
	}

	seAPI := cseClient.ServiceEndpoints()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cse_instance" {
			continue
		}

		instanceID := rs.Primary.ID

		_, err := seAPI.GetServiceEndpoint(instanceID)
		if err != nil {
			return nil
		}
	}

	return nil
}

func testAccCheckIBMCseInstance_basic(name string) string {
	return fmt.Sprintf(`
				resource "ibm_cse_instance" "%s" {
				  service = "terraform-1"
				  customer = "%s"
				  service_addresses = ["10.102.33.131", "10.102.33.133"]
				  region = "us-south"
				  data_centers = ["dal10", "dal13"]
				  tcp_ports = [8080, 80]
			    }`, "inst-"+name, "customer-"+name)
}
