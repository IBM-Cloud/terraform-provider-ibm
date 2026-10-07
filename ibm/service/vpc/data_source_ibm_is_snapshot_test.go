// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/vpc"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMISSnapshotDatasource_basic(t *testing.T) {
	var snapshot string
	snpName := "data.ibm_is_snapshot.ds_snapshot"
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	volname := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfsnapshotuat-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISSnapshotDestroy,
		Steps: []resource.TestStep{
			{
				Config: testDSCheckIBMISSnapshotConfig(vpcname, subnetname, sshname, publicKey, volname, name, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSnapshotExists("ibm_is_snapshot.testacc_snapshot", snapshot),
					resource.TestCheckResourceAttr(
						"ibm_is_snapshot.testacc_snapshot", "name", name1),
					resource.TestCheckResourceAttrSet(snpName, "href"),
					resource.TestCheckResourceAttrSet(snpName, "crn"),
					resource.TestCheckResourceAttrSet(snpName, "lifecycle_state"),
					resource.TestCheckResourceAttrSet(snpName, "encryption"),
					resource.TestCheckResourceAttrSet(snpName, "allowed_use.#"),
					resource.TestCheckResourceAttrSet(snpName, "allowed_use.0.bare_metal_server"),
					resource.TestCheckResourceAttrSet(snpName, "allowed_use.0.instance"),
					resource.TestCheckResourceAttrSet(snpName, "allowed_use.0.api_version"),
					// resource.TestCheckResourceAttrSet(snpName, "captured_at"), // Commented as the attribute is optional.
					resource.TestCheckResourceAttr(snpName, "software_attachments.#", "0"),
				),
			},
		},
	})
}
func TestAccIBMISSnapshotDatasource_404(t *testing.T) {
	snpId := "8843-5fr454ft-f6-4565-9555-5f889f5f3f7777"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISSnapshotDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testDSCheckIBMISSnapshotConfig404(snpId),
				ExpectError: regexp.MustCompile("Error fetching snapshot with id"),
			},
		},
	})
}
func TestAccIBMISSnapshotDatasource_serviceTags(t *testing.T) {
	var snapshot string
	snpName := "data.ibm_is_snapshot.ds_snapshot"
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	volname := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfsnapshotuat-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISSnapshotDestroy,
		Steps: []resource.TestStep{
			{
				Config: testDSCheckIBMISSnapshotConfig(vpcname, subnetname, sshname, publicKey, volname, name, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSnapshotExists("ibm_is_snapshot.testacc_snapshot", snapshot),
					resource.TestCheckResourceAttr(
						"ibm_is_snapshot.testacc_snapshot", "name", name1),
					resource.TestCheckResourceAttrSet(snpName, "href"),
					resource.TestCheckResourceAttrSet(snpName, "crn"),
					resource.TestCheckResourceAttrSet(snpName, "lifecycle_state"),
					resource.TestCheckResourceAttrSet(snpName, "encryption"),
					resource.TestCheckResourceAttrSet(snpName, "service_tags.#"),
					// resource.TestCheckResourceAttrSet(snpName, "captured_at"), // Commented as the attribute is optional.
				),
			},
		},
	})
}
func TestAccIBMISSnapshotDatasource_clonesbasic(t *testing.T) {
	var snapshot string
	snpName := "data.ibm_is_snapshot.ds_snapshot"
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	volname := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfsnapshotuat-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISSnapshotDestroy,
		Steps: []resource.TestStep{
			{
				Config: testDSCheckIBMISSnapshotClonesBasicConfig(vpcname, subnetname, sshname, publicKey, volname, name, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSnapshotExists("ibm_is_snapshot.testacc_snapshot", snapshot),
					resource.TestCheckResourceAttr(
						"ibm_is_snapshot.testacc_snapshot", "name", name1),
					resource.TestCheckResourceAttrSet(snpName, "href"),
					resource.TestCheckResourceAttrSet(snpName, "crn"),
					resource.TestCheckResourceAttrSet(snpName, "lifecycle_state"),
					resource.TestCheckResourceAttrSet(snpName, "encryption"),
					resource.TestCheckResourceAttrSet(snpName, "clones.#"),
					resource.TestCheckResourceAttr(snpName, "clones.#", "2"),
					resource.TestCheckResourceAttr(snpName, "clones.0", acc.ISZoneName),
					resource.TestCheckResourceAttr(snpName, "clones.1", acc.ISZoneName2),
					// resource.TestCheckResourceAttrSet(snpName, "captured_at"), // Commented as the attribute is optional.
				),
			},
		},
	})
}

