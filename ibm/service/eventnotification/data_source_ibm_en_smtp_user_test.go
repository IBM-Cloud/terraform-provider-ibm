// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventnotification_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMEnSMTPUserDataSourceBasic(t *testing.T) {
	smtpUserInstanceID := fmt.Sprintf("tf_instance_id_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_instance_id_%d", acctest.RandIntRange(10, 100))
	smtpconfigID := fmt.Sprintf("tf_instance_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMEnSMTPUserDataSourceConfigBasic(smtpUserInstanceID, description, smtpconfigID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_user.en_smtp_user_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_user.en_smtp_user_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_user.en_smtp_user_instance", "en_smtp_config_id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_user.en_smtp_user_instance", "user_id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_user.en_smtp_user_instance", "smtp_config_id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_user.en_smtp_user_instance", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_user.en_smtp_user_instance", "domain"),
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_user.en_smtp_user_instance", "username"),
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_user.en_smtp_user_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_user.en_smtp_user_instance", "updated_at"),
				),
			},
		},
	})
}

func TestAccIBMEnSMTPUserDataSourceAllArgs(t *testing.T) {
	smtpUserInstanceID := fmt.Sprintf("tf_instance_id_%d", acctest.RandIntRange(10, 100))
	smtpUserDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMEnSMTPUserDataSourceConfig(smtpUserInstanceID, smtpUserDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_user.en_smtp_user_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_user.en_smtp_user_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_user.en_smtp_user_instance", "en_smtp_config_id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_user.en_smtp_user_instance", "user_id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_user.en_smtp_user_instance", "smtp_config_id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_user.en_smtp_user_instance", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_user.en_smtp_user_instance", "domain"),
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_user.en_smtp_user_instance", "username"),
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_user.en_smtp_user_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_user.en_smtp_user_instance", "updated_at"),
				),
			},
		},
	})
}

func testAccCheckIBMEnSMTPUserDataSourceConfigBasic(smtpUserInstanceID string, description string, smtpconfigID string) string {
	return fmt.Sprintf(`
		resource "ibm_en_smtp_user" "en_smtp_user_instance" {
			instance_id = "%s"
			description   = "%s"
            smtp_config_id = "%s"
		}

		data "ibm_en_smtp_user" "en_smtp_user_instance" {
			instance_id = ibm_en_smtp_user.en_smtp_user_instance.instance_id
			en_smtp_config_id = "en_smtp_user_id"
			user_id = ibm_en_smtp_user.en_smtp_user_instance.user_id
		}
	`, smtpUserInstanceID, description, smtpconfigID)
}

func testAccCheckIBMEnSMTPUserDataSourceConfig(smtpUserInstanceID string, smtpUserDescription string) string {
	return fmt.Sprintf(`
		resource "ibm_en_smtp_user" "en_smtp_user_instance" {
			instance_id = "%s"
			description = "%s"
		}

		data "ibm_en_smtp_user" "en_smtp_user_instance" {
			instance_id = ibm_en_smtp_user.en_smtp_user_instance.instance_id
			en_smtp_user_id = "en_smtp_user_id"
			user_id = ibm_en_smtp_user.en_smtp_user_instance.user_id
		}
	`, smtpUserInstanceID, smtpUserDescription)
}
