// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package classicinfrastructure_test

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/services"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
)

func TestAccIBMComputeSSHKey_basic(t *testing.T) {
	var key datatypes.Security_Ssh_Key

	label1 := fmt.Sprintf("terraformsshuat_create_step_label_%d", acctest.RandIntRange(10, 100))
	label2 := fmt.Sprintf("terraformsshuat_update_step_label_%d", acctest.RandIntRange(10, 100))
	notes1 := fmt.Sprintf("terraformsshuat_create_step_notes_%d", acctest.RandIntRange(10, 100))
	notes2 := fmt.Sprintf("terraformsshuat_update_step_notes_%d", acctest.RandIntRange(10, 100))

	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMComputeSSHKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMComputeSSHKeyConfig(label1, notes1, publicKey),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeSSHKeyExists("ibm_compute_ssh_key.testacc_ssh_key", &key),
					resource.TestCheckResourceAttr(
						"ibm_compute_ssh_key.testacc_ssh_key", "label", label1),
					resource.TestCheckResourceAttr(
						"ibm_compute_ssh_key.testacc_ssh_key", "public_key", publicKey),
					resource.TestCheckResourceAttr(
						"ibm_compute_ssh_key.testacc_ssh_key", "notes", notes1),
				),
			},

			{
				Config: testAccCheckIBMComputeSSHKeyConfig(label2, notes2, publicKey),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeSSHKeyExists("ibm_compute_ssh_key.testacc_ssh_key", &key),
					resource.TestCheckResourceAttr(
						"ibm_compute_ssh_key.testacc_ssh_key", "label", label2),
					resource.TestCheckResourceAttr(
						"ibm_compute_ssh_key.testacc_ssh_key", "public_key", publicKey),
					resource.TestCheckResourceAttr(
						"ibm_compute_ssh_key.testacc_ssh_key", "notes", notes2),
				),
			},
		},
	})
}

func TestAccIBMComputeSSHKeyWithTag(t *testing.T) {
	var key datatypes.Security_Ssh_Key

	label1 := fmt.Sprintf("terraformsshuat_create_step_label_%d", acctest.RandIntRange(10, 100))
	notes1 := fmt.Sprintf("terraformsshuat_create_step_notes_%d", acctest.RandIntRange(10, 100))

	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMComputeSSHKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMComputeSSHKeyWithTag(label1, notes1, publicKey),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeSSHKeyExists("ibm_compute_ssh_key.testacc_ssh_key", &key),
					resource.TestCheckResourceAttr(
						"ibm_compute_ssh_key.testacc_ssh_key", "label", label1),
					resource.TestCheckResourceAttr(
						"ibm_compute_ssh_key.testacc_ssh_key", "public_key", publicKey),
					resource.TestCheckResourceAttr(
						"ibm_compute_ssh_key.testacc_ssh_key", "notes", notes1),
					resource.TestCheckResourceAttr(
						"ibm_compute_ssh_key.testacc_ssh_key", "tags.#", "2"),
				),
			},

			{
				Config: testAccCheckIBMComputeSSHKeyWithUpdatedTag(label1, notes1, publicKey),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeSSHKeyExists("ibm_compute_ssh_key.testacc_ssh_key", &key),
					resource.TestCheckResourceAttr(
						"ibm_compute_ssh_key.testacc_ssh_key", "label", label1),
					resource.TestCheckResourceAttr(
						"ibm_compute_ssh_key.testacc_ssh_key", "public_key", publicKey),
					resource.TestCheckResourceAttr(
						"ibm_compute_ssh_key.testacc_ssh_key", "notes", notes1),
					resource.TestCheckResourceAttr(
						"ibm_compute_ssh_key.testacc_ssh_key", "tags.#", "3"),
				),
			},
		},
	})
}

func testAccCheckIBMComputeSSHKeyDestroy(s *terraform.State) error {
	service := services.GetSecuritySshKeyService(acc.TestAccProvider.Meta().(conns.ClientSession).SoftLayerSession())

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_compute_ssh_key" {
			continue
		}

		keyID, _ := strconv.Atoi(rs.Primary.ID)

		// Try to find the key
		_, err := service.Id(keyID).GetObject()

		if err == nil {
			return fmt.Errorf("SSH key %d still exists", keyID)
		}
	}

	return nil
}

func testAccCheckIBMComputeSSHKeyExists(n string, key *datatypes.Security_Ssh_Key) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("[ERROR] Not  found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		keyID, _ := strconv.Atoi(rs.Primary.ID)

		service := services.GetSecuritySshKeyService(acc.TestAccProvider.Meta().(conns.ClientSession).SoftLayerSession())
		foundKey, err := service.Id(keyID).GetObject()

		if err != nil {
			return err
		}

		if strconv.Itoa(int(*foundKey.Id)) != rs.Primary.ID {
			return fmt.Errorf("Record %d not found", keyID)
		}

		*key = foundKey

		return nil
	}
}

func testAccCheckIBMComputeSSHKeyConfig(label, notes, publicKey string) string {
	return fmt.Sprintf(`
resource "ibm_compute_ssh_key" "testacc_ssh_key" {
    label = "%s"
    notes = "%s"
    public_key = "%s"
}`, label, notes, publicKey)

}

func testAccCheckIBMComputeSSHKeyWithTag(label, notes, publicKey string) string {
	return fmt.Sprintf(`
resource "ibm_compute_ssh_key" "testacc_ssh_key" {
    label = "%s"
    notes = "%s"
	public_key = "%s"
	tags = ["one", "two"]
}`, label, notes, publicKey)

}

func testAccCheckIBMComputeSSHKeyWithUpdatedTag(label, notes, publicKey string) string {
	return fmt.Sprintf(`
resource "ibm_compute_ssh_key" "testacc_ssh_key" {
    label = "%s"
    notes = "%s"
	public_key = "%s"
	tags = ["one", "two", "three"]
}`, label, notes, publicKey)

}
