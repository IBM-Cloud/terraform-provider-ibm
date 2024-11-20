// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"context"
	"errors"
	"fmt"
	"regexp"
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

func TestAccIBMPISPPPlacementGroupBasic(t *testing.T) {
	name := fmt.Sprintf("tfspp%d", acctest.RandIntRange(10, 100))
	policy := "affinity"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPISPPPlacementGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMPICreateSAPInstanceWithSPP(name, policy),
				ExpectError: regexp.MustCompile("\"pi_shared_processor_pool\": conflicts with pi_sap_profile_id"),
			},
			{
				Config:       testAccCheckIBMPISPPPlacementGroupConfig(name, policy),
				ResourceName: "ibm_pi_spp_placement_group.spp_placement_group",
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPISPPPlacementGroupExists("ibm_pi_spp_placement_group.spp_placement_group"),
					resource.TestCheckResourceAttr(
						"ibm_pi_spp_placement_group.spp_placement_group", "pi_spp_placement_group_name", name+"pg"),
					resource.TestCheckResourceAttr(
						"ibm_pi_spp_placement_group.spp_placement_group", "pi_spp_placement_group_policy", policy),
				),
			},
			{
				Config: testAccCheckIBMPISPPPlacementGroupAddMemberConfig(name, policy),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"ibm_pi_shared_processor_pool.spp_pool", "pi_shared_processor_pool_placement_group_id"),
					testAccCheckIBMPISPPPlacementGroupMemberExists("ibm_pi_spp_placement_group.spp_placement_group", "ibm_pi_shared_processor_pool.spp_pool"),
				),
			},
			{
				Config: testAccCheckIBMPISPPPlacementGroupUpdateMemberConfig(name, policy),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"ibm_pi_shared_processor_pool.spp_pool", "spp_placement_groups.0"),
					testAccCheckIBMPISPPPlacementGroupMemberExists("ibm_pi_spp_placement_group.spp_placement_group_another", "ibm_pi_shared_processor_pool.spp_pool"),
				),
			},
			{
				Config: testAccCheckIBMPISPPPlacementGroupRemoveMemberConfig(name, policy),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPISPPPlacementGroupMemberDoesNotExist("ibm_pi_spp_placement_group.spp_placement_group", "ibm_pi_shared_processor_pool.spp_pool"),
				),
			},
			{
				Config: testAccCheckIBMPICreateSPPInPlacementGroup(name, policy),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr(
						"ibm_pi_shared_processor_pool.spp_pool", "spp_placement_groups.0"),
					resource.TestCheckResourceAttrSet(
						"ibm_pi_shared_processor_pool.spp_pool_2", "pi_shared_processor_pool_placement_group_id"),
					testAccCheckIBMPISPPPlacementGroupMemberExistsFromSPPCreate("ibm_pi_spp_placement_group.spp_placement_group", "ibm_pi_shared_processor_pool.spp_pool_2"),
				),
			},
			{
				Config: testAccCheckIBMPIDeleteSPPPlacementGroup(name, policy),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPISPPPlacementGroupDelete("ibm_pi_spp_placement_group.spp_placement_group_another", "ibm_pi_shared_processor_pool.spp_pool", "ibm_pi_shared_processor_pool.spp_pool_2"),
				),
			},
			{
				Config: testAccCheckIBMPICreateInstanceWithSPP(name, policy),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIInstanceInSPP("ibm_pi_shared_processor_pool.spp_pool", "ibm_pi_instance.power_instance"),
				),
			},
		},
	})
}