func TestAccIBMISSnapshotDatasourceWithCatalogOffering(t *testing.T) {
	var snapshot string
	snpName := "data.ibm_is_snapshot.ds_snapshot"
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfsnapshotuat-%d", acctest.RandIntRange(10, 100))
	planCrn := acc.ISCatalogImagePlanCRN
	versionCrn := acc.ISCatalogImageOfferingCRN
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISSnapshotDestroy,
		Steps: []resource.TestStep{
			{
				Config: testDSCheckIBMISSnapshotConfigWithCatalogOffering(vpcname, subnetname, sshname, publicKey, name, name1, planCrn, versionCrn),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSnapshotExists("ibm_is_snapshot.testacc_snapshot", snapshot),
					resource.TestCheckResourceAttr(
						"ibm_is_snapshot.testacc_snapshot", "name", name1),
					resource.TestCheckResourceAttrSet(snpName, "catalog_offering.#"),
					resource.TestCheckResourceAttrSet(snpName, "catalog_offering.0.version_crn"),
					resource.TestCheckResourceAttrSet(snpName, "catalog_offering.0.plan_crn"),
					// lookup by name
					resource.TestCheckResourceAttrSet(snpName, "software_attachments.#"),
					resource.TestCheckResourceAttrSet(snpName, "software_attachments.0.id"),
					resource.TestCheckResourceAttrSet(snpName, "software_attachments.0.href"),
					resource.TestCheckResourceAttrSet(snpName, "software_attachments.0.name"),
					resource.TestCheckResourceAttr(snpName, "software_attachments.0.resource_type", "snapshot_software_attachment"),
					// lookup by identifier must return the same software attachments
					resource.TestCheckResourceAttrPair("data.ibm_is_snapshot.ds_snapshot_id", "software_attachments.#", snpName, "software_attachments.#"),
					resource.TestCheckResourceAttrPair("data.ibm_is_snapshot.ds_snapshot_id", "software_attachments.0.id", snpName, "software_attachments.0.id"),
				),
			},
		},
	})
}

func testDSCheckIBMISSnapshotConfigWithCatalogOffering(vpcname, subnetname, sshname, publicKey, name, sname, planCrn, versionCrn string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name           				= "%s"
		vpc             			= ibm_is_vpc.testacc_vpc.id
		zone            			= "%s"
		total_ipv4_address_count 	= 16
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  } 
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		boot_volume {
		  auto_delete_volume = false  
		}
		catalog_offering {
		  version_crn = "%s"
		  plan_crn    = "%s"
		}
	  }
	resource "ibm_is_snapshot" "testacc_snapshot" {
		name 			= "%s"
		source_volume 	= ibm_is_instance.testacc_instance.boot_volume.0.volume_id
		}
	data "ibm_is_snapshot" "ds_snapshot" {
		depends_on 	= [ibm_is_snapshot.testacc_snapshot]
		name 		= "%s"
	}
	data "ibm_is_snapshot" "ds_snapshot_id" {
		identifier 	= ibm_is_snapshot.testacc_snapshot.id
	}
