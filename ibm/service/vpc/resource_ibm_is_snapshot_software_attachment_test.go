// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/vpc"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMIsSnapshotSoftwareAttachmentBasic(t *testing.T) {
	var conf vpcv1.SnapshotSoftwareAttachment
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf-instance-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsSnapshotSoftwareAttachmentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsSnapshotSoftwareAttachmentConfigBasic(vpcname, subnetname, sshname, instanceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsSnapshotSoftwareAttachmentExists("ibm_is_snapshot_software_attachment.is_snapshot_software_attachment_instance", conf),
					resource.TestCheckResourceAttrPair("ibm_is_snapshot_software_attachment.is_snapshot_software_attachment_instance", "snapshot_id", "ibm_is_snapshot.test_snapshot", "id"),
					resource.TestCheckResourceAttrPair("ibm_is_snapshot_software_attachment.is_snapshot_software_attachment_instance", "snapshot_software_attachment_id", "data.ibm_is_snapshot_software_attachments.is_snapshot_software_attachments", "software_attachments.0.id"),
					// name is not set in config; it must be read back from the API (Optional + Computed).
					resource.TestCheckResourceAttrPair("ibm_is_snapshot_software_attachment.is_snapshot_software_attachment_instance", "name", "data.ibm_is_snapshot_software_attachments.is_snapshot_software_attachments", "software_attachments.0.name"),
					resource.TestCheckResourceAttrSet("ibm_is_snapshot_software_attachment.is_snapshot_software_attachment_instance", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_snapshot_software_attachment.is_snapshot_software_attachment_instance", "href"),
					resource.TestCheckResourceAttr("ibm_is_snapshot_software_attachment.is_snapshot_software_attachment_instance", "resource_type", "snapshot_software_attachment"),
					// ibm_is_snapshot and ibm_is_snapshot_software_attachments must return the same attachment.
					resource.TestCheckResourceAttrPair("data.ibm_is_snapshot.test_snapshot", "software_attachments.0.id", "data.ibm_is_snapshot_software_attachments.is_snapshot_software_attachments", "software_attachments.0.id"),
				),
			},
			// Re-applying the same config must not produce a diff for the computed name.
			resource.TestStep{
				Config:   testAccCheckIBMIsSnapshotSoftwareAttachmentConfigBasic(vpcname, subnetname, sshname, instanceName),
				PlanOnly: true,
			},
		},
	})
}

func TestAccIBMIsSnapshotSoftwareAttachmentAllArgs(t *testing.T) {
	var conf vpcv1.SnapshotSoftwareAttachment
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf-instance-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-name-%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf-name-upd-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsSnapshotSoftwareAttachmentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsSnapshotSoftwareAttachmentConfig(vpcname, subnetname, sshname, instanceName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsSnapshotSoftwareAttachmentExists("ibm_is_snapshot_software_attachment.is_snapshot_software_attachment_instance", conf),
					resource.TestCheckResourceAttrPair("ibm_is_snapshot_software_attachment.is_snapshot_software_attachment_instance", "snapshot_id", "ibm_is_snapshot.test_snapshot", "id"),
					resource.TestCheckResourceAttr("ibm_is_snapshot_software_attachment.is_snapshot_software_attachment_instance", "name", name),
					testAccCheckIBMIsSnapshotSoftwareAttachmentRemoteName("ibm_is_snapshot_software_attachment.is_snapshot_software_attachment_instance", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsSnapshotSoftwareAttachmentConfig(vpcname, subnetname, sshname, instanceName, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_snapshot_software_attachment.is_snapshot_software_attachment_instance", "name", nameUpdate),
					testAccCheckIBMIsSnapshotSoftwareAttachmentRemoteName("ibm_is_snapshot_software_attachment.is_snapshot_software_attachment_instance", nameUpdate),
				),
			},
			// Removing name from the config keeps the last name (Optional + Computed), it is not cleared.
			resource.TestStep{
				Config: testAccCheckIBMIsSnapshotSoftwareAttachmentConfigBasic(vpcname, subnetname, sshname, instanceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_snapshot_software_attachment.is_snapshot_software_attachment_instance", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_is_snapshot_software_attachment.is_snapshot_software_attachment_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Import with an ID that is not in <snapshot_id>/<snapshot_software_attachment_id> format must fail.
			resource.TestStep{
				ResourceName:  "ibm_is_snapshot_software_attachment.is_snapshot_software_attachment_instance",
				ImportState:   true,
				ImportStateId: "invalid-import-id",
				ExpectError:   regexp.MustCompile("does not contain /"),
			},
		},
	})
}

