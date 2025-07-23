// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMISImage_basic(t *testing.T) {
	var image string
	name := fmt.Sprintf("tfimg-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckImage(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkImageDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISImageConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISImageExists("ibm_is_image.isExampleImage", image),
					resource.TestCheckResourceAttr(
						"ibm_is_image.isExampleImage", "name", name),
					resource.TestCheckResourceAttrSet("ibm_is_image.isExampleImage", "user_data_format"),
				),
			},
		},
	})
}

func TestAccIBMISImage_accessTags(t *testing.T) {
	var image string
	name := fmt.Sprintf("tfimg-access-tags-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckImage(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkImageDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISImageAccessTagsConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISImageExists("ibm_is_image.isExampleImage", image),
					resource.TestCheckResourceAttr("ibm_is_image.isExampleImage", "name", name),
					resource.TestCheckResourceAttr("ibm_is_image.isExampleImage", "operating_system", acc.Image_operating_system),
					resource.TestCheckResourceAttrSet("ibm_is_image.isExampleImage", "user_data_format"),
					resource.TestCheckResourceAttrSet("ibm_is_image.isExampleImage", "status"),
					resource.TestCheckResourceAttrSet("ibm_is_image.isExampleImage", "visibility"),

					// Access tags validation
					resource.TestCheckResourceAttr("ibm_is_image.isExampleImage", "access_tags.#", "1"),
					resource.TestCheckResourceAttr("ibm_is_image.isExampleImage", "access_tags.0", "test:access"),
				),
			},
		},
	})
}
func TestAccIBMISImage_allowedUse(t *testing.T) {
	var image string
	name := fmt.Sprintf("tfimg-name-%d", acctest.RandIntRange(10, 100))
	apiVersion := "2025-07-02"
	bareMetalServer := "enable_secure_boot==true"
	instance := "enable_secure_boot==true"
	apiVersionUpdate := "2025-07-02"
	bareMetalServerUpdate := "true"
	instanceUpdate := "true"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckImage(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkImageDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISImageConfigAllowedUse(name, apiVersion, bareMetalServer, instance),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISImageExists("ibm_is_image.isExampleImage", image),
					resource.TestCheckResourceAttr(
						"ibm_is_image.isExampleImage", "name", name),
					resource.TestCheckResourceAttrSet("ibm_is_image.isExampleImage", "user_data_format"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_image.isExampleImage", "allowed_use.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_image.isExampleImage", "allowed_use.0.bare_metal_server"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_image.isExampleImage", "allowed_use.0.instance"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_image.isExampleImage", "allowed_use.0.api_version"),
					resource.TestCheckResourceAttr("ibm_is_image.isExampleImage", "allowed_use.0.bare_metal_server", bareMetalServer),
					resource.TestCheckResourceAttr("ibm_is_image.isExampleImage", "allowed_use.0.instance", instance),
					resource.TestCheckResourceAttr("ibm_is_image.isExampleImage", "allowed_use.0.api_version", apiVersion),
				),
			},
			{
				Config: testAccCheckIBMISImageConfigAllowedUse(name, apiVersionUpdate, bareMetalServerUpdate, instanceUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISImageExists("ibm_is_image.isExampleImage", image),
					resource.TestCheckResourceAttr(
						"ibm_is_image.isExampleImage", "name", name),
					resource.TestCheckResourceAttrSet("ibm_is_image.isExampleImage", "user_data_format"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_image.isExampleImage", "allowed_use.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_image.isExampleImage", "allowed_use.0.bare_metal_server"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_image.isExampleImage", "allowed_use.0.instance"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_image.isExampleImage", "allowed_use.0.api_version"),
					resource.TestCheckResourceAttr("ibm_is_image.isExampleImage", "allowed_use.0.bare_metal_server", bareMetalServerUpdate),
					resource.TestCheckResourceAttr("ibm_is_image.isExampleImage", "allowed_use.0.instance", instanceUpdate),
					resource.TestCheckResourceAttr("ibm_is_image.isExampleImage", "allowed_use.0.api_version", apiVersionUpdate),
				),
			},
		},
	})
}

func testAccCheckIBMISImageAccessTagsConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_image" "isExampleImage" {
			href             = "%s"
			name             = "%s"
			operating_system = "%s"
			access_tags      = ["test:access"]
		}
	`, acc.Image_cos_url, name, acc.Image_operating_system)
}

func TestAccIBMISImage_lifecycle(t *testing.T) {
	var image string
	name := fmt.Sprintf("tfimg-name-%d", acctest.RandIntRange(10, 100))
	deprecationAt := "2023-09-28T15:10:00.000Z"
	obsolescenceAt := "2023-11-28T15:10:00.000Z"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckImage(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkImageDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISImageLifecycleConfig(name, deprecationAt, obsolescenceAt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISImageExists("ibm_is_image.isExampleImage", image),
					resource.TestCheckResourceAttr(
						"ibm_is_image.isExampleImage", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_image.isExampleImage", "deprecation_at", deprecationAt),
					resource.TestCheckResourceAttr(
						"ibm_is_image.isExampleImage", "obsolescence_at", obsolescenceAt),
				),
			},
		},
	})
}
func TestAccIBMISImage_lifecycle_test_steps(t *testing.T) {
	var image string
	name := fmt.Sprintf("tfimg-name-%d", acctest.RandIntRange(10, 100))
	deprecationAt := "2023-09-28T15:10:00.000Z"
	obsolescenceAt := "2023-11-28T15:10:00.000Z"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckImage(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkImageDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISImageConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISImageExists("ibm_is_image.isExampleImage", image),
					resource.TestCheckResourceAttr(
						"ibm_is_image.isExampleImage", "name", name),
				),
			},
			{
				Config: testAccCheckIBMISImageLifecycleConfig(name, deprecationAt, obsolescenceAt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISImageExists("ibm_is_image.isExampleImage", image),
					resource.TestCheckResourceAttr(
						"ibm_is_image.isExampleImage", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_image.isExampleImage", "deprecation_at", deprecationAt),
					resource.TestCheckResourceAttr(
						"ibm_is_image.isExampleImage", "obsolescence_at", obsolescenceAt),
				),
			},
			{
				Config: testAccCheckIBMISImageLifecycleConfig(name, deprecationAt, "null"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISImageExists("ibm_is_image.isExampleImage", image),
					resource.TestCheckResourceAttr(
						"ibm_is_image.isExampleImage", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_image.isExampleImage", "status", "available"),
					resource.TestCheckResourceAttr(
						"ibm_is_image.isExampleImage", "deprecation_at", deprecationAt),
					resource.TestCheckResourceAttr(
						"ibm_is_image.isExampleImage", "obsolescence_at", "null"),
				),
			},
			{
				Config: testAccCheckIBMISImageLifecycleConfig(name, "null", "null"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISImageExists("ibm_is_image.isExampleImage", image),
					resource.TestCheckResourceAttr(
						"ibm_is_image.isExampleImage", "status", "available"),
					resource.TestCheckResourceAttr(
						"ibm_is_image.isExampleImage", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_image.isExampleImage", "deprecation_at", "null"),
					resource.TestCheckResourceAttr(
						"ibm_is_image.isExampleImage", "obsolescence_at", "null"),
				),
			},
		},
	})
}
func TestAccIBMISImage_lifecycle_test_deprecate_obsolete(t *testing.T) {
	var image string
	name := fmt.Sprintf("tfimg-name-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckImage(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkImageDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISImageConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISImageExists("ibm_is_image.isExampleImage", image),
					resource.TestCheckResourceAttr(
						"ibm_is_image.isExampleImage", "name", name),
				),
			},
			{
				Config: testAccCheckIBMISImageLifecycleDeprecateConfig(name, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISImageExists("ibm_is_image.isExampleImage", image),
					resource.TestCheckResourceAttr(
						"ibm_is_image.isExampleImage", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_image.isExampleImage", "status", "deprecated"),
				),
			},
			{
				Config: testAccCheckIBMISImageLifecycleObsoleteConfig(name, true, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISImageExists("ibm_is_image.isExampleImage", image),
					resource.TestCheckResourceAttr(
						"ibm_is_image.isExampleImage", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_image.isExampleImage", "status", "obsolete"),
				),
			},
		},
	})
}
func TestAccIBMISImage_error(t *testing.T) {
	name := fmt.Sprintf("tfimg-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckImage(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkImageDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMISImageError(name),
				ExpectError: regexp.MustCompile("is not attached to a virtual server instance"),
			},
			{
				Config:      testAccCheckIBMISImageError2(name),
				ExpectError: regexp.MustCompile("is not boot volume"),
			},
		},
	})
}

func TestAccIBMISImage_fromVolume(t *testing.T) {
	var image string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tfimg-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckImage(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkImageDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISImageConfig1(vpcname, subnetname, sshname, publicKey, instanceName, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISImageExists("ibm_is_image.isExampleImageFromVolume", image),
					resource.TestCheckResourceAttr(
						"ibm_is_image.isExampleImageFromVolume", "name", name),
				),
			},
		},
	})
}

func TestAccIBMISImage_encrypted(t *testing.T) {
	var image string
	name := fmt.Sprintf("tfimg-enc-name-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckEncryptedImage(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkImageDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISImageEncryptedConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISImageExists("ibm_is_image.isExampleImageEncrypted", image),
					resource.TestCheckResourceAttr(
						"ibm_is_image.isExampleImageEncrypted", "name", name),
				),
			},
		},
	})
}
func checkImageDestroy(s *terraform.State) error {

	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_image" {
			continue
		}

		getimgoptions := &vpcv1.GetImageOptions{
			ID: &rs.Primary.ID,
		}
		_, _, err := sess.GetImage(getimgoptions)
		if err == nil {
			return fmt.Errorf("Image still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISImageExists(n, image string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		getimgoptions := &vpcv1.GetImageOptions{
			ID: &rs.Primary.ID,
		}
		foundImage, _, err := sess.GetImage(getimgoptions)
		if err != nil {
			return err
		}
		image = *foundImage.ID

		return nil
	}
}

func testAccCheckIBMISImageConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_image" "isExampleImage" {
			href = "%s"
			name = "%s"
			operating_system = "%s"
		}
	`, acc.Image_cos_url, name, acc.Image_operating_system)
}

