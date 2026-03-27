// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.111.0-1bfb72c2-20260206-185521
 */

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/vpc"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMIsVolumeSoftwareAttachmentDataSourceBasic(t *testing.T) {
	volumeSoftwareAttachmentVolumeID := fmt.Sprintf("tf_volume_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVolumeSoftwareAttachmentDataSourceConfigBasic(volumeSoftwareAttachmentVolumeID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachment.is_volume_software_attachment_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachment.is_volume_software_attachment_instance", "volume_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachment.is_volume_software_attachment_instance", "is_volume_software_attachment_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachment.is_volume_software_attachment_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachment.is_volume_software_attachment_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachment.is_volume_software_attachment_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachment.is_volume_software_attachment_instance", "resource_type"),
				),
			},
		},
	})
}

func TestAccIBMIsVolumeSoftwareAttachmentDataSourceAllArgs(t *testing.T) {
	volumeSoftwareAttachmentVolumeID := fmt.Sprintf("tf_volume_id_%d", acctest.RandIntRange(10, 100))
	volumeSoftwareAttachmentName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVolumeSoftwareAttachmentDataSourceConfig(volumeSoftwareAttachmentVolumeID, volumeSoftwareAttachmentName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachment.is_volume_software_attachment_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachment.is_volume_software_attachment_instance", "volume_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachment.is_volume_software_attachment_instance", "is_volume_software_attachment_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachment.is_volume_software_attachment_instance", "catalog_offering.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachment.is_volume_software_attachment_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachment.is_volume_software_attachment_instance", "entitlement.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachment.is_volume_software_attachment_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachment.is_volume_software_attachment_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachment.is_volume_software_attachment_instance", "resource_type"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVolumeSoftwareAttachmentDataSourceConfigBasic(volumeSoftwareAttachmentVolumeID string) string {
	return fmt.Sprintf(`
		resource "ibm_is_volume_software_attachment" "is_volume_software_attachment_instance" {
			volume_id = "%s"
		}

		data "ibm_is_volume_software_attachment" "is_volume_software_attachment_instance" {
			volume_id = ibm_is_volume_software_attachment.is_volume_software_attachment_instance.volume_id
			is_volume_software_attachment_id = ibm_is_volume_software_attachment.is_volume_software_attachment_instance.is_volume_software_attachment_id
		}
	`, volumeSoftwareAttachmentVolumeID)
}

func testAccCheckIBMIsVolumeSoftwareAttachmentDataSourceConfig(volumeSoftwareAttachmentVolumeID string, volumeSoftwareAttachmentName string) string {
	return fmt.Sprintf(`
		resource "ibm_is_volume_software_attachment" "is_volume_software_attachment_instance" {
			volume_id = "%s"
			name = "%s"
		}

		data "ibm_is_volume_software_attachment" "is_volume_software_attachment_instance" {
			volume_id = ibm_is_volume_software_attachment.is_volume_software_attachment_instance.volume_id
			is_volume_software_attachment_id = ibm_is_volume_software_attachment.is_volume_software_attachment_instance.is_volume_software_attachment_id
		}
	`, volumeSoftwareAttachmentVolumeID, volumeSoftwareAttachmentName)
}

func TestDataSourceIBMIsVolumeSoftwareAttachmentVolumeSoftwareAttachmentCatalogOfferingToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsVolumeSoftwareAttachmentVolumeSoftwareAttachmentCatalogOfferingToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsVolumeSoftwareAttachmentCatalogOfferingVersionPlanReferenceToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsVolumeSoftwareAttachmentCatalogOfferingVersionPlanReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsVolumeSoftwareAttachmentDeletedToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.Deleted)
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	result, err := vpc.DataSourceIBMIsVolumeSoftwareAttachmentDeletedToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsVolumeSoftwareAttachmentCatalogOfferingVersionReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["crn"] = "crn:v1:bluemix:public:globalcatalog-collection:global:a/aa2432b1fa4d4ace891e9b80fc104e34:1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc:version:00111601-0ec5-41ac-b142-96d1e64e6442/ec66bec2-6a33-42d6-9323-26dd4dc8875d"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.CatalogOfferingVersionReference)
	model.CRN = core.StringPtr("crn:v1:bluemix:public:globalcatalog-collection:global:a/aa2432b1fa4d4ace891e9b80fc104e34:1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc:version:00111601-0ec5-41ac-b142-96d1e64e6442/ec66bec2-6a33-42d6-9323-26dd4dc8875d")

	result, err := vpc.DataSourceIBMIsVolumeSoftwareAttachmentCatalogOfferingVersionReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsVolumeSoftwareAttachmentVolumeSoftwareAttachmentEntitlementToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsVolumeSoftwareAttachmentVolumeSoftwareAttachmentEntitlementToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsVolumeSoftwareAttachmentVolumeSoftwareAttachmentEntitlementLicensableSoftwareToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["sku"] = "FC1-10-IDCLD-445-02-12"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.VolumeSoftwareAttachmentEntitlementLicensableSoftware)
	model.Sku = core.StringPtr("FC1-10-IDCLD-445-02-12")

	result, err := vpc.DataSourceIBMIsVolumeSoftwareAttachmentVolumeSoftwareAttachmentEntitlementLicensableSoftwareToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