func TestAccIBMIsSnapshotSoftwareAttachmentInvalidName(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: `
					resource "ibm_is_snapshot_software_attachment" "is_snapshot_software_attachment_instance" {
						snapshot_id                     = "r006-00000000-0000-0000-0000-000000000000"
						snapshot_software_attachment_id = "r006-00000000-0000-0000-0000-000000000001"
						name                            = "Invalid_Name"
					}
				`,
				PlanOnly:    true,
				ExpectError: regexp.MustCompile("should match regexp"),
			},
		},
	})
}

func testAccCheckIBMIsSnapshotSoftwareAttachmentBaseConfig(vpcname, subnetname, sshname, instanceName string) string {
	return testAccCheckIBMIsSoftwareAttachmentBaseConfig(vpcname, subnetname, sshname, instanceName)
}

func testAccCheckIBMIsSnapshotSoftwareAttachmentConfigBasic(vpcname, subnetname, sshname, instanceName string) string {
	return testAccCheckIBMIsSnapshotSoftwareAttachmentBaseConfig(vpcname, subnetname, sshname, instanceName) + `
		resource "ibm_is_snapshot_software_attachment" "is_snapshot_software_attachment_instance" {
			snapshot_id                     = ibm_is_snapshot.test_snapshot.id
			snapshot_software_attachment_id = data.ibm_is_snapshot.test_snapshot.software_attachments.0.id
		}
	`
}

func testAccCheckIBMIsSnapshotSoftwareAttachmentConfig(vpcname, subnetname, sshname, instanceName, name string) string {
	return testAccCheckIBMIsSnapshotSoftwareAttachmentBaseConfig(vpcname, subnetname, sshname, instanceName) + fmt.Sprintf(`
		resource "ibm_is_snapshot_software_attachment" "is_snapshot_software_attachment_instance" {
			snapshot_id                     = ibm_is_snapshot.test_snapshot.id
			snapshot_software_attachment_id = data.ibm_is_snapshot.test_snapshot.software_attachments.0.id
			name                            = "%s"
		}
	`, name)
}

func testAccCheckIBMIsSnapshotSoftwareAttachmentExists(n string, obj vpcv1.SnapshotSoftwareAttachment) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		if err != nil {
			return err
		}

		getSnapshotSoftwareAttachmentOptions := &vpcv1.GetSnapshotSoftwareAttachmentOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getSnapshotSoftwareAttachmentOptions.SetSnapshotID(parts[0])
		getSnapshotSoftwareAttachmentOptions.SetID(parts[1])

		snapshotSoftwareAttachment, _, err := vpcClient.GetSnapshotSoftwareAttachment(getSnapshotSoftwareAttachmentOptions)
		if err != nil {
			return err
		}

		obj = *snapshotSoftwareAttachment
		return nil
	}
}

