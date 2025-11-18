package resourcecontroller_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMResourceReclamationDelete_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccReclamationDeleteConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_resource_reclamation_delete.test", "reclamation_id", acc.ReclamationId),
					resource.TestCheckResourceAttrSet("ibm_resource_reclamation_delete.test", "entity_id"),
					resource.TestCheckResourceAttrSet("ibm_resource_reclamation_delete.test", "state"),
					resource.TestCheckResourceAttr("ibm_resource_reclamation_delete.test", "comment", "Terraform test permanent deletion"),
				),
			},
		},
	})
}

func TestAccIBMResourceReclamationDelete_withRequestBy(t *testing.T) {
	requestBy := "test-user@example.com"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccReclamationDeleteConfigWithRequestBy(requestBy),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_resource_reclamation_delete.test", "reclamation_id", acc.ReclamationId),
					resource.TestCheckResourceAttr("ibm_resource_reclamation_delete.test", "comment", "Deletion with request_by field"),
				),
			},
		},
	})
}

func testAccReclamationDeleteConfig() string {
	return fmt.Sprintf(`
resource "ibm_resource_reclamation_delete" "test" {
  reclamation_id = "%s"
  comment        = "Terraform test permanent deletion"
}
`, acc.ReclamationId)
}

func testAccReclamationDeleteConfigWithRequestBy(requestBy string) string {
	return fmt.Sprintf(`
resource "ibm_resource_reclamation_delete" "test" {
  reclamation_id = "%s"
  request_by     = "%s"
  comment        = "Deletion with request_by field"
}
`, acc.ReclamationId, requestBy)
}
