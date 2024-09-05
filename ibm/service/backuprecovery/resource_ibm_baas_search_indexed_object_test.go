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
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/backuprecovery"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func TestAccIbmBaasSearchIndexedObjectBasic(t *testing.T) {
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

func TestResourceIbmBaasSearchIndexedObjectCassandraOnPremSearchParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["cassandra_object_types"] = []string{"CassandraKeyspaces"}
		model["search_string"] = "testString"
		model["source_ids"] = []int64{int64(26)}

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.CassandraOnPremSearchParams)
	model.CassandraObjectTypes = []string{"CassandraKeyspaces"}
	model.SearchString = core.StringPtr("testString")
	model.SourceIds = []int64{int64(26)}

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectCassandraOnPremSearchParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSearchIndexedObjectCouchBaseOnPremSearchParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["couchbase_object_types"] = []string{"CouchbaseBuckets"}
		model["search_string"] = "testString"
		model["source_ids"] = []int64{int64(26)}

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.CouchBaseOnPremSearchParams)
	model.CouchbaseObjectTypes = []string{"CouchbaseBuckets"}
	model.SearchString = core.StringPtr("testString")
	model.SourceIds = []int64{int64(26)}

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectCouchBaseOnPremSearchParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSearchIndexedObjectSearchEmailRequestParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		o365SearchEmailsRequestParamsModel := make(map[string]interface{})
		o365SearchEmailsRequestParamsModel["domain_ids"] = []int64{int64(26)}
		o365SearchEmailsRequestParamsModel["mailbox_ids"] = []int64{int64(26)}

		model := make(map[string]interface{})
		model["attendees_addresses"] = []string{"testString"}
		model["bcc_recipient_addresses"] = []string{"testString"}
		model["cc_recipient_addresses"] = []string{"testString"}
		model["created_end_time_secs"] = int(26)
		model["created_start_time_secs"] = int(26)
		model["due_date_end_time_secs"] = int(26)
		model["due_date_start_time_secs"] = int(26)
		model["email_address"] = "testString"
		model["email_subject"] = "testString"
		model["first_name"] = "testString"
		model["folder_names"] = []string{"testString"}
		model["has_attachment"] = true
		model["last_modified_end_time_secs"] = int(26)
		model["last_modified_start_time_secs"] = int(26)
		model["last_name"] = "testString"
		model["middle_name"] = "testString"
		model["organizer_address"] = "testString"
		model["received_end_time_secs"] = int(26)
		model["received_start_time_secs"] = int(26)
		model["recipient_addresses"] = []string{"testString"}
		model["sender_address"] = "testString"
		model["source_environment"] = "kO365"
		model["task_status_types"] = []string{"NotStarted"}
		model["types"] = []string{"Email"}
		model["o365_params"] = []map[string]interface{}{o365SearchEmailsRequestParamsModel}

		assert.Equal(t, result, model)
	}

	o365SearchEmailsRequestParamsModel := new(backuprecoveryv1.O365SearchEmailsRequestParams)
	o365SearchEmailsRequestParamsModel.DomainIds = []int64{int64(26)}
	o365SearchEmailsRequestParamsModel.MailboxIds = []int64{int64(26)}

	model := new(backuprecoveryv1.SearchEmailRequestParams)
	model.AttendeesAddresses = []string{"testString"}
	model.BccRecipientAddresses = []string{"testString"}
	model.CcRecipientAddresses = []string{"testString"}
	model.CreatedEndTimeSecs = core.Int64Ptr(int64(26))
	model.CreatedStartTimeSecs = core.Int64Ptr(int64(26))
	model.DueDateEndTimeSecs = core.Int64Ptr(int64(26))
	model.DueDateStartTimeSecs = core.Int64Ptr(int64(26))
	model.EmailAddress = core.StringPtr("testString")
	model.EmailSubject = core.StringPtr("testString")
	model.FirstName = core.StringPtr("testString")
	model.FolderNames = []string{"testString"}
	model.HasAttachment = core.BoolPtr(true)
	model.LastModifiedEndTimeSecs = core.Int64Ptr(int64(26))
	model.LastModifiedStartTimeSecs = core.Int64Ptr(int64(26))
	model.LastName = core.StringPtr("testString")
	model.MiddleName = core.StringPtr("testString")
	model.OrganizerAddress = core.StringPtr("testString")
	model.ReceivedEndTimeSecs = core.Int64Ptr(int64(26))
	model.ReceivedStartTimeSecs = core.Int64Ptr(int64(26))
	model.RecipientAddresses = []string{"testString"}
	model.SenderAddress = core.StringPtr("testString")
	model.SourceEnvironment = core.StringPtr("kO365")
	model.TaskStatusTypes = []string{"NotStarted"}
	model.Types = []string{"Email"}
	model.O365Params = o365SearchEmailsRequestParamsModel

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectSearchEmailRequestParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSearchIndexedObjectO365SearchEmailsRequestParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["domain_ids"] = []int64{int64(26)}
		model["mailbox_ids"] = []int64{int64(26)}

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.O365SearchEmailsRequestParams)
	model.DomainIds = []int64{int64(26)}
	model.MailboxIds = []int64{int64(26)}

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectO365SearchEmailsRequestParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSearchIndexedObjectSearchExchangeObjectsRequestParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["search_string"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.SearchExchangeObjectsRequestParams)
	model.SearchString = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectSearchExchangeObjectsRequestParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSearchIndexedObjectSearchFileRequestParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["search_string"] = "testString"
		model["types"] = []string{"File"}
		model["source_environments"] = []string{"kVMware"}
		model["source_ids"] = []int64{int64(26)}
		model["object_ids"] = []int64{int64(26)}

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.SearchFileRequestParams)
	model.SearchString = core.StringPtr("testString")
	model.Types = []string{"File"}
	model.SourceEnvironments = []string{"kVMware"}
	model.SourceIds = []int64{int64(26)}
	model.ObjectIds = []int64{int64(26)}

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectSearchFileRequestParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSearchIndexedObjectHbaseOnPremSearchParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["hbase_object_types"] = []string{"HbaseNamespaces"}
		model["search_string"] = "testString"
		model["source_ids"] = []int64{int64(26)}

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.HbaseOnPremSearchParams)
	model.HbaseObjectTypes = []string{"HbaseNamespaces"}
	model.SearchString = core.StringPtr("testString")
	model.SourceIds = []int64{int64(26)}

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectHbaseOnPremSearchParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSearchIndexedObjectHDFSOnPremSearchParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["hdfs_types"] = []string{"HDFSFolders"}
		model["search_string"] = "testString"
		model["source_ids"] = []int64{int64(26)}

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.HDFSOnPremSearchParams)
	model.HdfsTypes = []string{"HDFSFolders"}
	model.SearchString = core.StringPtr("testString")
	model.SourceIds = []int64{int64(26)}

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectHDFSOnPremSearchParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSearchIndexedObjectHiveOnPremSearchParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["hive_object_types"] = []string{"HiveDatabases"}
		model["search_string"] = "testString"
		model["source_ids"] = []int64{int64(26)}

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.HiveOnPremSearchParams)
	model.HiveObjectTypes = []string{"HiveDatabases"}
	model.SearchString = core.StringPtr("testString")
	model.SourceIds = []int64{int64(26)}

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectHiveOnPremSearchParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSearchIndexedObjectMongoDbOnPremSearchParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["mongo_db_object_types"] = []string{"MongoDatabases"}
		model["search_string"] = "testString"
		model["source_ids"] = []int64{int64(26)}

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.MongoDbOnPremSearchParams)
	model.MongoDBObjectTypes = []string{"MongoDatabases"}
	model.SearchString = core.StringPtr("testString")
	model.SourceIds = []int64{int64(26)}

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectMongoDbOnPremSearchParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSearchIndexedObjectSearchMsGroupsRequestParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		searchEmailRequestParamsBaseModel := make(map[string]interface{})
		searchEmailRequestParamsBaseModel["attendees_addresses"] = []string{"testString"}
		searchEmailRequestParamsBaseModel["bcc_recipient_addresses"] = []string{"testString"}
		searchEmailRequestParamsBaseModel["cc_recipient_addresses"] = []string{"testString"}
		searchEmailRequestParamsBaseModel["created_end_time_secs"] = int(26)
		searchEmailRequestParamsBaseModel["created_start_time_secs"] = int(26)
		searchEmailRequestParamsBaseModel["due_date_end_time_secs"] = int(26)
		searchEmailRequestParamsBaseModel["due_date_start_time_secs"] = int(26)
		searchEmailRequestParamsBaseModel["email_address"] = "testString"
		searchEmailRequestParamsBaseModel["email_subject"] = "testString"
		searchEmailRequestParamsBaseModel["first_name"] = "testString"
		searchEmailRequestParamsBaseModel["folder_names"] = []string{"testString"}
		searchEmailRequestParamsBaseModel["has_attachment"] = true
		searchEmailRequestParamsBaseModel["last_modified_end_time_secs"] = int(26)
		searchEmailRequestParamsBaseModel["last_modified_start_time_secs"] = int(26)
		searchEmailRequestParamsBaseModel["last_name"] = "testString"
		searchEmailRequestParamsBaseModel["middle_name"] = "testString"
		searchEmailRequestParamsBaseModel["organizer_address"] = "testString"
		searchEmailRequestParamsBaseModel["received_end_time_secs"] = int(26)
		searchEmailRequestParamsBaseModel["received_start_time_secs"] = int(26)
		searchEmailRequestParamsBaseModel["recipient_addresses"] = []string{"testString"}
		searchEmailRequestParamsBaseModel["sender_address"] = "testString"
		searchEmailRequestParamsBaseModel["source_environment"] = "kO365"
		searchEmailRequestParamsBaseModel["task_status_types"] = []string{"NotStarted"}
		searchEmailRequestParamsBaseModel["types"] = []string{"Email"}

		o365SearchRequestParamsModel := make(map[string]interface{})
		o365SearchRequestParamsModel["domain_ids"] = []int64{int64(26)}
		o365SearchRequestParamsModel["group_ids"] = []int64{int64(26)}
		o365SearchRequestParamsModel["site_ids"] = []int64{int64(26)}
		o365SearchRequestParamsModel["teams_ids"] = []int64{int64(26)}
		o365SearchRequestParamsModel["user_ids"] = []int64{int64(26)}

		searchDocumentLibraryRequestParamsModel := make(map[string]interface{})
		searchDocumentLibraryRequestParamsModel["category_types"] = []string{"Document"}
		searchDocumentLibraryRequestParamsModel["creation_end_time_secs"] = int(26)
		searchDocumentLibraryRequestParamsModel["creation_start_time_secs"] = int(26)
		searchDocumentLibraryRequestParamsModel["include_files"] = true
		searchDocumentLibraryRequestParamsModel["include_folders"] = true
		searchDocumentLibraryRequestParamsModel["o365_params"] = []map[string]interface{}{o365SearchRequestParamsModel}
		searchDocumentLibraryRequestParamsModel["owner_names"] = []string{"testString"}
		searchDocumentLibraryRequestParamsModel["search_string"] = "testString"
		searchDocumentLibraryRequestParamsModel["size_bytes_lower_limit"] = int(26)
		searchDocumentLibraryRequestParamsModel["size_bytes_upper_limit"] = int(26)

		model := make(map[string]interface{})
		model["mailbox_params"] = []map[string]interface{}{searchEmailRequestParamsBaseModel}
		model["o365_params"] = []map[string]interface{}{o365SearchRequestParamsModel}
		model["site_params"] = []map[string]interface{}{searchDocumentLibraryRequestParamsModel}

		assert.Equal(t, result, model)
	}

	searchEmailRequestParamsBaseModel := new(backuprecoveryv1.SearchEmailRequestParamsBase)
	searchEmailRequestParamsBaseModel.AttendeesAddresses = []string{"testString"}
	searchEmailRequestParamsBaseModel.BccRecipientAddresses = []string{"testString"}
	searchEmailRequestParamsBaseModel.CcRecipientAddresses = []string{"testString"}
	searchEmailRequestParamsBaseModel.CreatedEndTimeSecs = core.Int64Ptr(int64(26))
	searchEmailRequestParamsBaseModel.CreatedStartTimeSecs = core.Int64Ptr(int64(26))
	searchEmailRequestParamsBaseModel.DueDateEndTimeSecs = core.Int64Ptr(int64(26))
	searchEmailRequestParamsBaseModel.DueDateStartTimeSecs = core.Int64Ptr(int64(26))
	searchEmailRequestParamsBaseModel.EmailAddress = core.StringPtr("testString")
	searchEmailRequestParamsBaseModel.EmailSubject = core.StringPtr("testString")
	searchEmailRequestParamsBaseModel.FirstName = core.StringPtr("testString")
	searchEmailRequestParamsBaseModel.FolderNames = []string{"testString"}
	searchEmailRequestParamsBaseModel.HasAttachment = core.BoolPtr(true)
	searchEmailRequestParamsBaseModel.LastModifiedEndTimeSecs = core.Int64Ptr(int64(26))
	searchEmailRequestParamsBaseModel.LastModifiedStartTimeSecs = core.Int64Ptr(int64(26))
	searchEmailRequestParamsBaseModel.LastName = core.StringPtr("testString")
	searchEmailRequestParamsBaseModel.MiddleName = core.StringPtr("testString")
	searchEmailRequestParamsBaseModel.OrganizerAddress = core.StringPtr("testString")
	searchEmailRequestParamsBaseModel.ReceivedEndTimeSecs = core.Int64Ptr(int64(26))
	searchEmailRequestParamsBaseModel.ReceivedStartTimeSecs = core.Int64Ptr(int64(26))
	searchEmailRequestParamsBaseModel.RecipientAddresses = []string{"testString"}
	searchEmailRequestParamsBaseModel.SenderAddress = core.StringPtr("testString")
	searchEmailRequestParamsBaseModel.SourceEnvironment = core.StringPtr("kO365")
	searchEmailRequestParamsBaseModel.TaskStatusTypes = []string{"NotStarted"}
	searchEmailRequestParamsBaseModel.Types = []string{"Email"}

	o365SearchRequestParamsModel := new(backuprecoveryv1.O365SearchRequestParams)
	o365SearchRequestParamsModel.DomainIds = []int64{int64(26)}
	o365SearchRequestParamsModel.GroupIds = []int64{int64(26)}
	o365SearchRequestParamsModel.SiteIds = []int64{int64(26)}
	o365SearchRequestParamsModel.TeamsIds = []int64{int64(26)}
	o365SearchRequestParamsModel.UserIds = []int64{int64(26)}

	searchDocumentLibraryRequestParamsModel := new(backuprecoveryv1.SearchDocumentLibraryRequestParams)
	searchDocumentLibraryRequestParamsModel.CategoryTypes = []string{"Document"}
	searchDocumentLibraryRequestParamsModel.CreationEndTimeSecs = core.Int64Ptr(int64(26))
	searchDocumentLibraryRequestParamsModel.CreationStartTimeSecs = core.Int64Ptr(int64(26))
	searchDocumentLibraryRequestParamsModel.IncludeFiles = core.BoolPtr(true)
	searchDocumentLibraryRequestParamsModel.IncludeFolders = core.BoolPtr(true)
	searchDocumentLibraryRequestParamsModel.O365Params = o365SearchRequestParamsModel
	searchDocumentLibraryRequestParamsModel.OwnerNames = []string{"testString"}
	searchDocumentLibraryRequestParamsModel.SearchString = core.StringPtr("testString")
	searchDocumentLibraryRequestParamsModel.SizeBytesLowerLimit = core.Int64Ptr(int64(26))
	searchDocumentLibraryRequestParamsModel.SizeBytesUpperLimit = core.Int64Ptr(int64(26))

	model := new(backuprecoveryv1.SearchMsGroupsRequestParams)
	model.MailboxParams = searchEmailRequestParamsBaseModel
	model.O365Params = o365SearchRequestParamsModel
	model.SiteParams = searchDocumentLibraryRequestParamsModel

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectSearchMsGroupsRequestParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSearchIndexedObjectSearchEmailRequestParamsBaseToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["attendees_addresses"] = []string{"testString"}
		model["bcc_recipient_addresses"] = []string{"testString"}
		model["cc_recipient_addresses"] = []string{"testString"}
		model["created_end_time_secs"] = int(26)
		model["created_start_time_secs"] = int(26)
		model["due_date_end_time_secs"] = int(26)
		model["due_date_start_time_secs"] = int(26)
		model["email_address"] = "testString"
		model["email_subject"] = "testString"
		model["first_name"] = "testString"
		model["folder_names"] = []string{"testString"}
		model["has_attachment"] = true
		model["last_modified_end_time_secs"] = int(26)
		model["last_modified_start_time_secs"] = int(26)
		model["last_name"] = "testString"
		model["middle_name"] = "testString"
		model["organizer_address"] = "testString"
		model["received_end_time_secs"] = int(26)
		model["received_start_time_secs"] = int(26)
		model["recipient_addresses"] = []string{"testString"}
		model["sender_address"] = "testString"
		model["source_environment"] = "kO365"
		model["task_status_types"] = []string{"NotStarted"}
		model["types"] = []string{"Email"}

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.SearchEmailRequestParamsBase)
	model.AttendeesAddresses = []string{"testString"}
	model.BccRecipientAddresses = []string{"testString"}
	model.CcRecipientAddresses = []string{"testString"}
	model.CreatedEndTimeSecs = core.Int64Ptr(int64(26))
	model.CreatedStartTimeSecs = core.Int64Ptr(int64(26))
	model.DueDateEndTimeSecs = core.Int64Ptr(int64(26))
	model.DueDateStartTimeSecs = core.Int64Ptr(int64(26))
	model.EmailAddress = core.StringPtr("testString")
	model.EmailSubject = core.StringPtr("testString")
	model.FirstName = core.StringPtr("testString")
	model.FolderNames = []string{"testString"}
	model.HasAttachment = core.BoolPtr(true)
	model.LastModifiedEndTimeSecs = core.Int64Ptr(int64(26))
	model.LastModifiedStartTimeSecs = core.Int64Ptr(int64(26))
	model.LastName = core.StringPtr("testString")
	model.MiddleName = core.StringPtr("testString")
	model.OrganizerAddress = core.StringPtr("testString")
	model.ReceivedEndTimeSecs = core.Int64Ptr(int64(26))
	model.ReceivedStartTimeSecs = core.Int64Ptr(int64(26))
	model.RecipientAddresses = []string{"testString"}
	model.SenderAddress = core.StringPtr("testString")
	model.SourceEnvironment = core.StringPtr("kO365")
	model.TaskStatusTypes = []string{"NotStarted"}
	model.Types = []string{"Email"}

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectSearchEmailRequestParamsBaseToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSearchIndexedObjectO365SearchRequestParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["domain_ids"] = []int64{int64(26)}
		model["group_ids"] = []int64{int64(26)}
		model["site_ids"] = []int64{int64(26)}
		model["teams_ids"] = []int64{int64(26)}
		model["user_ids"] = []int64{int64(26)}

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.O365SearchRequestParams)
	model.DomainIds = []int64{int64(26)}
	model.GroupIds = []int64{int64(26)}
	model.SiteIds = []int64{int64(26)}
	model.TeamsIds = []int64{int64(26)}
	model.UserIds = []int64{int64(26)}

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectO365SearchRequestParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSearchIndexedObjectSearchDocumentLibraryRequestParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		o365SearchRequestParamsModel := make(map[string]interface{})
		o365SearchRequestParamsModel["domain_ids"] = []int64{int64(26)}
		o365SearchRequestParamsModel["group_ids"] = []int64{int64(26)}
		o365SearchRequestParamsModel["site_ids"] = []int64{int64(26)}
		o365SearchRequestParamsModel["teams_ids"] = []int64{int64(26)}
		o365SearchRequestParamsModel["user_ids"] = []int64{int64(26)}

		model := make(map[string]interface{})
		model["category_types"] = []string{"Document"}
		model["creation_end_time_secs"] = int(26)
		model["creation_start_time_secs"] = int(26)
		model["include_files"] = true
		model["include_folders"] = true
		model["o365_params"] = []map[string]interface{}{o365SearchRequestParamsModel}
		model["owner_names"] = []string{"testString"}
		model["search_string"] = "testString"
		model["size_bytes_lower_limit"] = int(26)
		model["size_bytes_upper_limit"] = int(26)

		assert.Equal(t, result, model)
	}

	o365SearchRequestParamsModel := new(backuprecoveryv1.O365SearchRequestParams)
	o365SearchRequestParamsModel.DomainIds = []int64{int64(26)}
	o365SearchRequestParamsModel.GroupIds = []int64{int64(26)}
	o365SearchRequestParamsModel.SiteIds = []int64{int64(26)}
	o365SearchRequestParamsModel.TeamsIds = []int64{int64(26)}
	o365SearchRequestParamsModel.UserIds = []int64{int64(26)}

	model := new(backuprecoveryv1.SearchDocumentLibraryRequestParams)
	model.CategoryTypes = []string{"Document"}
	model.CreationEndTimeSecs = core.Int64Ptr(int64(26))
	model.CreationStartTimeSecs = core.Int64Ptr(int64(26))
	model.IncludeFiles = core.BoolPtr(true)
	model.IncludeFolders = core.BoolPtr(true)
	model.O365Params = o365SearchRequestParamsModel
	model.OwnerNames = []string{"testString"}
	model.SearchString = core.StringPtr("testString")
	model.SizeBytesLowerLimit = core.Int64Ptr(int64(26))
	model.SizeBytesUpperLimit = core.Int64Ptr(int64(26))

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectSearchDocumentLibraryRequestParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSearchIndexedObjectSearchMsTeamsRequestParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		o365TeamsChannelsSearchRequestParamsModel := make(map[string]interface{})
		o365TeamsChannelsSearchRequestParamsModel["channel_email"] = "testString"
		o365TeamsChannelsSearchRequestParamsModel["channel_id"] = "testString"
		o365TeamsChannelsSearchRequestParamsModel["channel_name"] = "testString"
		o365TeamsChannelsSearchRequestParamsModel["include_private_channels"] = true
		o365TeamsChannelsSearchRequestParamsModel["include_public_channels"] = true

		o365SearchRequestParamsModel := make(map[string]interface{})
		o365SearchRequestParamsModel["domain_ids"] = []int64{int64(26)}
		o365SearchRequestParamsModel["group_ids"] = []int64{int64(26)}
		o365SearchRequestParamsModel["site_ids"] = []int64{int64(26)}
		o365SearchRequestParamsModel["teams_ids"] = []int64{int64(26)}
		o365SearchRequestParamsModel["user_ids"] = []int64{int64(26)}

		model := make(map[string]interface{})
		model["category_types"] = []string{"Document"}
		model["channel_names"] = []string{"testString"}
		model["channel_params"] = []map[string]interface{}{o365TeamsChannelsSearchRequestParamsModel}
		model["creation_end_time_secs"] = int(26)
		model["creation_start_time_secs"] = int(26)
		model["o365_params"] = []map[string]interface{}{o365SearchRequestParamsModel}
		model["owner_names"] = []string{"testString"}
		model["search_string"] = "testString"
		model["size_bytes_lower_limit"] = int(26)
		model["size_bytes_upper_limit"] = int(26)
		model["types"] = []string{"Channel"}

		assert.Equal(t, result, model)
	}

	o365TeamsChannelsSearchRequestParamsModel := new(backuprecoveryv1.O365TeamsChannelsSearchRequestParams)
	o365TeamsChannelsSearchRequestParamsModel.ChannelEmail = core.StringPtr("testString")
	o365TeamsChannelsSearchRequestParamsModel.ChannelID = core.StringPtr("testString")
	o365TeamsChannelsSearchRequestParamsModel.ChannelName = core.StringPtr("testString")
	o365TeamsChannelsSearchRequestParamsModel.IncludePrivateChannels = core.BoolPtr(true)
	o365TeamsChannelsSearchRequestParamsModel.IncludePublicChannels = core.BoolPtr(true)

	o365SearchRequestParamsModel := new(backuprecoveryv1.O365SearchRequestParams)
	o365SearchRequestParamsModel.DomainIds = []int64{int64(26)}
	o365SearchRequestParamsModel.GroupIds = []int64{int64(26)}
	o365SearchRequestParamsModel.SiteIds = []int64{int64(26)}
	o365SearchRequestParamsModel.TeamsIds = []int64{int64(26)}
	o365SearchRequestParamsModel.UserIds = []int64{int64(26)}

	model := new(backuprecoveryv1.SearchMsTeamsRequestParams)
	model.CategoryTypes = []string{"Document"}
	model.ChannelNames = []string{"testString"}
	model.ChannelParams = o365TeamsChannelsSearchRequestParamsModel
	model.CreationEndTimeSecs = core.Int64Ptr(int64(26))
	model.CreationStartTimeSecs = core.Int64Ptr(int64(26))
	model.O365Params = o365SearchRequestParamsModel
	model.OwnerNames = []string{"testString"}
	model.SearchString = core.StringPtr("testString")
	model.SizeBytesLowerLimit = core.Int64Ptr(int64(26))
	model.SizeBytesUpperLimit = core.Int64Ptr(int64(26))
	model.Types = []string{"Channel"}

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectSearchMsTeamsRequestParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSearchIndexedObjectO365TeamsChannelsSearchRequestParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["channel_email"] = "testString"
		model["channel_id"] = "testString"
		model["channel_name"] = "testString"
		model["include_private_channels"] = true
		model["include_public_channels"] = true

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.O365TeamsChannelsSearchRequestParams)
	model.ChannelEmail = core.StringPtr("testString")
	model.ChannelID = core.StringPtr("testString")
	model.ChannelName = core.StringPtr("testString")
	model.IncludePrivateChannels = core.BoolPtr(true)
	model.IncludePublicChannels = core.BoolPtr(true)

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectO365TeamsChannelsSearchRequestParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSearchIndexedObjectSearchPublicFolderRequestParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["search_string"] = "testString"
		model["types"] = []string{"Calendar"}
		model["has_attachment"] = true
		model["sender_address"] = "testString"
		model["recipient_addresses"] = []string{"testString"}
		model["cc_recipient_addresses"] = []string{"testString"}
		model["bcc_recipient_addresses"] = []string{"testString"}
		model["received_start_time_secs"] = int(26)
		model["received_end_time_secs"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.SearchPublicFolderRequestParams)
	model.SearchString = core.StringPtr("testString")
	model.Types = []string{"Calendar"}
	model.HasAttachment = core.BoolPtr(true)
	model.SenderAddress = core.StringPtr("testString")
	model.RecipientAddresses = []string{"testString"}
	model.CcRecipientAddresses = []string{"testString"}
	model.BccRecipientAddresses = []string{"testString"}
	model.ReceivedStartTimeSecs = core.Int64Ptr(int64(26))
	model.ReceivedEndTimeSecs = core.Int64Ptr(int64(26))

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectSearchPublicFolderRequestParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSearchIndexedObjectSearchSfdcRecordsRequestParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["mutation_types"] = []string{"All"}
		model["object_name"] = "testString"
		model["query_string"] = "testString"
		model["snapshot_id"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.SearchSfdcRecordsRequestParams)
	model.MutationTypes = []string{"All"}
	model.ObjectName = core.StringPtr("testString")
	model.QueryString = core.StringPtr("testString")
	model.SnapshotID = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectSearchSfdcRecordsRequestParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSearchIndexedObjectUdaOnPremSearchParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["search_string"] = "testString"
		model["source_ids"] = []int64{int64(26)}

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.UdaOnPremSearchParams)
	model.SearchString = core.StringPtr("testString")
	model.SourceIds = []int64{int64(26)}

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectUdaOnPremSearchParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSearchIndexedObjectMapToCassandraOnPremSearchParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.CassandraOnPremSearchParams) {
		model := new(backuprecoveryv1.CassandraOnPremSearchParams)
		model.CassandraObjectTypes = []string{"CassandraKeyspaces"}
		model.SearchString = core.StringPtr("testString")
		model.SourceIds = []int64{int64(26)}

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["cassandra_object_types"] = []interface{}{"CassandraKeyspaces"}
	model["search_string"] = "testString"
	model["source_ids"] = []interface{}{int(26)}

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectMapToCassandraOnPremSearchParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSearchIndexedObjectMapToCouchBaseOnPremSearchParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.CouchBaseOnPremSearchParams) {
		model := new(backuprecoveryv1.CouchBaseOnPremSearchParams)
		model.CouchbaseObjectTypes = []string{"CouchbaseBuckets"}
		model.SearchString = core.StringPtr("testString")
		model.SourceIds = []int64{int64(26)}

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["couchbase_object_types"] = []interface{}{"CouchbaseBuckets"}
	model["search_string"] = "testString"
	model["source_ids"] = []interface{}{int(26)}

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectMapToCouchBaseOnPremSearchParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSearchIndexedObjectMapToSearchEmailRequestParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.SearchEmailRequestParams) {
		o365SearchEmailsRequestParamsModel := new(backuprecoveryv1.O365SearchEmailsRequestParams)
		o365SearchEmailsRequestParamsModel.DomainIds = []int64{int64(26)}
		o365SearchEmailsRequestParamsModel.MailboxIds = []int64{int64(26)}

		model := new(backuprecoveryv1.SearchEmailRequestParams)
		model.AttendeesAddresses = []string{"testString"}
		model.BccRecipientAddresses = []string{"testString"}
		model.CcRecipientAddresses = []string{"testString"}
		model.CreatedEndTimeSecs = core.Int64Ptr(int64(26))
		model.CreatedStartTimeSecs = core.Int64Ptr(int64(26))
		model.DueDateEndTimeSecs = core.Int64Ptr(int64(26))
		model.DueDateStartTimeSecs = core.Int64Ptr(int64(26))
		model.EmailAddress = core.StringPtr("testString")
		model.EmailSubject = core.StringPtr("testString")
		model.FirstName = core.StringPtr("testString")
		model.FolderNames = []string{"testString"}
		model.HasAttachment = core.BoolPtr(true)
		model.LastModifiedEndTimeSecs = core.Int64Ptr(int64(26))
		model.LastModifiedStartTimeSecs = core.Int64Ptr(int64(26))
		model.LastName = core.StringPtr("testString")
		model.MiddleName = core.StringPtr("testString")
		model.OrganizerAddress = core.StringPtr("testString")
		model.ReceivedEndTimeSecs = core.Int64Ptr(int64(26))
		model.ReceivedStartTimeSecs = core.Int64Ptr(int64(26))
		model.RecipientAddresses = []string{"testString"}
		model.SenderAddress = core.StringPtr("testString")
		model.SourceEnvironment = core.StringPtr("kO365")
		model.TaskStatusTypes = []string{"NotStarted"}
		model.Types = []string{"Email"}
		model.O365Params = o365SearchEmailsRequestParamsModel

		assert.Equal(t, result, model)
	}

	o365SearchEmailsRequestParamsModel := make(map[string]interface{})
	o365SearchEmailsRequestParamsModel["domain_ids"] = []interface{}{int(26)}
	o365SearchEmailsRequestParamsModel["mailbox_ids"] = []interface{}{int(26)}

	model := make(map[string]interface{})
	model["attendees_addresses"] = []interface{}{"testString"}
	model["bcc_recipient_addresses"] = []interface{}{"testString"}
	model["cc_recipient_addresses"] = []interface{}{"testString"}
	model["created_end_time_secs"] = int(26)
	model["created_start_time_secs"] = int(26)
	model["due_date_end_time_secs"] = int(26)
	model["due_date_start_time_secs"] = int(26)
	model["email_address"] = "testString"
	model["email_subject"] = "testString"
	model["first_name"] = "testString"
	model["folder_names"] = []interface{}{"testString"}
	model["has_attachment"] = true
	model["last_modified_end_time_secs"] = int(26)
	model["last_modified_start_time_secs"] = int(26)
	model["last_name"] = "testString"
	model["middle_name"] = "testString"
	model["organizer_address"] = "testString"
	model["received_end_time_secs"] = int(26)
	model["received_start_time_secs"] = int(26)
	model["recipient_addresses"] = []interface{}{"testString"}
	model["sender_address"] = "testString"
	model["source_environment"] = "kO365"
	model["task_status_types"] = []interface{}{"NotStarted"}
	model["types"] = []interface{}{"Email"}
	model["o365_params"] = []interface{}{o365SearchEmailsRequestParamsModel}

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectMapToSearchEmailRequestParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSearchIndexedObjectMapToO365SearchEmailsRequestParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.O365SearchEmailsRequestParams) {
		model := new(backuprecoveryv1.O365SearchEmailsRequestParams)
		model.DomainIds = []int64{int64(26)}
		model.MailboxIds = []int64{int64(26)}

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["domain_ids"] = []interface{}{int(26)}
	model["mailbox_ids"] = []interface{}{int(26)}

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectMapToO365SearchEmailsRequestParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSearchIndexedObjectMapToSearchExchangeObjectsRequestParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.SearchExchangeObjectsRequestParams) {
		model := new(backuprecoveryv1.SearchExchangeObjectsRequestParams)
		model.SearchString = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["search_string"] = "testString"

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectMapToSearchExchangeObjectsRequestParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSearchIndexedObjectMapToSearchFileRequestParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.SearchFileRequestParams) {
		model := new(backuprecoveryv1.SearchFileRequestParams)
		model.SearchString = core.StringPtr("testString")
		model.Types = []string{"File"}
		model.SourceEnvironments = []string{"kVMware"}
		model.SourceIds = []int64{int64(26)}
		model.ObjectIds = []int64{int64(26)}

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["search_string"] = "testString"
	model["types"] = []interface{}{"File"}
	model["source_environments"] = []interface{}{"kVMware"}
	model["source_ids"] = []interface{}{int(26)}
	model["object_ids"] = []interface{}{int(26)}

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectMapToSearchFileRequestParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSearchIndexedObjectMapToHbaseOnPremSearchParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.HbaseOnPremSearchParams) {
		model := new(backuprecoveryv1.HbaseOnPremSearchParams)
		model.HbaseObjectTypes = []string{"HbaseNamespaces"}
		model.SearchString = core.StringPtr("testString")
		model.SourceIds = []int64{int64(26)}

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["hbase_object_types"] = []interface{}{"HbaseNamespaces"}
	model["search_string"] = "testString"
	model["source_ids"] = []interface{}{int(26)}

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectMapToHbaseOnPremSearchParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSearchIndexedObjectMapToHDFSOnPremSearchParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.HDFSOnPremSearchParams) {
		model := new(backuprecoveryv1.HDFSOnPremSearchParams)
		model.HdfsTypes = []string{"HDFSFolders"}
		model.SearchString = core.StringPtr("testString")
		model.SourceIds = []int64{int64(26)}

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["hdfs_types"] = []interface{}{"HDFSFolders"}
	model["search_string"] = "testString"
	model["source_ids"] = []interface{}{int(26)}

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectMapToHDFSOnPremSearchParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSearchIndexedObjectMapToHiveOnPremSearchParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.HiveOnPremSearchParams) {
		model := new(backuprecoveryv1.HiveOnPremSearchParams)
		model.HiveObjectTypes = []string{"HiveDatabases"}
		model.SearchString = core.StringPtr("testString")
		model.SourceIds = []int64{int64(26)}

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["hive_object_types"] = []interface{}{"HiveDatabases"}
	model["search_string"] = "testString"
	model["source_ids"] = []interface{}{int(26)}

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectMapToHiveOnPremSearchParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSearchIndexedObjectMapToMongoDbOnPremSearchParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.MongoDbOnPremSearchParams) {
		model := new(backuprecoveryv1.MongoDbOnPremSearchParams)
		model.MongoDBObjectTypes = []string{"MongoDatabases"}
		model.SearchString = core.StringPtr("testString")
		model.SourceIds = []int64{int64(26)}

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["mongo_db_object_types"] = []interface{}{"MongoDatabases"}
	model["search_string"] = "testString"
	model["source_ids"] = []interface{}{int(26)}

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectMapToMongoDbOnPremSearchParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSearchIndexedObjectMapToSearchMsGroupsRequestParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.SearchMsGroupsRequestParams) {
		searchEmailRequestParamsBaseModel := new(backuprecoveryv1.SearchEmailRequestParamsBase)
		searchEmailRequestParamsBaseModel.AttendeesAddresses = []string{"testString"}
		searchEmailRequestParamsBaseModel.BccRecipientAddresses = []string{"testString"}
		searchEmailRequestParamsBaseModel.CcRecipientAddresses = []string{"testString"}
		searchEmailRequestParamsBaseModel.CreatedEndTimeSecs = core.Int64Ptr(int64(26))
		searchEmailRequestParamsBaseModel.CreatedStartTimeSecs = core.Int64Ptr(int64(26))
		searchEmailRequestParamsBaseModel.DueDateEndTimeSecs = core.Int64Ptr(int64(26))
		searchEmailRequestParamsBaseModel.DueDateStartTimeSecs = core.Int64Ptr(int64(26))
		searchEmailRequestParamsBaseModel.EmailAddress = core.StringPtr("testString")
		searchEmailRequestParamsBaseModel.EmailSubject = core.StringPtr("testString")
		searchEmailRequestParamsBaseModel.FirstName = core.StringPtr("testString")
		searchEmailRequestParamsBaseModel.FolderNames = []string{"testString"}
		searchEmailRequestParamsBaseModel.HasAttachment = core.BoolPtr(true)
		searchEmailRequestParamsBaseModel.LastModifiedEndTimeSecs = core.Int64Ptr(int64(26))
		searchEmailRequestParamsBaseModel.LastModifiedStartTimeSecs = core.Int64Ptr(int64(26))
		searchEmailRequestParamsBaseModel.LastName = core.StringPtr("testString")
		searchEmailRequestParamsBaseModel.MiddleName = core.StringPtr("testString")
		searchEmailRequestParamsBaseModel.OrganizerAddress = core.StringPtr("testString")
		searchEmailRequestParamsBaseModel.ReceivedEndTimeSecs = core.Int64Ptr(int64(26))
		searchEmailRequestParamsBaseModel.ReceivedStartTimeSecs = core.Int64Ptr(int64(26))
		searchEmailRequestParamsBaseModel.RecipientAddresses = []string{"testString"}
		searchEmailRequestParamsBaseModel.SenderAddress = core.StringPtr("testString")
		searchEmailRequestParamsBaseModel.SourceEnvironment = core.StringPtr("kO365")
		searchEmailRequestParamsBaseModel.TaskStatusTypes = []string{"NotStarted"}
		searchEmailRequestParamsBaseModel.Types = []string{"Email"}

		o365SearchRequestParamsModel := new(backuprecoveryv1.O365SearchRequestParams)
		o365SearchRequestParamsModel.DomainIds = []int64{int64(26)}
		o365SearchRequestParamsModel.GroupIds = []int64{int64(26)}
		o365SearchRequestParamsModel.SiteIds = []int64{int64(26)}
		o365SearchRequestParamsModel.TeamsIds = []int64{int64(26)}
		o365SearchRequestParamsModel.UserIds = []int64{int64(26)}

		searchDocumentLibraryRequestParamsModel := new(backuprecoveryv1.SearchDocumentLibraryRequestParams)
		searchDocumentLibraryRequestParamsModel.CategoryTypes = []string{"Document"}
		searchDocumentLibraryRequestParamsModel.CreationEndTimeSecs = core.Int64Ptr(int64(26))
		searchDocumentLibraryRequestParamsModel.CreationStartTimeSecs = core.Int64Ptr(int64(26))
		searchDocumentLibraryRequestParamsModel.IncludeFiles = core.BoolPtr(true)
		searchDocumentLibraryRequestParamsModel.IncludeFolders = core.BoolPtr(true)
		searchDocumentLibraryRequestParamsModel.O365Params = o365SearchRequestParamsModel
		searchDocumentLibraryRequestParamsModel.OwnerNames = []string{"testString"}
		searchDocumentLibraryRequestParamsModel.SearchString = core.StringPtr("testString")
		searchDocumentLibraryRequestParamsModel.SizeBytesLowerLimit = core.Int64Ptr(int64(26))
		searchDocumentLibraryRequestParamsModel.SizeBytesUpperLimit = core.Int64Ptr(int64(26))

		model := new(backuprecoveryv1.SearchMsGroupsRequestParams)
		model.MailboxParams = searchEmailRequestParamsBaseModel
		model.O365Params = o365SearchRequestParamsModel
		model.SiteParams = searchDocumentLibraryRequestParamsModel

		assert.Equal(t, result, model)
	}

	searchEmailRequestParamsBaseModel := make(map[string]interface{})
	searchEmailRequestParamsBaseModel["attendees_addresses"] = []interface{}{"testString"}
	searchEmailRequestParamsBaseModel["bcc_recipient_addresses"] = []interface{}{"testString"}
	searchEmailRequestParamsBaseModel["cc_recipient_addresses"] = []interface{}{"testString"}
	searchEmailRequestParamsBaseModel["created_end_time_secs"] = int(26)
	searchEmailRequestParamsBaseModel["created_start_time_secs"] = int(26)
	searchEmailRequestParamsBaseModel["due_date_end_time_secs"] = int(26)
	searchEmailRequestParamsBaseModel["due_date_start_time_secs"] = int(26)
	searchEmailRequestParamsBaseModel["email_address"] = "testString"
	searchEmailRequestParamsBaseModel["email_subject"] = "testString"
	searchEmailRequestParamsBaseModel["first_name"] = "testString"
	searchEmailRequestParamsBaseModel["folder_names"] = []interface{}{"testString"}
	searchEmailRequestParamsBaseModel["has_attachment"] = true
	searchEmailRequestParamsBaseModel["last_modified_end_time_secs"] = int(26)
	searchEmailRequestParamsBaseModel["last_modified_start_time_secs"] = int(26)
	searchEmailRequestParamsBaseModel["last_name"] = "testString"
	searchEmailRequestParamsBaseModel["middle_name"] = "testString"
	searchEmailRequestParamsBaseModel["organizer_address"] = "testString"
	searchEmailRequestParamsBaseModel["received_end_time_secs"] = int(26)
	searchEmailRequestParamsBaseModel["received_start_time_secs"] = int(26)
	searchEmailRequestParamsBaseModel["recipient_addresses"] = []interface{}{"testString"}
	searchEmailRequestParamsBaseModel["sender_address"] = "testString"
	searchEmailRequestParamsBaseModel["source_environment"] = "kO365"
	searchEmailRequestParamsBaseModel["task_status_types"] = []interface{}{"NotStarted"}
	searchEmailRequestParamsBaseModel["types"] = []interface{}{"Email"}

	o365SearchRequestParamsModel := make(map[string]interface{})
	o365SearchRequestParamsModel["domain_ids"] = []interface{}{int(26)}
	o365SearchRequestParamsModel["group_ids"] = []interface{}{int(26)}
	o365SearchRequestParamsModel["site_ids"] = []interface{}{int(26)}
	o365SearchRequestParamsModel["teams_ids"] = []interface{}{int(26)}
	o365SearchRequestParamsModel["user_ids"] = []interface{}{int(26)}

	searchDocumentLibraryRequestParamsModel := make(map[string]interface{})
	searchDocumentLibraryRequestParamsModel["category_types"] = []interface{}{"Document"}
	searchDocumentLibraryRequestParamsModel["creation_end_time_secs"] = int(26)
	searchDocumentLibraryRequestParamsModel["creation_start_time_secs"] = int(26)
	searchDocumentLibraryRequestParamsModel["include_files"] = true
	searchDocumentLibraryRequestParamsModel["include_folders"] = true
	searchDocumentLibraryRequestParamsModel["o365_params"] = []interface{}{o365SearchRequestParamsModel}
	searchDocumentLibraryRequestParamsModel["owner_names"] = []interface{}{"testString"}
	searchDocumentLibraryRequestParamsModel["search_string"] = "testString"
	searchDocumentLibraryRequestParamsModel["size_bytes_lower_limit"] = int(26)
	searchDocumentLibraryRequestParamsModel["size_bytes_upper_limit"] = int(26)

	model := make(map[string]interface{})
	model["mailbox_params"] = []interface{}{searchEmailRequestParamsBaseModel}
	model["o365_params"] = []interface{}{o365SearchRequestParamsModel}
	model["site_params"] = []interface{}{searchDocumentLibraryRequestParamsModel}

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectMapToSearchMsGroupsRequestParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSearchIndexedObjectMapToSearchEmailRequestParamsBase(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.SearchEmailRequestParamsBase) {
		model := new(backuprecoveryv1.SearchEmailRequestParamsBase)
		model.AttendeesAddresses = []string{"testString"}
		model.BccRecipientAddresses = []string{"testString"}
		model.CcRecipientAddresses = []string{"testString"}
		model.CreatedEndTimeSecs = core.Int64Ptr(int64(26))
		model.CreatedStartTimeSecs = core.Int64Ptr(int64(26))
		model.DueDateEndTimeSecs = core.Int64Ptr(int64(26))
		model.DueDateStartTimeSecs = core.Int64Ptr(int64(26))
		model.EmailAddress = core.StringPtr("testString")
		model.EmailSubject = core.StringPtr("testString")
		model.FirstName = core.StringPtr("testString")
		model.FolderNames = []string{"testString"}
		model.HasAttachment = core.BoolPtr(true)
		model.LastModifiedEndTimeSecs = core.Int64Ptr(int64(26))
		model.LastModifiedStartTimeSecs = core.Int64Ptr(int64(26))
		model.LastName = core.StringPtr("testString")
		model.MiddleName = core.StringPtr("testString")
		model.OrganizerAddress = core.StringPtr("testString")
		model.ReceivedEndTimeSecs = core.Int64Ptr(int64(26))
		model.ReceivedStartTimeSecs = core.Int64Ptr(int64(26))
		model.RecipientAddresses = []string{"testString"}
		model.SenderAddress = core.StringPtr("testString")
		model.SourceEnvironment = core.StringPtr("kO365")
		model.TaskStatusTypes = []string{"NotStarted"}
		model.Types = []string{"Email"}

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["attendees_addresses"] = []interface{}{"testString"}
	model["bcc_recipient_addresses"] = []interface{}{"testString"}
	model["cc_recipient_addresses"] = []interface{}{"testString"}
	model["created_end_time_secs"] = int(26)
	model["created_start_time_secs"] = int(26)
	model["due_date_end_time_secs"] = int(26)
	model["due_date_start_time_secs"] = int(26)
	model["email_address"] = "testString"
	model["email_subject"] = "testString"
	model["first_name"] = "testString"
	model["folder_names"] = []interface{}{"testString"}
	model["has_attachment"] = true
	model["last_modified_end_time_secs"] = int(26)
	model["last_modified_start_time_secs"] = int(26)
	model["last_name"] = "testString"
	model["middle_name"] = "testString"
	model["organizer_address"] = "testString"
	model["received_end_time_secs"] = int(26)
	model["received_start_time_secs"] = int(26)
	model["recipient_addresses"] = []interface{}{"testString"}
	model["sender_address"] = "testString"
	model["source_environment"] = "kO365"
	model["task_status_types"] = []interface{}{"NotStarted"}
	model["types"] = []interface{}{"Email"}

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectMapToSearchEmailRequestParamsBase(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSearchIndexedObjectMapToO365SearchRequestParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.O365SearchRequestParams) {
		model := new(backuprecoveryv1.O365SearchRequestParams)
		model.DomainIds = []int64{int64(26)}
		model.GroupIds = []int64{int64(26)}
		model.SiteIds = []int64{int64(26)}
		model.TeamsIds = []int64{int64(26)}
		model.UserIds = []int64{int64(26)}

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["domain_ids"] = []interface{}{int(26)}
	model["group_ids"] = []interface{}{int(26)}
	model["site_ids"] = []interface{}{int(26)}
	model["teams_ids"] = []interface{}{int(26)}
	model["user_ids"] = []interface{}{int(26)}

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectMapToO365SearchRequestParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSearchIndexedObjectMapToSearchDocumentLibraryRequestParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.SearchDocumentLibraryRequestParams) {
		o365SearchRequestParamsModel := new(backuprecoveryv1.O365SearchRequestParams)
		o365SearchRequestParamsModel.DomainIds = []int64{int64(26)}
		o365SearchRequestParamsModel.GroupIds = []int64{int64(26)}
		o365SearchRequestParamsModel.SiteIds = []int64{int64(26)}
		o365SearchRequestParamsModel.TeamsIds = []int64{int64(26)}
		o365SearchRequestParamsModel.UserIds = []int64{int64(26)}

		model := new(backuprecoveryv1.SearchDocumentLibraryRequestParams)
		model.CategoryTypes = []string{"Document"}
		model.CreationEndTimeSecs = core.Int64Ptr(int64(26))
		model.CreationStartTimeSecs = core.Int64Ptr(int64(26))
		model.IncludeFiles = core.BoolPtr(true)
		model.IncludeFolders = core.BoolPtr(true)
		model.O365Params = o365SearchRequestParamsModel
		model.OwnerNames = []string{"testString"}
		model.SearchString = core.StringPtr("testString")
		model.SizeBytesLowerLimit = core.Int64Ptr(int64(26))
		model.SizeBytesUpperLimit = core.Int64Ptr(int64(26))

		assert.Equal(t, result, model)
	}

	o365SearchRequestParamsModel := make(map[string]interface{})
	o365SearchRequestParamsModel["domain_ids"] = []interface{}{int(26)}
	o365SearchRequestParamsModel["group_ids"] = []interface{}{int(26)}
	o365SearchRequestParamsModel["site_ids"] = []interface{}{int(26)}
	o365SearchRequestParamsModel["teams_ids"] = []interface{}{int(26)}
	o365SearchRequestParamsModel["user_ids"] = []interface{}{int(26)}

	model := make(map[string]interface{})
	model["category_types"] = []interface{}{"Document"}
	model["creation_end_time_secs"] = int(26)
	model["creation_start_time_secs"] = int(26)
	model["include_files"] = true
	model["include_folders"] = true
	model["o365_params"] = []interface{}{o365SearchRequestParamsModel}
	model["owner_names"] = []interface{}{"testString"}
	model["search_string"] = "testString"
	model["size_bytes_lower_limit"] = int(26)
	model["size_bytes_upper_limit"] = int(26)

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectMapToSearchDocumentLibraryRequestParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSearchIndexedObjectMapToSearchMsTeamsRequestParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.SearchMsTeamsRequestParams) {
		o365TeamsChannelsSearchRequestParamsModel := new(backuprecoveryv1.O365TeamsChannelsSearchRequestParams)
		o365TeamsChannelsSearchRequestParamsModel.ChannelEmail = core.StringPtr("testString")
		o365TeamsChannelsSearchRequestParamsModel.ChannelID = core.StringPtr("testString")
		o365TeamsChannelsSearchRequestParamsModel.ChannelName = core.StringPtr("testString")
		o365TeamsChannelsSearchRequestParamsModel.IncludePrivateChannels = core.BoolPtr(true)
		o365TeamsChannelsSearchRequestParamsModel.IncludePublicChannels = core.BoolPtr(true)

		o365SearchRequestParamsModel := new(backuprecoveryv1.O365SearchRequestParams)
		o365SearchRequestParamsModel.DomainIds = []int64{int64(26)}
		o365SearchRequestParamsModel.GroupIds = []int64{int64(26)}
		o365SearchRequestParamsModel.SiteIds = []int64{int64(26)}
		o365SearchRequestParamsModel.TeamsIds = []int64{int64(26)}
		o365SearchRequestParamsModel.UserIds = []int64{int64(26)}

		model := new(backuprecoveryv1.SearchMsTeamsRequestParams)
		model.CategoryTypes = []string{"Document"}
		model.ChannelNames = []string{"testString"}
		model.ChannelParams = o365TeamsChannelsSearchRequestParamsModel
		model.CreationEndTimeSecs = core.Int64Ptr(int64(26))
		model.CreationStartTimeSecs = core.Int64Ptr(int64(26))
		model.O365Params = o365SearchRequestParamsModel
		model.OwnerNames = []string{"testString"}
		model.SearchString = core.StringPtr("testString")
		model.SizeBytesLowerLimit = core.Int64Ptr(int64(26))
		model.SizeBytesUpperLimit = core.Int64Ptr(int64(26))
		model.Types = []string{"Channel"}

		assert.Equal(t, result, model)
	}

	o365TeamsChannelsSearchRequestParamsModel := make(map[string]interface{})
	o365TeamsChannelsSearchRequestParamsModel["channel_email"] = "testString"
	o365TeamsChannelsSearchRequestParamsModel["channel_id"] = "testString"
	o365TeamsChannelsSearchRequestParamsModel["channel_name"] = "testString"
	o365TeamsChannelsSearchRequestParamsModel["include_private_channels"] = true
	o365TeamsChannelsSearchRequestParamsModel["include_public_channels"] = true

	o365SearchRequestParamsModel := make(map[string]interface{})
	o365SearchRequestParamsModel["domain_ids"] = []interface{}{int(26)}
	o365SearchRequestParamsModel["group_ids"] = []interface{}{int(26)}
	o365SearchRequestParamsModel["site_ids"] = []interface{}{int(26)}
	o365SearchRequestParamsModel["teams_ids"] = []interface{}{int(26)}
	o365SearchRequestParamsModel["user_ids"] = []interface{}{int(26)}

	model := make(map[string]interface{})
	model["category_types"] = []interface{}{"Document"}
	model["channel_names"] = []interface{}{"testString"}
	model["channel_params"] = []interface{}{o365TeamsChannelsSearchRequestParamsModel}
	model["creation_end_time_secs"] = int(26)
	model["creation_start_time_secs"] = int(26)
	model["o365_params"] = []interface{}{o365SearchRequestParamsModel}
	model["owner_names"] = []interface{}{"testString"}
	model["search_string"] = "testString"
	model["size_bytes_lower_limit"] = int(26)
	model["size_bytes_upper_limit"] = int(26)
	model["types"] = []interface{}{"Channel"}

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectMapToSearchMsTeamsRequestParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSearchIndexedObjectMapToO365TeamsChannelsSearchRequestParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.O365TeamsChannelsSearchRequestParams) {
		model := new(backuprecoveryv1.O365TeamsChannelsSearchRequestParams)
		model.ChannelEmail = core.StringPtr("testString")
		model.ChannelID = core.StringPtr("testString")
		model.ChannelName = core.StringPtr("testString")
		model.IncludePrivateChannels = core.BoolPtr(true)
		model.IncludePublicChannels = core.BoolPtr(true)

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["channel_email"] = "testString"
	model["channel_id"] = "testString"
	model["channel_name"] = "testString"
	model["include_private_channels"] = true
	model["include_public_channels"] = true

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectMapToO365TeamsChannelsSearchRequestParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSearchIndexedObjectMapToSearchPublicFolderRequestParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.SearchPublicFolderRequestParams) {
		model := new(backuprecoveryv1.SearchPublicFolderRequestParams)
		model.SearchString = core.StringPtr("testString")
		model.Types = []string{"Calendar"}
		model.HasAttachment = core.BoolPtr(true)
		model.SenderAddress = core.StringPtr("testString")
		model.RecipientAddresses = []string{"testString"}
		model.CcRecipientAddresses = []string{"testString"}
		model.BccRecipientAddresses = []string{"testString"}
		model.ReceivedStartTimeSecs = core.Int64Ptr(int64(26))
		model.ReceivedEndTimeSecs = core.Int64Ptr(int64(26))

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["search_string"] = "testString"
	model["types"] = []interface{}{"Calendar"}
	model["has_attachment"] = true
	model["sender_address"] = "testString"
	model["recipient_addresses"] = []interface{}{"testString"}
	model["cc_recipient_addresses"] = []interface{}{"testString"}
	model["bcc_recipient_addresses"] = []interface{}{"testString"}
	model["received_start_time_secs"] = int(26)
	model["received_end_time_secs"] = int(26)

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectMapToSearchPublicFolderRequestParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSearchIndexedObjectMapToSearchSfdcRecordsRequestParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.SearchSfdcRecordsRequestParams) {
		model := new(backuprecoveryv1.SearchSfdcRecordsRequestParams)
		model.MutationTypes = []string{"All"}
		model.ObjectName = core.StringPtr("testString")
		model.QueryString = core.StringPtr("testString")
		model.SnapshotID = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["mutation_types"] = []interface{}{"All"}
	model["object_name"] = "testString"
	model["query_string"] = "testString"
	model["snapshot_id"] = "testString"

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectMapToSearchSfdcRecordsRequestParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSearchIndexedObjectMapToUdaOnPremSearchParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.UdaOnPremSearchParams) {
		model := new(backuprecoveryv1.UdaOnPremSearchParams)
		model.SearchString = core.StringPtr("testString")
		model.SourceIds = []int64{int64(26)}

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["search_string"] = "testString"
	model["source_ids"] = []interface{}{int(26)}

	result, err := backuprecovery.ResourceIbmBaasSearchIndexedObjectMapToUdaOnPremSearchParams(model)
	assert.Nil(t, err)
	checkResult(result)
}
