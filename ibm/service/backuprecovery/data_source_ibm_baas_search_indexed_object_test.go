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

func TestAccIbmBaasSearchIndexedObjectDataSourceBasic(t *testing.T) {
	var conf backuprecoveryv1.SearchIndexedObjectsResponse
	xIbmTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	objectType := "Emails"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBaasSearchIndexedObjectDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasSearchIndexedObjectConfigBasic(xIbmTenantID, objectType),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBaasSearchIndexedObjectExists("ibm_baas_search_indexed_object.baas_search_indexed_object_instance", conf),
					resource.TestCheckResourceAttr("ibm_baas_search_indexed_object.baas_search_indexed_object_instance", "x_ibm_tenant_id", xIbmTenantID),
					resource.TestCheckResourceAttr("ibm_baas_search_indexed_object.baas_search_indexed_object_instance", "object_type", objectType),
				),
			},
		},
	})
}

func TestAccIbmBaasSearchIndexedObjectAllArgs(t *testing.T) {
	var conf backuprecoveryv1.SearchIndexedObjectsResponse
	xIbmTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	tenantID := fmt.Sprintf("tf_tenant_id_%d", acctest.RandIntRange(10, 100))
	includeTenants := "false"
	paginationCookie := fmt.Sprintf("tf_pagination_cookie_%d", acctest.RandIntRange(10, 100))
	count := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	objectType := "Emails"
	useCachedData := "false"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBaasSearchIndexedObjectDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasSearchIndexedObjectConfig(xIbmTenantID, tenantID, includeTenants, paginationCookie, count, objectType, useCachedData),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBaasSearchIndexedObjectExists("ibm_baas_search_indexed_object.baas_search_indexed_object_instance", conf),
					resource.TestCheckResourceAttr("ibm_baas_search_indexed_object.baas_search_indexed_object_instance", "x_ibm_tenant_id", xIbmTenantID),
					resource.TestCheckResourceAttr("ibm_baas_search_indexed_object.baas_search_indexed_object_instance", "tenant_id", tenantID),
					resource.TestCheckResourceAttr("ibm_baas_search_indexed_object.baas_search_indexed_object_instance", "include_tenants", includeTenants),
					resource.TestCheckResourceAttr("ibm_baas_search_indexed_object.baas_search_indexed_object_instance", "pagination_cookie", paginationCookie),
					resource.TestCheckResourceAttr("ibm_baas_search_indexed_object.baas_search_indexed_object_instance", "count", count),
					resource.TestCheckResourceAttr("ibm_baas_search_indexed_object.baas_search_indexed_object_instance", "object_type", objectType),
					resource.TestCheckResourceAttr("ibm_baas_search_indexed_object.baas_search_indexed_object_instance", "use_cached_data", useCachedData),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_baas_search_indexed_object.baas_search_indexed_object",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmBaasSearchIndexedObjectConfigBasic(xIbmTenantID string, objectType string) string {
	return fmt.Sprintf(`
		resource "ibm_baas_search_indexed_object" "baas_search_indexed_object_instance" {
			x_ibm_tenant_id = "%s"
			object_type = "%s"
		}
	`, xIbmTenantID, objectType)
}

