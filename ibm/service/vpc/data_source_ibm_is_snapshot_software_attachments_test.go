// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.111.0-1bfb72c2-20260206-185521
 */

package vpc_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/vpc"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMIsSnapshotSoftwareAttachmentsDataSourceBasic(t *testing.T) {
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf-instance-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsSnapshotSoftwareAttachmentBaseConfig(vpcname, subnetname, sshname, instanceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_software_attachments.is_snapshot_software_attachments_instance", "id"),
					resource.TestCheckResourceAttrPair("data.ibm_is_snapshot_software_attachments.is_snapshot_software_attachments_instance", "snapshot_id", "ibm_is_snapshot.testacc_snapshot", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_software_attachments.is_snapshot_software_attachments_instance", "software_attachments.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_software_attachments.is_snapshot_software_attachments_instance", "software_attachments.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_software_attachments.is_snapshot_software_attachments_instance", "software_attachments.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_software_attachments.is_snapshot_software_attachments_instance", "software_attachments.0.href"),
					resource.TestCheckResourceAttr("data.ibm_is_snapshot_software_attachments.is_snapshot_software_attachments_instance", "software_attachments.0.resource_type", "snapshot_software_attachment"),
				),
			},
		},
	})
}

func TestAccIBMIsSnapshotSoftwareAttachmentsDataSourceAllArgs(t *testing.T) {
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf-instance-%d", acctest.RandIntRange(10, 100))
	snapshotSoftwareAttachmentName := fmt.Sprintf("tf-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsSnapshotSoftwareAttachmentsDataSourceConfig(vpcname, subnetname, sshname, instanceName, snapshotSoftwareAttachmentName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_software_attachments.is_snapshot_software_attachments_after_rename", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_software_attachments.is_snapshot_software_attachments_after_rename", "software_attachments.#"),
					resource.TestCheckResourceAttrPair("data.ibm_is_snapshot_software_attachments.is_snapshot_software_attachments_after_rename", "software_attachments.0.id", "ibm_is_snapshot_software_attachment.is_snapshot_software_attachment_instance", "snapshot_software_attachment_id"),
					resource.TestCheckResourceAttr("data.ibm_is_snapshot_software_attachments.is_snapshot_software_attachments_after_rename", "software_attachments.0.name", snapshotSoftwareAttachmentName),
				),
			},
		},
	})
}

