// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"testing"
	"time"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/scc-go-sdk/v3/posturemanagementv2"
)

func TestAccIBMSccPostureCollectorsBasic(t *testing.T) {
	var conf posturemanagementv2.Collector
	name := fmt.Sprintf("tf_name_%d", time.Now().UnixNano())
	isPublic := "true"
	managedBy := "customer"
	nameUpdate := fmt.Sprintf("tf_name_%d", time.Now().UnixNano())
	isPublicUpdate := "true"
	managedByUpdate := "customer"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSccPostureCollectorsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSccPostureCollectorsConfigBasic(name, isPublic, managedBy),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSccPostureCollectorsExists("ibm_scc_posture_collector.collectors", conf),
					resource.TestCheckResourceAttr("ibm_scc_posture_collector.collectors", "name", name),
					resource.TestCheckResourceAttr("ibm_scc_posture_collector.collectors", "is_public", isPublic),
					resource.TestCheckResourceAttr("ibm_scc_posture_collector.collectors", "managed_by", managedBy),
				),
			},
			{
				Config: testAccCheckIBMSccPostureCollectorsConfigBasic(nameUpdate, isPublicUpdate, managedByUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_scc_posture_collector.collectors", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_scc_posture_collector.collectors", "is_public", isPublicUpdate),
					resource.TestCheckResourceAttr("ibm_scc_posture_collector.collectors", "managed_by", managedByUpdate),
				),
			},
		},
	})
}

func TestAccIBMCollectorsAllArgs(t *testing.T) {
	var conf posturemanagementv2.Collector
	name := fmt.Sprintf("tf_name_%d", time.Now().UnixNano())
	isPublic := "false"
	managedBy := "ibm"
	description := fmt.Sprintf("tf_description_%d", time.Now().UnixNano())
	passphrase := fmt.Sprintf("tf_passphrase_%d", time.Now().UnixNano())
	isUbiImage := "true"
	nameUpdate := fmt.Sprintf("tf_name_%d", time.Now().UnixNano())
	isPublicUpdate := "true"
	managedByUpdate := "customer"
	descriptionUpdate := fmt.Sprintf("tf_description_%d", time.Now().UnixNano())
	passphraseUpdate := fmt.Sprintf("tf_passphrase_%d", time.Now().UnixNano())
	isUbiImageUpdate := "false"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSccPostureCollectorsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSccPostureCollectorsConfig(name, isPublic, managedBy, description, passphrase, isUbiImage),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSccPostureCollectorsExists("ibm_scc_posture_collector.collectors", conf),
					resource.TestCheckResourceAttr("ibm_scc_posture_collector.collectors", "name", name),
					resource.TestCheckResourceAttr("ibm_scc_posture_collector.collectors", "is_public", isPublic),
					resource.TestCheckResourceAttr("ibm_scc_posture_collector.collectors", "managed_by", managedBy),
					resource.TestCheckResourceAttr("ibm_scc_posture_collector.collectors", "description", description),
					resource.TestCheckResourceAttr("ibm_scc_posture_collector.collectors", "passphrase", passphrase),
					resource.TestCheckResourceAttr("ibm_scc_posture_collector.collectors", "is_ubi_image", isUbiImage),
				),
			},
			{
				Config: testAccCheckIBMSccPostureCollectorsConfig(nameUpdate, isPublicUpdate, managedByUpdate, descriptionUpdate, passphraseUpdate, isUbiImageUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_scc_posture_collector.collectors", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_scc_posture_collector.collectors", "is_public", isPublicUpdate),
					resource.TestCheckResourceAttr("ibm_scc_posture_collector.collectors", "managed_by", managedByUpdate),
					resource.TestCheckResourceAttr("ibm_scc_posture_collector.collectors", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_scc_posture_collector.collectors", "passphrase", passphraseUpdate),
					resource.TestCheckResourceAttr("ibm_scc_posture_collector.collectors", "is_ubi_image", isUbiImageUpdate),
				),
			},
			{
				ResourceName:      "ibm_scc_posture_collector.collectors",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMSccPostureCollectorsConfigBasic(name string, isPublic string, managedBy string) string {
	return fmt.Sprintf(`

		resource "ibm_scc_posture_collector" "collectors" {
			name = "%s"
			is_public = %s
			managed_by = "%s"
		}
	`, name, isPublic, managedBy)
}

func testAccCheckIBMSccPostureCollectorsConfig(name string, isPublic string, managedBy string, description string, passphrase string, isUbiImage string) string {
	return fmt.Sprintf(`

		resource "ibm_scc_posture_collector" "collectors" {
			name = "%s"
			is_public = %s
			managed_by = "%s"
			description = "%s"
			passphrase = "%s"
			is_ubi_image = %s
		}
	`, name, isPublic, managedBy, description, passphrase, isUbiImage)
}

func testAccCheckIBMSccPostureCollectorsExists(n string, obj posturemanagementv2.Collector) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		postureManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PostureManagementV2()
		if err != nil {
			return err
		}

		listCollectorsOptions := &posturemanagementv2.ListCollectorsOptions{}

		userDetails, err := acc.TestAccProvider.Meta().(conns.ClientSession).BluemixUserDetails()
		if err != nil {
			return err
		}

		listCollectorsOptions.SetAccountID(userDetails.UserAccount)

		newCollector, _, err := postureManagementClient.ListCollectors(listCollectorsOptions)
		if err != nil {
			return err
		}
		fmt.Println(rs)
		obj = (newCollector.Collectors[0])
		return nil
	}
}

func testAccCheckIBMSccPostureCollectorsDestroy(s *terraform.State) error {
	postureManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PostureManagementV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_scc_posture_collector" {
			continue
		}

		listCollectorsOptions := &posturemanagementv2.ListCollectorsOptions{}

		userDetails, err := acc.TestAccProvider.Meta().(conns.ClientSession).BluemixUserDetails()
		if err != nil {
			return err
		}

		listCollectorsOptions.SetAccountID(userDetails.UserAccount)

		// Try to find the key
		_, response, err := postureManagementClient.ListCollectors(listCollectorsOptions)

		if err == nil {
			return err //fmt.Errorf("collectors still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("[ERROR] Error checking for collectors (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
