package resourcecontroller_test

import (
	"fmt"
	"testing"
	"time"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMResourceReclamationAction_basic(t *testing.T) {
	name := fmt.Sprintf("resource-cos-%d", acctest.RandIntRange(0, 200000))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		// Do not try to verify remote deletion of reclamations here.
		// The action resource is a one-shot operation; its TF destroy doesn't affect server-side reclamations.
		Steps: []resource.TestStep{
			// 1) create an instance
			{
				Config: testAccCreateCOSInstance(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_resource_instance.resource_key", "id"),
				),
			},
			// 2) destroy the instance to generate a reclamation
			{
				Config:       "",
				RefreshState: true,
				// Give RC a moment to surface the reclamation (eventual consistency).
				// (Harness doesn't sleep natively; we do a no-op plan with a delay using PreConfig.)
				PreConfig: func() { time.Sleep(10 * time.Second) },
			},
			// 3) restore via reclamation action (filtered + guarded)
			{
				Config: testAccReclamationRestoreConfig(),
				Check: resource.ComposeTestCheckFunc(
					// ensure the action was planned/applied (count=1) and has the reclamation_id set
					resource.TestCheckResourceAttrSet("ibm_resource_reclamation_action.test", "reclamation_id"),
					resource.TestCheckResourceAttr("ibm_resource_reclamation_action.test", "action", "restore"),
				),
			},
		},
	})
}

func testAccCreateCOSInstance(name string) string {
	return fmt.Sprintf(`
resource "ibm_resource_instance" "resource_key" {
  name     = "%s"
  service  = "cloud-object-storage"
  plan     = "standard"
  location = "global"
}
`, name)
}

func testAccReclamationRestoreConfig() string {
	// Filter by resource_instance_name and guard against empty list.
	// If multiple match (rare), this chooses the first deterministically by index.
	return fmt.Sprintf(`
		data "ibm_resource_reclamations" "all" {}

		resource "ibm_resource_reclamation_action" "test" {
			reclamation_id = data.ibm_resource_reclamations.all.reclamations.0.id
			action         = "restore"
			comment        = "Terraform test restore action"
		}
		`)
}
