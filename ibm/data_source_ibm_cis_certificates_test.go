/* IBM Confidential
*  Object Code Only Source Materials
*  5747-SM3
*  (c) Copyright IBM Corp. 2017,2021
*
*  The source code for this program is not published or otherwise divested
*  of its trade secrets, irrespective of what has been deposited with the
*  U.S. Copyright Office. */

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMCisCertificatesDataSource_basic(t *testing.T) {
	node := "data.ibm_cis_certificates.test"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisCertificatesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "certificates.0.status"),
				),
			},
		},
	})
}

func testAccCheckIBMCisCertificatesDataSourceConfig() string {
	// status filter defaults to empty
	return testAccCheckCisCertificateOrderConfigBasic() + fmt.Sprintf(`
	data "ibm_cis_certificates" "test" {
		cis_id    = ibm_cis_certificate_order.test.cis_id
		domain_id = ibm_cis_certificate_order.test.domain_id
	}`)
}
