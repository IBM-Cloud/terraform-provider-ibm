// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
)

func TestAccIBMPIKey_basic(t *testing.T) {
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	name := fmt.Sprintf("tf-pi-sshkey-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMPIKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIKeyConfig(publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIKeyExists("ibm_pi_key.key"),
					resource.TestCheckResourceAttr(
						"ibm_pi_key.key", "pi_key_name", name),
				),
			},
		},
	})
}
func testAccCheckIBMPIKeyDestroy(s *terraform.State) error {

	sess, err := testAccProvider.Meta().(ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_key" {
			continue
		}
		parts, err := idParts(rs.Primary.ID)
		powerinstanceid := parts[0]
		sshkeyC := st.NewIBMPIKeyClient(sess, powerinstanceid)
		_, err = sshkeyC.Get(parts[1], powerinstanceid)
		if err == nil {
			return fmt.Errorf("PI key still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}
func testAccCheckIBMPIKeyExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		fmt.Println("siv ", s.RootModule().Resources)
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, err := testAccProvider.Meta().(ClientSession).IBMPISession()
		if err != nil {
			return err
		}
		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		powerinstanceid := parts[0]
		client := st.NewIBMPIKeyClient(sess, powerinstanceid)

		key, err := client.Get(parts[1], powerinstanceid)
		if err != nil {
			return err
		}
		parts[1] = *key.Name
		return nil

	}
}

func testAccCheckIBMPIKeyConfig(publicKey, name string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_key" "key" {
			pi_cloud_instance_id = "%s"
			pi_key_name          = "%s"
			pi_ssh_key           = "%s"
		  }
	`, pi_cloud_instance_id, name, publicKey)
}
