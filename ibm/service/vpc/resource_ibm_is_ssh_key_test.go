// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMISSSHKey_basic(t *testing.T) {
	var key string
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	name := fmt.Sprintf("tfssh-createname-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkKeyDestroy,
		Steps: []resource.TestStep{
			{
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
func TestAccIBMISSSHKey_Newlinebasic(t *testing.T) {
	var key, key1 string
	publicKeyTrim := strings.TrimSpace(`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDKd3e4uoENwyLGDxpUyxkeC008r6JOHGkeF4HxEUrE2ZkTXvuOyRaF8Utyv12U5Jvpf/NQVGmdlG3rQPY5VRELthr8mhNWexY5WYo/zXSZNjHCjozpL101bcxfNG498y6uv6UoTk6geJcmckVjOYY/2T9F2B1q6dQxIsYIghjfZBFM6+wA136Nx0nof2lZiK7KAIzIlgUY3g3hhno0x5FmJHM9waoHXFLgQA0psz8XUcSt2Zr0JGFOm5U6HV/tvoP4AVB5YrhRatHry2Ulfh4acy0wswgRM0zieU0U/nLJbCgDVLwZyABEC6WcLTfrkkI53I9oYb8XyeWpyRFQLfT7AIIjfgT7q0q4gzSKTJSDR85SHhOmC/bhCDBuJ9s1ICyschF1y8lPjL/maxweNorg3RfuqsZZdmNHR9RKSxt7CPM98Z6yu5wMWRVLC6Ux5MGp6m0mDIJOfaZla6uvp8d/G6cjWCdU5eCeBh6XdQn4UDXwEB/s86lpgbDsPLMCleP8J+w8uZPQA1KZ+uWGBjoswhtOCa6bU/6ZuTqGpQVOGjVOWUGOq/ocvR03ucj6fBKViFWxV75ABXfJLarKkkIMlv9IeJ05NZG6kQjiCRN4T2I0gd9lAm0YqcITEqcN4Wgbm1z2zPwvMyWuMCW3LY4932JHKQkCEXgGBAtsnrXhZw==`)
	publicKeyTrim1 := strings.TrimSpace(`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR`)
	name := fmt.Sprintf("tfssh-createname-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfssh-withsignaturename-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISKeyNewlineConfig(name, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISKeyExists("ibm_is_ssh_key.isExampleKey", key),
					testAccCheckIBMISKeyExists("ibm_is_ssh_key.isExampleKeyWithSignature", key1),
					resource.TestCheckResourceAttr(
						"ibm_is_ssh_key.isExampleKey", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_ssh_key.isExampleKey", "public_key", publicKeyTrim),
					resource.TestCheckResourceAttr(
						"ibm_is_ssh_key.isExampleKeyWithSignature", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_ssh_key.isExampleKeyWithSignature", "public_key", publicKeyTrim1),
				),
			},
		},
	})
}

func checkKeyDestroy(s *terraform.State) error {
	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
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

	return nil
}

func testAccCheckIBMISKeyExists(n, keyID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		getkeyoptions := &vpcv1.GetKeyOptions{
			ID: &rs.Primary.ID,
		}
		foundkey, _, err := sess.GetKey(getkeyoptions)
		if err != nil {
			return err
		}
		keyID = *foundkey.ID
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
func testAccCheckIBMISKeyNewlineConfig(name, name1 string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ssh_key" "isExampleKey" {
			name = "%s"
			public_key = <<-EOT
            ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDKd3e4uoENwyLGDxpUyxkeC008r6JOHGkeF4HxEUrE2ZkTXvuOyRaF8Utyv12U5Jvpf/NQVGmdlG3rQPY5VRELthr8mhNWexY5WYo/zXSZNjHCjozpL101bcxfNG498y6uv6UoTk6geJcmckVjOYY/2T9F2B1q6dQxIsYIghjfZBFM6+wA136Nx0nof2lZiK7KAIzIlgUY3g3hhno0x5FmJHM9waoHXFLgQA0psz8XUcSt2Zr0JGFOm5U6HV/tvoP4AVB5YrhRatHry2Ulfh4acy0wswgRM0zieU0U/nLJbCgDVLwZyABEC6WcLTfrkkI53I9oYb8XyeWpyRFQLfT7AIIjfgT7q0q4gzSKTJSDR85SHhOmC/bhCDBuJ9s1ICyschF1y8lPjL/maxweNorg3RfuqsZZdmNHR9RKSxt7CPM98Z6yu5wMWRVLC6Ux5MGp6m0mDIJOfaZla6uvp8d/G6cjWCdU5eCeBh6XdQn4UDXwEB/s86lpgbDsPLMCleP8J+w8uZPQA1KZ+uWGBjoswhtOCa6bU/6ZuTqGpQVOGjVOWUGOq/ocvR03ucj6fBKViFWxV75ABXfJLarKkkIMlv9IeJ05NZG6kQjiCRN4T2I0gd9lAm0YqcITEqcN4Wgbm1z2zPwvMyWuMCW3LY4932JHKQkCEXgGBAtsnrXhZw==
        EOT
		}
		resource "ibm_is_ssh_key" "isExampleKeyWithSignature" {
			name = "%s"
			public_key = <<-EOT
            ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR  abc@abc-MacBook-Pro.local
        EOT
		}
	`, name, name1)
}
