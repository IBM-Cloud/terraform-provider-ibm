// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM-Cloud/secrets-manager-mt-go-sdk/secretsmanagerv1"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
)

func TestAccIbmSecretGroupBasic(t *testing.T) {
	var conf secretsmanagerv1.SecretGroup
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSecretGroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSecretGroupConfigBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSecretGroupExists("ibm_secret_group.secret_group", conf),
					resource.TestCheckResourceAttr("ibm_secret_group.secret_group", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmSecretGroupConfigBasic(nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_secret_group.secret_group", "name", nameUpdate),
				),
			},
		},
	})
}

func TestAccIbmSecretGroupAllArgs(t *testing.T) {
	var conf secretsmanagerv1.SecretGroup
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSecretGroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSecretGroupConfig(name, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSecretGroupExists("ibm_secret_group.secret_group", conf),
					resource.TestCheckResourceAttr("ibm_secret_group.secret_group", "name", name),
					resource.TestCheckResourceAttr("ibm_secret_group.secret_group", "description", description),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmSecretGroupConfig(nameUpdate, descriptionUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_secret_group.secret_group", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_secret_group.secret_group", "description", descriptionUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_secret_group.secret_group",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmSecretGroupConfigBasic(name string) string {
	return fmt.Sprintf(`

		resource "ibm_secret_group" "secret_group" {
			name = "%s"
		}
	`, name)
}

func testAccCheckIbmSecretGroupConfig(name string, description string) string {
	return fmt.Sprintf(`

		resource "ibm_secret_group" "secret_group" {
			name = "%s"
			description = "%s"
		}
	`, name, description)
}

func testAccCheckIbmSecretGroupExists(n string, obj secretsmanagerv1.SecretGroup) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		secretsManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerV1()
		if err != nil {
			return err
		}

		getSecretGroupOptions := &secretsmanagerv1.GetSecretGroupOptions{}

		getSecretGroupOptions.SetID(rs.Primary.ID)

		secretGroup, _, err := secretsManagerClient.GetSecretGroup(getSecretGroupOptions)
		if err != nil {
			return err
		}

		obj = *secretGroup
		return nil
	}
}

func testAccCheckIbmSecretGroupDestroy(s *terraform.State) error {
	secretsManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_secret_group" {
			continue
		}

		getSecretGroupOptions := &secretsmanagerv1.GetSecretGroupOptions{}

		getSecretGroupOptions.SetID(rs.Primary.ID)

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