`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, name, acc.InstanceProfileName, acc.ISZoneName, versionCrn, planCrn, sname, sname)
}

func testDSCheckIBMISSnapshotConfig(vpcname, subnetname, sshname, publicKey, volname, name, sname string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name           					= "%s"
		vpc             				= ibm_is_vpc.testacc_vpc.id
		zone            				= "%s"
		total_ipv4_address_count 		= 16
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  } 
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    	= "%s"
		image   	= "%s"
		profile 	= "%s"
		primary_network_interface {
		  subnet    = ibm_is_subnet.testacc_subnet.id
		}
		vpc  		= ibm_is_vpc.testacc_vpc.id
		zone 		= "%s"
		keys 		= [ibm_is_ssh_key.testacc_sshkey.id]
		network_interfaces {
		  subnet 	= ibm_is_subnet.testacc_subnet.id
		  name   	= "eth1"
		}
	  }
	resource "ibm_is_snapshot" "testacc_snapshot" {
		name 			= "%s"
		source_volume 	= ibm_is_instance.testacc_instance.volume_attachments[0].volume_id
		}
	data "ibm_is_snapshot" "ds_snapshot" {
		depends_on 	= [ibm_is_snapshot.testacc_snapshot]
		name 		= "%s"
	}
`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName, sname, sname)
}
func testDSCheckIBMISSnapshotConfig404(sid string) string {
	return fmt.Sprintf(`
	data "ibm_is_snapshot" "ds_snapshot" {
		identifier 		= "%s"
	}
`, sid)
}
func testDSCheckIBMISSnapshotClonesBasicConfig(vpcname, subnetname, sshname, publicKey, volname, name, sname string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name           					= "%s"
		vpc             				= ibm_is_vpc.testacc_vpc.id
		zone            				= "%s"
		total_ipv4_address_count 		= 16
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  } 
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    	= "%s"
		image   	= "%s"
		profile 	= "%s"
		primary_network_interface {
		  subnet    = ibm_is_subnet.testacc_subnet.id
		}
		vpc  		= ibm_is_vpc.testacc_vpc.id
		zone 		= "%s"
		keys 		= [ibm_is_ssh_key.testacc_sshkey.id]
		network_interfaces {
		  subnet 	= ibm_is_subnet.testacc_subnet.id
		  name   	= "eth1"
		}
	  }
	resource "ibm_is_snapshot" "testacc_snapshot" {
		name 			= "%s"
		source_volume 	= ibm_is_instance.testacc_instance.volume_attachments[0].volume_id
		clones			= ["%s", "%s"]
		}
	data "ibm_is_snapshot" "ds_snapshot" {
		depends_on 	= [ibm_is_snapshot.testacc_snapshot]
		name 		= "%s"
	}
`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName, sname, acc.ISZoneName, acc.ISZoneName2, sname)
}

func TestDataSourceIBMIsSnapshotSnapshotSoftwareAttachmentReferenceToMap(t *testing.T) {
	href := "https://us-south.iaas.cloud.ibm.com/v1/snapshots/r006-7ec86020-1c6e-4889-b3f0-a15f2e50f87e/software_attachments/r006-a569e8ae-3254-495e-ae75-86bb08e2c4d1"

	// Active reference: "deleted" must be left out.
	model := new(vpcv1.SnapshotSoftwareAttachmentReference)
	model.Href = core.StringPtr(href)
	model.ID = core.StringPtr("r006-a569e8ae-3254-495e-ae75-86bb08e2c4d1")
	model.Name = core.StringPtr("my-software-attachment")
	model.ResourceType = core.StringPtr("snapshot_software_attachment")

	result, err := vpc.DataSourceIBMIsSnapshotSnapshotSoftwareAttachmentReferenceToMap(model)
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{
		"href":          href,
		"id":            "r006-a569e8ae-3254-495e-ae75-86bb08e2c4d1",
		"name":          "my-software-attachment",
		"resource_type": "snapshot_software_attachment",
	}, result)

	// Deleted reference: "deleted" is a single element list with more_info.
	deletedModel := new(vpcv1.Deleted)
	deletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")
	model.Deleted = deletedModel

	result, err = vpc.DataSourceIBMIsSnapshotSnapshotSoftwareAttachmentReferenceToMap(model)
	assert.Nil(t, err)
	assert.Equal(t, []map[string]interface{}{{"more_info": "https://cloud.ibm.com/apidocs/vpc#deleted-resources"}}, result["deleted"])
}
