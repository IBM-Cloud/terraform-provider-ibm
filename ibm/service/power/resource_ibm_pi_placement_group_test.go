// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
)

func TestAccIBMPIPlacementGroupBasic(t *testing.T) {
	name := fmt.Sprintf("tf-pi-placement-group-%d", acctest.RandIntRange(10, 100))
	policy := "affinity"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIPlacementGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config:       testAccCheckIBMPIPlacementGroupConfig(name, policy),
				ResourceName: "ibm_pi_placement_group.power_placement_group",
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIPlacementGroupExists("ibm_pi_placement_group.power_placement_group"),
					resource.TestCheckResourceAttr(
						"ibm_pi_placement_group.power_placement_group", "pi_placement_group_name", name),
					resource.TestCheckResourceAttr(
						"ibm_pi_placement_group.power_placement_group", "pi_placement_group_policy", policy),
					resource.TestCheckNoResourceAttr(
						"ibm_pi_placement_group.power_placement_group", "members"),
				),
			},
			{
				Config: testAccCheckIBMPIPlacementGroupAddMemberConfig(name, policy),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"ibm_pi_instance.power_instance", "pi_placement_group_id"),
					testAccCheckIBMPIPlacementGroupMemberExists("ibm_pi_placement_group.power_placement_group", "ibm_pi_instance.power_instance"),
				),
			},
			{
				Config: testAccCheckIBMPIPlacementGroupUpdateMemberConfig(name, policy),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"ibm_pi_instance.power_instance", "pi_placement_group_id"),
					testAccCheckIBMPIPlacementGroupMemberExists("ibm_pi_placement_group.power_placement_group_another", "ibm_pi_instance.power_instance"),
				),
			},
			{
				Config: testAccCheckIBMPIPlacementGroupRemoveMemberConfig(name, policy),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIPlacementGroupMemberDoesNotExist("ibm_pi_placement_group.power_placement_group", "ibm_pi_instance.power_instance"),
				),
			},
			{
				Config: testAccCheckIBMPICreateInstanceInPlacementGroup(name, policy, "tinytest-1x4"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"ibm_pi_instance.power_instance", "pi_placement_group_id"),
					resource.TestCheckResourceAttrSet(
						"ibm_pi_instance.power_instance_in_pg", "pi_placement_group_id"),
					resource.TestCheckResourceAttrSet(
						"ibm_pi_instance.sap_power_instance", "pi_placement_group_id"),
					testAccCheckIBMPIPlacementGroupMemberExistsFromInstanceCreate("ibm_pi_placement_group.power_placement_group", "ibm_pi_instance.power_instance", "ibm_pi_instance.power_instance_in_pg"),
					testAccCheckIBMPIPlacementGroupMemberExists("ibm_pi_placement_group.power_placement_group", "ibm_pi_instance.sap_power_instance"),
				),
			},
			{
				Config: testAccCheckIBMPIDeletePlacementGroup(name, policy, "tinytest-1x4"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIPlacementGroupDelete("ibm_pi_placement_group.power_placement_group", "ibm_pi_instance.power_instance", "ibm_pi_instance.power_instance_in_pg"),
				),
			},
		},
	})
}

