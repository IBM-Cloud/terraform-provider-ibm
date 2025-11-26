// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMIsSshKeysDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsSshKeysDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_ssh_keys.is_ssh_keys", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ssh_keys.is_ssh_keys", "keys.#"),
				),
			},
		},
	})
}
func TestAccIBMIsSshKeysDataSourceComprehensive(t *testing.T) {
	name1 := fmt.Sprintf("tfssh-name-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsSshKeysDataSourceConfigComprehensive(publicKey, name1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_ssh_key.key", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_ssh_key.key", "tags.#", "3"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ssh_keys.is_ssh_keys", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ssh_keys.is_ssh_keys", "keys.#"),
					// Check all fields in the first key in the list
					resource.TestCheckResourceAttrSet("data.ibm_is_ssh_keys.is_ssh_keys", "keys.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ssh_keys.is_ssh_keys", "keys.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ssh_keys.is_ssh_keys", "keys.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ssh_keys.is_ssh_keys", "keys.0.fingerprint"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ssh_keys.is_ssh_keys", "keys.0.length"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ssh_keys.is_ssh_keys", "keys.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ssh_keys.is_ssh_keys", "keys.0.public_key"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ssh_keys.is_ssh_keys", "keys.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ssh_keys.is_ssh_keys", "keys.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ssh_keys.is_ssh_keys", "keys.0.resource_group.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ssh_keys.is_ssh_keys", "keys.0.resource_group.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ssh_keys.is_ssh_keys", "keys.0.resource_group.0.href"),
				),
			},
		},
	})
}
func testAccCheckIBMIsSshKeysDataSourceConfigComprehensive(publicKey, name1 string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ssh_key" "key" {
			name 		= "%s"
			public_key 	= "%s"
			tags		= ["test:1", "test:2", "test:3"]
		}
		data "ibm_is_ssh_keys" "is_ssh_keys" {
			depends_on = [ibm_is_ssh_key.key]
		}
	`, name1, publicKey)
}
func TestAccIBMIsSshKeysDataSourceTags(t *testing.T) {
	name1 := fmt.Sprintf("tfssh-name-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsSshKeysDataSourceConfigTags(publicKey, name1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_ssh_key.key", "name", name1),
					resource.TestCheckResourceAttrSet(
						"ibm_is_ssh_key.key", "crn"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_ssh_key.key", "fingerprint"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_ssh_key.key", "id"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_ssh_key.key", "length"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ssh_keys.is_ssh_keys", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ssh_keys.is_ssh_keys", "keys.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ssh_keys.is_ssh_keys", "keys.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ssh_keys.is_ssh_keys", "keys.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ssh_keys.is_ssh_keys", "keys.0.fingerprint"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ssh_keys.is_ssh_keys", "keys.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ssh_keys.is_ssh_keys", "keys.0.tags.#"),
					resource.TestCheckResourceAttr("data.ibm_is_ssh_keys.is_ssh_keys", "keys.0.tags.#", "3"),
				),
			},
		},
	})
}

func testAccCheckIBMIsSshKeysDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_is_ssh_keys" "is_ssh_keys" {
		}
	`)
}
func testAccCheckIBMIsSshKeysDataSourceConfigTags(publicKey, name1 string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ssh_key" "key" {
			name 		= "%s"
			public_key 	= "%s"
			tags		= ["test:1", "test:2", "test:3"]
		}
		data "ibm_is_ssh_keys" "is_ssh_keys" {
			depends_on = [ ibm_is_ssh_key.key ]
		}
	`, name1, publicKey)
}