func testAccCheckIBMPISPPPlacementGroupDestroy(s *terraform.State) error {

	sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_spp_placement_group" {
			continue
		}
		parts, _ := flex.IdParts(rs.Primary.ID)
		cloudpoolid := parts[0]
		placementGroupC := st.NewIBMPISPPPlacementGroupClient(context.Background(), sess, cloudpoolid)
		_, err = placementGroupC.Get(parts[1])
		if err == nil {
			return fmt.Errorf("PI SPP placement group still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}
func testAccCheckIBMPISPPPlacementGroupExists(n string) resource.TestCheckFunc {
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
		cloudpoolid := parts[0]
		client := st.NewIBMPISPPPlacementGroupClient(context.Background(), sess, cloudpoolid)

		placementGroup, err := client.Get(parts[1])
		if err != nil {
			return err
		}
		parts[1] = *placementGroup.ID
		return nil
	}
}

func testAccCheckIBMPISPPPlacementGroupMemberExists(sppPG string, pool string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		pgResource, ok := s.RootModule().Resources[sppPG]

		if !ok {
			return fmt.Errorf("Not found: %s", sppPG)
		}

		if pgResource.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		// refresh placement group info since a spp should be in the placement group
		sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
		if err != nil {
			return err
		}
		parts, err := flex.IdParts(pgResource.Primary.ID)
		if err != nil {
			return err
		}
		cloudpoolid := parts[0]
		client := st.NewIBMPISPPPlacementGroupClient(context.Background(), sess, cloudpoolid)

		pg, err := client.Get(parts[1])
		if err != nil {
			return err
		}

		poolResource, ok := s.RootModule().Resources[pool]
		if !ok {
			return fmt.Errorf("Not found: %s", pool)
		}
		poolName := poolResource.Primary.Attributes["pi_shared_processor_pool_name"]

		var isPoolFound bool = false
		for _, x := range pg.MemberSharedProcessorPools {
			if x == poolName {
				isPoolFound = true
				break
			}
		}
		if !isPoolFound {
			return fmt.Errorf("Expected pool ID %s in the PG members field but found %s", poolName, strings.Join(pg.MemberSharedProcessorPools[:], ","))
		}
		return nil
	}
}

func testAccCheckIBMPISPPPlacementGroupMemberDoesNotExist(n string, pool string) resource.TestCheckFunc {
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
		cloudpoolid := parts[0]
		client := st.NewIBMPISPPPlacementGroupClient(context.Background(), sess, cloudpoolid)

		pg, err := client.Get(parts[1])
		if err != nil {
			return err
		}

		poolrs, ok := s.RootModule().Resources[pool]
		if !ok {
			return fmt.Errorf("Not found: %s", pool)
		}
		instanccParts, err := flex.IdParts(poolrs.Primary.ID)
		if err != nil {
			return err
		}
		if len(pg.MemberSharedProcessorPools) > 0 {
			return fmt.Errorf("Expected pool ID %s to be removed so that the PG members field is empty but foumd %s", instanccParts[1], pg.MemberSharedProcessorPools[0])
		}

		return nil
	}
}

func containsMemberPool(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func testAccCheckIBMPISPPPlacementGroupMemberExistsFromSPPCreate(n string, pool string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		// refresh placement group info since a pool should be in the placement group
		sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
		if err != nil {
			return err
		}
		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		cloudpoolid := parts[0]
		client := st.NewIBMPISPPPlacementGroupClient(context.Background(), sess, cloudpoolid)

		pg, err := client.Get(parts[1])
		if err != nil {
			return err
		}

		poolrs, ok := s.RootModule().Resources[pool]
		if !ok {
			return fmt.Errorf("Not found: %s", pool)
		}
		poolName := poolrs.Primary.Attributes["pi_shared_processor_pool_name"]

		if !containsMemberPool(pg.MemberSharedProcessorPools, poolName) {
			return fmt.Errorf("Expected pool %s in the PG members field", poolName)
		}
		return nil
	}
}

