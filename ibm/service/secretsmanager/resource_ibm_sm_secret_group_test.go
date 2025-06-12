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
	resourceName := "ibm_sm_secret_group.sm_secret_group"
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
					testAccCheckIbmSmSecretGroupExists(resourceName, name, ""),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmSmSecretGroupConfigBasic(nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", nameUpdate),
				),
			},
		},
	})
}

func TestAccIbmSmSecretGroupAllArgs(t *testing.T) {
	resourceName := "ibm_sm_secret_group.sm_secret_group"
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
					testAccCheckIbmSmSecretGroupExists(resourceName, name, description),
					resource.TestCheckResourceAttrSet(resourceName, "secret_group_id"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", description),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmSmSecretGroupConfig(nameUpdate, descriptionUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmSecretGroupExists(resourceName, nameUpdate, descriptionUpdate),
					resource.TestCheckResourceAttr(resourceName, "name", nameUpdate),
					resource.TestCheckResourceAttr(resourceName, "description", descriptionUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      resourceName,
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
	if description == "" {
		return testAccCheckIbmSmSecretGroupConfigBasic(name)
	}
	return fmt.Sprintf(`

		resource "ibm_sm_secret_group" "sm_secret_group" {
			instance_id   = "%s"
			region        = "%s"
			name = "%s"
			description = "%s"
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, name, description)
}

func testAccCheckIbmSmSecretGroupExists(n string, expectedName, expectedDescription string) resource.TestCheckFunc {

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

		if err := verifyAttr(*secretGroup.Name, expectedName, "group name"); err != nil {
			return err
		}
		if err := verifyAttr(*secretGroup.Description, expectedDescription, "group description"); err != nil {
			return err
		}
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