// testAccCheckIBMIsSnapshotSoftwareAttachmentRemoteName verifies that the name was
// actually patched on the API side, not only stored in Terraform state.
func testAccCheckIBMIsSnapshotSoftwareAttachmentRemoteName(n string, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		if err != nil {
			return err
		}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getSnapshotSoftwareAttachmentOptions := &vpcv1.GetSnapshotSoftwareAttachmentOptions{}
		getSnapshotSoftwareAttachmentOptions.SetSnapshotID(parts[0])
		getSnapshotSoftwareAttachmentOptions.SetID(parts[1])

		snapshotSoftwareAttachment, _, err := vpcClient.GetSnapshotSoftwareAttachment(getSnapshotSoftwareAttachmentOptions)
		if err != nil {
			return err
		}
		if snapshotSoftwareAttachment.Name == nil || *snapshotSoftwareAttachment.Name != name {
			return fmt.Errorf("SnapshotSoftwareAttachment %s has name %v, expected %s", rs.Primary.ID, snapshotSoftwareAttachment.Name, name)
		}
		return nil
	}
}

// testAccCheckIBMIsSnapshotSoftwareAttachmentDestroy runs after all resources in the
// test are destroyed. Destroying ibm_is_snapshot_software_attachment only removes it
// from state, but the attachment goes away together with its snapshot, which is
// destroyed in the same test, so by now the attachment must no longer exist.
func testAccCheckIBMIsSnapshotSoftwareAttachmentDestroy(s *terraform.State) error {
	vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_snapshot_software_attachment" {
			continue
		}

		getSnapshotSoftwareAttachmentOptions := &vpcv1.GetSnapshotSoftwareAttachmentOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getSnapshotSoftwareAttachmentOptions.SetSnapshotID(parts[0])
		getSnapshotSoftwareAttachmentOptions.SetID(parts[1])

		// Try to find the key
		_, response, err := vpcClient.GetSnapshotSoftwareAttachment(getSnapshotSoftwareAttachmentOptions)

		if err == nil {
			return fmt.Errorf("SnapshotSoftwareAttachment still exists: %s", rs.Primary.ID)
		} else if response == nil || response.StatusCode != 404 {
			return fmt.Errorf("Error checking for SnapshotSoftwareAttachment (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

// TestResourceIBMIsSnapshotSoftwareAttachmentSchema checks that the resource is
// registered with the provider and that its arguments have the expected
// Required / Optional / Computed / ForceNew behavior and validation.
func TestResourceIBMIsSnapshotSoftwareAttachmentSchema(t *testing.T) {
	r, ok := acc.TestAccProvider.ResourcesMap["ibm_is_snapshot_software_attachment"]
	assert.True(t, ok, "ibm_is_snapshot_software_attachment must be registered in the provider")
	if !ok {
		return
	}
	assert.NotNil(t, r.Importer)

	snapshotID := r.Schema["snapshot_id"]
	assert.True(t, snapshotID.Required)
	assert.True(t, snapshotID.ForceNew)

	attachmentID := r.Schema["snapshot_software_attachment_id"]
	assert.NotNil(t, attachmentID)
	assert.True(t, attachmentID.Required)
	assert.True(t, attachmentID.ForceNew)

	name := r.Schema["name"]
	assert.True(t, name.Optional)
	assert.True(t, name.Computed)
	assert.False(t, name.ForceNew)
	assert.NotNil(t, name.ValidateFunc)
	_, errs := name.ValidateFunc("my-software-attachment", "name")
	assert.Empty(t, errs)
	_, errs = name.ValidateFunc("Invalid_Name", "name")
	assert.NotEmpty(t, errs)

	for _, computed := range []string{"catalog_offering", "created_at", "entitlement", "href", "resource_type"} {
		assert.True(t, r.Schema[computed].Computed, computed)
		assert.False(t, r.Schema[computed].Optional, computed)
	}

	_, ok = acc.TestAccProvider.DataSourcesMap["ibm_is_snapshot_software_attachment"]
	assert.True(t, ok, "data source ibm_is_snapshot_software_attachment must be registered in the provider")
	_, ok = acc.TestAccProvider.DataSourcesMap["ibm_is_snapshot_software_attachments"]
	assert.True(t, ok, "data source ibm_is_snapshot_software_attachments must be registered in the provider")
}

func TestResourceIBMIsSnapshotSoftwareAttachmentPatchAsPatch(t *testing.T) {
	r := vpc.ResourceIBMIsSnapshotSoftwareAttachment()

	// name set: it is sent in the patch.
	d := schema.TestResourceDataRaw(t, r.Schema, map[string]interface{}{
		"snapshot_id":                     "r006-00000000-0000-0000-0000-000000000000",
		"snapshot_software_attachment_id": "r006-00000000-0000-0000-0000-000000000001",
		"name":                            "my-software-attachment",
	})
	patchVals := &vpcv1.SnapshotSoftwareAttachmentPatch{Name: core.StringPtr("my-software-attachment")}
	patch := vpc.ResourceIBMIsSnapshotSoftwareAttachmentSnapshotSoftwareAttachmentPatchAsPatch(patchVals, d)
	assert.Equal(t, core.StringPtr("my-software-attachment"), patch["name"])

	// name not set and not changed: the key is left out of the patch.
	d = schema.TestResourceDataRaw(t, r.Schema, map[string]interface{}{
		"snapshot_id":                     "r006-00000000-0000-0000-0000-000000000000",
		"snapshot_software_attachment_id": "r006-00000000-0000-0000-0000-000000000001",
	})
	patch = vpc.ResourceIBMIsSnapshotSoftwareAttachmentSnapshotSoftwareAttachmentPatchAsPatch(&vpcv1.SnapshotSoftwareAttachmentPatch{}, d)
	_, exists := patch["name"]
	assert.False(t, exists)
}

func TestResourceIBMIsSnapshotSoftwareAttachmentSnapshotSoftwareAttachmentCatalogOfferingToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		deletedModel := make(map[string]interface{})
		deletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		catalogOfferingVersionPlanReferenceModel := make(map[string]interface{})
		catalogOfferingVersionPlanReferenceModel["crn"] = "crn:v1:bluemix:public:globalcatalog-collection:global:a/aa2432b1fa4d4ace891e9b80fc104e34:51c9e0db-2911-45a6-adb0-ac5332d27cf2:plan:sw.51c9e0db-2911-45a6-adb0-ac5332d27cf2.772c0dbe-aa62-482e-adbe-a3fc20101e0e"
		catalogOfferingVersionPlanReferenceModel["deleted"] = []map[string]interface{}{deletedModel}

		catalogOfferingVersionReferenceModel := make(map[string]interface{})
		catalogOfferingVersionReferenceModel["crn"] = "crn:v1:bluemix:public:globalcatalog-collection:global:a/aa2432b1fa4d4ace891e9b80fc104e34:1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc:version:00111601-0ec5-41ac-b142-96d1e64e6442/ec66bec2-6a33-42d6-9323-26dd4dc8875d"

		model := make(map[string]interface{})
		model["plan"] = []map[string]interface{}{catalogOfferingVersionPlanReferenceModel}
		model["version"] = []map[string]interface{}{catalogOfferingVersionReferenceModel}

		assert.Equal(t, result, model)
	}

	deletedModel := new(vpcv1.Deleted)
	deletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	catalogOfferingVersionPlanReferenceModel := new(vpcv1.CatalogOfferingVersionPlanReference)
	catalogOfferingVersionPlanReferenceModel.CRN = core.StringPtr("crn:v1:bluemix:public:globalcatalog-collection:global:a/aa2432b1fa4d4ace891e9b80fc104e34:51c9e0db-2911-45a6-adb0-ac5332d27cf2:plan:sw.51c9e0db-2911-45a6-adb0-ac5332d27cf2.772c0dbe-aa62-482e-adbe-a3fc20101e0e")
	catalogOfferingVersionPlanReferenceModel.Deleted = deletedModel

	catalogOfferingVersionReferenceModel := new(vpcv1.CatalogOfferingVersionReference)
	catalogOfferingVersionReferenceModel.CRN = core.StringPtr("crn:v1:bluemix:public:globalcatalog-collection:global:a/aa2432b1fa4d4ace891e9b80fc104e34:1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc:version:00111601-0ec5-41ac-b142-96d1e64e6442/ec66bec2-6a33-42d6-9323-26dd4dc8875d")

	model := new(vpcv1.SnapshotSoftwareAttachmentCatalogOffering)
	model.Plan = catalogOfferingVersionPlanReferenceModel
	model.Version = catalogOfferingVersionReferenceModel

	result, err := vpc.ResourceIBMIsSnapshotSoftwareAttachmentSnapshotSoftwareAttachmentCatalogOfferingToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsSnapshotSoftwareAttachmentCatalogOfferingVersionPlanReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		deletedModel := make(map[string]interface{})
		deletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		model := make(map[string]interface{})
		model["crn"] = "crn:v1:bluemix:public:globalcatalog-collection:global:a/aa2432b1fa4d4ace891e9b80fc104e34:51c9e0db-2911-45a6-adb0-ac5332d27cf2:plan:sw.51c9e0db-2911-45a6-adb0-ac5332d27cf2.772c0dbe-aa62-482e-adbe-a3fc20101e0e"
		model["deleted"] = []map[string]interface{}{deletedModel}

		assert.Equal(t, result, model)
	}

	deletedModel := new(vpcv1.Deleted)
	deletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	model := new(vpcv1.CatalogOfferingVersionPlanReference)
	model.CRN = core.StringPtr("crn:v1:bluemix:public:globalcatalog-collection:global:a/aa2432b1fa4d4ace891e9b80fc104e34:51c9e0db-2911-45a6-adb0-ac5332d27cf2:plan:sw.51c9e0db-2911-45a6-adb0-ac5332d27cf2.772c0dbe-aa62-482e-adbe-a3fc20101e0e")
	model.Deleted = deletedModel

	result, err := vpc.ResourceIBMIsSnapshotSoftwareAttachmentCatalogOfferingVersionPlanReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsSnapshotSoftwareAttachmentDeletedToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.Deleted)
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	result, err := vpc.ResourceIBMIsSnapshotSoftwareAttachmentDeletedToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsSnapshotSoftwareAttachmentCatalogOfferingVersionReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["crn"] = "crn:v1:bluemix:public:globalcatalog-collection:global:a/aa2432b1fa4d4ace891e9b80fc104e34:1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc:version:00111601-0ec5-41ac-b142-96d1e64e6442/ec66bec2-6a33-42d6-9323-26dd4dc8875d"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.CatalogOfferingVersionReference)
	model.CRN = core.StringPtr("crn:v1:bluemix:public:globalcatalog-collection:global:a/aa2432b1fa4d4ace891e9b80fc104e34:1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc:version:00111601-0ec5-41ac-b142-96d1e64e6442/ec66bec2-6a33-42d6-9323-26dd4dc8875d")

	result, err := vpc.ResourceIBMIsSnapshotSoftwareAttachmentCatalogOfferingVersionReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsSnapshotSoftwareAttachmentSnapshotSoftwareAttachmentEntitlementToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		snapshotSoftwareAttachmentEntitlementLicensableSoftwareModel := make(map[string]interface{})
		snapshotSoftwareAttachmentEntitlementLicensableSoftwareModel["sku"] = "FC1-10-IDCLD-445-02-12"

		model := make(map[string]interface{})
		model["licensable_software"] = []map[string]interface{}{snapshotSoftwareAttachmentEntitlementLicensableSoftwareModel}

		assert.Equal(t, result, model)
	}

	snapshotSoftwareAttachmentEntitlementLicensableSoftwareModel := new(vpcv1.SnapshotSoftwareAttachmentEntitlementLicensableSoftware)
	snapshotSoftwareAttachmentEntitlementLicensableSoftwareModel.Sku = core.StringPtr("FC1-10-IDCLD-445-02-12")

	model := new(vpcv1.SnapshotSoftwareAttachmentEntitlement)
	model.LicensableSoftware = []vpcv1.SnapshotSoftwareAttachmentEntitlementLicensableSoftware{*snapshotSoftwareAttachmentEntitlementLicensableSoftwareModel}

	result, err := vpc.ResourceIBMIsSnapshotSoftwareAttachmentSnapshotSoftwareAttachmentEntitlementToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsSnapshotSoftwareAttachmentSnapshotSoftwareAttachmentEntitlementLicensableSoftwareToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["sku"] = "FC1-10-IDCLD-445-02-12"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.SnapshotSoftwareAttachmentEntitlementLicensableSoftware)
	model.Sku = core.StringPtr("FC1-10-IDCLD-445-02-12")

	result, err := vpc.ResourceIBMIsSnapshotSoftwareAttachmentSnapshotSoftwareAttachmentEntitlementLicensableSoftwareToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsSnapshotSoftwareAttachmentCatalogOfferingWithoutPlanToMap(t *testing.T) {
	// A free catalog offering version has no billing plan: "plan" must be left out, not set to an empty list.
	catalogOfferingVersionReferenceModel := new(vpcv1.CatalogOfferingVersionReference)
	catalogOfferingVersionReferenceModel.CRN = core.StringPtr("crn:v1:bluemix:public:globalcatalog-collection:global:a/aa2432b1fa4d4ace891e9b80fc104e34:1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc:version:00111601-0ec5-41ac-b142-96d1e64e6442/ec66bec2-6a33-42d6-9323-26dd4dc8875d")

	model := new(vpcv1.SnapshotSoftwareAttachmentCatalogOffering)
	model.Version = catalogOfferingVersionReferenceModel

	result, err := vpc.ResourceIBMIsSnapshotSoftwareAttachmentSnapshotSoftwareAttachmentCatalogOfferingToMap(model)
	assert.Nil(t, err)
	_, hasPlan := result["plan"]
	assert.False(t, hasPlan)
	assert.Equal(t, []map[string]interface{}{{"crn": *catalogOfferingVersionReferenceModel.CRN}}, result["version"])
}