func testAccCheckIBMPISPPPlacementGroupDelete(n string, pool string, newPool string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
		if err != nil {
			return err
		}

		poolrs, ok := s.RootModule().Resources[pool]
		if !ok {
			return fmt.Errorf("Not found: %s", pool)
		}
		poolParts, err := flex.IdParts(poolrs.Primary.ID)
		if err != nil {
			return err
		}

		newpoolrs, ok := s.RootModule().Resources[newPool]
		if !ok {
			return fmt.Errorf("Not found: %s", newPool)
		}
		newpoolParts, err := flex.IdParts(newpoolrs.Primary.ID)
		if err != nil {
			return err
		}
		cloudpoolid := poolParts[0]
		spp_client := st.NewIBMPISharedProcessorPoolClient(context.Background(), sess, cloudpoolid)

		pool, err := spp_client.Get(poolParts[1])
		if err != nil {
			return err
		}

		if len(pool.SharedProcessorPool.SharedProcessorPoolPlacementGroups) > 0 {
			return fmt.Errorf("Expected no spp placement group ID in the spp placement groups array but found %s", *pool.SharedProcessorPool.SharedProcessorPoolPlacementGroups[0].ID)
		}
		newpool, err := spp_client.Get(newpoolParts[1])
		if err != nil {
			return err
		}
		if len(newpool.SharedProcessorPool.SharedProcessorPoolPlacementGroups) > 0 {
			return fmt.Errorf("Expected no spp placement group ID in the spp placement groups array but found %s", *newpool.SharedProcessorPool.SharedProcessorPoolPlacementGroups[0].ID)
		}
		return nil
	}
}

func testAccCheckIBMPIInstanceInSPP(spp string, instance string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		sppResource, ok := s.RootModule().Resources[spp]

		if !ok {
			return fmt.Errorf("Not found: %s", spp)
		}

		if sppResource.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		// refresh shared processor pool info since a instance should be in the spp
		sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
		if err != nil {
			return err
		}
		parts, err := flex.IdParts(sppResource.Primary.ID)
		if err != nil {
			return err
		}
		cloudpoolid := parts[0]
		client := st.NewIBMPISharedProcessorPoolClient(context.Background(), sess, cloudpoolid)

		sppFromSB, err := client.Get(parts[1])
		if err != nil {
			return err
		}

		instanceResource, ok := s.RootModule().Resources[instance]
		if !ok {
			return fmt.Errorf("Instance not found: %s", instance)
		}
		instanceName := instanceResource.Primary.Attributes["pi_instance_name"]

		var isInstanceFoundInSPPServersList bool = false
		for _, s := range sppFromSB.Servers {
			if s.Name == instanceName {
				isInstanceFoundInSPPServersList = true
				break
			}
		}
		if !isInstanceFoundInSPPServersList {
			return fmt.Errorf("Expected instance name %s in the SPP servers object but found %v", instanceName, sppFromSB.Servers)
		}
		return nil
	}
}

func testAccCheckIBMPISPPPlacementGroupConfig(name string, policy string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_shared_processor_pool" "spp_pool" {
			pi_cloud_instance_id  = "%[1]s"
			pi_shared_processor_pool_name = "%[2]s"
			pi_shared_processor_pool_host_group       = "s922"
			pi_shared_processor_pool_reserved_cores = "2"
		}
		resource "ibm_pi_spp_placement_group" "spp_placement_group" {
			pi_cloud_instance_id      = "%[1]s"
			pi_spp_placement_group_name   = "%[2]spg"
			pi_spp_placement_group_policy = "%[3]s"
		}`, acc.Pi_cloud_instance_id, name, policy)
}

func testAccCheckIBMPISPPPlacementGroupAddMemberConfig(name string, policy string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_shared_processor_pool" "spp_pool" {
			pi_cloud_instance_id  = "%[1]s"
			pi_shared_processor_pool_name = "%[2]s"
			pi_shared_processor_pool_host_group       = "s922"
			pi_shared_processor_pool_reserved_cores = "2"
			pi_shared_processor_pool_placement_group_id = ibm_pi_spp_placement_group.spp_placement_group.spp_placement_group_id
			spp_placement_groups = [ibm_pi_spp_placement_group.spp_placement_group.spp_placement_group_id]
		}
		resource "ibm_pi_spp_placement_group" "spp_placement_group" {
			pi_cloud_instance_id      = "%[1]s"
			pi_spp_placement_group_name   = "%[2]spg"
			pi_spp_placement_group_policy = "%[3]s"
		}`, acc.Pi_cloud_instance_id, name, policy)
}