func testAccCheckIBMPIPlacementGroupDestroy(s *terraform.State) error {

	sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_placement_group" {
			continue
		}
		parts, _ := flex.IdParts(rs.Primary.ID)
		cloudinstanceid := parts[0]
		placementGroupC := st.NewIBMPIPlacementGroupClient(context.Background(), sess, cloudinstanceid)
		_, err = placementGroupC.Get(parts[1])
		if err == nil {
			return fmt.Errorf("PI placement group still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}
func testAccCheckIBMPIPlacementGroupExists(n string) resource.TestCheckFunc {
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
		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		cloudinstanceid := parts[0]
		client := st.NewIBMPIPlacementGroupClient(context.Background(), sess, cloudinstanceid)

		placementGroup, err := client.Get(parts[1])
		if err != nil {
			return err
		}
		parts[1] = *placementGroup.ID
		return nil
	}
}

func testAccCheckIBMPIPlacementGroupMemberExists(n string, instance string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		// refresh placement group info since a server should be in the placement group
		sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
		if err != nil {
			return err
		}
		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		cloudinstanceid := parts[0]
		client := st.NewIBMPIPlacementGroupClient(context.Background(), sess, cloudinstanceid)

		pg, err := client.Get(parts[1])
		if err != nil {
			return err
		}

		instancers, ok := s.RootModule().Resources[instance]
		if !ok {
			return fmt.Errorf("Not found: %s", instance)
		}
		instanceParts, err := flex.IdParts(instancers.Primary.ID)
		if err != nil {
			return err
		}
		var isInstanceFound bool = false
		for _, x := range pg.Members {
			if x == instanceParts[1] {
				isInstanceFound = true
				break
			}
		}
		if !isInstanceFound {
			return fmt.Errorf("Expected server ID %s in the PG members field but found %s", instanceParts[1], strings.Join(pg.Members[:], ","))
		}
		return nil
	}
}

func testAccCheckIBMPIPlacementGroupMemberDoesNotExist(n string, instance string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		// refresh placement group info since a server should be in the placement group
		sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
		if err != nil {
			return err
		}
		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		cloudinstanceid := parts[0]
		client := st.NewIBMPIPlacementGroupClient(context.Background(), sess, cloudinstanceid)

		pg, err := client.Get(parts[1])
		if err != nil {
			return err
		}

		instancers, ok := s.RootModule().Resources[instance]
		if !ok {
			return fmt.Errorf("Not found: %s", instance)
		}
		instanccParts, err := flex.IdParts(instancers.Primary.ID)
		if err != nil {
			return err
		}
		if len(pg.Members) > 0 {
			return fmt.Errorf("Expected server ID %s to be removed so that the PG members field is empty but foumd %s", instanccParts[1], pg.Members[0])
		}

		return nil
	}
}

func containsMember(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func testAccCheckIBMPIPlacementGroupMemberExistsFromInstanceCreate(n string, instance string, newInstance string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		// refresh placement group info since a server should be in the placement group
		sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
		if err != nil {
			return err
		}
		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		cloudinstanceid := parts[0]
		client := st.NewIBMPIPlacementGroupClient(context.Background(), sess, cloudinstanceid)

		pg, err := client.Get(parts[1])
		if err != nil {
			return err
		}

		instancers, ok := s.RootModule().Resources[instance]
		if !ok {
			return fmt.Errorf("Not found: %s", instance)
		}
		instanceParts, err := flex.IdParts(instancers.Primary.ID)
		if err != nil {
			return err
		}

		newinstancers, ok := s.RootModule().Resources[newInstance]
		if !ok {
			return fmt.Errorf("Not found: %s", newInstance)
		}
		newinstanceParts, err := flex.IdParts(newinstancers.Primary.ID)
		if err != nil {
			return err
		}

		if !containsMember(pg.Members, instanceParts[1]) {
			return fmt.Errorf("Expected server ID %s in the PG members field", instanceParts[1])
		}
		if !containsMember(pg.Members, newinstanceParts[1]) {
			return fmt.Errorf("Expected new server ID %s in the PG members field", newinstanceParts[1])
		}
		return nil
	}
}

func testAccCheckIBMPIPlacementGroupDelete(n string, instance string, newInstance string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
		if err != nil {
			return err
		}

		instancers, ok := s.RootModule().Resources[instance]
		if !ok {
			return fmt.Errorf("Not found: %s", instance)
		}
		instanceParts, err := flex.IdParts(instancers.Primary.ID)
		if err != nil {
			return err
		}

		newinstancers, ok := s.RootModule().Resources[newInstance]
		if !ok {
			return fmt.Errorf("Not found: %s", newInstance)
		}
		newinstanceParts, err := flex.IdParts(newinstancers.Primary.ID)
		if err != nil {
			return err
		}
		cloudinstanceid := instanceParts[0]
		inst_client := st.NewIBMPIInstanceClient(context.Background(), sess, cloudinstanceid)

		instance, err := inst_client.Get(instanceParts[1])
		if err != nil {
			return err
		}

		if *instance.PlacementGroup != "none" {
			return fmt.Errorf("Expected no placement group ID in the PVM instance data but found %s", *instance.PlacementGroup)
		}
		newinstance, err := inst_client.Get(newinstanceParts[1])
		if err != nil {
			return err
		}
		if *newinstance.PlacementGroup != "none" {
			return fmt.Errorf("Expected no placement group ID in the PVM instance data but found %s", *newinstance.PlacementGroup)
		}
		return nil
	}
}

