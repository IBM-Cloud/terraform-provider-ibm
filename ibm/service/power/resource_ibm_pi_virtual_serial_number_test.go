// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
)

func TestAccIBMPIVirtualSerialNumberBasic(t *testing.T) {
	resLocator := "ibm_pi_virtual_serial_number.power_virtual_serial_number"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMPIVirtualSerialNumberBasicConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIVirtualSerialNumberExists(resLocator),
					resource.TestCheckResourceAttrSet(resLocator, "id"),
				),
			},
		},
	})
}

func TestAccIBMPIVirtualSerialNumberWithInstance(t *testing.T) {
	resLocator := "ibm_pi_virtual_serial_number.power_virtual_serial_number"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIVirtualSerialNumberDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMPIVirtualSerialNumberWithInstanceConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIVirtualSerialNumberExists(resLocator),
					resource.TestCheckResourceAttrSet(resLocator, "id"),
					resource.TestCheckResourceAttrSet(resLocator, "pi_serial"),
					resource.TestCheckResourceAttr(resLocator, "pi_description", "TF test description"),
					resource.TestCheckResourceAttrSet(resLocator, "pi_instance_id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIVirtualSerialNumberExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return flex.FmtErrorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
		if err != nil {
			return err
		}
		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		cloudInstanceId := parts[0]

		client := st.NewIBMPIVSNClient(context.Background(), sess, cloudInstanceId)

		_, err = client.Get(parts[1])
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckIBMPIVirtualSerialNumberDestroy(s *terraform.State) error {
	sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_virtual_serial_number" {
			continue
		}
		parts, _ := flex.IdParts(rs.Primary.ID)
		cloudInstanceId := parts[0]
		vsnClient := st.NewIBMPIVSNClient(context.Background(), sess, cloudInstanceId)
		_, err = vsnClient.Get(parts[1])
		if err == nil {
			return flex.FmtErrorf("PI virtual serial number still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccIBMPIVirtualSerialNumberBasicConfig() string {
	return fmt.Sprintf(`
		resource "ibm_pi_virtual_serial_number" "power_virtual_serial_number" {
			pi_cloud_instance_id 			= "%[1]s"
			pi_description       			= "TF test description"
			pi_serial            			= "%[2]s"
		}
	`, acc.Pi_cloud_instance_id, acc.Pi_virtual_serial_number)
}

func testAccIBMPIVirtualSerialNumberWithInstanceConfig() string {
	return fmt.Sprintf(`
		resource "ibm_pi_virtual_serial_number" "power_virtual_serial_number" {
			pi_cloud_instance_id            = "%[1]s"
			pi_description   				= "TF test description"
			pi_instance_id                  = "%[2]s"
			pi_retain_virtual_serial_number = false
			pi_serial 						= "auto-assign"
		}
	`, acc.Pi_cloud_instance_id, acc.Pi_instance_name)
}
