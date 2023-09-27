// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/scc-go-sdk/v5/securityandcompliancecenterapiv3"
)

func TestAccIbmSccProviderTypeInstanceBasic(t *testing.T) {
	var conf securityandcompliancecenterapiv3.ProviderTypeInstanceItem
	providerTypeAttributes := os.Getenv("IBMCLOUD_SCC_PROVIDER_TYPE_ATTRIBUTES")
	name := fmt.Sprintf("tf_provider_type_instance_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_provider_type_instance_name_%d", acctest.RandIntRange(10, 100))
	instanceID, ok := os.LookupEnv("IBMCLOUD_SCC_INSTANCE_ID")
	if !ok {
		t.Logf("Missing the env var IBMCLOUD_SCC_INSTANCE_ID.")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSccProviderTypeInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSccProviderTypeInstanceConfigBasic(instanceID, name, providerTypeAttributes),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSccProviderTypeInstanceExists("ibm_scc_provider_type_instance.scc_provider_type_instance_wlp", conf),
					resource.TestCheckResourceAttr("ibm_scc_provider_type_instance.scc_provider_type_instance_wlp", "name", name),
				),
			},
			{
				Config: testAccCheckIbmSccProviderTypeInstanceConfigBasic(instanceID, nameUpdate, providerTypeAttributes),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_scc_provider_type_instance.scc_provider_type_instance_wlp", "name", nameUpdate),
				),
			},
		},
	})
}

func TestAccIbmSccProviderTypeInstanceAllArgs(t *testing.T) {
	var conf securityandcompliancecenterapiv3.ProviderTypeInstanceItem
	providerTypeAttributes := os.Getenv("IBMCLOUD_SCC_PROVIDER_TYPE_ATTRIBUTES")
	name := fmt.Sprintf("tf_provider_type_instance_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_provider_type_instance_name_%d", acctest.RandIntRange(10, 100))

	instanceID, ok := os.LookupEnv("IBMCLOUD_SCC_INSTANCE_ID")
	if !ok {
		t.Logf("Missing the env var IBMCLOUD_SCC_INSTANCE_ID.")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSccProviderTypeInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSccProviderTypeInstanceConfig(instanceID, name, providerTypeAttributes),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSccProviderTypeInstanceExists("ibm_scc_provider_type_instance.scc_provider_type_instance_wlp", conf),
					resource.TestCheckResourceAttr("ibm_scc_provider_type_instance.scc_provider_type_instance_wlp", "name", name),
				),
			},
			{
				Config: testAccCheckIbmSccProviderTypeInstanceConfig(instanceID, nameUpdate, providerTypeAttributes),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_scc_provider_type_instance.scc_provider_type_instance_wlp", "name", nameUpdate),
				),
			},
			{
				ResourceName:      "ibm_scc_provider_type_instance.scc_provider_type_instance_wlp",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmSccProviderTypeInstanceConfigBasic(instanceID string, name string, attributes string) string {
	return fmt.Sprintf(`
		resource "ibm_scc_provider_type_instance" "scc_provider_type_instance_wlp" {
			instance_id = "%s"
			provider_type_id = "afa2476ecfa5f09af248492fe991b4d1"
			name = "%s"
			attributes = %s
		}
	`, instanceID, name, attributes)
}

func testAccCheckIbmSccProviderTypeInstanceConfig(instanceID string, name string, attributes string) string {
	return fmt.Sprintf(`
		resource "ibm_scc_provider_type_instance" "scc_provider_type_instance_wlp" {
			instance_id = "%s"
			provider_type_id = "afa2476ecfa5f09af248492fe991b4d1"
			name = "%s"
			attributes = %s
		}
	`, instanceID, name, attributes)
}

func testAccCheckIbmSccProviderTypeInstanceExists(n string, obj securityandcompliancecenterapiv3.ProviderTypeInstanceItem) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		securityAndComplianceCenterApIsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecurityAndComplianceCenterV3()
		if err != nil {
			return err
		}

		getProviderTypeInstanceOptions := &securityandcompliancecenterapiv3.GetProviderTypeInstanceOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getProviderTypeInstanceOptions.SetInstanceID(parts[0])
		getProviderTypeInstanceOptions.SetProviderTypeID(parts[1])
		getProviderTypeInstanceOptions.SetProviderTypeInstanceID(parts[2])

		providerTypeInstanceItem, _, err := securityAndComplianceCenterApIsClient.GetProviderTypeInstance(getProviderTypeInstanceOptions)
		if err != nil {
			return err
		}

		obj = *providerTypeInstanceItem
		return nil
	}
}

func testAccCheckIbmSccProviderTypeInstanceDestroy(s *terraform.State) error {
	securityAndComplianceCenterApIsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecurityAndComplianceCenterV3()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_scc_provider_type_instance" {
			continue
		}

		getProviderTypeInstanceOptions := &securityandcompliancecenterapiv3.GetProviderTypeInstanceOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getProviderTypeInstanceOptions.SetInstanceID(parts[0])
		getProviderTypeInstanceOptions.SetProviderTypeID(parts[1])
		getProviderTypeInstanceOptions.SetProviderTypeInstanceID(parts[2])

		// Try to find the key
		_, response, err := securityAndComplianceCenterApIsClient.GetProviderTypeInstance(getProviderTypeInstanceOptions)

		if err == nil {
			return fmt.Errorf("scc_provider_type_instance still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for scc_provider_type_instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
