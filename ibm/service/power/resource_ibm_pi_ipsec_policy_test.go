// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
)

func TestAccIBMPIIPSecPolicyBasic(t *testing.T) {
	policyRes := "ibm_pi_ipsec_policy.policy"
	name := fmt.Sprintf("tf-pi-ipsec-policy-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIIPSecPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIIPSecPolicyConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIIPSecPolicyExists(policyRes),
					resource.TestCheckResourceAttr(policyRes, "pi_policy_name", name),
					resource.TestCheckResourceAttrSet(policyRes, "policy_id"),
					resource.TestCheckResourceAttr(policyRes, "pi_policy_authentication", "hmac-sha-256-128"),
				),
			},
		},
	})
}
func testAccCheckIBMPIIPSecPolicyDestroy(s *terraform.State) error {
	sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_ipsec_policy" {
			continue
		}
		cloudInstanceID, policyID, err := splitID(rs.Primary.ID)
		if err != nil {
			return err
		}
		client := st.NewIBMPIVpnPolicyClient(context.Background(), sess, cloudInstanceID)
		_, err = client.GetIPSecPolicy(policyID)
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

		sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
		if err != nil {
			return err
		}
		cloudInstanceID, policyID, err := splitID(rs.Primary.ID)
		if err != nil {
			return err
		}
		client := st.NewIBMPIVpnPolicyClient(context.Background(), sess, cloudInstanceID)

		_, err = client.GetIPSecPolicy(policyID)
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
		pi_policy_authentication = "hmac-sha-256-128"
	}
	`, acc.Pi_cloud_instance_id, name)
}
