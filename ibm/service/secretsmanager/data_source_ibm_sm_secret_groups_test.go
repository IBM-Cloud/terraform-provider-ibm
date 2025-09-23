// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSmSecretGroupsDataSourceBasic(t *testing.T) {
	secretGroupName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmSecretGroupsDataSourceConfigBasic(secretGroupName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_sm_secret_groups.sm_secret_groups", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_secret_groups.sm_secret_groups", "secret_groups.#"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_secret_groups.sm_secret_groups", "total_count"),
				),
			},
		},
	})
}

func TestAccIbmSmSecretGroupsDataSourceAllArgs(t *testing.T) {
	secretGroupName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	secretGroupDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmSecretGroupsDataSourceConfig(secretGroupName, secretGroupDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_sm_secret_groups.sm_secret_groups", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_secret_groups.sm_secret_groups", "secret_groups.#"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_secret_groups.sm_secret_groups", "secret_groups.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_secret_groups.sm_secret_groups", "total_count"),
				),
			},
		},
	})
}

func testAccCheckIbmSmSecretGroupsDataSourceConfigBasic(secretGroupName string) string {
	return fmt.Sprintf(`
		resource "ibm_sm_secret_group" "sm_secret_group_instance_1" {
			instance_id   = "%s"
			region        = "%s"
			name = "%s_1"
		}
		
		resource "ibm_sm_secret_group" "sm_secret_group_instance_2" {
			instance_id   = "%s"
			region        = "%s"
			name = "%s_2"
		}

		data "ibm_sm_secret_groups" "sm_secret_groups" {
			depends_on = [
				ibm_sm_secret_group.sm_secret_group_instance_1,
				ibm_sm_secret_group.sm_secret_group_instance_2
			]
			instance_id   = "%s"
			region        = "%s"
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, secretGroupName, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, secretGroupName, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}

func testAccCheckIbmSmSecretGroupsDataSourceConfig(secretGroupName string, secretGroupDescription string) string {
	return fmt.Sprintf(`
		resource "ibm_sm_secret_group" "sm_secret_group_instance_1" {
			instance_id   = "%s"
			region        = "%s"
			name = "%s_1"
			description = "%s"
		}
		resource "ibm_sm_secret_group" "sm_secret_group_instance_2" {
			instance_id   = "%s"
			region        = "%s"			
			name = "%s_2"
			description = "%s"
		}

		data "ibm_sm_secret_groups" "sm_secret_groups" {
			depends_on = [
				ibm_sm_secret_group.sm_secret_group_instance_1,
				ibm_sm_secret_group.sm_secret_group_instance_2
			]
			instance_id   = "%s"
			region        = "%s"
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, secretGroupName, secretGroupDescription, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, secretGroupName, secretGroupDescription, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}
