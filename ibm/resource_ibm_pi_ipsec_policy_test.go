// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
)

func TestAccIBMPIIPSecPolicyBasic(t *testing.T) {
	policyRes := "ibm_pi_ipsec_policy.policy"
	name := fmt.Sprintf("tf-pi-ipsec-policy-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMPIIPSecPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIIPSecPolicyConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIIPSecPolicyExists(policyRes),
					resource.TestCheckResourceAttr(policyRes, "pi_policy_name", name),
					resource.TestCheckResourceAttrSet(policyRes, "policy_id"),
					resource.TestCheckResourceAttr(policyRes, "pi_policy_authentication", "hmac-md5-96"),
				),
			},
		},
	})
}
func testAccCheckIBMPIIPSecPolicyDestroy(s *terraform.State) error {
	sess, err := testAccProvider.Meta().(ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_ipsec_policy" {
			continue
		}
		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		cloudInstanceID := parts[0]
		policyID := parts[1]
		client := st.NewIBMPIVpnPolicyClient(sess, cloudInstanceID)
		_, err = client.GetIPSecPolicy(policyID, cloudInstanceID)
		if err == nil {
			return fmt.Errorf("ipsec policy still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}
func testAccCheckIBMPIIPSecPolicyExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
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
		cloudInstanceID := parts[0]
		policyID := parts[1]
		client := st.NewIBMPIVpnPolicyClient(sess, cloudInstanceID)

		_, err = client.GetIPSecPolicy(policyID, cloudInstanceID)
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckIBMPIIPSecPolicyConfig(name string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_ipsec_policy" "policy" {
		pi_cloud_instance_id = "%s"
		pi_policy_name = "%s"
		pi_policy_dh_group = 1
		pi_policy_encryption = "3des-cbc"
		pi_policy_key_lifetime = 180
		pi_policy_pfs = true
		pi_policy_authentication = "hmac-md5-96"
	}
	`, pi_cloud_instance_id, name)
}
