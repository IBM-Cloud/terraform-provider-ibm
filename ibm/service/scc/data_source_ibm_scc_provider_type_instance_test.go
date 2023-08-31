// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSccProviderTypeInstanceDataSourceBasic(t *testing.T) {
	providerTypeInstanceItemProviderTypeID := fmt.Sprintf("tf_provider_type_id_%d", acctest.RandIntRange(10, 100))
	providerTypeInstanceItemName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSccProviderTypeInstanceDataSourceConfigBasic(providerTypeInstanceItemProviderTypeID, providerTypeInstanceItemName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_type_instance.scc_provider_type_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_type_instance.scc_provider_type_instance", "provider_type_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_type_instance.scc_provider_type_instance", "provider_type_instance_id"),
				),
			},
		},
	})
}

func TestAccIbmSccProviderTypeInstanceDataSourceAllArgs(t *testing.T) {
	providerTypeInstanceItemProviderTypeID := fmt.Sprintf("tf_provider_type_id_%d", acctest.RandIntRange(10, 100))
	providerTypeInstanceItemName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSccProviderTypeInstanceDataSourceConfig(providerTypeInstanceItemProviderTypeID, providerTypeInstanceItemName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_type_instance.scc_provider_type_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_type_instance.scc_provider_type_instance", "provider_type_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_type_instance.scc_provider_type_instance", "provider_type_instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_type_instance.scc_provider_type_instance", "provider_type_instance_item_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_type_instance.scc_provider_type_instance", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_type_instance.scc_provider_type_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_type_instance.scc_provider_type_instance", "attributes.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_type_instance.scc_provider_type_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_type_instance.scc_provider_type_instance", "updated_at"),
				),
			},
		},
	})
}

func testAccCheckIbmSccProviderTypeInstanceDataSourceConfigBasic(providerTypeInstanceItemProviderTypeID string, providerTypeInstanceItemName string) string {
	return fmt.Sprintf(`
		resource "ibm_scc_provider_type_instance" "scc_provider_type_instance_instance" {
			provider_type_id = "provider_type_id"
			name = "workload-protection-instance-1"
			attributes = {"wp_crn":"crn:v1:staging:public:sysdig-secure:eu-gb:a/14q5SEnVIbwxzvP4AWPCjr2dJg5BAvPb:d1461d1ae-df1eee12fa81812e0-12-aa259::"}
		}

		resource "ibm_scc_provider_type_instance" "scc_provider_type_instance_instance" {
			provider_type_id = "%s"
			name = "%s"
			attributes = {"wp_crn":"crn:v1:staging:public:sysdig-secure:eu-gb:a/14q5SEnVIbwxzvP4AWPCjr2dJg5BAvPb:d1461d1ae-df1eee12fa81812e0-12-aa259::"}
		}

		data "ibm_scc_provider_type_instance" "scc_provider_type_instance_instance" {
			provider_type_id = ibm_scc_provider_type_instance.scc_provider_type_instance.provider_type_id
			provider_type_instance_id = ibm_scc_provider_type_instance.scc_provider_type_instance_instance.providerTypeInstanceItem_id
		}
	`, providerTypeInstanceItemProviderTypeID, providerTypeInstanceItemName)
}

func testAccCheckIbmSccProviderTypeInstanceDataSourceConfig(providerTypeInstanceItemProviderTypeID string, providerTypeInstanceItemName string) string {
	return fmt.Sprintf(`
		resource "ibm_scc_provider_type_instance" "scc_provider_type_instance_instance" {
			provider_type_id = "provider_type_id"
			name = "workload-protection-instance-1"
			attributes = {"wp_crn":"crn:v1:staging:public:sysdig-secure:eu-gb:a/14q5SEnVIbwxzvP4AWPCjr2dJg5BAvPb:d1461d1ae-df1eee12fa81812e0-12-aa259::"}
		}

		resource "ibm_scc_provider_type_instance" "scc_provider_type_instance_instance" {
			provider_type_id = "%s"
			name = "%s"
			attributes = {"wp_crn":"crn:v1:staging:public:sysdig-secure:eu-gb:a/14q5SEnVIbwxzvP4AWPCjr2dJg5BAvPb:d1461d1ae-df1eee12fa81812e0-12-aa259::"}
		}

		data "ibm_scc_provider_type_instance" "scc_provider_type_instance_instance" {
			provider_type_id = ibm_scc_provider_type_instance.scc_provider_type_instance.provider_type_id
			provider_type_instance_id = ibm_scc_provider_type_instance.scc_provider_type_instance_instance.providerTypeInstanceItem_id
		}
	`, providerTypeInstanceItemProviderTypeID, providerTypeInstanceItemName)
}