func TestResourceIBMIsSnapshotSoftwareAttachmentPlanWithoutDeletedToMap(t *testing.T) {
	model := new(vpcv1.CatalogOfferingVersionPlanReference)
	model.CRN = core.StringPtr("crn:v1:bluemix:public:globalcatalog-collection:global:a/aa2432b1fa4d4ace891e9b80fc104e34:51c9e0db-2911-45a6-adb0-ac5332d27cf2:plan:sw.51c9e0db-2911-45a6-adb0-ac5332d27cf2.772c0dbe-aa62-482e-adbe-a3fc20101e0e")

	result, err := vpc.ResourceIBMIsSnapshotSoftwareAttachmentCatalogOfferingVersionPlanReferenceToMap(model)
	assert.Nil(t, err)
	_, hasDeleted := result["deleted"]
	assert.False(t, hasDeleted)
	assert.Equal(t, *model.CRN, result["crn"])
}

func TestResourceIBMIsSnapshotSoftwareAttachmentEntitlementEmptyToMap(t *testing.T) {
	// No licensable software: an empty list, not nil, so the attribute count is 0.
	model := new(vpcv1.SnapshotSoftwareAttachmentEntitlement)
	model.LicensableSoftware = []vpcv1.SnapshotSoftwareAttachmentEntitlementLicensableSoftware{}

	result, err := vpc.ResourceIBMIsSnapshotSoftwareAttachmentSnapshotSoftwareAttachmentEntitlementToMap(model)
	assert.Nil(t, err)
	assert.Equal(t, []map[string]interface{}{}, result["licensable_software"])
}
