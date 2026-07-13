// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/vpc"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMIsInstanceSoftwareAttachmentBasic(t *testing.T) {
	var conf vpcv1.InstanceSoftwareAttachment
	instanceID := fmt.Sprintf("tf_instance_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsInstanceSoftwareAttachmentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsInstanceSoftwareAttachmentConfigBasic(instanceID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsInstanceSoftwareAttachmentExists("ibm_is_instance_software_attachment.is_instance_software_attachment_instance", conf),
					resource.TestCheckResourceAttr("ibm_is_instance_software_attachment.is_instance_software_attachment_instance", "instance_id", instanceID),
				),
			},
		},
	})
}

func TestAccIBMIsInstanceSoftwareAttachmentAllArgs(t *testing.T) {
	var conf vpcv1.InstanceSoftwareAttachment
	instanceID := fmt.Sprintf("tf_instance_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsInstanceSoftwareAttachmentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsInstanceSoftwareAttachmentConfig(instanceID, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsInstanceSoftwareAttachmentExists("ibm_is_instance_software_attachment.is_instance_software_attachment_instance", conf),
					resource.TestCheckResourceAttr("ibm_is_instance_software_attachment.is_instance_software_attachment_instance", "instance_id", instanceID),
					resource.TestCheckResourceAttr("ibm_is_instance_software_attachment.is_instance_software_attachment_instance", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsInstanceSoftwareAttachmentConfig(instanceID, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_instance_software_attachment.is_instance_software_attachment_instance", "instance_id", instanceID),
					resource.TestCheckResourceAttr("ibm_is_instance_software_attachment.is_instance_software_attachment_instance", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_is_instance_software_attachment.is_instance_software_attachment_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIsInstanceSoftwareAttachmentConfigBasic(instanceID string) string {
	return fmt.Sprintf(`
		resource "ibm_is_instance_software_attachment" "is_instance_software_attachment_instance" {
			instance_id = "%s"
		}
	`, instanceID)
}

func testAccCheckIBMIsInstanceSoftwareAttachmentConfig(instanceID string, name string) string {
	return fmt.Sprintf(`

		resource "ibm_is_instance_software_attachment" "is_instance_software_attachment_instance" {
			instance_id = "%s"
			name = "%s"
		}
	`, instanceID, name)
}

func testAccCheckIBMIsInstanceSoftwareAttachmentExists(n string, obj vpcv1.InstanceSoftwareAttachment) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		if err != nil {
			return err
		}

		getInstanceSoftwareAttachmentOptions := &vpcv1.GetInstanceSoftwareAttachmentOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getInstanceSoftwareAttachmentOptions.SetInstanceID(parts[0])
		getInstanceSoftwareAttachmentOptions.SetID(parts[1])

		instanceSoftwareAttachment, _, err := vpcClient.GetInstanceSoftwareAttachment(getInstanceSoftwareAttachmentOptions)
		if err != nil {
			return err
		}

		obj = *instanceSoftwareAttachment
		return nil
	}
}

func testAccCheckIBMIsInstanceSoftwareAttachmentDestroy(s *terraform.State) error {
	vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_instance_software_attachment" {
			continue
		}

		getInstanceSoftwareAttachmentOptions := &vpcv1.GetInstanceSoftwareAttachmentOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getInstanceSoftwareAttachmentOptions.SetInstanceID(parts[0])
		getInstanceSoftwareAttachmentOptions.SetID(parts[1])

		// Try to find the key
		_, response, err := vpcClient.GetInstanceSoftwareAttachment(getInstanceSoftwareAttachmentOptions)

		if err == nil {
			return fmt.Errorf("InstanceSoftwareAttachment still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for InstanceSoftwareAttachment (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentCatalogOfferingToMap(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentCatalogOfferingToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsInstanceSoftwareAttachmentCatalogOfferingVersionPlanReferenceToMap(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsInstanceSoftwareAttachmentCatalogOfferingVersionPlanReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsInstanceSoftwareAttachmentDeletedToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.Deleted)
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	result, err := vpc.ResourceIBMIsInstanceSoftwareAttachmentDeletedToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsInstanceSoftwareAttachmentCatalogOfferingVersionReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["crn"] = "crn:v1:bluemix:public:globalcatalog-collection:global:a/aa2432b1fa4d4ace891e9b80fc104e34:1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc:version:00111601-0ec5-41ac-b142-96d1e64e6442/ec66bec2-6a33-42d6-9323-26dd4dc8875d"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.CatalogOfferingVersionReference)
	model.CRN = core.StringPtr("crn:v1:bluemix:public:globalcatalog-collection:global:a/aa2432b1fa4d4ace891e9b80fc104e34:1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc:version:00111601-0ec5-41ac-b142-96d1e64e6442/ec66bec2-6a33-42d6-9323-26dd4dc8875d")

	result, err := vpc.ResourceIBMIsInstanceSoftwareAttachmentCatalogOfferingVersionReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentEntitlementToMap(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentEntitlementToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentEntitlementLicensedSoftwareToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["sku"] = "FC1-10-IDCLD-445-02-12"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.InstanceSoftwareAttachmentEntitlementLicensedSoftware)
	model.Sku = core.StringPtr("FC1-10-IDCLD-445-02-12")

	result, err := vpc.ResourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentEntitlementLicensedSoftwareToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentLifecycleReasonToMap(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentLifecycleReasonToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentOfferingInstanceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["crn"] = "crn:v1:bluemix:public:globalcatalog-instance:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34:oi-07a7_26fb923a-873f-43bf-94e5-93431235ad5f-31323517-65cc-49cf-9281-13ee3399b747::"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.InstanceSoftwareAttachmentOfferingInstance)
	model.CRN = core.StringPtr("crn:v1:bluemix:public:globalcatalog-instance:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34:oi-07a7_26fb923a-873f-43bf-94e5-93431235ad5f-31323517-65cc-49cf-9281-13ee3399b747::")

	result, err := vpc.ResourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentOfferingInstanceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