func testAccCheckIBMPISPPPlacementGroupUpdateMemberConfig(name string, policy string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_shared_processor_pool" "spp_pool" {
			pi_cloud_instance_id  = "%[1]s"
			pi_shared_processor_pool_name = "%[2]s"
			pi_shared_processor_pool_host_group       = "s922"
			pi_shared_processor_pool_reserved_cores = "2"
			spp_placement_groups = [ibm_pi_spp_placement_group.spp_placement_group_another.spp_placement_group_id]
		}
		resource "ibm_pi_spp_placement_group" "spp_placement_group" {
			pi_cloud_instance_id      = "%[1]s"
			pi_spp_placement_group_name   = "%[2]spg"
			pi_spp_placement_group_policy = "%[3]s"
		}
		resource "ibm_pi_spp_placement_group" "spp_placement_group_another" {
			pi_cloud_instance_id      = "%[1]s"
			pi_spp_placement_group_name   = "%[2]s2pg"
			pi_spp_placement_group_policy = "%[3]s"
		}`, acc.Pi_cloud_instance_id, name, policy)
}

func testAccCheckIBMPISPPPlacementGroupRemoveMemberConfig(name string, policy string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_shared_processor_pool" "spp_pool" {
			pi_cloud_instance_id  = "%[1]s"
			pi_shared_processor_pool_name = "%[2]s"
			pi_shared_processor_pool_host_group       = "s922"
			pi_shared_processor_pool_reserved_cores = "2"
			spp_placement_groups = []
		}
		resource "ibm_pi_spp_placement_group" "spp_placement_group" {
			pi_cloud_instance_id      = "%[1]s"
			pi_spp_placement_group_name   = "%[2]spg"
			pi_spp_placement_group_policy = "%[3]s"
		}
		resource "ibm_pi_spp_placement_group" "spp_placement_group_another" {
			pi_cloud_instance_id      = "%[1]s"
			pi_spp_placement_group_name   = "%[2]s2pg"
			pi_spp_placement_group_policy = "%[3]s"
		}`, acc.Pi_cloud_instance_id, name, policy)
}