func testAccCheckIBMISImageConfigAllowedUse(name, apiVersion, bareMetalServer, instance string) string {
	return fmt.Sprintf(`
		resource "ibm_is_image" "isExampleImage" {
			href = "%s"
			name = "%s"
			operating_system = "%s"
			allowed_use {
   				api_version       = "%s"
    			bare_metal_server = "%s"
    			instance          = "%s"
  			}
		}
	`, acc.Image_cos_url, name, acc.Image_operating_system, apiVersion, bareMetalServer, instance)
}

func testAccCheckIBMISImageLifecycleConfig(name, deprecationAt, obsolescenceAt string) string {
	return fmt.Sprintf(`
		resource "ibm_is_image" "isExampleImage" {
			href = "%s"
			name = "%s"
			operating_system = "%s"
			deprecation_at = "%s"
			obsolescence_at = "%s"
		}
	`, acc.Image_cos_url, name, acc.Image_operating_system, deprecationAt, obsolescenceAt)
}
func testAccCheckIBMISImageLifecycleDeprecateConfig(name string, deprecate bool) string {
	return fmt.Sprintf(`
		resource "ibm_is_image" "isExampleImage" {
			href = "%s"
			name = "%s"
			operating_system = "%s"
			deprecate = %t
		}
	`, acc.Image_cos_url, name, acc.Image_operating_system, deprecate)
}
func testAccCheckIBMISImageLifecycleObsoleteConfig(name string, deprecate, obsolete bool) string {
	return fmt.Sprintf(`
		resource "ibm_is_image" "isExampleImage" {
			href = "%s"
			name = "%s"
			operating_system = "%s"
			deprecate = %t
			obsolete = %t
		}
	`, acc.Image_cos_url, name, acc.Image_operating_system, deprecate, obsolete)
}
func testAccCheckIBMISImageConfig1(vpcname, subnetname, sshname, publicKey, instanceName, name string) string {
	return fmt.Sprintf(`
		  resource "ibm_is_vpc" "testacc_vpc" {
			name = "%s"
		  }
		  
		  resource "ibm_is_subnet" "testacc_subnet" {
			name            		 = "%s"
			vpc             		 = ibm_is_vpc.testacc_vpc.id
			zone            		 = "%s"
			total_ipv4_address_count = 32
		  }
		  
		  resource "ibm_is_ssh_key" "testacc_sshkey" {
			name       = "%s"
			public_key = "%s"
		  }
		  
		  resource "ibm_is_instance" "testacc_instance" {
			name    = "%s"
			image   = "%s"
			profile = "%s"
			primary_network_interface {
			  subnet     = ibm_is_subnet.testacc_subnet.id
			}
			vpc  = ibm_is_vpc.testacc_vpc.id
			zone = "%s"
			keys = [ibm_is_ssh_key.testacc_sshkey.id]
			network_interfaces {
			  subnet = ibm_is_subnet.testacc_subnet.id
			  name   = "eth1"
			}
		  }
		  
		  resource "ibm_is_image" "isExampleImageFromVolume" {
			name = "%s"
			source_volume = ibm_is_instance.testacc_instance.volume_attachments.0.volume_id
			timeouts {
				create = "45m"
			}
		  }`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, instanceName, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName, name)
}
func testAccCheckIBMISImageError(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_image" "isExampleImageFromVolume" {
			name = "%s"
			source_volume = "%s"
			timeouts {
				create = "45m"
			}
		}`, name, acc.VSIUnattachedBootVolumeID)
}
func testAccCheckIBMISImageError2(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_image" "isExampleImageFromVolume" {
			name = "%s"
			source_volume = "%s"
			timeouts {
				create = "45m"
			}
		}`, name, acc.VSIDataVolumeID)
}

func testAccCheckIBMISImageEncryptedConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_image" "isExampleImageEncrypted" {
			encrypted_data_key = "%s"
  			encryption_key = "%s"
			href = "%s"
			name = "%s"
			operating_system = "%s"
		}
		`, acc.IsImageEncryptedDataKey, acc.IsImageEncryptionKey, acc.Image_cos_url_encrypted, name, acc.Image_operating_system)
}
