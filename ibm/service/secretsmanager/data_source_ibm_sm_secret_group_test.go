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

func TestAccIbmSmSecretGroupDataSourceBasic(t *testing.T) {
	secretGroupName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmSecretGroupDataSourceConfigBasic(secretGroupName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_sm_secret_group.sm_secret_group", "secret_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_secret_group.sm_secret_group", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_secret_group.sm_secret_group", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_secret_group.sm_secret_group", "updated_at"),
				),
			},
		},
	})
}

func TestAccIbmSmSecretGroupDataSourceAllArgs(t *testing.T) {
	secretGroupName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	secretGroupDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmSecretGroupDataSourceConfig(secretGroupName, secretGroupDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_sm_secret_group.sm_secret_group", "secret_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_secret_group.sm_secret_group", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_secret_group.sm_secret_group", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_secret_group.sm_secret_group", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_secret_group.sm_secret_group", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_secret_group.sm_secret_group", "updated_at"),
				),
			},
		},
	})
}

func testAccCheckIbmSmSecretGroupDataSourceConfigBasic(secretGroupName string) string {
	return fmt.Sprintf(`
		resource "ibm_sm_secret_group" "sm_secret_group_instance" {
			instance_id   = "%s"
			region        = "%s"
			name = "%s"
		}

		data "ibm_sm_secret_group" "sm_secret_group" {
			instance_id   = "%s"
			region        = "%s"
			secret_group_id = ibm_sm_secret_group.sm_secret_group_instance.secret_group_id
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, secretGroupName, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}

func testAccCheckIbmSmSecretGroupDataSourceConfig(secretGroupName string, secretGroupDescription string) string {
	return fmt.Sprintf(`
		resource "ibm_sm_secret_group" "sm_secret_group_instance" {
			instance_id   = "%s"
			region        = "%s"
			name = "%s"
			description = "%s"
		}

		data "ibm_sm_secret_group" "sm_secret_group" {
			instance_id   = "%s"
			region        = "%s"
			secret_group_id = ibm_sm_secret_group.sm_secret_group_instance.secret_group_id
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, secretGroupName, secretGroupDescription, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}
