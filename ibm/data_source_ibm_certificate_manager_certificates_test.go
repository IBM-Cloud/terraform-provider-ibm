package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMCertificateManagerCertificatesDataSource_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCertificateManagerCertificatesDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_certificate_manager_certificates.certs", "certificate_manager_instance_id"),
				),
			},
		},
	})
}

func testAccCheckIBMCertificateManagerCertificatesDataSourceConfig_basic() string {
	return fmt.Sprintf(`
	data "ibm_resource_instance" "cm" {
		name     = "testname"
		location = "us-south"
		service  = "cloudcerts"
	}
	data "ibm_certificate_manager_certificates" "certs"{
		certificate_manager_instance_id=data.ibm_resource_instance.cm.id
	}
	  
	`)
}