func testAccCheckIbmBaasSearchIndexedObjectConfig(xIbmTenantID string, tenantID string, includeTenants string, paginationCookie string, count string, objectType string, useCachedData string) string {
	return fmt.Sprintf(`

		resource "ibm_baas_search_indexed_object" "baas_search_indexed_object_instance" {
			x_ibm_tenant_id = "%s"
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
			cassandra_params {
				cassandra_object_types = [ "CassandraKeyspaces" ]
				search_string = "search_string"
				source_ids = [ 1 ]
			}
			couchbase_params {
				couchbase_object_types = [ "CouchbaseBuckets" ]
				search_string = "search_string"
				source_ids = [ 1 ]
			}
			email_params {
				attendees_addresses = [ "attendeesAddresses" ]
				bcc_recipient_addresses = [ "bccRecipientAddresses" ]
				cc_recipient_addresses = [ "ccRecipientAddresses" ]
				created_end_time_secs = 1
				created_start_time_secs = 1
				due_date_end_time_secs = 1
				due_date_start_time_secs = 1
				email_address = "email_address"
				email_subject = "email_subject"
				first_name = "first_name"
				folder_names = [ "folderNames" ]
				has_attachment = true
				last_modified_end_time_secs = 1
				last_modified_start_time_secs = 1
				last_name = "last_name"
				middle_name = "middle_name"
				organizer_address = "organizer_address"
				received_end_time_secs = 1
				received_start_time_secs = 1
				recipient_addresses = [ "recipientAddresses" ]
				sender_address = "sender_address"
				source_environment = "kO365"
				task_status_types = [ "NotStarted" ]
				types = [ "Email" ]
				o365_params {
					domain_ids = [ 1 ]
					mailbox_ids = [ 1 ]
				}
			}
			exchange_params {
				search_string = "search_string"
			}
			file_params {
				search_string = "search_string"
				types = [ "File" ]
				source_environments = [ "kVMware" ]
				source_ids = [ 1 ]
				object_ids = [ 1 ]
			}
			hbase_params {
				hbase_object_types = [ "HbaseNamespaces" ]
				search_string = "search_string"
				source_ids = [ 1 ]
			}
			hdfs_params {
				hdfs_types = [ "HDFSFolders" ]
				search_string = "search_string"
				source_ids = [ 1 ]
			}
			hive_params {
				hive_object_types = [ "HiveDatabases" ]
				search_string = "search_string"
				source_ids = [ 1 ]
			}
			mongodb_params {
				mongo_db_object_types = [ "MongoDatabases" ]
				search_string = "search_string"
				source_ids = [ 1 ]
			}
			ms_groups_params {
				mailbox_params {
					attendees_addresses = [ "attendeesAddresses" ]
					bcc_recipient_addresses = [ "bccRecipientAddresses" ]
					cc_recipient_addresses = [ "ccRecipientAddresses" ]
					created_end_time_secs = 1
					created_start_time_secs = 1
					due_date_end_time_secs = 1
					due_date_start_time_secs = 1
					email_address = "email_address"
					email_subject = "email_subject"
					first_name = "first_name"
					folder_names = [ "folderNames" ]
					has_attachment = true
					last_modified_end_time_secs = 1
					last_modified_start_time_secs = 1
					last_name = "last_name"
					middle_name = "middle_name"
					organizer_address = "organizer_address"
					received_end_time_secs = 1
					received_start_time_secs = 1
					recipient_addresses = [ "recipientAddresses" ]
					sender_address = "sender_address"
					source_environment = "kO365"
					task_status_types = [ "NotStarted" ]
					types = [ "Email" ]
				}
				o365_params {
					domain_ids = [ 1 ]
					group_ids = [ 1 ]
					site_ids = [ 1 ]
					teams_ids = [ 1 ]
					user_ids = [ 1 ]
				}
				site_params {
					category_types = [ "Document" ]
					creation_end_time_secs = 1
					creation_start_time_secs = 1
					include_files = true
					include_folders = true
					o365_params {
						domain_ids = [ 1 ]
						group_ids = [ 1 ]
						site_ids = [ 1 ]
						teams_ids = [ 1 ]
						user_ids = [ 1 ]
					}
					owner_names = [ "ownerNames" ]
					search_string = "search_string"
					size_bytes_lower_limit = 1
					size_bytes_upper_limit = 1
				}
			}
			ms_teams_params {
				category_types = [ "Document" ]
				channel_names = [ "channelNames" ]
				channel_params {
					channel_email = "channel_email"
					channel_id = "channel_id"
					channel_name = "channel_name"
					include_private_channels = true
					include_public_channels = true
				}
				creation_end_time_secs = 1
				creation_start_time_secs = 1
				o365_params {
					domain_ids = [ 1 ]
					group_ids = [ 1 ]
					site_ids = [ 1 ]
					teams_ids = [ 1 ]
					user_ids = [ 1 ]
				}
				owner_names = [ "ownerNames" ]
				search_string = "search_string"
				size_bytes_lower_limit = 1
				size_bytes_upper_limit = 1
				types = [ "Channel" ]
			}
			one_drive_params {
				category_types = [ "Document" ]
				creation_end_time_secs = 1
				creation_start_time_secs = 1
				include_files = true
				include_folders = true
				o365_params {
					domain_ids = [ 1 ]
					group_ids = [ 1 ]
					site_ids = [ 1 ]
					teams_ids = [ 1 ]
					user_ids = [ 1 ]
				}
				owner_names = [ "ownerNames" ]
				search_string = "search_string"
				size_bytes_lower_limit = 1
				size_bytes_upper_limit = 1
			}
			public_folder_params {
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
			sfdc_params {
				mutation_types = [ "All" ]
				object_name = "object_name"
				query_string = "query_string"
				snapshot_id = "snapshot_id"
			}
			sharepoint_params {
				category_types = [ "Document" ]
				creation_end_time_secs = 1
				creation_start_time_secs = 1
				include_files = true
				include_folders = true
				o365_params {
					domain_ids = [ 1 ]
					group_ids = [ 1 ]
					site_ids = [ 1 ]
					teams_ids = [ 1 ]
					user_ids = [ 1 ]
				}
				owner_names = [ "ownerNames" ]
				search_string = "search_string"
				size_bytes_lower_limit = 1
				size_bytes_upper_limit = 1
			}
			uda_params {
				search_string = "search_string"
				source_ids = [ 1 ]
			}
		}
	`, xIbmTenantID, tenantID, includeTenants, paginationCookie, count, objectType, useCachedData)
}

func testAccCheckIbmBaasSearchIndexedObjectExists(n string, obj backuprecoveryv1.SearchIndexedObjectsResponse) resource.TestCheckFunc {

	return nil
}

func testAccCheckIbmBaasSearchIndexedObjectDestroy(s *terraform.State) error {
	backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_baas_search_indexed_object" {
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
