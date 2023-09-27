// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSccProviderTypeInstanceDataSourceBasic(t *testing.T) {
	providerTypeInstanceName := fmt.Sprintf("tf_provider_type_instance_name_%d", acctest.RandIntRange(10, 100))
	providerTypeInstanceAttributes := os.Getenv("IBMCLOUD_SCC_PROVIDER_TYPE_ATTRIBUTES")
	instanceID, ok := os.LookupEnv("IBMCLOUD_SCC_INSTANCE_ID")
	if !ok {
		t.Logf("Missing the env var IBMCLOUD_SCC_INSTANCE_ID.")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckSccInstanceID(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSccProviderTypeInstanceDataSourceConfigBasic(instanceID, providerTypeInstanceName, providerTypeInstanceAttributes),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_type_instance.scc_provider_type_instance_tf", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_type_instance.scc_provider_type_instance_tf", "provider_type_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_type_instance.scc_provider_type_instance_tf", "provider_type_instance_id"),
				),
			},
		},
	})
}

func TestAccIbmSccProviderTypeInstanceDataSourceAllArgs(t *testing.T) {
	providerTypeInstanceName := fmt.Sprintf("tf_provider_type_instance_name_%d", acctest.RandIntRange(10, 100))
	providerTypeInstanceAttributes := os.Getenv("IBMCLOUD_SCC_PROVIDER_TYPE_ATTRIBUTES")
	instanceID, ok := os.LookupEnv("IBMCLOUD_SCC_INSTANCE_ID")
	if !ok {
		t.Logf("Missing the env var IBMCLOUD_SCC_INSTANCE_ID.")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckSccInstanceID(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSccProviderTypeInstanceDataSourceConfig(instanceID, providerTypeInstanceName, providerTypeInstanceAttributes),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_type_instance.scc_provider_type_instance_tf", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_type_instance.scc_provider_type_instance_tf", "provider_type_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_type_instance.scc_provider_type_instance_tf", "provider_type_instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_type_instance.scc_provider_type_instance_tf", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_type_instance.scc_provider_type_instance_tf", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_type_instance.scc_provider_type_instance_tf", "attributes.%"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_type_instance.scc_provider_type_instance_tf", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_type_instance.scc_provider_type_instance_tf", "updated_at"),
				),
			},
		},
	})
}

func testAccCheckIbmSccProviderTypeInstanceDataSourceConfigBasic(instanceID string, providerTypeInstanceName string, providerTypeInstanceAttributes string) string {
	return fmt.Sprintf(`
		resource "ibm_scc_provider_type_instance" "scc_provider_type_instance" {
			instance_id = "%s"
			provider_type_id = "afa2476ecfa5f09af248492fe991b4d1"
			name = "%s"
			attributes = %s
		}

		data "ibm_scc_provider_type_instance" "scc_provider_type_instance_tf" {
			instance_id = resource.ibm_scc_provider_type_instance.scc_provider_type_instance.instance_id
			provider_type_id = ibm_scc_provider_type_instance.scc_provider_type_instance.provider_type_id
			provider_type_instance_id = ibm_scc_provider_type_instance.scc_provider_type_instance.provider_type_instance_id
		}
	`, instanceID, providerTypeInstanceName, providerTypeInstanceAttributes)
}

func testAccCheckIbmSccProviderTypeInstanceDataSourceConfig(instanceID string, providerTypeInstanceName string, providerTypeInstanceAttributes string) string {
	return fmt.Sprintf(`
		resource "ibm_scc_provider_type_instance" "scc_provider_type_instance" {
			instance_id = "%s"
			provider_type_id = "afa2476ecfa5f09af248492fe991b4d1"
			name = "%s"
			attributes = %s
		}

		data "ibm_scc_provider_type_instance" "scc_provider_type_instance_tf" {
			instance_id = resource.ibm_scc_provider_type_instance.scc_provider_type_instance.instance_id
			provider_type_id = ibm_scc_provider_type_instance.scc_provider_type_instance.provider_type_id
			provider_type_instance_id = ibm_scc_provider_type_instance.scc_provider_type_instance.provider_type_instance_id
		}
	`, instanceID, providerTypeInstanceName, providerTypeInstanceAttributes)
}
