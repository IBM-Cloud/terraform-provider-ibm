// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMPIKey_basic(t *testing.T) {
	keyRes := "ibm_pi_key.key"
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	name := fmt.Sprintf("tf-pi-sshkey-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIKeyConfig(publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIKeyExists(keyRes),
					resource.TestCheckResourceAttr(keyRes, "pi_key_name", name),
					resource.TestCheckResourceAttr(keyRes, "primary_workspace", "true"),
					resource.TestCheckResourceAttrSet(keyRes, "creation_date"),
					resource.TestCheckResourceAttr(keyRes, "key", publicKey),
					resource.TestCheckResourceAttr(keyRes, "name", name),
					resource.TestCheckResourceAttr(keyRes, "pi_visibility", "workspace"),
				),
			},
		},
	})
}

func testAccCheckIBMPIKeyConfig(publicKey, name string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_key" "key" {
			pi_cloud_instance_id = "%[1]s"
			pi_key_name          = "%[2]s"
			pi_ssh_key           = "%[3]s"
		}`, acc.Pi_cloud_instance_id, name, publicKey)
}

func TestAccIBMPIKeyAccount(t *testing.T) {
	keyRes := "ibm_pi_key.key"
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	name := fmt.Sprintf("tf-pi-sshkey-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIKeyAccountConfig(publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIKeyExists(keyRes),
					resource.TestCheckResourceAttr(keyRes, "pi_key_name", name),
					resource.TestCheckResourceAttr(keyRes, "primary_workspace", "true"),
					resource.TestCheckResourceAttrSet(keyRes, "creation_date"),
					resource.TestCheckResourceAttr(keyRes, "key", publicKey),
					resource.TestCheckResourceAttr(keyRes, "name", name),
					resource.TestCheckResourceAttr(keyRes, "pi_visibility", "account"),
				),
			},
		},
	})
}

func testAccCheckIBMPIKeyAccountConfig(publicKey, name string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_key" "key" {
			pi_cloud_instance_id = "%[1]s"
			pi_key_name          = "%[2]s"
			pi_ssh_key           = "%[3]s"
			pi_visibility        = "account"
		}`, acc.Pi_cloud_instance_id, name, publicKey)
}

func TestAccIBMPIKeyUpdate(t *testing.T) {
	keyRes := "ibm_pi_key.key"
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	description := "test_description"
	descriptionUpdated := "test_description_updated"
	name := fmt.Sprintf("tf-pi-sshkey-%d", acctest.RandIntRange(10, 100))
	nameUpdated := name + "updated"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIKeyUpdateConfig(publicKey, name, description, "workspace"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIKeyExists(keyRes),
					resource.TestCheckResourceAttr(keyRes, "pi_key_name", name),
					resource.TestCheckResourceAttr(keyRes, "primary_workspace", "true"),
					resource.TestCheckResourceAttrSet(keyRes, "creation_date"),
					resource.TestCheckResourceAttr(keyRes, "key", publicKey),
					resource.TestCheckResourceAttr(keyRes, "name", name),
					resource.TestCheckResourceAttr(keyRes, "pi_visibility", "workspace"),
					resource.TestCheckResourceAttr(keyRes, "pi_description", description),
				),
			},
			{
				Config: testAccCheckIBMPIKeyUpdateConfig(publicKey, nameUpdated, descriptionUpdated, "account"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIKeyExists(keyRes),
					resource.TestCheckResourceAttr(keyRes, "pi_key_name", nameUpdated),
					resource.TestCheckResourceAttr(keyRes, "primary_workspace", "true"),
					resource.TestCheckResourceAttrSet(keyRes, "creation_date"),
					resource.TestCheckResourceAttr(keyRes, "key", publicKey),
					resource.TestCheckResourceAttr(keyRes, "name", nameUpdated),
					resource.TestCheckResourceAttr(keyRes, "pi_visibility", "account"),
					resource.TestCheckResourceAttr(keyRes, "pi_description", descriptionUpdated),
				),
			},
		},
	})
}

func testAccCheckIBMPIKeyUpdateConfig(publicKey, name string, description string, visibility string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_key" "key" {
			pi_cloud_instance_id = "%[1]s"
			pi_key_name          = "%[2]s"
			pi_ssh_key           = "%[3]s"
			pi_description       = "%[4]s"
			pi_visibility        = "%[5]s"
		}`, acc.Pi_cloud_instance_id, name, publicKey, description, visibility)
}

func testAccCheckIBMPIKeyDestroy(s *terraform.State) error {
	sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_key" {
			continue
		}
		cloudInstanceID, key, err := splitID(rs.Primary.ID)
		if err != nil {
			return err
		}
		sshkeyC := instance.NewIBMPISSHKeyClient(context.Background(), sess, cloudInstanceID)
		_, err = sshkeyC.Get(key)
		if err == nil {
			return fmt.Errorf("PI key still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckIBMPIKeyExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
		if err != nil {
			return err
		}

		cloudInstanceID, key, err := splitID(rs.Primary.ID)
		if err != nil {
			return err
		}

		client := instance.NewIBMPISSHKeyClient(context.Background(), sess, cloudInstanceID)
		_, err = client.Get(key)
		if err != nil {
			return err
		}
		return nil
	}
}
