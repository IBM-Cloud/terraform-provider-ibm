package resourcecontroller_test

import (
	"fmt"
	"testing"
	"time"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMResourceReclamationDelete_basic(t *testing.T) {
	name := fmt.Sprintf("resource-cos-%d", acctest.RandIntRange(0, 200000))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		// Do not try to verify remote deletion of reclamations here.
		// The delete resource is a one-shot operation; its TF destroy doesn't affect server-side reclamations.
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
				PreConfig: func() { time.Sleep(10 * time.Second) },
			},
			// 3) permanently delete via reclamation delete resource
			{
				Config: testAccReclamationDeleteConfig(),
				Check: resource.ComposeTestCheckFunc(
					// ensure the delete action was planned/applied and has the reclamation_id set
					resource.TestCheckResourceAttrSet("ibm_resource_reclamation_delete.test", "reclamation_id"),
					resource.TestCheckResourceAttrSet("ibm_resource_reclamation_delete.test", "entity_id"),
					resource.TestCheckResourceAttrSet("ibm_resource_reclamation_delete.test", "state"),
					resource.TestCheckResourceAttr("ibm_resource_reclamation_delete.test", "comment", "Terraform test permanent deletion"),
				),
			},
		},
	})
}

func TestAccIBMResourceReclamationDelete_withRequestBy(t *testing.T) {
	name := fmt.Sprintf("resource-cos-%d", acctest.RandIntRange(0, 200000))
	requestBy := "test-user@example.com"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
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
				PreConfig:    func() { time.Sleep(10 * time.Second) },
			},
			// 3) permanently delete with request_by field
			{
				Config: testAccReclamationDeleteConfigWithRequestBy(requestBy),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_resource_reclamation_delete.test", "reclamation_id"),
					resource.TestCheckResourceAttr("ibm_resource_reclamation_delete.test", "comment", "Deletion with request_by field"),
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

func testAccReclamationDeleteConfig() string {
	// Filter by getting the first available reclamation and permanently delete it
	return `
		data "ibm_resource_reclamations" "all" {}

		resource "ibm_resource_reclamation_delete" "test" {
			reclamation_id = data.ibm_resource_reclamations.all.reclamations.0.id
			comment        = "Terraform test permanent deletion"
		}
	`
}

func testAccReclamationDeleteConfigWithRequestBy(requestBy string) string {
	return fmt.Sprintf(`
		data "ibm_resource_reclamations" "all" {}

		resource "ibm_resource_reclamation_delete" "test" {
			reclamation_id = data.ibm_resource_reclamations.all.reclamations.0.id
			request_by     = "%s"
			comment        = "Deletion with request_by field"
		}
	`, requestBy)
}
