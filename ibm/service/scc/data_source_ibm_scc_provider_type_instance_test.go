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

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSccProviderTypeInstanceDataSourceConfigBasic(providerTypeInstanceName, providerTypeInstanceAttributes),
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

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSccProviderTypeInstanceDataSourceConfig(providerTypeInstanceName, providerTypeInstanceAttributes),
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

func testAccCheckIbmSccProviderTypeInstanceDataSourceConfigBasic(providerTypeInstanceName string, providerTypeInstanceAttributes string) string {
	return fmt.Sprintf(`
		resource "ibm_scc_provider_type_instance" "scc_provider_type_instance" {
			provider_type_id = "afa2476ecfa5f09af248492fe991b4d1"
			name = "%s"
			attributes = %s
		}

		data "ibm_scc_provider_type_instance" "scc_provider_type_instance_tf" {
			provider_type_id = ibm_scc_provider_type_instance.scc_provider_type_instance.provider_type_id
			provider_type_instance_id = ibm_scc_provider_type_instance.scc_provider_type_instance.provider_type_instance_id
		}
	`, providerTypeInstanceName, providerTypeInstanceAttributes)
}

func testAccCheckIbmSccProviderTypeInstanceDataSourceConfig(providerTypeInstanceName string, providerTypeInstanceAttributes string) string {
	return fmt.Sprintf(`
		resource "ibm_scc_provider_type_instance" "scc_provider_type_instance" {
			provider_type_id = "afa2476ecfa5f09af248492fe991b4d1"
			name = "%s"
			attributes = %s
		}

		data "ibm_scc_provider_type_instance" "scc_provider_type_instance_tf" {
			provider_type_id = ibm_scc_provider_type_instance.scc_provider_type_instance.provider_type_id
			provider_type_instance_id = ibm_scc_provider_type_instance.scc_provider_type_instance.provider_type_instance_id
		}
	`, providerTypeInstanceName, providerTypeInstanceAttributes)
}
