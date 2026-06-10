// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.111.0-1bfb72c2-20260206-185521
 */

package vpc_test

import (
	"fmt"
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

func TestAccIBMIsVolumeSoftwareAttachmentsDataSourceBasic(t *testing.T) {
	volumeSoftwareAttachmentVolumeID := fmt.Sprintf("tf_volume_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVolumeSoftwareAttachmentsDataSourceConfigBasic(volumeSoftwareAttachmentVolumeID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachments.is_volume_software_attachments_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachments.is_volume_software_attachments_instance", "volume_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachments.is_volume_software_attachments_instance", "software_attachments.#"),
				),
			},
		},
	})
}

func TestAccIBMIsVolumeSoftwareAttachmentsDataSourceAllArgs(t *testing.T) {
	volumeSoftwareAttachmentVolumeID := fmt.Sprintf("tf_volume_id_%d", acctest.RandIntRange(10, 100))
	volumeSoftwareAttachmentName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVolumeSoftwareAttachmentsDataSourceConfig(volumeSoftwareAttachmentVolumeID, volumeSoftwareAttachmentName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachments.is_volume_software_attachments_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachments.is_volume_software_attachments_instance", "volume_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachments.is_volume_software_attachments_instance", "software_attachments.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachments.is_volume_software_attachments_instance", "software_attachments.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachments.is_volume_software_attachments_instance", "software_attachments.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachments.is_volume_software_attachments_instance", "software_attachments.0.id"),
					resource.TestCheckResourceAttr("data.ibm_is_volume_software_attachments.is_volume_software_attachments_instance", "software_attachments.0.name", volumeSoftwareAttachmentName),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachments.is_volume_software_attachments_instance", "software_attachments.0.resource_type"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVolumeSoftwareAttachmentsDataSourceConfigBasic(volumeSoftwareAttachmentVolumeID string) string {
	return fmt.Sprintf(`
		resource "ibm_is_volume_software_attachment" "is_volume_software_attachment_instance" {
			volume_id = "%s"
		}

		data "ibm_is_volume_software_attachments" "is_volume_software_attachments_instance" {
			volume_id = ibm_is_volume_software_attachment.is_volume_software_attachment_instance.volume_id
		}
	`, volumeSoftwareAttachmentVolumeID)
}

func testAccCheckIBMIsVolumeSoftwareAttachmentsDataSourceConfig(volumeSoftwareAttachmentVolumeID string, volumeSoftwareAttachmentName string) string {
	return fmt.Sprintf(`
		resource "ibm_is_volume_software_attachment" "is_volume_software_attachment_instance" {
			volume_id = "%s"
			name = "%s"
		}

		data "ibm_is_volume_software_attachments" "is_volume_software_attachments_instance" {
			volume_id = ibm_is_volume_software_attachment.is_volume_software_attachment_instance.volume_id
		}
	`, volumeSoftwareAttachmentVolumeID, volumeSoftwareAttachmentName)
}

func TestDataSourceIBMIsVolumeSoftwareAttachmentsVolumeSoftwareAttachmentToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		deletedModel := make(map[string]interface{})
		deletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		catalogOfferingVersionPlanReferenceModel := make(map[string]interface{})
		catalogOfferingVersionPlanReferenceModel["crn"] = "crn:v1:bluemix:public:globalcatalog-collection:global:a/aa2432b1fa4d4ace891e9b80fc104e34:51c9e0db-2911-45a6-adb0-ac5332d27cf2:plan:sw.51c9e0db-2911-45a6-adb0-ac5332d27cf2.772c0dbe-aa62-482e-adbe-a3fc20101e0e"
		catalogOfferingVersionPlanReferenceModel["deleted"] = []map[string]interface{}{deletedModel}

		catalogOfferingVersionReferenceModel := make(map[string]interface{})
		catalogOfferingVersionReferenceModel["crn"] = "crn:v1:bluemix:public:globalcatalog-collection:global:a/aa2432b1fa4d4ace891e9b80fc104e34:51c9e0db-2911-45a6-adb0-ac5332d27cf2:plan:sw.51c9e0db-2911-45a6-adb0-ac5332d27cf2.772c0dbe-aa62-482e-adbe-a3fc20101e0e"

		volumeSoftwareAttachmentCatalogOfferingModel := make(map[string]interface{})
		volumeSoftwareAttachmentCatalogOfferingModel["plan"] = []map[string]interface{}{catalogOfferingVersionPlanReferenceModel}
		volumeSoftwareAttachmentCatalogOfferingModel["version"] = []map[string]interface{}{catalogOfferingVersionReferenceModel}

		volumeSoftwareAttachmentEntitlementLicensableSoftwareModel := make(map[string]interface{})
		volumeSoftwareAttachmentEntitlementLicensableSoftwareModel["sku"] = "FC1-10-IDCLD-445-02-12"

		volumeSoftwareAttachmentEntitlementModel := make(map[string]interface{})
		volumeSoftwareAttachmentEntitlementModel["licensable_software"] = []map[string]interface{}{volumeSoftwareAttachmentEntitlementLicensableSoftwareModel}

		model := make(map[string]interface{})
		model["catalog_offering"] = []map[string]interface{}{volumeSoftwareAttachmentCatalogOfferingModel}
		model["created_at"] = "2020-03-12T12:34:56Z"
		model["entitlement"] = []map[string]interface{}{volumeSoftwareAttachmentEntitlementModel}
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/volumes/0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e/software_attachments/r006-a569e8ae-3254-495e-ae75-86bb08e2c4d1"
		model["id"] = "0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
		model["name"] = "my-software-attachment"
		model["resource_type"] = "volume_software_attachment"

		assert.Equal(t, result, model)
	}

	deletedModel := new(vpcv1.Deleted)
	deletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	catalogOfferingVersionPlanReferenceModel := new(vpcv1.CatalogOfferingVersionPlanReference)
	catalogOfferingVersionPlanReferenceModel.CRN = core.StringPtr("crn:v1:bluemix:public:globalcatalog-collection:global:a/aa2432b1fa4d4ace891e9b80fc104e34:51c9e0db-2911-45a6-adb0-ac5332d27cf2:plan:sw.51c9e0db-2911-45a6-adb0-ac5332d27cf2.772c0dbe-aa62-482e-adbe-a3fc20101e0e")
	catalogOfferingVersionPlanReferenceModel.Deleted = deletedModel

	catalogOfferingVersionReferenceModel := new(vpcv1.CatalogOfferingVersionReference)
	catalogOfferingVersionReferenceModel.CRN = core.StringPtr("crn:v1:bluemix:public:globalcatalog-collection:global:a/aa2432b1fa4d4ace891e9b80fc104e34:51c9e0db-2911-45a6-adb0-ac5332d27cf2:plan:sw.51c9e0db-2911-45a6-adb0-ac5332d27cf2.772c0dbe-aa62-482e-adbe-a3fc20101e0e")

	volumeSoftwareAttachmentCatalogOfferingModel := new(vpcv1.VolumeSoftwareAttachmentCatalogOffering)
	volumeSoftwareAttachmentCatalogOfferingModel.Plan = catalogOfferingVersionPlanReferenceModel
	volumeSoftwareAttachmentCatalogOfferingModel.Version = catalogOfferingVersionReferenceModel

	volumeSoftwareAttachmentEntitlementLicensableSoftwareModel := new(vpcv1.VolumeSoftwareAttachmentEntitlementLicensableSoftware)
	volumeSoftwareAttachmentEntitlementLicensableSoftwareModel.Sku = core.StringPtr("FC1-10-IDCLD-445-02-12")

	volumeSoftwareAttachmentEntitlementModel := new(vpcv1.VolumeSoftwareAttachmentEntitlement)
	volumeSoftwareAttachmentEntitlementModel.LicensableSoftware = []vpcv1.VolumeSoftwareAttachmentEntitlementLicensableSoftware{*volumeSoftwareAttachmentEntitlementLicensableSoftwareModel}

	model := new(vpcv1.VolumeSoftwareAttachment)
	model.CatalogOffering = volumeSoftwareAttachmentCatalogOfferingModel
	model.CreatedAt = CreateMockDateTime("2020-03-12T12:34:56Z")
	model.Entitlement = volumeSoftwareAttachmentEntitlementModel
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/volumes/0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e/software_attachments/r006-a569e8ae-3254-495e-ae75-86bb08e2c4d1")
	model.ID = core.StringPtr("0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e")
	model.Name = core.StringPtr("my-software-attachment")
	model.ResourceType = core.StringPtr("volume_software_attachment")

	result, err := vpc.DataSourceIBMIsVolumeSoftwareAttachmentsVolumeSoftwareAttachmentToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsVolumeSoftwareAttachmentsVolumeSoftwareAttachmentCatalogOfferingToMap(t *testing.T) {
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

	model := new(vpcv1.VolumeSoftwareAttachmentCatalogOffering)
	model.Plan = catalogOfferingVersionPlanReferenceModel
	model.Version = catalogOfferingVersionReferenceModel

	result, err := vpc.DataSourceIBMIsVolumeSoftwareAttachmentsVolumeSoftwareAttachmentCatalogOfferingToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsVolumeSoftwareAttachmentsCatalogOfferingVersionPlanReferenceToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsVolumeSoftwareAttachmentsCatalogOfferingVersionPlanReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsVolumeSoftwareAttachmentsDeletedToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.Deleted)
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	result, err := vpc.DataSourceIBMIsVolumeSoftwareAttachmentsDeletedToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsVolumeSoftwareAttachmentsCatalogOfferingVersionReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["crn"] = "crn:v1:bluemix:public:globalcatalog-collection:global:a/aa2432b1fa4d4ace891e9b80fc104e34:1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc:version:00111601-0ec5-41ac-b142-96d1e64e6442/ec66bec2-6a33-42d6-9323-26dd4dc8875d"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.CatalogOfferingVersionReference)
	model.CRN = core.StringPtr("crn:v1:bluemix:public:globalcatalog-collection:global:a/aa2432b1fa4d4ace891e9b80fc104e34:1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc:version:00111601-0ec5-41ac-b142-96d1e64e6442/ec66bec2-6a33-42d6-9323-26dd4dc8875d")

	result, err := vpc.DataSourceIBMIsVolumeSoftwareAttachmentsCatalogOfferingVersionReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsVolumeSoftwareAttachmentsVolumeSoftwareAttachmentEntitlementToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		volumeSoftwareAttachmentEntitlementLicensableSoftwareModel := make(map[string]interface{})
		volumeSoftwareAttachmentEntitlementLicensableSoftwareModel["sku"] = "FC1-10-IDCLD-445-02-12"

		model := make(map[string]interface{})
		model["licensable_software"] = []map[string]interface{}{volumeSoftwareAttachmentEntitlementLicensableSoftwareModel}

		assert.Equal(t, result, model)
	}

	volumeSoftwareAttachmentEntitlementLicensableSoftwareModel := new(vpcv1.VolumeSoftwareAttachmentEntitlementLicensableSoftware)
	volumeSoftwareAttachmentEntitlementLicensableSoftwareModel.Sku = core.StringPtr("FC1-10-IDCLD-445-02-12")

	model := new(vpcv1.VolumeSoftwareAttachmentEntitlement)
	model.LicensableSoftware = []vpcv1.VolumeSoftwareAttachmentEntitlementLicensableSoftware{*volumeSoftwareAttachmentEntitlementLicensableSoftwareModel}

	result, err := vpc.DataSourceIBMIsVolumeSoftwareAttachmentsVolumeSoftwareAttachmentEntitlementToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsVolumeSoftwareAttachmentsVolumeSoftwareAttachmentEntitlementLicensableSoftwareToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["sku"] = "FC1-10-IDCLD-445-02-12"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.VolumeSoftwareAttachmentEntitlementLicensableSoftware)
	model.Sku = core.StringPtr("FC1-10-IDCLD-445-02-12")

	result, err := vpc.DataSourceIBMIsVolumeSoftwareAttachmentsVolumeSoftwareAttachmentEntitlementLicensableSoftwareToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