func testAccCheckIBMPIPlacementGroupConfig(name string, policy string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_key" "key" {
			pi_cloud_instance_id = "%[1]s"
			pi_key_name          = "%[2]s"
			pi_ssh_key           = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDKUt7bk9yLBZFC187bQJFuLaBZONKFYjeIGCZj5mN0OvaJdJqPN2Mbx9Ui42Y5vrLE7SipG5c94BS/fYf7e2LvsQ+xaU1VQnMvP6XS8emoyKR6q/YzD60MkvkSopwTAgpyf6CpfCsKE5Yclbrsc1HIP16bjSgOapfgaVuEDXifn27i1fP1QRYhosY7YkfSKjyJQihxnFH1sONdl4JspJDC5rp8wZ4E7jSXyaZh6QIMbMBEvKoE8+/8CUgT3EWWndIOIMuPQtills3X3jDojTt722OBW1qETPahYDDEmN00R1Q1Q8V8pfVi1XG+ESLzY93gC8hV+/lWIoIvSEazwkfi7/5kludrZG1RhCGbOffGo3DkrmtqaBbKbjrTh/ZbY0GzHPXqccfW/KIhk6xlmoR0wF9LYPtFuzTkqnHF/tHi8EXPHI5XVv9m01kMLkoUqtWVXP2O7ZM7EwrJ+1TyJqLTrzbKMUbn52GqNuTSFJCAgEVc3XrvIRFjTL1/b428mS9JV5kCfRVLmDAUtPjuaQg1wmI/W97gZCF8IoF4JXWTEQP8IIb2opLxvEoBggsZpiFOtjsr9A914i/Tyd4T4KlvfkavJXqkzQoj29oZZPt10gt2ywwXPvV6usM1iofATB+YtX6vl8wUDaqvEyC8d4OTnSVkPZnFxTG3lhY4cDwa/w== tedfordt@us.ibm.com"
		}

		resource "ibm_pi_instance" "power_instance" {
			pi_processors         = "0.25"
			pi_proc_type          = "shared"
			pi_memory             = "2"
			pi_key_pair_name      = ibm_pi_key.key.key_id
			pi_image_id           = "%[4]s"
			pi_sys_type           = "e980"
			pi_instance_name      = "%[2]s"
			pi_cloud_instance_id  = "%[1]s"
			pi_storage_type       = "tier3"
			pi_network {
				network_id = "%[5]s"
			}
		}
		resource "ibm_pi_placement_group" "power_placement_group" {
			pi_cloud_instance_id      = "%[1]s"
			pi_placement_group_name   = "%[2]s"
			pi_placement_group_policy = "%[3]s"
		}
	`, acc.Pi_cloud_instance_id, name, policy, acc.Pi_image, acc.Pi_network_name)
}

func testAccCheckIBMPIPlacementGroupAddMemberConfig(name string, policy string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_key" "key" {
			pi_cloud_instance_id = "%[1]s"
			pi_key_name          = "%[2]s"
			pi_ssh_key           = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDKUt7bk9yLBZFC187bQJFuLaBZONKFYjeIGCZj5mN0OvaJdJqPN2Mbx9Ui42Y5vrLE7SipG5c94BS/fYf7e2LvsQ+xaU1VQnMvP6XS8emoyKR6q/YzD60MkvkSopwTAgpyf6CpfCsKE5Yclbrsc1HIP16bjSgOapfgaVuEDXifn27i1fP1QRYhosY7YkfSKjyJQihxnFH1sONdl4JspJDC5rp8wZ4E7jSXyaZh6QIMbMBEvKoE8+/8CUgT3EWWndIOIMuPQtills3X3jDojTt722OBW1qETPahYDDEmN00R1Q1Q8V8pfVi1XG+ESLzY93gC8hV+/lWIoIvSEazwkfi7/5kludrZG1RhCGbOffGo3DkrmtqaBbKbjrTh/ZbY0GzHPXqccfW/KIhk6xlmoR0wF9LYPtFuzTkqnHF/tHi8EXPHI5XVv9m01kMLkoUqtWVXP2O7ZM7EwrJ+1TyJqLTrzbKMUbn52GqNuTSFJCAgEVc3XrvIRFjTL1/b428mS9JV5kCfRVLmDAUtPjuaQg1wmI/W97gZCF8IoF4JXWTEQP8IIb2opLxvEoBggsZpiFOtjsr9A914i/Tyd4T4KlvfkavJXqkzQoj29oZZPt10gt2ywwXPvV6usM1iofATB+YtX6vl8wUDaqvEyC8d4OTnSVkPZnFxTG3lhY4cDwa/w== tedfordt@us.ibm.com"
		}

		resource "ibm_pi_instance" "power_instance" {
			pi_processors         = "0.25"
			pi_proc_type          = "shared"
			pi_memory             = "2"
			pi_key_pair_name      = ibm_pi_key.key.key_id
			pi_image_id           = "%[4]s"
			pi_sys_type           = "e980"
			pi_instance_name      = "%[2]s"
			pi_cloud_instance_id  = "%[1]s"
			pi_storage_type       = "tier3"
			pi_network {
				network_id = "%[5]s"
			}
			pi_placement_group_id = ibm_pi_placement_group.power_placement_group.placement_group_id
		}

		resource "ibm_pi_placement_group" "power_placement_group" {
			pi_cloud_instance_id      = "%[1]s"
			pi_placement_group_name   = "%[2]s"
			pi_placement_group_policy = "%[3]s"
		}
	`, acc.Pi_cloud_instance_id, name, policy, acc.Pi_image, acc.Pi_network_name)
}

