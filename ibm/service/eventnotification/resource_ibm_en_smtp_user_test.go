// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventnotification_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/event-notifications-go-admin-sdk/eventnotificationsv1"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
)

func TestAccIBMEnSMTPUserBasic(t *testing.T) {
	var conf eventnotificationsv1.SMTPUser
	instanceID := fmt.Sprintf("tf_instance_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMEnSMTPUserDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMEnSMTPUserConfigBasic(instanceID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMEnSMTPUserExists("ibm_en_smtp_user.en_smtp_user_instance", conf),
					resource.TestCheckResourceAttr("ibm_en_smtp_user.en_smtp_user_instance", "instance_id", instanceID),
				),
			},
		},
	})
}

func TestAccIBMEnSMTPUserAllArgs(t *testing.T) {
	var conf eventnotificationsv1.SMTPUser
	instanceID := fmt.Sprintf("tf_instance_id_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMEnSMTPUserDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMEnSMTPUserConfig(instanceID, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMEnSMTPUserExists("ibm_en_smtp_user.en_smtp_user_instance", conf),
					resource.TestCheckResourceAttr("ibm_en_smtp_user.en_smtp_user_instance", "instance_id", instanceID),
					resource.TestCheckResourceAttr("ibm_en_smtp_user.en_smtp_user_instance", "description", description),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMEnSMTPUserConfig(instanceID, descriptionUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_en_smtp_user.en_smtp_user_instance", "instance_id", instanceID),
					resource.TestCheckResourceAttr("ibm_en_smtp_user.en_smtp_user_instance", "description", descriptionUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_en_smtp_user.en_smtp_user",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMEnSMTPUserConfigBasic(instanceID string) string {
	return fmt.Sprintf(`
		resource "ibm_en_smtp_user" "en_smtp_user_instance" {
			instance_id = "%s"
		}
	`, instanceID)
}

func testAccCheckIBMEnSMTPUserConfig(instanceID string, description string) string {
	return fmt.Sprintf(`

		resource "ibm_en_smtp_user" "en_smtp_user_instance" {
			instance_id = "%s"
			description = "%s"
		}
	`, instanceID, description)
}

func testAccCheckIBMEnSMTPUserExists(n string, obj eventnotificationsv1.SMTPUser) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		eventNotificationsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).EventNotificationsApiV1()
		if err != nil {
			return err
		}

		getSMTPUserOptions := &eventnotificationsv1.GetSMTPUserOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getSMTPUserOptions.SetInstanceID(parts[0])
		getSMTPUserOptions.SetID(parts[1])
		getSMTPUserOptions.SetUserID(parts[2])

		smtpUser, _, err := eventNotificationsClient.GetSMTPUser(getSMTPUserOptions)
		if err != nil {
			return err
		}

		obj = *smtpUser
		return nil
	}
}

func testAccCheckIBMEnSMTPUserDestroy(s *terraform.State) error {
	eventNotificationsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).EventNotificationsApiV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_en_smtp_user" {
			continue
		}

		getSMTPUserOptions := &eventnotificationsv1.GetSMTPUserOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getSMTPUserOptions.SetInstanceID(parts[0])
		getSMTPUserOptions.SetID(parts[1])
		getSMTPUserOptions.SetUserID(parts[2])

		// Try to find the key
		_, response, err := eventNotificationsClient.GetSMTPUser(getSMTPUserOptions)

		if err == nil {
			return fmt.Errorf("en_smtp_user still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for en_smtp_user (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