func testAccCheckIBMPICreateSPPInPlacementGroup(name string, policy string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_shared_processor_pool" "spp_pool" {
			pi_cloud_instance_id  = "%[1]s"
			pi_shared_processor_pool_name = "%[2]s"
			pi_shared_processor_pool_host_group       = "s922"
			pi_shared_processor_pool_reserved_cores = "2"
			spp_placement_groups = []
		}
		resource "ibm_pi_shared_processor_pool" "spp_pool_2" {
			pi_cloud_instance_id  = "%[1]s"
			pi_shared_processor_pool_name = "%[2]s2"
			pi_shared_processor_pool_host_group       = "e980"
			pi_shared_processor_pool_reserved_cores = "1"
			pi_shared_processor_pool_placement_group_id = ibm_pi_spp_placement_group.spp_placement_group.spp_placement_group_id
			spp_placement_groups = [ibm_pi_spp_placement_group.spp_placement_group.spp_placement_group_id]
		}
		resource "ibm_pi_spp_placement_group" "spp_placement_group" {
			pi_cloud_instance_id      = "%[1]s"
			pi_spp_placement_group_name   = "%[2]spg"
			pi_spp_placement_group_policy = "%[3]s"
		}
		resource "ibm_pi_spp_placement_group" "spp_placement_group_another" {
			pi_cloud_instance_id      = "%[1]s"
			pi_spp_placement_group_name   = "%[2]s2pg"
			pi_spp_placement_group_policy = "%[3]s"
		}`, acc.Pi_cloud_instance_id, name, policy)
}

func testAccCheckIBMPIDeleteSPPPlacementGroup(name string, policy string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_shared_processor_pool" "spp_pool" {
			pi_cloud_instance_id  = "%[1]s"
			pi_shared_processor_pool_name = "%[2]s"
			pi_shared_processor_pool_host_group       = "s922"
			pi_shared_processor_pool_reserved_cores = "2"
		}
		resource "ibm_pi_shared_processor_pool" "spp_pool_2" {
			pi_cloud_instance_id  = "%[1]s"
			pi_shared_processor_pool_name = "%[2]s2"
			pi_shared_processor_pool_host_group       = "e980"
			pi_shared_processor_pool_reserved_cores = "1"
		}
		resource "ibm_pi_spp_placement_group" "spp_placement_group" {
			pi_cloud_instance_id      = "%[1]s"
			pi_spp_placement_group_name   = "%[2]spg"
			pi_spp_placement_group_policy = "%[3]s"
		}
		resource "ibm_pi_spp_placement_group" "spp_placement_group_another" {
			pi_cloud_instance_id      = "%[1]s"
			pi_spp_placement_group_name   = "%[2]s2pg"
			pi_spp_placement_group_policy = "%[3]s"
		}`, acc.Pi_cloud_instance_id, name, policy)
}