func testAccCheckIBMPIPlacementGroupUpdateMemberConfig(name string, policy string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_key" "key" {
			pi_cloud_instance_id = "%[1]s"
			pi_key_name          = "%[2]s"
			pi_ssh_key           = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDKUt7bk9yLBZFC187bQJFuLaBZONKFYjeIGCZj5mN0OvaJdJqPN2Mbx9Ui42Y5vrLE7SipG5c94BS/fYf7e2LvsQ+xaU1VQnMvP6XS8emoyKR6q/YzD60MkvkSopwTAgpyf6CpfCsKE5Yclbrsc1HIP16bjSgOapfgaVuEDXifn27i1fP1QRYhosY7YkfSKjyJQihxnFH1sONdl4JspJDC5rp8wZ4E7jSXyaZh6QIMbMBEvKoE8+/8CUgT3EWWndIOIMuPQtills3X3jDojTt722OBW1qETPahYDDEmN00R1Q1Q8V8pfVi1XG+ESLzY93gC8hV+/lWIoIvSEazwkfi7/5kludrZG1RhCGbOffGo3DkrmtqaBbKbjrTh/ZbY0GzHPXqccfW/KIhk6xlmoR0wF9LYPtFuzTkqnHF/tHi8EXPHI5XVv9m01kMLkoUqtWVXP2O7ZM7EwrJ+1TyJqLTrzbKMUbn52GqNuTSFJCAgEVc3XrvIRFjTL1/b428mS9JV5kCfRVLmDAUtPjuaQg1wmI/W97gZCF8IoF4JXWTEQP8IIb2opLxvEoBggsZpiFOtjsr9A914i/Tyd4T4KlvfkavJXqkzQoj29oZZPt10gt2ywwXPvV6usM1iofATB+YtX6vl8wUDaqvEyC8d4OTnSVkPZnFxTG3lhY4cDwa/w== tedfordt@us.ibm.com"
		}

		resource "ibm_pi_instance" "power_instance" {
			pi_processors         = "0.25"
			pi_proc_type          = "shared"
			pi_memory             = "2"
			pi_key_pair_name      = ibm_pi_key.key.key_id
			pi_image_id           = "%[4]s"
			pi_sys_type           = "e980"
			pi_instance_name      = "%[2]s"
			pi_cloud_instance_id  = "%[1]s"
			pi_storage_type       = "tier3"
			pi_network {
				network_id = "%[5]s"
			}
			pi_placement_group_id = ibm_pi_placement_group.power_placement_group_another.placement_group_id
		}

		resource "ibm_pi_placement_group" "power_placement_group" {
			pi_cloud_instance_id      = "%[1]s"
			pi_placement_group_name   = "%[2]s"
			pi_placement_group_policy = "%[3]s"
		}

		resource "ibm_pi_placement_group" "power_placement_group_another" {
			pi_cloud_instance_id      = "%[1]s"
			pi_placement_group_name   = "%[2]s-2"
			pi_placement_group_policy = "%[3]s"
		}
	`, acc.Pi_cloud_instance_id, name, policy, acc.Pi_image, acc.Pi_network_name)
}

func testAccCheckIBMPIPlacementGroupRemoveMemberConfig(name string, policy string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_key" "key" {
			pi_cloud_instance_id = "%[1]s"
			pi_key_name          = "%[2]s"
			pi_ssh_key           = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDKUt7bk9yLBZFC187bQJFuLaBZONKFYjeIGCZj5mN0OvaJdJqPN2Mbx9Ui42Y5vrLE7SipG5c94BS/fYf7e2LvsQ+xaU1VQnMvP6XS8emoyKR6q/YzD60MkvkSopwTAgpyf6CpfCsKE5Yclbrsc1HIP16bjSgOapfgaVuEDXifn27i1fP1QRYhosY7YkfSKjyJQihxnFH1sONdl4JspJDC5rp8wZ4E7jSXyaZh6QIMbMBEvKoE8+/8CUgT3EWWndIOIMuPQtills3X3jDojTt722OBW1qETPahYDDEmN00R1Q1Q8V8pfVi1XG+ESLzY93gC8hV+/lWIoIvSEazwkfi7/5kludrZG1RhCGbOffGo3DkrmtqaBbKbjrTh/ZbY0GzHPXqccfW/KIhk6xlmoR0wF9LYPtFuzTkqnHF/tHi8EXPHI5XVv9m01kMLkoUqtWVXP2O7ZM7EwrJ+1TyJqLTrzbKMUbn52GqNuTSFJCAgEVc3XrvIRFjTL1/b428mS9JV5kCfRVLmDAUtPjuaQg1wmI/W97gZCF8IoF4JXWTEQP8IIb2opLxvEoBggsZpiFOtjsr9A914i/Tyd4T4KlvfkavJXqkzQoj29oZZPt10gt2ywwXPvV6usM1iofATB+YtX6vl8wUDaqvEyC8d4OTnSVkPZnFxTG3lhY4cDwa/w== tedfordt@us.ibm.com"
		}

		resource "ibm_pi_instance" "power_instance" {
			pi_processors         = "0.25"
			pi_proc_type          = "shared"
			pi_memory             = "2"
			pi_key_pair_name      = ibm_pi_key.key.key_id
			pi_image_id           = "%[4]s"
			pi_sys_type           = "e980"
			pi_instance_name      = "%[2]s"
			pi_cloud_instance_id  = "%[1]s"
			pi_storage_type       = "tier3"
			pi_network {
				network_id = "%[5]s"
			}
			pi_placement_group_id = ""
		}
	
		resource "ibm_pi_placement_group" "power_placement_group" {
			pi_cloud_instance_id      = "%[1]s"
			pi_placement_group_name   = "%[2]s"
			pi_placement_group_policy = "%[3]s"
		}

		resource "ibm_pi_placement_group" "power_placement_group_another" {
			pi_cloud_instance_id      = "%[1]s"
			pi_placement_group_name   = "%[2]s-2"
			pi_placement_group_policy = "%[3]s"
		}
	`, acc.Pi_cloud_instance_id, name, policy, acc.Pi_image, acc.Pi_network_name)
}

