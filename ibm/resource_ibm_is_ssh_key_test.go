// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/IBM/vpc-go-sdk/vpcclassicv1"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMISSSHKey_basic(t *testing.T) {
	var key string
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	name := fmt.Sprintf("tfssh-createname-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: checkKeyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISKeyConfig(publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISKeyExists("ibm_is_ssh_key.isExampleKey", key),
					resource.TestCheckResourceAttr(
						"ibm_is_ssh_key.isExampleKey", "name", name),
				),
			},
		},
	})
}

func checkKeyDestroy(s *terraform.State) error {
	userDetails, _ := testAccProvider.Meta().(ClientSession).BluemixUserDetails()

	if userDetails.generation == 1 {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcClassicV1API()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_is_ssh_key" {
				continue
			}

			getkeyoptions := &vpcclassicv1.GetKeyOptions{
				ID: &rs.Primary.ID,
			}
			_, _, err := sess.GetKey(getkeyoptions)
			if err == nil {
				return fmt.Errorf("key still exists: %s", rs.Primary.ID)
			}
		}
	} else {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_is_ssh_key" {
				continue
			}

			getkeyoptions := &vpcv1.GetKeyOptions{
				ID: &rs.Primary.ID,
			}
			_, _, err := sess.GetKey(getkeyoptions)
			if err == nil {
				return fmt.Errorf("key still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}

func testAccCheckIBMISKeyExists(n, keyID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		fmt.Println("siv ", s.RootModule().Resources)
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		userDetails, _ := testAccProvider.Meta().(ClientSession).BluemixUserDetails()

		if userDetails.generation == 1 {
			sess, _ := testAccProvider.Meta().(ClientSession).VpcClassicV1API()
			getkeyoptions := &vpcclassicv1.GetKeyOptions{
				ID: &rs.Primary.ID,
			}
			foundkey, _, err := sess.GetKey(getkeyoptions)
			if err != nil {
				return err
			}
			keyID = *foundkey.ID
		} else {
			sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
			getkeyoptions := &vpcv1.GetKeyOptions{
				ID: &rs.Primary.ID,
			}
			foundkey, _, err := sess.GetKey(getkeyoptions)
			if err != nil {
				return err
			}
			keyID = *foundkey.ID
		}
		return nil
	}
}

func testAccCheckIBMISKeyConfig(publicKey, name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ssh_key" "isExampleKey" {
			name = "%s"
			public_key = "%s"
		}
	`, name, publicKey)
}
