// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMPISPPbasic(t *testing.T) {
	name := fmt.Sprintf("tf_pi_spp_%d", acctest.RandIntRange(10, 100))
	sppRes := "ibm_pi_shared_processor_pool.power_shared_processor_pool"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPISPPDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPISPPConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPISPPExists(sppRes),
					resource.TestCheckResourceAttr(sppRes, "pi_shared_processor_pool_name", name),
				),
			},
		},
	})
}

func TestAccIBMPISPPUserTags(t *testing.T) {
	name := fmt.Sprintf("tf_pi_spp_%d", acctest.RandIntRange(10, 100))
	sppRes := "ibm_pi_shared_processor_pool.power_shared_processor_pool"
	userTagsString := `["env:dev","test_tag"]`
	userTagsStringUpdated := `["env:dev","test_tag","test_tag2"]`
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPISPPDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPISPPUserTagsConfig(name, userTagsString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPISPPExists(sppRes),
					resource.TestCheckResourceAttr(sppRes, "pi_shared_processor_pool_name", name),
					resource.TestCheckResourceAttr(sppRes, "pi_user_tags.#", "2"),
					resource.TestCheckTypeSetElemAttr(sppRes, "pi_user_tags.*", "env:dev"),
					resource.TestCheckTypeSetElemAttr(sppRes, "pi_user_tags.*", "test_tag"),
				),
			},
			{
				Config: testAccCheckIBMPISPPUserTagsConfig(name, userTagsStringUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPISPPExists(sppRes),
					resource.TestCheckResourceAttr(sppRes, "pi_shared_processor_pool_name", name),
					resource.TestCheckResourceAttr(sppRes, "pi_user_tags.#", "3"),
					resource.TestCheckTypeSetElemAttr(sppRes, "pi_user_tags.*", "env:dev"),
					resource.TestCheckTypeSetElemAttr(sppRes, "pi_user_tags.*", "test_tag"),
					resource.TestCheckTypeSetElemAttr(sppRes, "pi_user_tags.*", "test_tag2"),
				),
			},
		},
	})
}

func testAccCheckIBMPISPPDestroy(s *terraform.State) error {
	sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_shared_processor_pool" {
			continue
		}
		cloudInstanceID, sppID, err := splitID(rs.Primary.ID)
		if err != nil {
			return err
		}
		sppC := instance.NewIBMPISharedProcessorPoolClient(context.Background(), sess, cloudInstanceID)
		spp, err := sppC.Get(sppID)
		if err == nil {
			return fmt.Errorf("PI SPP still exists: %s", *spp.SharedProcessorPool.ID)
		}
	}

	return nil
}

func testAccCheckIBMPISPPExists(n string) resource.TestCheckFunc {
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
		cloudInstanceID, sppID, err := splitID(rs.Primary.ID)
		if err != nil {
			return err
		}
		client := instance.NewIBMPISharedProcessorPoolClient(context.Background(), sess, cloudInstanceID)

		_, err = client.Get(sppID)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckIBMPISPPConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_shared_processor_pool" "power_shared_processor_pool" {
			pi_cloud_instance_id                    = "%[2]s"
			pi_shared_processor_pool_host_group     = "s922"
			pi_shared_processor_pool_name           = "%[1]s"
			pi_shared_processor_pool_reserved_cores = "1"
		}`, name, acc.Pi_cloud_instance_id)
}

func testAccCheckIBMPISPPUserTagsConfig(name string, userTagsString string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_shared_processor_pool" "power_shared_processor_pool" {
			pi_cloud_instance_id                    = "%[2]s"
			pi_shared_processor_pool_host_group     = "s922"
			pi_shared_processor_pool_name           = "%[1]s"
			pi_shared_processor_pool_reserved_cores = "1"
			pi_user_tags                            = %[3]s
		}`, name, acc.Pi_cloud_instance_id, userTagsString)
}
