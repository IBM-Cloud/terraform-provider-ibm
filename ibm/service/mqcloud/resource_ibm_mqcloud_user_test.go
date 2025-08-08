// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package mqcloud_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/mqcloud-go-sdk/mqcloudv1"
)

func TestAccIbmMqcloudUserBasic(t *testing.T) {
	var conf mqcloudv1.UserDetails
	serviceInstanceGuid := acc.MqcloudDeploymentID
	name := fmt.Sprintf("tf-name-%d", acctest.RandIntRange(10, 100))
	email := fmt.Sprintf("tf_email_%d@ibm.com", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf-name-%d", acctest.RandIntRange(101, 200))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckMqcloud(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmMqcloudUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmMqcloudUserConfigBasic(serviceInstanceGuid, name, email),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmMqcloudUserExists("ibm_mqcloud_user.mqcloud_user_instance", conf),
					resource.TestCheckResourceAttr("ibm_mqcloud_user.mqcloud_user_instance", "service_instance_guid", serviceInstanceGuid),
					resource.TestCheckResourceAttr("ibm_mqcloud_user.mqcloud_user_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_mqcloud_user.mqcloud_user_instance", "email", email),
				),
			},
			{
				Config: testAccCheckIbmMqcloudUserConfigBasic(serviceInstanceGuid, nameUpdate, email),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_mqcloud_user.mqcloud_user_instance", "service_instance_guid", serviceInstanceGuid),
					resource.TestCheckResourceAttr("ibm_mqcloud_user.mqcloud_user_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_mqcloud_user.mqcloud_user_instance", "email", email),
				),
			},
			{
				ResourceName:      "ibm_mqcloud_user.mqcloud_user_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmMqcloudUserConfigBasic(serviceInstanceGuid string, name string, email string) string {
	return fmt.Sprintf(`
		resource "ibm_mqcloud_user" "mqcloud_user_instance" {
			service_instance_guid = "%s"
			name = "%s"
			email = "%s"
		}
	`, serviceInstanceGuid, name, email)
}

func testAccCheckIbmMqcloudUserExists(n string, obj mqcloudv1.UserDetails) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		mqcloudClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).MqcloudV1()
		if err != nil {
			return err
		}

		getUserOptions := &mqcloudv1.GetUserOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getUserOptions.SetServiceInstanceGuid(parts[0])
		getUserOptions.SetUserID(parts[1])

		userDetails, _, err := mqcloudClient.GetUser(getUserOptions)
		if err != nil {
			return err
		}

		obj = *userDetails
		return nil
	}
}

func testAccCheckIbmMqcloudUserDestroy(s *terraform.State) error {
	mqcloudClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).MqcloudV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_mqcloud_user" {
			continue
		}

		getUserOptions := &mqcloudv1.GetUserOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getUserOptions.SetServiceInstanceGuid(parts[0])
		getUserOptions.SetUserID(parts[1])

		// Try to find the key
		_, response, err := mqcloudClient.GetUser(getUserOptions)

		if err == nil {
			return fmt.Errorf("mqcloud_user still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for mqcloud_user (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