func testAccCheckIBMPICreateInstanceWithSPP(name string, policy string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_shared_processor_pool" "spp_pool" {
			pi_cloud_instance_id  = "%[1]s"
			pi_shared_processor_pool_name = "%[2]s"
			pi_shared_processor_pool_host_group       = "s922"
			pi_shared_processor_pool_reserved_cores = "2"
		}
		resource "ibm_pi_shared_processor_pool" "spp_pool_2" {
			pi_cloud_instance_id  = "%[1]s"
			pi_shared_processor_pool_name = "%[2]s2"
			pi_shared_processor_pool_host_group       = "e980"
			pi_shared_processor_pool_reserved_cores = "1"
		}
		resource "ibm_pi_spp_placement_group" "spp_placement_group" {
			pi_cloud_instance_id      = "%[1]s"
			pi_spp_placement_group_name   = "%[2]spg"
			pi_spp_placement_group_policy = "%[3]s"
		}
		resource "ibm_pi_spp_placement_group" "spp_placement_group_another" {
			pi_cloud_instance_id      = "%[1]s"
			pi_spp_placement_group_name   = "%[2]s2pg"
			pi_spp_placement_group_policy = "%[3]s"
		}

		resource "ibm_pi_key" "key" {
			pi_cloud_instance_id = "%[1]s"
			pi_key_name          = "%[2]s"
			pi_ssh_key           = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR"
		}
		data "ibm_pi_image" "power_image" {
			pi_image_name        = "%[4]s"
			pi_cloud_instance_id = "%[1]s"
		}
		data "ibm_pi_network" "power_networks" {
			pi_cloud_instance_id = "%[1]s"
			pi_network_name      = "%[5]s"
		}
		resource "ibm_pi_volume" "power_volume" {
			pi_volume_size       = 20
			pi_volume_name       = "%[2]s"
			pi_volume_shareable  = true
			pi_volume_pool       = data.ibm_pi_image.power_image.storage_pool
			pi_cloud_instance_id = "%[1]s"
		}
		resource "ibm_pi_instance" "power_instance" {
			pi_memory             = "2"
			pi_processors         = "0.25"
			pi_instance_name      = "%[2]s"
			pi_proc_type          = "shared"
			pi_image_id           = data.ibm_pi_image.power_image.id
			pi_sys_type           = "s922"
			pi_cloud_instance_id  = "%[1]s"
			pi_storage_pool       = data.ibm_pi_image.power_image.storage_pool
			pi_volume_ids         = [ibm_pi_volume.power_volume.volume_id]
			pi_network {
				network_id = data.ibm_pi_network.power_networks.id
			}
			pi_shared_processor_pool = ibm_pi_shared_processor_pool.spp_pool.pi_shared_processor_pool_name
		}`, acc.Pi_cloud_instance_id, name, policy, acc.Pi_image, acc.Pi_network_name)
}

func testAccCheckIBMPICreateSAPInstanceWithSPP(name string, policy string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_shared_processor_pool" "spp_pool" {
			pi_cloud_instance_id  = "%[1]s"
			pi_shared_processor_pool_name = "%[2]s"
			pi_shared_processor_pool_host_group       = "s922"
			pi_shared_processor_pool_reserved_cores = "2"
		}
		resource "ibm_pi_shared_processor_pool" "spp_pool_2" {
			pi_cloud_instance_id  = "%[1]s"
			pi_shared_processor_pool_name = "%[2]s2"
			pi_shared_processor_pool_host_group       = "e980"
			pi_shared_processor_pool_reserved_cores = "1"
		}
		resource "ibm_pi_spp_placement_group" "spp_placement_group" {
			pi_cloud_instance_id      = "%[1]s"
			pi_spp_placement_group_name   = "%[2]spg"
			pi_spp_placement_group_policy = "%[3]s"
		}
		resource "ibm_pi_spp_placement_group" "spp_placement_group_another" {
			pi_cloud_instance_id      = "%[1]s"
			pi_spp_placement_group_name   = "%[2]s2pg"
			pi_spp_placement_group_policy = "%[3]s"
		}

		resource "ibm_pi_key" "key" {
			pi_cloud_instance_id = "%[1]s"
			pi_key_name          = "%[2]s"
			pi_ssh_key           = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR"
		}
		data "ibm_pi_image" "power_image" {
			pi_image_name        = "%[4]s"
			pi_cloud_instance_id = "%[1]s"
		}
		data "ibm_pi_network" "power_networks" {
			pi_cloud_instance_id = "%[1]s"
			pi_network_name      = "%[5]s"
		}
		resource "ibm_pi_volume" "power_volume" {
			pi_volume_size       = 20
			pi_volume_name       = "%[2]s"
			pi_volume_shareable  = true
			pi_volume_pool       = data.ibm_pi_image.power_image.storage_pool
			pi_cloud_instance_id = "%[1]s"
		}
		resource "ibm_pi_instance" "power_instance" {
			pi_memory             = "2"
			pi_processors         = "0.25"
			pi_instance_name      = "%[2]s"
			pi_proc_type          = "shared"
			pi_image_id           = data.ibm_pi_image.power_image.id
			pi_sys_type           = "s922"
			pi_cloud_instance_id  = "%[1]s"
			pi_storage_pool       = data.ibm_pi_image.power_image.storage_pool
			pi_volume_ids         = [ibm_pi_volume.power_volume.volume_id]
			pi_network {
				network_id = data.ibm_pi_network.power_networks.id
			}
			pi_shared_processor_pool = ibm_pi_shared_processor_pool.spp_pool.pi_shared_processor_pool_name
		}
		resource "ibm_pi_instance" "sap" {
			pi_cloud_instance_id  	= "%[1]s"
			pi_instance_name      	= "%[2]sSAP"
			pi_sap_profile_id       = "%[7]s"
			pi_image_id           	= "%[6]s"
			pi_storage_type			= "tier1"
			pi_network {
				network_id = data.ibm_pi_network.power_networks.id
			}
			pi_health_status		= "OK"
			pi_shared_processor_pool = ibm_pi_shared_processor_pool.spp_pool_2.pi_shared_processor_pool_name
		}`, acc.Pi_cloud_instance_id, name, policy, acc.Pi_image, acc.Pi_network_name, acc.Pi_sap_image, acc.PiSAPProfileID)
}
