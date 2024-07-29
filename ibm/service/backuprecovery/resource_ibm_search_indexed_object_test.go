// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func TestAccIbmSearchIndexedObjectBasic(t *testing.T) {
	var conf backuprecoveryv1.SearchIndexedObjectsOptions
	objectType := "Files"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSearchIndexedObjectDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSearchIndexedObjectConfigBasic(objectType),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSearchIndexedObjectExists("ibm_search_indexed_object.search_indexed_object_instance", conf),
					resource.TestCheckResourceAttr("ibm_search_indexed_object.search_indexed_object_instance", "object_type", objectType),
				),
			},
		},
	})
}

func TestAccIbmSearchIndexedObjectAllArgs(t *testing.T) {
	var conf backuprecoveryv1.SearchIndexedObjectsOptions
	tenantID := fmt.Sprintf("tf_tenant_id_%d", acctest.RandIntRange(10, 100))
	includeTenants := "false"
	paginationCookie := fmt.Sprintf("tf_pagination_cookie_%d", acctest.RandIntRange(10, 100))
	count := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	objectType := "Files"
	useCachedData := "false"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSearchIndexedObjectDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSearchIndexedObjectConfig(tenantID, includeTenants, paginationCookie, count, objectType, useCachedData),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSearchIndexedObjectExists("ibm_search_indexed_object.search_indexed_object_instance", conf),
					resource.TestCheckResourceAttr("ibm_search_indexed_object.search_indexed_object_instance", "tenant_id", tenantID),
					resource.TestCheckResourceAttr("ibm_search_indexed_object.search_indexed_object_instance", "include_tenants", includeTenants),
					resource.TestCheckResourceAttr("ibm_search_indexed_object.search_indexed_object_instance", "pagination_cookie", paginationCookie),
					resource.TestCheckResourceAttr("ibm_search_indexed_object.search_indexed_object_instance", "count", count),
					resource.TestCheckResourceAttr("ibm_search_indexed_object.search_indexed_object_instance", "object_type", objectType),
					resource.TestCheckResourceAttr("ibm_search_indexed_object.search_indexed_object_instance", "use_cached_data", useCachedData),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_search_indexed_object.search_indexed_object",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmSearchIndexedObjectConfigBasic(objectType string) string {
	return fmt.Sprintf(`
		resource "ibm_search_indexed_object" "search_indexed_object_instance" {
			object_type = "%s"
		}
	`, objectType)
}

func testAccCheckIbmSearchIndexedObjectConfig(tenantID string, includeTenants string, paginationCookie string, count string, objectType string, useCachedData string) string {
	return fmt.Sprintf(`

		resource "ibm_search_indexed_object" "search_indexed_object_instance" {
			protection_group_ids = "FIXME"
			storage_domain_ids = "FIXME"
			tenant_id = "%s"
			include_tenants = %s
			tags = "FIXME"
			snapshot_tags = "FIXME"
			must_have_tag_ids = "FIXME"
			might_have_tag_ids = "FIXME"
			must_have_snapshot_tag_ids = "FIXME"
			might_have_snapshot_tag_ids = "FIXME"
			pagination_cookie = "%s"
			count = %s
			object_type = "%s"
			use_cached_data = %s
			files {
				search_string = "search_string"
				types = [ "File" ]
				source_environments = [ "kSQL" ]
				source_ids = [ 1 ]
				object_ids = [ 1 ]
			}
			public_folders {
				search_string = "search_string"
				types = [ "Calendar" ]
				has_attachment = true
				sender_address = "sender_address"
				recipient_addresses = [ "recipientAddresses" ]
				cc_recipient_addresses = [ "ccRecipientAddresses" ]
				bcc_recipient_addresses = [ "bccRecipientAddresses" ]
				received_start_time_secs = 1
				received_end_time_secs = 1
			}
		}
	`, tenantID, includeTenants, paginationCookie, count, objectType, useCachedData)
}

func testAccCheckIbmSearchIndexedObjectExists(n string, obj backuprecoveryv1.SearchIndexedObjectsOptions) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		// rs, ok := s.RootModule().Resources[n]
		// if !ok {
		// 	return fmt.Errorf("Not found: %s", n)
		// }

		// backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
		// if err != nil {
		// 	return err
		// }

		// getRecoveryByIdOptions := &backuprecoveryv1.GetRecoveryByIdOptions{}

		// getRecoveryByIdOptions.SetID(rs.Primary.ID)

		// searchIndexedObjectsRequest, _, err := backupRecoveryClient.GetRecoveryByID(getRecoveryByIdOptions)
		// if err != nil {
		// 	return err
		// }

		// obj = *searchIndexedObjectsRequest
		return nil
	}
}

func testAccCheckIbmSearchIndexedObjectDestroy(s *terraform.State) error {
	backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_search_indexed_object" {
			continue
		}

		getRecoveryByIdOptions := &backuprecoveryv1.GetRecoveryByIdOptions{}

		getRecoveryByIdOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := backupRecoveryClient.GetRecoveryByID(getRecoveryByIdOptions)

		if err == nil {
			return fmt.Errorf("Common Search Indexed Objects Params still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for Common Search Indexed Objects Params (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
