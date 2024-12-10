// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package classicinfrastructure_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/services"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
)

func TestAccIBMSecurityGroup_basic(t *testing.T) {
	var sg datatypes.Network_SecurityGroup

	name1 := fmt.Sprintf("terraformsguat_create_step_name_%d", acctest.RandIntRange(10, 100))
	desc1 := fmt.Sprintf("terraformsguat_create_step_desc_%d", acctest.RandIntRange(10, 100))
	name2 := fmt.Sprintf("terraformsguat_create_step_name_%d", acctest.RandIntRange(10, 100))
	desc2 := fmt.Sprintf("terraformsguat_create_step_desc_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSecurityGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSecurityGroupConfig(name1, desc1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMSecurityGroupExists("ibm_security_group.testacc_security_group", &sg),
					resource.TestCheckResourceAttr(
						"ibm_security_group.testacc_security_group", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_security_group.testacc_security_group", "description", desc1),
				),
			},

			{
				Config: testAccCheckIBMSecurityGroupConfig(name2, desc2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMSecurityGroupExists("ibm_security_group.testacc_security_group", &sg),
					resource.TestCheckResourceAttr(
						"ibm_security_group.testacc_security_group", "name", name2),
					resource.TestCheckResourceAttr(
						"ibm_security_group.testacc_security_group", "description", desc2),
				),
			},
		},
	})
}

func testAccCheckIBMSecurityGroupDestroy(s *terraform.State) error {
	service := services.GetNetworkSecurityGroupService(acc.TestAccProvider.Meta().(conns.ClientSession).SoftLayerSession())

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_security_group" {
			continue
		}

		sgID, _ := strconv.Atoi(rs.Primary.ID)

		// Try to find the key
		_, err := service.Id(sgID).GetObject()

		if err == nil {
			return fmt.Errorf("Security Group %d still exists", sgID)
		}
	}

	return nil
}

func testAccCheckIBMSecurityGroupExists(n string, sg *datatypes.Network_SecurityGroup) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("[ERROR] Not  found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sgID, _ := strconv.Atoi(rs.Primary.ID)

		service := services.GetNetworkSecurityGroupService(acc.TestAccProvider.Meta().(conns.ClientSession).SoftLayerSession())
		foundSG, err := service.Id(sgID).GetObject()

		if err != nil {
			return err
		}

		if strconv.Itoa(int(*foundSG.Id)) != rs.Primary.ID {
			return fmt.Errorf("Record %d not found", sgID)
		}

		*sg = foundSG

		return nil
	}
}

func testAccCheckIBMSecurityGroupConfig(name, description string) string {
	return fmt.Sprintf(`
resource "ibm_security_group" "testacc_security_group" {
    name = "%s"
    description = "%s"
}`, name, description)

}
