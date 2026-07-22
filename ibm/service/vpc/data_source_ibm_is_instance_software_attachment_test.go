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

func TestAccIBMIsInstanceSoftwareAttachmentDataSourceBasic(t *testing.T) {
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf-instance-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsInstanceSoftwareAttachmentDataSourceConfigBasic(vpcname, subnetname, sshname, instanceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_software_attachment.is_instance_software_attachment_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_software_attachment.is_instance_software_attachment_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_software_attachment.is_instance_software_attachment_instance", "instance_software_attachment_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_software_attachment.is_instance_software_attachment_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_software_attachment.is_instance_software_attachment_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_software_attachment.is_instance_software_attachment_instance", "lifecycle_reasons.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_software_attachment.is_instance_software_attachment_instance", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_software_attachment.is_instance_software_attachment_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_software_attachment.is_instance_software_attachment_instance", "resource_type"),
				),
			},
		},
	})
}

func TestAccIBMIsInstanceSoftwareAttachmentDataSourceAllArgs(t *testing.T) {
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf-instance-%d", acctest.RandIntRange(10, 100))
	instanceSoftwareAttachmentName := fmt.Sprintf("tf-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsInstanceSoftwareAttachmentDataSourceConfig(vpcname, subnetname, sshname, instanceName, instanceSoftwareAttachmentName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_software_attachment.is_instance_software_attachment_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_software_attachment.is_instance_software_attachment_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_software_attachment.is_instance_software_attachment_instance", "instance_software_attachment_id"),
					resource.TestCheckResourceAttr("data.ibm_is_instance_software_attachment.is_instance_software_attachment_instance", "name", instanceSoftwareAttachmentName),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_software_attachment.is_instance_software_attachment_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_software_attachment.is_instance_software_attachment_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_software_attachment.is_instance_software_attachment_instance", "lifecycle_reasons.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_software_attachment.is_instance_software_attachment_instance", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_software_attachment.is_instance_software_attachment_instance", "resource_type"),
				),
			},
		},
	})
}

func testAccCheckIBMIsInstanceSoftwareAttachmentDataSourceConfigBasic(vpcname, subnetname, sshname, instanceName string) string {
	return testAccCheckIBMIsInstanceSoftwareAttachmentBaseConfig(vpcname, subnetname, sshname, instanceName) + `
		resource "ibm_is_instance_software_attachment" "is_instance_software_attachment_instance" {
			instance_id                        = ibm_is_instance.testacc_instance.id
			instance_software_attachment_id = ibm_is_instance.testacc_instance.software_attachments.0.id
		}

		data "ibm_is_instance_software_attachment" "is_instance_software_attachment_instance" {
			instance_id                        = ibm_is_instance_software_attachment.is_instance_software_attachment_instance.instance_id
			instance_software_attachment_id = ibm_is_instance_software_attachment.is_instance_software_attachment_instance.instance_software_attachment_id
		}
	`
}

func testAccCheckIBMIsInstanceSoftwareAttachmentDataSourceConfig(vpcname, subnetname, sshname, instanceName, instanceSoftwareAttachmentName string) string {
	return testAccCheckIBMIsInstanceSoftwareAttachmentBaseConfig(vpcname, subnetname, sshname, instanceName) + fmt.Sprintf(`
		resource "ibm_is_instance_software_attachment" "is_instance_software_attachment_instance" {
			instance_id                        = ibm_is_instance.testacc_instance.id
			instance_software_attachment_id = ibm_is_instance.testacc_instance.software_attachments.0.id
			name                               = "%s"
		}

		data "ibm_is_instance_software_attachment" "is_instance_software_attachment_instance" {
			instance_id                        = ibm_is_instance_software_attachment.is_instance_software_attachment_instance.instance_id
			instance_software_attachment_id = ibm_is_instance_software_attachment.is_instance_software_attachment_instance.instance_software_attachment_id
		}
	`, instanceSoftwareAttachmentName)
}

func TestDataSourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentCatalogOfferingToMap(t *testing.T) {
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

	model := new(vpcv1.InstanceSoftwareAttachmentCatalogOffering)
	model.Plan = catalogOfferingVersionPlanReferenceModel
	model.Version = catalogOfferingVersionReferenceModel

	result, err := vpc.DataSourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentCatalogOfferingToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsInstanceSoftwareAttachmentCatalogOfferingVersionPlanReferenceToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsInstanceSoftwareAttachmentCatalogOfferingVersionPlanReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsInstanceSoftwareAttachmentDeletedToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.Deleted)
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	result, err := vpc.DataSourceIBMIsInstanceSoftwareAttachmentDeletedToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsInstanceSoftwareAttachmentCatalogOfferingVersionReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["crn"] = "crn:v1:bluemix:public:globalcatalog-collection:global:a/aa2432b1fa4d4ace891e9b80fc104e34:1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc:version:00111601-0ec5-41ac-b142-96d1e64e6442/ec66bec2-6a33-42d6-9323-26dd4dc8875d"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.CatalogOfferingVersionReference)
	model.CRN = core.StringPtr("crn:v1:bluemix:public:globalcatalog-collection:global:a/aa2432b1fa4d4ace891e9b80fc104e34:1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc:version:00111601-0ec5-41ac-b142-96d1e64e6442/ec66bec2-6a33-42d6-9323-26dd4dc8875d")

	result, err := vpc.DataSourceIBMIsInstanceSoftwareAttachmentCatalogOfferingVersionReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentEntitlementToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		instanceSoftwareAttachmentEntitlementLicensedSoftwareModel := make(map[string]interface{})
		instanceSoftwareAttachmentEntitlementLicensedSoftwareModel["sku"] = "FC1-10-IDCLD-445-02-12"

		model := make(map[string]interface{})
		model["licensed_software"] = []map[string]interface{}{instanceSoftwareAttachmentEntitlementLicensedSoftwareModel}

		assert.Equal(t, result, model)
	}

	instanceSoftwareAttachmentEntitlementLicensedSoftwareModel := new(vpcv1.InstanceSoftwareAttachmentEntitlementLicensedSoftware)
	instanceSoftwareAttachmentEntitlementLicensedSoftwareModel.Sku = core.StringPtr("FC1-10-IDCLD-445-02-12")

	model := new(vpcv1.InstanceSoftwareAttachmentEntitlement)
	model.LicensedSoftware = []vpcv1.InstanceSoftwareAttachmentEntitlementLicensedSoftware{*instanceSoftwareAttachmentEntitlementLicensedSoftwareModel}

	result, err := vpc.DataSourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentEntitlementToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentEntitlementLicensedSoftwareToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["sku"] = "FC1-10-IDCLD-445-02-12"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.InstanceSoftwareAttachmentEntitlementLicensedSoftware)
	model.Sku = core.StringPtr("FC1-10-IDCLD-445-02-12")

	result, err := vpc.DataSourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentEntitlementLicensedSoftwareToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentLifecycleReasonToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["code"] = "pending_registration"
		model["message"] = "The software instance is being registered with Resource Controller"
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#failed-registration"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.InstanceSoftwareAttachmentLifecycleReason)
	model.Code = core.StringPtr("pending_registration")
	model.Message = core.StringPtr("The software instance is being registered with Resource Controller")
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#failed-registration")

	result, err := vpc.DataSourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentLifecycleReasonToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentOfferingInstanceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["crn"] = "crn:v1:bluemix:public:globalcatalog-instance:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34:oi-07a7_26fb923a-873f-43bf-94e5-93431235ad5f-31323517-65cc-49cf-9281-13ee3399b747::"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.InstanceSoftwareAttachmentOfferingInstance)
	model.CRN = core.StringPtr("crn:v1:bluemix:public:globalcatalog-instance:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34:oi-07a7_26fb923a-873f-43bf-94e5-93431235ad5f-31323517-65cc-49cf-9281-13ee3399b747::")

	result, err := vpc.DataSourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentOfferingInstanceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