func testAccCheckIBMPICreateInstanceInPlacementGroup(name string, policy string, sapProfile string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_key" "key" {
			pi_cloud_instance_id = "%[1]s"
			pi_key_name          = "%[2]s"
			pi_ssh_key           = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDKUt7bk9yLBZFC187bQJFuLaBZONKFYjeIGCZj5mN0OvaJdJqPN2Mbx9Ui42Y5vrLE7SipG5c94BS/fYf7e2LvsQ+xaU1VQnMvP6XS8emoyKR6q/YzD60MkvkSopwTAgpyf6CpfCsKE5Yclbrsc1HIP16bjSgOapfgaVuEDXifn27i1fP1QRYhosY7YkfSKjyJQihxnFH1sONdl4JspJDC5rp8wZ4E7jSXyaZh6QIMbMBEvKoE8+/8CUgT3EWWndIOIMuPQtills3X3jDojTt722OBW1qETPahYDDEmN00R1Q1Q8V8pfVi1XG+ESLzY93gC8hV+/lWIoIvSEazwkfi7/5kludrZG1RhCGbOffGo3DkrmtqaBbKbjrTh/ZbY0GzHPXqccfW/KIhk6xlmoR0wF9LYPtFuzTkqnHF/tHi8EXPHI5XVv9m01kMLkoUqtWVXP2O7ZM7EwrJ+1TyJqLTrzbKMUbn52GqNuTSFJCAgEVc3XrvIRFjTL1/b428mS9JV5kCfRVLmDAUtPjuaQg1wmI/W97gZCF8IoF4JXWTEQP8IIb2opLxvEoBggsZpiFOtjsr9A914i/Tyd4T4KlvfkavJXqkzQoj29oZZPt10gt2ywwXPvV6usM1iofATB+YtX6vl8wUDaqvEyC8d4OTnSVkPZnFxTG3lhY4cDwa/w== tedfordt@us.ibm.com"
		}

		resource "ibm_pi_instance" "power_instance" {
			pi_processors         = "0.25"
			pi_proc_type          = "shared"
			pi_memory             = "2"
			pi_key_pair_name      = ibm_pi_key.key.key_id
			pi_image_id           = "%[4]s"
			pi_sys_type           = "e980"
			pi_instance_name      = "%[2]s"
			pi_cloud_instance_id  = "%[1]s"
			pi_storage_type       = "tier3"
			pi_network {
				network_id = "%[7]s"
			}
			pi_placement_group_id = ibm_pi_placement_group.power_placement_group.placement_group_id
		}

		resource "ibm_pi_instance" "power_instance_in_pg" {
			pi_processors         = "0.25"
			pi_proc_type          = "shared"
			pi_memory             = "2"
			pi_key_pair_name      = ibm_pi_key.key.key_id
			pi_image_id           = "%[4]s"
			pi_sys_type           = "e980"
			pi_instance_name      = "%[2]s-2"
			pi_cloud_instance_id  = "%[1]s"
			pi_network {
				network_id = "%[7]s"
			}
			pi_placement_group_id = ibm_pi_placement_group.power_placement_group.placement_group_id
		}

		resource "ibm_pi_instance" "sap_power_instance" {
			pi_cloud_instance_id  	= "%[1]s"
			pi_instance_name      	= "sap-%[2]s"
			pi_sap_profile_id       = "%[5]s"
			pi_image_id           	= "%[6]s"
			pi_storage_type			= "tier1"
			pi_network {
				network_id = "%[7]s"
			}
			pi_placement_group_id = ibm_pi_placement_group.power_placement_group.placement_group_id
			depends_on = [    ibm_pi_instance.power_instance_in_pg  ]
		}

		resource "ibm_pi_placement_group" "power_placement_group" {
			pi_cloud_instance_id      = "%[1]s"
			pi_placement_group_name   = "%[2]s"
			pi_placement_group_policy = "%[3]s"
		}

		resource "ibm_pi_placement_group" "power_placement_group_another" {
			pi_cloud_instance_id      = "%[1]s"
			pi_placement_group_name   = "%[2]s-2"
			pi_placement_group_policy = "%[3]s"
		}
	`, acc.Pi_cloud_instance_id, name, policy, acc.Pi_image, sapProfile, acc.Pi_sap_image, acc.Pi_network_name)
}

func testAccCheckIBMPIDeletePlacementGroup(name string, policy string, sapProfile string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_key" "key" {
			pi_cloud_instance_id = "%[1]s"
			pi_key_name          = "%[2]s"
			pi_ssh_key           = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDKUt7bk9yLBZFC187bQJFuLaBZONKFYjeIGCZj5mN0OvaJdJqPN2Mbx9Ui42Y5vrLE7SipG5c94BS/fYf7e2LvsQ+xaU1VQnMvP6XS8emoyKR6q/YzD60MkvkSopwTAgpyf6CpfCsKE5Yclbrsc1HIP16bjSgOapfgaVuEDXifn27i1fP1QRYhosY7YkfSKjyJQihxnFH1sONdl4JspJDC5rp8wZ4E7jSXyaZh6QIMbMBEvKoE8+/8CUgT3EWWndIOIMuPQtills3X3jDojTt722OBW1qETPahYDDEmN00R1Q1Q8V8pfVi1XG+ESLzY93gC8hV+/lWIoIvSEazwkfi7/5kludrZG1RhCGbOffGo3DkrmtqaBbKbjrTh/ZbY0GzHPXqccfW/KIhk6xlmoR0wF9LYPtFuzTkqnHF/tHi8EXPHI5XVv9m01kMLkoUqtWVXP2O7ZM7EwrJ+1TyJqLTrzbKMUbn52GqNuTSFJCAgEVc3XrvIRFjTL1/b428mS9JV5kCfRVLmDAUtPjuaQg1wmI/W97gZCF8IoF4JXWTEQP8IIb2opLxvEoBggsZpiFOtjsr9A914i/Tyd4T4KlvfkavJXqkzQoj29oZZPt10gt2ywwXPvV6usM1iofATB+YtX6vl8wUDaqvEyC8d4OTnSVkPZnFxTG3lhY4cDwa/w== tedfordt@us.ibm.com"
		}

		resource "ibm_pi_instance" "power_instance" {
			pi_processors         = "0.25"
			pi_proc_type          = "shared"
			pi_memory             = "2"
			pi_key_pair_name      = ibm_pi_key.key.key_id
			pi_image_id           = "%[4]s"
			pi_sys_type           = "e980"
			pi_instance_name      = "%[2]s"
			pi_cloud_instance_id  = "%[1]s"
			pi_storage_type       = "tier3"
			pi_network {
				network_id = "%[7]s"
			}
		}

		resource "ibm_pi_instance" "power_instance_in_pg" {
			pi_processors         = "0.25"
			pi_proc_type          = "shared"
			pi_memory             = "2"
			pi_key_pair_name      = ibm_pi_key.key.key_id
			pi_image_id           = "%[4]s"
			pi_sys_type           = "e980"
			pi_instance_name      = "%[2]s-2"
			pi_cloud_instance_id  = "%[1]s"
			pi_network {
				network_id = "%[7]s"
			}
		}

		resource "ibm_pi_instance" "sap_power_instance" {
			pi_cloud_instance_id  	= "%[1]s"
			pi_instance_name      	= "sap-%[2]s"
			pi_sap_profile_id       = "%[5]s"
			pi_image_id           	= "%[6]s"
			pi_storage_type			= "tier1"
			pi_network {
				network_id = "%[7]s"
			}
			depends_on = [    ibm_pi_instance.power_instance_in_pg  ]
		}

		resource "ibm_pi_placement_group" "power_placement_group_another" {
			pi_cloud_instance_id      = "%[1]s"
			pi_placement_group_name   = "%[2]s-2"
			pi_placement_group_policy = "%[3]s"
		}
	`, acc.Pi_cloud_instance_id, name, policy, acc.Pi_image, sapProfile, acc.Pi_sap_image, acc.Pi_network_name)
}
