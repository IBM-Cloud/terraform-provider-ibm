// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/secrets-manager-go-sdk/v2/secretsmanagerv2"
)

func TestAccIbmSmSecretGroupBasic(t *testing.T) {
	var conf secretsmanagerv2.SecretGroup
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmSecretGroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmSecretGroupConfigBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmSecretGroupExists("ibm_sm_secret_group.sm_secret_group", conf),
					resource.TestCheckResourceAttr("ibm_sm_secret_group.sm_secret_group", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmSmSecretGroupConfigBasic(nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_sm_secret_group.sm_secret_group", "name", nameUpdate),
				),
			},
		},
	})
}

func TestAccIbmSmSecretGroupAllArgs(t *testing.T) {
	var conf secretsmanagerv2.SecretGroup
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmSecretGroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmSecretGroupConfig(name, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmSecretGroupExists("ibm_sm_secret_group.sm_secret_group", conf),
					resource.TestCheckResourceAttr("ibm_sm_secret_group.sm_secret_group", "name", name),
					resource.TestCheckResourceAttr("ibm_sm_secret_group.sm_secret_group", "description", description),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmSmSecretGroupConfig(nameUpdate, descriptionUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_sm_secret_group.sm_secret_group", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_sm_secret_group.sm_secret_group", "description", descriptionUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_sm_secret_group.sm_secret_group",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmSmSecretGroupConfigBasic(name string) string {
	return fmt.Sprintf(`

		resource "ibm_sm_secret_group" "sm_secret_group" {
			instance_id   = "%s"
			region        = "%s"
			name = "%s"
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, name)
}

func testAccCheckIbmSmSecretGroupConfig(name string, description string) string {
	return fmt.Sprintf(`

		resource "ibm_sm_secret_group" "sm_secret_group" {
			instance_id   = "%s"
			region        = "%s"
			name = "%s"
			description = "%s"
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, name, description)
}

func testAccCheckIbmSmSecretGroupExists(n string, obj secretsmanagerv2.SecretGroup) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		secretsManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerV2()
		if err != nil {
			return err
		}

		secretsManagerClient = getClientWithInstanceEndpointTest(secretsManagerClient)

		getSecretGroupOptions := &secretsmanagerv2.GetSecretGroupOptions{}

		id := strings.Split(rs.Primary.ID, "/")
		secretGroupId := id[2]
		getSecretGroupOptions.SetID(secretGroupId)

		secretGroup, _, err := secretsManagerClient.GetSecretGroup(getSecretGroupOptions)
		if err != nil {
			return err
		}

		obj = *secretGroup
		return nil
	}
}

func testAccCheckIbmSmSecretGroupDestroy(s *terraform.State) error {
	secretsManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerV2()
	if err != nil {
		return err
	}

	secretsManagerClient = getClientWithInstanceEndpointTest(secretsManagerClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_sm_secret_group" {
			continue
		}

		getSecretGroupOptions := &secretsmanagerv2.GetSecretGroupOptions{}

		id := strings.Split(rs.Primary.ID, "/")
		secretGroupId := id[2]
		getSecretGroupOptions.SetID(secretGroupId)

		// Try to find the key
		_, response, err := secretsManagerClient.GetSecretGroup(getSecretGroupOptions)

		if err == nil {
			return fmt.Errorf("SecretGroup still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for SecretGroup (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
