---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_search_indexed_object"
description: |-
  Manages Common Search Indexed Objects Params.
subcategory: "IBM Backup Recovery"
---

# ibm_backup_recovery_search_indexed_object

Provides a read-only data source to retrieve information about Common Search Indexed Objects Params.

## Example Usage

```hcl
data "ibm_backup_recovery_search_indexed_object" "backup_recovery_search_indexed_object_instance" {
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
  object_type = "Emails"
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
  x_ibm_tenant_id = "x_ibm_tenant_id"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `cassandra_params` - (Optional, Forces new resource, List) Parameters required to search Cassandra on a cluster.
Nested schema for **cassandra_params**:
	* `cassandra_object_types` - (Required, List) Specifies one or more Cassandra object types to be searched.
	  * Constraints: Allowable list items are: `CassandraKeyspaces`, `CassandraTables`.
	* `search_string` - (Required, String) Specifies the search string to search the Cassandra Objects.
	* `source_ids` - (Optional, List) Specifies a list of source ids. Only files found in these sources will be returned.
* `couchbase_params` - (Optional, Forces new resource, List) Parameters required to search CouchBase on a cluster.
Nested schema for **couchbase_params**:
	* `couchbase_object_types` - (Required, List) Specifies Couchbase object types be searched. For Couchbase it can only be set to 'CouchbaseBuckets'.
	  * Constraints: Allowable list items are: `CouchbaseBuckets`.
	* `search_string` - (Required, String) Specifies the search string to search the Couchbase Objects.
	* `source_ids` - (Optional, List) Specifies a list of source ids. Only files found in these sources will be returned.
* `count` - (Optional, Forces new resource, Integer) Specifies the number of indexed objects to be fetched for the specified pagination cookie.
* `email_params` - (Optional, Forces new resource, List) Specifies the request parameters to search for emails and email folders.
Nested schema for **email_params**:
	* `attendees_addresses` - (Optional, List) Filters the calendar items which have specified email addresses as attendees.
	* `bcc_recipient_addresses` - (Optional, List) Filters the emails which are sent to specified email addresses in BCC.
	* `cc_recipient_addresses` - (Optional, List) Filters the emails which are sent to specified email addresses in CC.
	* `created_end_time_secs` - (Optional, Integer) Specifies the end time in Unix timestamp epoch in seconds where the created time of the email/item is less than specified value.
	* `created_start_time_secs` - (Optional, Integer) Specifies the start time in Unix timestamp epoch in seconds where the created time of the email/item is more than specified value.
	* `due_date_end_time_secs` - (Optional, Integer) Specifies the end time in Unix timestamp epoch in seconds where the last modification time of the email/item is less than specified value.
	* `due_date_start_time_secs` - (Optional, Integer) Specifies the start time in Unix timestamp epoch in seconds where the last modification time of the email/item is more than specified value.
	* `email_address` - (Optional, String) Filters the contact items which have specified text in email address.
	* `email_subject` - (Optional, String) Filters the emails which have the specified text in its subject.
	* `first_name` - (Optional, String) Filters the contacts with specified text in first name.
	* `folder_names` - (Optional, List) Filters the emails which are categorized to specified folders.
	* `has_attachment` - (Optional, Boolean) Filters the emails which have attachment.
	* `last_modified_end_time_secs` - (Optional, Integer) Specifies the end time in Unix timestamp epoch in seconds where the last modification time of the email/item is less than specified value.
	* `last_modified_start_time_secs` - (Optional, Integer) Specifies the start time in Unix timestamp epoch in seconds where the last modification time of the email/item is more than specified value.
	* `last_name` - (Optional, String) Filters the contacts with specified text in last name.
	* `middle_name` - (Optional, String) Filters the contacts with specified text in middle name.
	* `o365_params` - (Optional, List) Specifies email search request params specific to O365 environment.
	Nested schema for **o365_params**:
		* `domain_ids` - (Optional, List) Specifies the domain Ids in which mailboxes are registered.
		* `mailbox_ids` - (Optional, List) Specifies the mailbox Ids which contains the emails/folders.
	* `organizer_address` - (Optional, String) Filters the calendar items which are organized by specified User's email address.
	* `received_end_time_secs` - (Optional, Integer) Specifies the end time in Unix timestamp epoch in seconds where the received time of the email is less than specified value.
	* `received_start_time_secs` - (Optional, Integer) Specifies the start time in Unix timestamp epoch in seconds where the received time of the email is more than specified value.
	* `recipient_addresses` - (Optional, List) Filters the emails which are sent to specified email addresses.
	* `sender_address` - (Optional, String) Filters the emails which are received from specified User's email address.
	* `source_environment` - (Optional, String) Specifies the source environment.
	  * Constraints: Allowable values are: `kO365`.
	* `task_status_types` - (Optional, List) Specifies a list of task item status types. Task items having status within the given types will be returned.
	  * Constraints: Allowable list items are: `NotStarted`, `InProgress`, `Completed`, `WaitingOnOthers`, `Deferred`.
	* `types` - (Optional, List) Specifies a list of mailbox item types. Only items within the given types will be returned.
	  * Constraints: Allowable list items are: `Email`, `Folder`, `Calendar`, `Contact`, `Task`, `Note`.
* `exchange_params` - (Optional, Forces new resource, List) Specifies the parameters which are specific for searching Exchange mailboxes.
Nested schema for **exchange_params**:
	* `search_string` - (Required, String) Specifies the search string to search the Exchange Objects.
* `file_params` - (Optional, Forces new resource, List) Specifies the request parameters to search for files and file folders.
Nested schema for **file_params**:
	* `object_ids` - (Optional, List) Specifies a list of object ids. Only files found in these objects will be returned.
	* `search_string` - (Optional, String) Specifies the search string to filter the files. User can specify a wildcard character '*' as a suffix to a string where all files name are matched with the prefix string.
	* `source_environments` - (Optional, List) Specifies a list of the source environments. Only files from these types of source will be returned.
	  * Constraints: Allowable list items are: `kVMware`, `kHyperV`, `kSQL`, `kView`, `kRemoteAdapter`, `kPhysical`, `kPhysicalFiles`, `kPure`, `kIbmFlashSystem`, `kAzure`, `kNetapp`, `kGenericNas`, `kAcropolis`, `kIsilon`, `kGPFS`, `kKVM`, `kAWS`, `kExchange`, `kOracle`, `kGCP`, `kFlashBlade`, `kO365`, `kHyperFlex`, `kKubernetes`, `kElastifile`, `kSAPHANA`, `kUDA`, `kSfdc`.
	* `source_ids` - (Optional, List) Specifies a list of source ids. Only files found in these sources will be returned.
	* `types` - (Optional, List) Specifies a list of file types. Only files within the given types will be returned.
	  * Constraints: Allowable list items are: `File`, `Directory`, `Symlink`.
* `hbase_params` - (Optional, Forces new resource, List) Parameters required to search Hbase on a cluster.
Nested schema for **hbase_params**:
	* `hbase_object_types` - (Required, List) Specifies one or more Hbase object types be searched.
	  * Constraints: Allowable list items are: `HbaseNamespaces`, `HbaseTables`.
	* `search_string` - (Required, String) Specifies the search string to search the Hbase Objects.
	* `source_ids` - (Optional, List) Specifies a list of source ids. Only files found in these sources will be returned.
* `hdfs_params` - (Optional, Forces new resource, List) Parameters required to search HDFS on a cluster.
Nested schema for **hdfs_params**:
	* `hdfs_types` - (Required, List) Specifies types as Folders or Files or both to be searched.
	  * Constraints: Allowable list items are: `HDFSFolders`, `HDFSFiles`.
	* `search_string` - (Required, String) Specifies the search string to search the HDFS Folders and Files.
	* `source_ids` - (Optional, List) Specifies a list of source ids. Only files found in these sources will be returned.
* `hive_params` - (Optional, Forces new resource, List) Parameters required to search Hive on a cluster.
Nested schema for **hive_params**:
	* `hive_object_types` - (Required, List) Specifies one or more Hive object types be searched.
	  * Constraints: Allowable list items are: `HiveDatabases`, `HiveTables`, `HivePartitions`.
	* `search_string` - (Required, String) Specifies the search string to search the Hive Objects.
	* `source_ids` - (Optional, List) Specifies a list of source ids. Only files found in these sources will be returned.
* `include_tenants` - (Optional, Forces new resource, Boolean) If true, the response will include objects which belongs to all tenants which the current user has permission to see. Default value is false.
  * Constraints: The default value is `false`.
* `might_have_snapshot_tag_ids` - (Optional, Forces new resource, List) Specifies list of snapshot tags, one or more of which might be present in the document. These are OR'ed together and the resulting criteria AND'ed with the rest of the query.
  * Constraints: The list items must match regular expression `/^\\d+:\\d+:[A-Z0-9-]+$/`.
* `might_have_tag_ids` - (Optional, Forces new resource, List) Specifies list of tags, one or more of which might be present in the document. These are OR'ed together and the resulting criteria AND'ed with the rest of the query.
  * Constraints: The list items must match regular expression `/^\\d+:\\d+:[A-Z0-9-]+$/`.
* `mongodb_params` - (Optional, Forces new resource, List) Parameters required to search Mongo DB on a cluster.
Nested schema for **mongodb_params**:
	* `mongo_db_object_types` - (Required, List) Specifies one or more MongoDB object types be searched.
	  * Constraints: Allowable list items are: `MongoDatabases`, `MongoCollections`.
	* `search_string` - (Required, String) Specifies the search string to search the MongoDB Objects.
	* `source_ids` - (Optional, List) Specifies a list of source ids. Only files found in these sources will be returned.
* `ms_groups_params` - (Optional, Forces new resource, List) Specifies the request params to search for Groups items.
Nested schema for **ms_groups_params**:
	* `mailbox_params` - (Optional, List) Specifies the request parameters to search for mailbox items and folders.
	Nested schema for **mailbox_params**:
		* `attendees_addresses` - (Optional, List) Filters the calendar items which have specified email addresses as attendees.
		* `bcc_recipient_addresses` - (Optional, List) Filters the emails which are sent to specified email addresses in BCC.
		* `cc_recipient_addresses` - (Optional, List) Filters the emails which are sent to specified email addresses in CC.
		* `created_end_time_secs` - (Optional, Integer) Specifies the end time in Unix timestamp epoch in seconds where the created time of the email/item is less than specified value.
		* `created_start_time_secs` - (Optional, Integer) Specifies the start time in Unix timestamp epoch in seconds where the created time of the email/item is more than specified value.
		* `due_date_end_time_secs` - (Optional, Integer) Specifies the end time in Unix timestamp epoch in seconds where the last modification time of the email/item is less than specified value.
		* `due_date_start_time_secs` - (Optional, Integer) Specifies the start time in Unix timestamp epoch in seconds where the last modification time of the email/item is more than specified value.
		* `email_address` - (Optional, String) Filters the contact items which have specified text in email address.
		* `email_subject` - (Optional, String) Filters the emails which have the specified text in its subject.
		* `first_name` - (Optional, String) Filters the contacts with specified text in first name.
		* `folder_names` - (Optional, List) Filters the emails which are categorized to specified folders.
		* `has_attachment` - (Optional, Boolean) Filters the emails which have attachment.
		* `last_modified_end_time_secs` - (Optional, Integer) Specifies the end time in Unix timestamp epoch in seconds where the last modification time of the email/item is less than specified value.
		* `last_modified_start_time_secs` - (Optional, Integer) Specifies the start time in Unix timestamp epoch in seconds where the last modification time of the email/item is more than specified value.
		* `last_name` - (Optional, String) Filters the contacts with specified text in last name.
		* `middle_name` - (Optional, String) Filters the contacts with specified text in middle name.
		* `organizer_address` - (Optional, String) Filters the calendar items which are organized by specified User's email address.
		* `received_end_time_secs` - (Optional, Integer) Specifies the end time in Unix timestamp epoch in seconds where the received time of the email is less than specified value.
		* `received_start_time_secs` - (Optional, Integer) Specifies the start time in Unix timestamp epoch in seconds where the received time of the email is more than specified value.
		* `recipient_addresses` - (Optional, List) Filters the emails which are sent to specified email addresses.
		* `sender_address` - (Optional, String) Filters the emails which are received from specified User's email address.
		* `source_environment` - (Optional, String) Specifies the source environment.
		  * Constraints: Allowable values are: `kO365`.
		* `task_status_types` - (Optional, List) Specifies a list of task item status types. Task items having status within the given types will be returned.
		  * Constraints: Allowable list items are: `NotStarted`, `InProgress`, `Completed`, `WaitingOnOthers`, `Deferred`.
		* `types` - (Optional, List) Specifies a list of mailbox item types. Only items within the given types will be returned.
		  * Constraints: Allowable list items are: `Email`, `Folder`, `Calendar`, `Contact`, `Task`, `Note`.
	* `o365_params` - (Optional, List) Specifies O365 specific params search request params to search for indexed items.
	Nested schema for **o365_params**:
		* `domain_ids` - (Optional, List) Specifies the domain Ids in which indexed items are searched.
		* `group_ids` - (Optional, List) Specifies the Group ids across which the indexed items needs to be searched.
		* `site_ids` - (Optional, List) Specifies the Sharepoint site ids across which the indexed items needs to be searched.
		* `teams_ids` - (Optional, List) Specifies the Teams ids across which the indexed items needs to be searched.
		* `user_ids` - (Optional, List) Specifies the user ids across which the indexed items needs to be searched.
	* `site_params` - (Optional, List) Specifies the request parameters to search for files/folders in document libraries.
	Nested schema for **site_params**:
		* `category_types` - (Optional, List) Specifies a list of document library types. Only items within the given types will be returned.
		  * Constraints: Allowable list items are: `Document`, `Excel`, `Powerpoint`, `Image`, `OneNote`.
		* `creation_end_time_secs` - (Optional, Integer) Specifies the end time in Unix timestamp epoch in seconds when the file/folder is created.
		* `creation_start_time_secs` - (Optional, Integer) Specifies the start time in Unix timestamp epoch in seconds when the file/folder is created.
		* `include_files` - (Optional, Boolean) Specifies whether to include files in the response. Default is true.
		  * Constraints: The default value is `true`.
		* `include_folders` - (Optional, Boolean) Specifies whether to include folders in the response. Default is true.
		  * Constraints: The default value is `true`.
		* `o365_params` - (Optional, List) Specifies O365 specific params search request params to search for indexed items.
		Nested schema for **o365_params**:
			* `domain_ids` - (Optional, List) Specifies the domain Ids in which indexed items are searched.
			* `group_ids` - (Optional, List) Specifies the Group ids across which the indexed items needs to be searched.
			* `site_ids` - (Optional, List) Specifies the Sharepoint site ids across which the indexed items needs to be searched.
			* `teams_ids` - (Optional, List) Specifies the Teams ids across which the indexed items needs to be searched.
			* `user_ids` - (Optional, List) Specifies the user ids across which the indexed items needs to be searched.
		* `owner_names` - (Optional, List) Specifies the list of owner names to filter on owner of the file/folder.
		* `search_string` - (Optional, String) Specifies the search string to filter the files/folders. User can specify a wildcard character '*' as a suffix to a string where all item names are matched with the prefix string.
		* `size_bytes_lower_limit` - (Optional, Integer) Specifies the minimum size of the file in bytes.
		* `size_bytes_upper_limit` - (Optional, Integer) Specifies the maximum size of the file in bytes.
* `ms_teams_params` - (Optional, Forces new resource, List) Specifies the request params to search for Teams items.
Nested schema for **ms_teams_params**:
	* `category_types` - (Optional, List) Specifies a list of teams files types. Only items within the given types will be returned.
	  * Constraints: Allowable list items are: `Document`, `Excel`, `Powerpoint`, `Image`, `OneNote`.
	* `channel_names` - (Optional, List) Specifies the list of channel names to filter while doing search for files.
	* `channel_params` - (Optional, List) Specifies the request parameters related to channels for Microsoft365 teams.
	Nested schema for **channel_params**:
		* `channel_email` - (Optional, String) Specifies the email id of the channel.
		* `channel_id` - (Optional, String) Specifies the unique id of the channel.
		* `channel_name` - (Optional, String) Specifies the name of the channel. Only items within the specified channel will be returned.
		* `include_private_channels` - (Optional, Boolean) Specifies whether to include private channels in the response. Default is true.
		  * Constraints: The default value is `true`.
		* `include_public_channels` - (Optional, Boolean) Specifies whether to include public channels in the response. Default is true.
		  * Constraints: The default value is `true`.
	* `creation_end_time_secs` - (Optional, Integer) Specifies the end time in Unix timestamp epoch in seconds when the item is created.
	* `creation_start_time_secs` - (Optional, Integer) Specifies the start time in Unix timestamp epoch in seconds when the item is created.
	* `o365_params` - (Optional, List) Specifies O365 specific params search request params to search for indexed items.
	Nested schema for **o365_params**:
		* `domain_ids` - (Optional, List) Specifies the domain Ids in which indexed items are searched.
		* `group_ids` - (Optional, List) Specifies the Group ids across which the indexed items needs to be searched.
		* `site_ids` - (Optional, List) Specifies the Sharepoint site ids across which the indexed items needs to be searched.
		* `teams_ids` - (Optional, List) Specifies the Teams ids across which the indexed items needs to be searched.
		* `user_ids` - (Optional, List) Specifies the user ids across which the indexed items needs to be searched.
	* `owner_names` - (Optional, List) Specifies the list of owner email ids to filter on owner of the item.
	* `search_string` - (Optional, String) Specifies the search string to filter the items. User can specify a wildcard character '*' as a suffix to a string where all item names are matched with the prefix string.
	* `size_bytes_lower_limit` - (Optional, Integer) Specifies the minimum size of the item in bytes.
	* `size_bytes_upper_limit` - (Optional, Integer) Specifies the maximum size of the item in bytes.
	* `types` - (Optional, List) Specifies a list of Teams item types. Only items within the given types will be returned.
	  * Constraints: Allowable list items are: `Channel`, `Chat`, `Conversation`, `File`, `Folder`.
* `must_have_snapshot_tag_ids` - (Optional, Forces new resource, List) Specifies snapshot tags which must be all present in the document.
  * Constraints: The list items must match regular expression `/^\\d+:\\d+:[A-Z0-9-]+$/`.
* `must_have_tag_ids` - (Optional, Forces new resource, List) Specifies tags which must be all present in the document.
  * Constraints: The list items must match regular expression `/^\\d+:\\d+:[A-Z0-9-]+$/`.
* `object_type` - (Required, Forces new resource, String) Specifies the object type to be searched for.
  * Constraints: Allowable values are: `Emails`, `Files`, `CassandraObjects`, `CouchbaseObjects`, `HbaseObjects`, `HiveObjects`, `MongoObjects`, `HDFSObjects`, `ExchangeObjects`, `PublicFolders`, `GroupsObjects`, `TeamsObjects`, `SharepointObjects`, `OneDriveObjects`, `UdaObjects`, `SfdcRecords`.
* `one_drive_params` - (Optional, Forces new resource, List) Specifies the request parameters to search for files/folders in document libraries.
Nested schema for **one_drive_params**:
	* `category_types` - (Optional, List) Specifies a list of document library types. Only items within the given types will be returned.
	  * Constraints: Allowable list items are: `Document`, `Excel`, `Powerpoint`, `Image`, `OneNote`.
	* `creation_end_time_secs` - (Optional, Integer) Specifies the end time in Unix timestamp epoch in seconds when the file/folder is created.
	* `creation_start_time_secs` - (Optional, Integer) Specifies the start time in Unix timestamp epoch in seconds when the file/folder is created.
	* `include_files` - (Optional, Boolean) Specifies whether to include files in the response. Default is true.
	  * Constraints: The default value is `true`.
	* `include_folders` - (Optional, Boolean) Specifies whether to include folders in the response. Default is true.
	  * Constraints: The default value is `true`.
	* `o365_params` - (Optional, List) Specifies O365 specific params search request params to search for indexed items.
	Nested schema for **o365_params**:
		* `domain_ids` - (Optional, List) Specifies the domain Ids in which indexed items are searched.
		* `group_ids` - (Optional, List) Specifies the Group ids across which the indexed items needs to be searched.
		* `site_ids` - (Optional, List) Specifies the Sharepoint site ids across which the indexed items needs to be searched.
		* `teams_ids` - (Optional, List) Specifies the Teams ids across which the indexed items needs to be searched.
		* `user_ids` - (Optional, List) Specifies the user ids across which the indexed items needs to be searched.
	* `owner_names` - (Optional, List) Specifies the list of owner names to filter on owner of the file/folder.
	* `search_string` - (Optional, String) Specifies the search string to filter the files/folders. User can specify a wildcard character '*' as a suffix to a string where all item names are matched with the prefix string.
	* `size_bytes_lower_limit` - (Optional, Integer) Specifies the minimum size of the file in bytes.
	* `size_bytes_upper_limit` - (Optional, Integer) Specifies the maximum size of the file in bytes.
* `pagination_cookie` - (Optional, Forces new resource, String) Specifies the pagination cookie with which subsequent parts of the response can be fetched.
* `protection_group_ids` - (Optional, Forces new resource, List) Specifies a list of Protection Group ids to filter the indexed objects. If specified, the objects indexed by specified Protection Group ids will be returned.
* `public_folder_params` - (Optional, Forces new resource, List) Specifies the request parameters to search for Public Folder items.
Nested schema for **public_folder_params**:
	* `bcc_recipient_addresses` - (Optional, List) Filters the public folder items which are sent to specified email addresses in BCC.
	  * Constraints: The list items must match regular expression `/^\\S+@\\S+.\\S+$/`.
	* `cc_recipient_addresses` - (Optional, List) Filters the public folder items which are sent to specified email addresses in CC.
	  * Constraints: The list items must match regular expression `/^\\S+@\\S+.\\S+$/`.
	* `has_attachment` - (Optional, Boolean) Filters the public folder items which have attachment.
	* `received_end_time_secs` - (Optional, Integer) Specifies the end time in Unix timestamp epoch in seconds where the received time of the public folder items is less than specified value.
	* `received_start_time_secs` - (Optional, Integer) Specifies the start time in Unix timestamp epoch in seconds where the received time of the public folder item is more than specified value.
	* `recipient_addresses` - (Optional, List) Filters the public folder items which are sent to specified email addresses.
	  * Constraints: The list items must match regular expression `/^\\S+@\\S+.\\S+$/`.
	* `search_string` - (Optional, String) Specifies the search string to filter the items. User can specify a wildcard character '*' as a suffix to a string where all item names are matched with the prefix string.
	* `sender_address` - (Optional, String) Filters the public folder items which are received from specified user's email address.
	  * Constraints: The value must match regular expression `/^\\S+@\\S+.\\S+$/`.
	* `types` - (Optional, List) Specifies a list of public folder item types. Only items within the given types will be returned.
	  * Constraints: Allowable list items are: `Calendar`, `Contact`, `Post`, `Folder`, `Task`, `Journal`, `Note`.
* `sfdc_params` - (Optional, Forces new resource, List) Specifies the parameters which are specific for searching Salesforce records.
Nested schema for **sfdc_params**:
	* `mutation_types` - (Required, List) Specifies a list of mutuation types for an object.
	  * Constraints: Allowable list items are: `All`, `Added`, `Removed`, `Changed`.
	* `object_name` - (Required, String) Specifies the name of the object.
	* `query_string` - (Optional, String) Specifies the query string to search records. Query string can be one or multiples clauses joined together by 'AND' or 'OR' claused.
	* `snapshot_id` - (Required, String) Specifies the id of the snapshot for the object.
* `sharepoint_params` - (Optional, Forces new resource, List) Specifies the request parameters to search for files/folders in document libraries.
Nested schema for **sharepoint_params**:
	* `category_types` - (Optional, List) Specifies a list of document library types. Only items within the given types will be returned.
	  * Constraints: Allowable list items are: `Document`, `Excel`, `Powerpoint`, `Image`, `OneNote`.
	* `creation_end_time_secs` - (Optional, Integer) Specifies the end time in Unix timestamp epoch in seconds when the file/folder is created.
	* `creation_start_time_secs` - (Optional, Integer) Specifies the start time in Unix timestamp epoch in seconds when the file/folder is created.
	* `include_files` - (Optional, Boolean) Specifies whether to include files in the response. Default is true.
	  * Constraints: The default value is `true`.
	* `include_folders` - (Optional, Boolean) Specifies whether to include folders in the response. Default is true.
	  * Constraints: The default value is `true`.
	* `o365_params` - (Optional, List) Specifies O365 specific params search request params to search for indexed items.
	Nested schema for **o365_params**:
		* `domain_ids` - (Optional, List) Specifies the domain Ids in which indexed items are searched.
		* `group_ids` - (Optional, List) Specifies the Group ids across which the indexed items needs to be searched.
		* `site_ids` - (Optional, List) Specifies the Sharepoint site ids across which the indexed items needs to be searched.
		* `teams_ids` - (Optional, List) Specifies the Teams ids across which the indexed items needs to be searched.
		* `user_ids` - (Optional, List) Specifies the user ids across which the indexed items needs to be searched.
	* `owner_names` - (Optional, List) Specifies the list of owner names to filter on owner of the file/folder.
	* `search_string` - (Optional, String) Specifies the search string to filter the files/folders. User can specify a wildcard character '*' as a suffix to a string where all item names are matched with the prefix string.
	* `size_bytes_lower_limit` - (Optional, Integer) Specifies the minimum size of the file in bytes.
	* `size_bytes_upper_limit` - (Optional, Integer) Specifies the maximum size of the file in bytes.
* `snapshot_tags` - (Optional, Forces new resource, List) "This field is deprecated. Please use mightHaveSnapshotTagIds.".
* `storage_domain_ids` - (Optional, Forces new resource, List) Specifies the Storage Domain ids to filter indexed objects for which Protection Groups are writing data to Cohesity Views on the specified Storage Domains.
* `tags` - (Optional, Forces new resource, List) "This field is deprecated. Please use mightHaveTagIds.".
* `tenant_id` - (Optional, Forces new resource, String) TenantId contains id of the tenant for which objects are to be returned.
* `uda_params` - (Optional, Forces new resource, List) Parameters required to search Universal Data Adapter objects.
Nested schema for **uda_params**:
	* `search_string` - (Required, String) Specifies the search string to search the Universal Data Adapter Objects.
	* `source_ids` - (Optional, List) Specifies a list of source ids. Only files found in these sources will be returned.
* `use_cached_data` - (Optional, Forces new resource, Boolean) Specifies whether we can serve the GET request from the read replica cache. There is a lag of 15 seconds between the read replica and primary data source.
* `x_ibm_tenant_id` - (Required, Forces new resource, String) Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.