// A snapshot taken from a volume without software attachments has no software
// attachments. The list must be empty (count 0), not an error.
func TestAccIBMIsSnapshotSoftwareAttachmentsDataSourceEmpty(t *testing.T) {
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf-instance-%d", acctest.RandIntRange(10, 100))
	volname := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))
	snapshotName := fmt.Sprintf("tf-snapshot-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				// The instance boots from a stock image, so its volumes and their snapshots carry no software attachments.
				Config: testDSCheckIBMISSnapshotConfig(vpcname, subnetname, sshname, publicKey, volname, instanceName, snapshotName) + `
					data "ibm_is_snapshot_software_attachments" "is_snapshot_software_attachments_instance" {
						snapshot_id = ibm_is_snapshot.testacc_snapshot.id
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_snapshot_software_attachments.is_snapshot_software_attachments_instance", "id"),
					resource.TestCheckResourceAttr("data.ibm_is_snapshot_software_attachments.is_snapshot_software_attachments_instance", "software_attachments.#", "0"),
				),
			},
		},
	})
}

func testAccCheckIBMIsSnapshotSoftwareAttachmentsDataSourceConfig(vpcname, subnetname, sshname, instanceName, snapshotSoftwareAttachmentName string) string {
	return testAccCheckIBMIsSnapshotSoftwareAttachmentConfig(vpcname, subnetname, sshname, instanceName, snapshotSoftwareAttachmentName) + `
		data "ibm_is_snapshot_software_attachments" "is_snapshot_software_attachments_after_rename" {
			snapshot_id = ibm_is_snapshot_software_attachment.is_snapshot_software_attachment_instance.snapshot_id
		}
	`
}

func TestDataSourceIBMIsSnapshotSoftwareAttachmentsSnapshotSoftwareAttachmentMinimalToMap(t *testing.T) {
	// catalog_offering and entitlement may be absent (for example when the
	// software attachment is not stable yet): those keys must be left out.
	model := new(vpcv1.SnapshotSoftwareAttachment)
	model.CreatedAt = CreateMockDateTime("2020-03-12T12:34:56Z")
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/snapshots/0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e/software_attachments/r006-a569e8ae-3254-495e-ae75-86bb08e2c4d1")
	model.ID = core.StringPtr("0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")
	model.Name = core.StringPtr("my-software-attachment")
	model.ResourceType = core.StringPtr("snapshot_software_attachment")

	result, err := vpc.DataSourceIBMIsSnapshotSoftwareAttachmentsSnapshotSoftwareAttachmentToMap(model)
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{
		"created_at":    "2020-03-12T12:34:56.000Z",
		"href":          *model.Href,
		"id":            *model.ID,
		"name":          *model.Name,
		"resource_type": *model.ResourceType,
	}, result)
}

func TestDataSourceIBMIsSnapshotSoftwareAttachmentsSnapshotSoftwareAttachmentToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		deletedModel := make(map[string]interface{})
		deletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		catalogOfferingVersionPlanReferenceModel := make(map[string]interface{})
		catalogOfferingVersionPlanReferenceModel["crn"] = "crn:v1:bluemix:public:globalcatalog-collection:global:a/aa2432b1fa4d4ace891e9b80fc104e34:51c9e0db-2911-45a6-adb0-ac5332d27cf2:plan:sw.51c9e0db-2911-45a6-adb0-ac5332d27cf2.772c0dbe-aa62-482e-adbe-a3fc20101e0e"
		catalogOfferingVersionPlanReferenceModel["deleted"] = []map[string]interface{}{deletedModel}

		catalogOfferingVersionReferenceModel := make(map[string]interface{})
		catalogOfferingVersionReferenceModel["crn"] = "crn:v1:bluemix:public:globalcatalog-collection:global:a/aa2432b1fa4d4ace891e9b80fc104e34:51c9e0db-2911-45a6-adb0-ac5332d27cf2:plan:sw.51c9e0db-2911-45a6-adb0-ac5332d27cf2.772c0dbe-aa62-482e-adbe-a3fc20101e0e"

		snapshotSoftwareAttachmentCatalogOfferingModel := make(map[string]interface{})
		snapshotSoftwareAttachmentCatalogOfferingModel["plan"] = []map[string]interface{}{catalogOfferingVersionPlanReferenceModel}
		snapshotSoftwareAttachmentCatalogOfferingModel["version"] = []map[string]interface{}{catalogOfferingVersionReferenceModel}

		snapshotSoftwareAttachmentEntitlementLicensableSoftwareModel := make(map[string]interface{})
		snapshotSoftwareAttachmentEntitlementLicensableSoftwareModel["sku"] = "FC1-10-IDCLD-445-02-12"

		snapshotSoftwareAttachmentEntitlementModel := make(map[string]interface{})
		snapshotSoftwareAttachmentEntitlementModel["licensable_software"] = []map[string]interface{}{snapshotSoftwareAttachmentEntitlementLicensableSoftwareModel}

		model := make(map[string]interface{})
		model["catalog_offering"] = []map[string]interface{}{snapshotSoftwareAttachmentCatalogOfferingModel}
		model["created_at"] = "2020-03-12T12:34:56.000Z"
		model["entitlement"] = []map[string]interface{}{snapshotSoftwareAttachmentEntitlementModel}
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/snapshots/0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e/software_attachments/r006-a569e8ae-3254-495e-ae75-86bb08e2c4d1"
		model["id"] = "0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
		model["name"] = "my-software-attachment"
		model["resource_type"] = "snapshot_software_attachment"

		assert.Equal(t, result, model)
	}

	deletedModel := new(vpcv1.Deleted)
	deletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	catalogOfferingVersionPlanReferenceModel := new(vpcv1.CatalogOfferingVersionPlanReference)
	catalogOfferingVersionPlanReferenceModel.CRN = core.StringPtr("crn:v1:bluemix:public:globalcatalog-collection:global:a/aa2432b1fa4d4ace891e9b80fc104e34:51c9e0db-2911-45a6-adb0-ac5332d27cf2:plan:sw.51c9e0db-2911-45a6-adb0-ac5332d27cf2.772c0dbe-aa62-482e-adbe-a3fc20101e0e")
	catalogOfferingVersionPlanReferenceModel.Deleted = deletedModel

	catalogOfferingVersionReferenceModel := new(vpcv1.CatalogOfferingVersionReference)
	catalogOfferingVersionReferenceModel.CRN = core.StringPtr("crn:v1:bluemix:public:globalcatalog-collection:global:a/aa2432b1fa4d4ace891e9b80fc104e34:51c9e0db-2911-45a6-adb0-ac5332d27cf2:plan:sw.51c9e0db-2911-45a6-adb0-ac5332d27cf2.772c0dbe-aa62-482e-adbe-a3fc20101e0e")

	snapshotSoftwareAttachmentCatalogOfferingModel := new(vpcv1.SnapshotSoftwareAttachmentCatalogOffering)
	snapshotSoftwareAttachmentCatalogOfferingModel.Plan = catalogOfferingVersionPlanReferenceModel
	snapshotSoftwareAttachmentCatalogOfferingModel.Version = catalogOfferingVersionReferenceModel

	snapshotSoftwareAttachmentEntitlementLicensableSoftwareModel := new(vpcv1.SnapshotSoftwareAttachmentEntitlementLicensableSoftware)
	snapshotSoftwareAttachmentEntitlementLicensableSoftwareModel.Sku = core.StringPtr("FC1-10-IDCLD-445-02-12")

	snapshotSoftwareAttachmentEntitlementModel := new(vpcv1.SnapshotSoftwareAttachmentEntitlement)
	snapshotSoftwareAttachmentEntitlementModel.LicensableSoftware = []vpcv1.SnapshotSoftwareAttachmentEntitlementLicensableSoftware{*snapshotSoftwareAttachmentEntitlementLicensableSoftwareModel}

	model := new(vpcv1.SnapshotSoftwareAttachment)
	model.CatalogOffering = snapshotSoftwareAttachmentCatalogOfferingModel
	model.CreatedAt = CreateMockDateTime("2020-03-12T12:34:56Z")
	model.Entitlement = snapshotSoftwareAttachmentEntitlementModel
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/snapshots/0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e/software_attachments/r006-a569e8ae-3254-495e-ae75-86bb08e2c4d1")
	model.ID = core.StringPtr("0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")
	model.Name = core.StringPtr("my-software-attachment")
	model.ResourceType = core.StringPtr("snapshot_software_attachment")

	result, err := vpc.DataSourceIBMIsSnapshotSoftwareAttachmentsSnapshotSoftwareAttachmentToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsSnapshotSoftwareAttachmentsSnapshotSoftwareAttachmentCatalogOfferingToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsSnapshotSoftwareAttachmentsSnapshotSoftwareAttachmentCatalogOfferingToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsSnapshotSoftwareAttachmentsCatalogOfferingVersionPlanReferenceToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsSnapshotSoftwareAttachmentsCatalogOfferingVersionPlanReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsSnapshotSoftwareAttachmentsDeletedToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.Deleted)
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	result, err := vpc.DataSourceIBMIsSnapshotSoftwareAttachmentsDeletedToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsSnapshotSoftwareAttachmentsCatalogOfferingVersionReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["crn"] = "crn:v1:bluemix:public:globalcatalog-collection:global:a/aa2432b1fa4d4ace891e9b80fc104e34:1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc:version:00111601-0ec5-41ac-b142-96d1e64e6442/ec66bec2-6a33-42d6-9323-26dd4dc8875d"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.CatalogOfferingVersionReference)
	model.CRN = core.StringPtr("crn:v1:bluemix:public:globalcatalog-collection:global:a/aa2432b1fa4d4ace891e9b80fc104e34:1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc:version:00111601-0ec5-41ac-b142-96d1e64e6442/ec66bec2-6a33-42d6-9323-26dd4dc8875d")

	result, err := vpc.DataSourceIBMIsSnapshotSoftwareAttachmentsCatalogOfferingVersionReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsSnapshotSoftwareAttachmentsSnapshotSoftwareAttachmentEntitlementToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsSnapshotSoftwareAttachmentsSnapshotSoftwareAttachmentEntitlementToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsSnapshotSoftwareAttachmentsSnapshotSoftwareAttachmentEntitlementLicensableSoftwareToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["sku"] = "FC1-10-IDCLD-445-02-12"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.SnapshotSoftwareAttachmentEntitlementLicensableSoftware)
	model.Sku = core.StringPtr("FC1-10-IDCLD-445-02-12")

	result, err := vpc.DataSourceIBMIsSnapshotSoftwareAttachmentsSnapshotSoftwareAttachmentEntitlementLicensableSoftwareToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
