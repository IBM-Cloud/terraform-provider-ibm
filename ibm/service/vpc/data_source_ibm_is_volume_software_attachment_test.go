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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/vpc"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMIsVolumeSoftwareAttachmentDataSourceBasic(t *testing.T) {
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf-instance-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVolumeSoftwareAttachmentDataSourceConfigBasic(vpcname, subnetname, sshname, instanceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachment.is_volume_software_attachment", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachment.is_volume_software_attachment", "volume_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachment.is_volume_software_attachment", "volume_software_attachment_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachment.is_volume_software_attachment", "catalog_offering.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachment.is_volume_software_attachment", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachment.is_volume_software_attachment", "entitlement.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachment.is_volume_software_attachment", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachment.is_volume_software_attachment", "name"),
					resource.TestCheckResourceAttr("data.ibm_is_volume_software_attachment.is_volume_software_attachment", "resource_type", "volume_software_attachment"),
				),
			},
		},
	})
}

func TestAccIBMIsVolumeSoftwareAttachmentDataSourceAllArgs(t *testing.T) {
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf-instance-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVolumeSoftwareAttachmentDataSourceConfig(vpcname, subnetname, sshname, instanceName, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachment.is_volume_software_attachment", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachment.is_volume_software_attachment", "volume_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachment.is_volume_software_attachment", "volume_software_attachment_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachment.is_volume_software_attachment", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_software_attachment.is_volume_software_attachment", "href"),
					resource.TestCheckResourceAttr("data.ibm_is_volume_software_attachment.is_volume_software_attachment", "name", name),
					resource.TestCheckResourceAttr("data.ibm_is_volume_software_attachment.is_volume_software_attachment", "resource_type", "volume_software_attachment"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVolumeSoftwareAttachmentDataSourceConfigBasic(vpcname, subnetname, sshname, instanceName string) string {
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	return fmt.Sprintf(`
		resource "ibm_is_vpc" "test_vpc" {
			name = "%s"
		}

		resource "ibm_is_subnet" "test_subnet" {
			name                     = "%s"
			vpc                      = ibm_is_vpc.test_vpc.id
			zone                     = "%s"
			total_ipv4_address_count = 64
		}

		resource "ibm_is_ssh_key" "test_key" {
			name       = "%s"
			public_key = "%s"
		}

		resource "ibm_is_instance" "test_instance" {
			name    = "%s"
			profile = "%s"
			catalog_offering {
				version_crn = "%s"
				plan_crn    = "%s"
			}
			vpc  = ibm_is_vpc.test_vpc.id
			zone = ibm_is_subnet.test_subnet.zone
			keys = [ibm_is_ssh_key.test_key.id]

			primary_network_attachment {
				virtual_network_interface {
					subnet = ibm_is_subnet.test_subnet.id
				}
			}
		}

		data "ibm_is_volume" "test_volume" {
			identifier = ibm_is_instance.test_instance.boot_volume[0].volume_id
		}

		data "ibm_is_volume_software_attachment" "is_volume_software_attachment" {
			volume_id                     = ibm_is_instance.test_instance.boot_volume[0].volume_id
			volume_software_attachment_id = data.ibm_is_volume.test_volume.software_attachments.0.id
		}
	`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, instanceName, acc.InstanceProfileName, acc.ISCatalogImageOfferingCRN, acc.ISCatalogImagePlanCRN)
}

func testAccCheckIBMIsVolumeSoftwareAttachmentDataSourceConfig(vpcname, subnetname, sshname, instanceName, name string) string {
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	return fmt.Sprintf(`
		resource "ibm_is_vpc" "test_vpc" {
			name = "%s"
		}

		resource "ibm_is_subnet" "test_subnet" {
			name                     = "%s"
			vpc                      = ibm_is_vpc.test_vpc.id
			zone                     = "%s"
			total_ipv4_address_count = 64
		}

		resource "ibm_is_ssh_key" "test_key" {
			name       = "%s"
			public_key = "%s"
		}

		resource "ibm_is_instance" "test_instance" {
			name    = "%s"
			profile = "%s"
			catalog_offering {
				version_crn = "%s"
				plan_crn    = "%s"
			}
			vpc  = ibm_is_vpc.test_vpc.id
			zone = ibm_is_subnet.test_subnet.zone
			keys = [ibm_is_ssh_key.test_key.id]

			primary_network_attachment {
				virtual_network_interface {
					subnet = ibm_is_subnet.test_subnet.id
				}
			}
		}

		data "ibm_is_volume" "test_volume" {
			identifier = ibm_is_instance.test_instance.boot_volume[0].volume_id
		}

		resource "ibm_is_volume_software_attachment" "is_volume_software_attachment_instance" {
			volume_id                     = ibm_is_instance.test_instance.boot_volume[0].volume_id
			volume_software_attachment_id = data.ibm_is_volume.test_volume.software_attachments.0.id
			name                          = "%s"
		}

		data "ibm_is_volume_software_attachment" "is_volume_software_attachment" {
			depends_on = [ibm_is_volume_software_attachment.is_volume_software_attachment_instance]

			volume_id                     = ibm_is_instance.test_instance.boot_volume[0].volume_id
			volume_software_attachment_id = data.ibm_is_volume.test_volume.software_attachments.0.id
		}
	`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, instanceName, acc.InstanceProfileName, acc.ISCatalogImageOfferingCRN, acc.ISCatalogImagePlanCRN, name)
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
