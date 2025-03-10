// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISSSHKeyDatasource_basic(t *testing.T) {
	name1 := fmt.Sprintf("tfssh-name-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testDSCheckIBMISSSHKeyConfig(publicKey, name1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.ibm_is_ssh_key.ds_key", "name", name1),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_ssh_key.ds_key", "crn"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_ssh_key.ds_key", "fingerprint"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_ssh_key.ds_key", "id"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_ssh_key.ds_key", "length"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_ssh_key.ds_key", "public_key"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_ssh_key.ds_key", "type"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_ssh_key.ds_key", "tags.#"),
					resource.TestCheckResourceAttr(
						"data.ibm_is_ssh_key.ds_key", "tags.#", "3"),
				),
			},
		},
	})
}
func TestAccIBMISSSHKeyDatasource_basicidname(t *testing.T) {
	name1 := fmt.Sprintf("tfssh-name-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testDSCheckIBMISSSHKeyIdNameConfig(publicKey, name1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.ibm_is_ssh_key.ds_key", "name", name1),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_ssh_key.ds_key", "crn"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_ssh_key.ds_key", "fingerprint"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_ssh_key.ds_key", "id"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_ssh_key.ds_key", "length"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_ssh_key.ds_key", "public_key"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_ssh_key.ds_key", "type"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_ssh_key.ds_key", "tags.#"),
					resource.TestCheckResourceAttr(
						"data.ibm_is_ssh_key.ds_key", "tags.#", "3"),
					resource.TestCheckResourceAttr(
						"data.ibm_is_ssh_key.ds_key_id", "name", name1),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_ssh_key.ds_key_id", "crn"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_ssh_key.ds_key_id", "fingerprint"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_ssh_key.ds_key_id", "id"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_ssh_key.ds_key_id", "length"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_ssh_key.ds_key_id", "public_key"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_ssh_key.ds_key_id", "type"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_ssh_key.ds_key_id", "tags.#"),
					resource.TestCheckResourceAttr(
						"data.ibm_is_ssh_key.ds_key_id", "tags.#", "3"),
				),
			},
		},
	})
}

func TestAccIBMISSSHKeyDatasource_CreatedAtHref(t *testing.T) {
	name1 := fmt.Sprintf("tfssh-name-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testDSCheckIBMISSSHKeyConfig(publicKey, name1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.ibm_is_ssh_key.ds_key", "name", name1),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_ssh_key.ds_key", "crn"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_ssh_key.ds_key", "created_at"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_ssh_key.ds_key", "fingerprint"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_ssh_key.ds_key", "href"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_ssh_key.ds_key", "id"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_ssh_key.ds_key", "length"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_ssh_key.ds_key", "public_key"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_ssh_key.ds_key", "type"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_ssh_key.ds_key", "tags.#"),
					resource.TestCheckResourceAttr(
						"data.ibm_is_ssh_key.ds_key", "tags.#", "3"),
				),
			},
		},
	})
}

func testDSCheckIBMISSSHKeyConfig(publicKey, name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ssh_key" "key" {
			name 		= "%s"
			public_key 	= "%s"
			tags		= ["test:1", "test:2", "test:3"]
		}
		data "ibm_is_ssh_key" "ds_key" {
		    name = "${ibm_is_ssh_key.key.name}"
		}`, name, publicKey)
}
func testDSCheckIBMISSSHKeyIdNameConfig(publicKey, name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ssh_key" "key" {
			name 		= "%s"
			public_key 	= "%s"
			tags		= ["test:1", "test:2", "test:3"]
		}
		data "ibm_is_ssh_key" "ds_key" {
		    name = "${ibm_is_ssh_key.key.name}"
		}
		data "ibm_is_ssh_key" "ds_key_id" {
		    id = "${ibm_is_ssh_key.key.id}"
		}
			
		`, name, publicKey)
}
