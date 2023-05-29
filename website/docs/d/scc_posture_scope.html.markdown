---
layout: "ibm"
page_title: "IBM : ibm_scc_posture_scope"
description: |-
  Get information about scope
subcategory: "Security and Compliance Center"
---

# ibm_scc_posture_scope

Provides a read-only data source for scope. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_scc_posture_scope" "scope" {
	scope_id = "scope_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `id` - (Required, Forces new resource, String) The id for the given API.
  * Constraints: The maximum length is `20` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9]*$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `scope_id` - The unique identifier of the scope.
* `cloud_type` - (Optional, String) Stores the value of scope_cloud_type .Will be displayed only when value exists.

* `cloud_type_id` - (Optional, Integer) Stores the value of scope_cloud_type_id .Will be displayed only when value exists.

* `collectors` - (Optional, List) Stores the value of collectors .Will be displayed only when value exists.
Nested scheme for **collectors**:
	* `approved_internet_gateway_ip` - (Optional, String) The approved internet gateway ip of the collector. This field will be populated only when collector is installed.
	* `approved_local_gateway_ip` - (Optional, String) The approved local gateway ip of the collector. This field will be populated only when collector is installed.
	* `collector_version` - (Optional, String) The collector version. This field is populated when collector is installed.
	* `created_at` - (Required, String) The ISO Date/Time the collector was created.
	* `created_by` - (Required, String) The id of the user that created the collector.
	* `credential_public_key` - (Optional, String) The credential public key.
	* `description` - (Required, String) The description of the collector.
	* `display_name` - (Required, String) The user-friendly name of the collector.
	* `enabled` - (Required, Boolean) Identifies whether the collector is enabled or not(deleted).
	* `failure_count` - (Required, Integer) The number of times the collector has failed.
	* `hostname` - (Optional, String) The collector host name. This field will be populated when collector is installed.This will have fully qualified domain name.
	* `id` - (Required, String) The id of the collector.
	* `image_version` - (Optional, String) The image version of the collector. This field is populated when collector is installed. ".
	* `install_path` - (Optional, String) The installation path of the collector. This field will be populated when collector is installed.The value will be folder path.
	* `is_public` - (Required, Boolean) Determines whether the collector endpoint is accessible on a public network.If set to `true`, the collector connects to resources in your account over a public network. If set to `false`, the collector connects to resources by using a private IP that is accessible only through the IBM Cloud private network.
	* `is_ubi_image` - (Optional, Boolean) Determines whether the collector has a Ubi image.
	* `last_failed_internet_gateway_ip` - (Optional, String) The failed internet gateway ip of the collector.
	* `last_failed_local_gateway_ip` - (Optional, String) The failed local gateway ip. This field will be populated only when collector is installed.
	* `last_heartbeat` - (Optional, String) Stores the heartbeat time of a controller . This value exists when collector is installed and running.
	* `managed_by` - (Required, String) The entity that manages the collector.
	  * Constraints: Allowable values are: `ibm`, `customer`.
	* `name` - (Required, String) The name of the collector.
	* `public_key` - (Optional, String) The public key of the collector.Will be used for ssl communciation between collector and orchestrator .This will be populated when collector is installed.
	* `registration_code` - (Required, String) The registration code of the collector.This is will be used for initial authentication during installation of collector.
	* `reset_reason` - (Optional, String) The reason for the collector reset .User resets the collector with a reason for reset. The reason entered by the user is saved in this field .
	* `reset_time` - (Optional, String) The ISO Date/Time of the collector reset. This value will be populated when a collector is reset. The data-time when the reset event is occured is captured in this field.
	* `status` - (Required, String) The status of collector.
	  * Constraints: Allowable values are: `ready_to_install`, `core_downloaded`, `approval_required`, `approved_download_in_progress`, `approved_install_in_progress`, `install_in_progress`, `installed`, `installed_credentials_required`, `installed_assigning_credentials`, `active`, `unable_to_connect`, `waiting_for_upgrade`, `suspended`, `installation_failed`.
	* `status_description` - (Required, String) The collector status.
	* `trial_expiry` - (Optional, String) The trial expiry. This holds the expiry date of registration_code. This field will be populated when collector is installed.
	* `type` - (Required, String) The type of the collector.
	  * Constraints: Allowable values are: `restricted`, `unrestricted`.
	* `updated_at` - (Required, String) The ISO Date/Time the collector was modified.
	* `updated_by` - (Required, String) The id of the user that modified the collector.
	* `use_private_endpoint` - (Required, Boolean) Whether the collector should use a public or private endpoint. This value is generated based on is_public field value during collector creation. If is_public is set to true, this value will be false.

* `collectors_by_type` - (Optional, Map) Stores the value of collectors_by_type .Will be displayed only when value exists.

* `correlation_id` - (Optional, String) A correlation_Id is created when a scope is created and discovery task is triggered or when a validation is triggered on a Scope. This is used to get the status of the task(discovery or validation).

* `created_at` - (Optional, String) Stores the value of scope_created_on .Will be displayed only when value exists.

* `created_by` - (Optional, String) Stores the value of scope_created_by .Will be displayed only when value exists.

* `credential_attributes` - (Optional, String) Stores the value of scope_credential_attributes .Will be displayed only when value exists.

* `credentials_by_sub_categeory_type` - (Optional, Map) Stores the value of scope_credentials_by_sub_categeory_type .Will be displayed only when value exists.

* `credentials_by_type` - (Optional, Map) Stores the value of scope_credentials_by_type .Will be displayed only when value exists.

* `description` - (Optional, String) Stores the value of scope_description .Will be displayed only when value exists.

* `discovery_method` - (Optional, String) Stores the value of scope_discovery_method .Will be displayed only when value exists.

* `discovery_methods` - (Optional, List) Stores the value of scope_discovery_methods .Will be displayed only when value exists.

* `discovery_setting_id` - (Optional, Integer) Stores the value of scope_discovery_setting_id .Will be displayed only when value exists.

* `enabled` - (Optional, Boolean) Stores the value of scope_enabled .Will be displayed only when value exists.

* `env_sub_category` - (Optional, String) Stores the value of scope_env_sub_category .Will be displayed only when value exists.

* `file_format` - (Optional, String) Stores the value of scope_file_format .Will be displayed only when value exists.

* `file_type` - (Optional, String) Stores the value of scope_file_type .Will be displayed only when value exists.

* `first_level_scoped_data` - (Optional, List) Stores the value of scope_first_level_scoped_data .Will be displayed only when value exists.
Nested scheme for **first_level_scoped_data**:
	* `scope` - (Optional, String) Stores the value of scope .
	* `scope_changed` - (Optional, Boolean) Stores the value of  scope_changed .
	* `scope_children` - (Optional, Map) Stores the value of scope_children .
	* `scope_collector_id` - (Optional, Integer) Stores the value of scope_collector_id .
	* `scope_discovery_status` - (Optional, Map) Stores the value of scope_discovery_status .
	* `scope_drift` - (Optional, String) Stores the value of  scope_drift .
	* `scope_fact_status` - (Optional, Map) Stores the value of scope_fact_status .
	* `scope_facts` - (Optional, String) Stores the value of scope_facts .
	* `scope_id` - (Optional, String) Stores the value of scope_id .
	* `scope_init_scope` - (Optional, String) Stores the value of scope_init_scope .
	* `scope_list_members` - (Optional, Map) Stores the value of scope_list_members .
	* `scope_new_found` - (Optional, Boolean) Stores the value of scope_new_found .
	* `scope_object` - (Optional, String) Stores the value of  scope_object .
	* `scope_overlay` - (Optional, String) Stores the value of scope_overlay .
	* `scope_parse_status` - (Optional, String) Stores the value of scope_parse_status .
	* `scope_properties` - (Optional, String) Stores the value of  scope_properties .
	* `scope_resource` - (Optional, String) Stores the value of scope_resource .
	* `scope_resource_attributes` - (Optional, Map) Stores the value of scope_resource_attributes .
	* `scope_resource_category` - (Optional, String) Stores the value of scope_resource_category .
	* `scope_resource_type` - (Optional, String) Stores the value of scope_resource_type .
	* `scope_transformed_facts` - (Optional, Map) Stores the value of scope_transformed_facts .

* `include_new_eagerly` - (Optional, Boolean) Stores the value of scope_include_new_eagerly .Will be displayed only when value exists.

* `interval` - (Optional, Integer) Stores the value of scope_freq .Will be displayed only when value exists.

* `is_discovery_scheduled` - (Optional, Boolean) Stores the value of scope_is_discovery_scheduled .Will be displayed only when value exists.

* `last_discover_completed_time` - (Optional, String) Stores the value of scope_last_discover_completed_time .Will be displayed only when value exists.

* `last_discover_start_time` - (Optional, String) Stores the value of scope_last_discover_start_time .Will be displayed only when value exists.

* `last_successful_discover_completed_time` - (Optional, String) Stores the value of scope_last_successful_discover_completed_time .Will be displayed only when value exists.

* `last_successful_discover_start_time` - (Optional, String) Stores the value of scope_last_successful_discover_start_time .Will be displayed only when value exists.

* `modified_at` - (Optional, String) Stores the value of scope_modified_on .Will be displayed only when value exists.

* `modified_by` - (Optional, String) Stores the value of scope_modified_by .Will be displayed only when value exists.

* `name` - (Required, String) Stores the value of scope_name .

* `org_id` - (Optional, Integer) Stores the value of scope_org_id .Will be displayed only when value exists.

* `partner_uuid` - (Optional, String) Stores the value of partner_uuid .Will be displayed only when value exists.

* `region_names` - (Optional, String) Stores the value of scope_region_names .Will be displayed only when value exists.

* `resource_groups` - (Optional, String) Stores the value of scope_resource_groups .Will be displayed only when value exists.

* `status` - (Optional, String) Stores the value of scope_status .Will be displayed only when value exists.
  * Constraints: Allowable values are: `pending`, `discovery_started`, `discovery_completed`, `error_in_discover`, `gateway_aborted`, `controller_aborted`, `not_accepted`, `waiting_for_refine`, `validation_started`, `validation_completed`, `sent_to_collector`, `discovery_in_progress`, `validation_in_progress`, `error_in_validation`, `discovery_result_posted_with_error`, `discovery_result_posted_no_error`, `validation_result_posted_with_error`, `validation_result_posted_no_error`, `fact_collection_started`, `fact_collection_in_progress`, `fact_collection_completed`, `error_in_fact_collection`, `fact_validation_started`, `fact_validation_in_progress`, `fact_validation_completed`, `error_in_fact_validation`, `abort_task_request_received`, `error_in_abort_task_request`, `abort_task_request_completed`, `user_aborted`, `abort_task_request_failed`, `cve_validation_started`, `cve_validation_completed`, `cve_validation_error`, `eol_validation_started`, `eol_validation_completed`, `eol_validation_error`, `cve_regular_validation_started`, `cve_regular_validation_completed`, `cve_regular_validation_error`, `eol_regular_validation_started`, `eol_regular_validation_completed`, `eol_regular_validation_error`, `cert_validation_started`, `cert_validation_completed`, `cert_validation_error`, `cert_regular_validation_started`, `cert_regular_validation_completed`, `cert_regular_validation_error`, `remediation_started`, `remediation_in_progress`, `error_in_remediation`, `remediation_completed`, `inventory_started`, `inventory_in_progress`, `inventory_completed`, `error_in_inventory`, `inventory_completed_with_error`, `location_change_aborted`.

* `status_msg` - (Optional, String) Stores the value of scope_status_msg .Will be displayed only when value exists.

* `status_updated_time` - (Optional, String) Stores the value of scope_status_updated_time .Will be displayed only when value exists.

* `sub_categories_by_type` - (Optional, Map) Stores the value of scope_sub_categories_by_type .Will be displayed only when value exists.

* `subset_selected` - (Optional, Boolean) Stores the value of scope_subset_selected .Will be displayed only when value exists.

* `task_type` - (Optional, String) Stores the value of scope_task_type .Will be displayed only when value exists.
  * Constraints: Allowable values are: `nop`, `discover`, `evidence`, `factcollection`, `script`, `tldiscover`, `subsetvalidate`, `factvalidation`, `aborttasks`, `cve_validation`, `eol_validation`, `cve_regular_validation`, `eol_regular_validation`, `cert_regular_validation`, `cert_validation`, `remediation`, `inventory`.

* `tasks` - (Optional, List) Stores the value of scope_tasks .Will be displayed only when value exists.
Nested scheme for **tasks**:
	* `task_created_by` - (Optional, String) Stores the value of task_created_by .
	* `task_derived_status` - (Optional, String) Stores the value of task_derived_status .
	  * Constraints: Allowable values are: `pending`, `discovery_started`, `discovery_completed`, `error_in_discover`, `gateway_aborted`, `controller_aborted`, `not_accepted`, `waiting_for_refine`, `validation_started`, `validation_completed`, `sent_to_collector`, `discovery_in_progress`, `validation_in_progress`, `error_in_validation`, `discovery_result_posted_with_error`, `discovery_result_posted_no_error`, `validation_result_posted_with_error`, `validation_result_posted_no_error`, `fact_collection_started`, `fact_collection_in_progress`, `fact_collection_completed`, `error_in_fact_collection`, `fact_validation_started`, `fact_validation_in_progress`, `fact_validation_completed`, `error_in_fact_validation`, `abort_task_request_received`, `error_in_abort_task_request`, `abort_task_request_completed`, `user_aborted`, `abort_task_request_failed`, `cve_validation_started`, `cve_validation_completed`, `cve_validation_error`, `eol_validation_started`, `eol_validation_completed`, `eol_validation_error`, `cve_regular_validation_started`, `cve_regular_validation_completed`, `cve_regular_validation_error`, `eol_regular_validation_started`, `eol_regular_validation_completed`, `eol_regular_validation_error`, `cert_validation_started`, `cert_validation_completed`, `cert_validation_error`, `cert_regular_validation_started`, `cert_regular_validation_completed`, `cert_regular_validation_error`, `remediation_started`, `remediation_in_progress`, `error_in_remediation`, `remediation_completed`, `inventory_started`, `inventory_in_progress`, `inventory_completed`, `error_in_inventory`, `inventory_completed_with_error`, `location_change_aborted`.
	* `task_discover_id` - (Optional, Integer) Stores the value of task_discover_id .
	* `task_gateway_id` - (Optional, Integer) Stores the value of task_gateway_id .
	* `task_gateway_name` - (Optional, String) Stores the value of task_gateway_name .
	* `task_gateway_schema_id` - (Optional, Integer) Stores the value of task_gateway_schema_id .
	* `task_id` - (Optional, Integer) Stores the value of task_id .
	* `task_logs` - (Optional, List) Stores the value of task_logs .
	Nested scheme for **task_logs**:
	* `task_schema_name` - (Optional, String) Stores the value of task_schema_name .
	* `task_start_time` - (Optional, Integer) Stores the value of task_start_time .
	* `task_status` - (Optional, String) Stores the value of task_status .
	  * Constraints: Allowable values are: `pending`, `discovery_started`, `discovery_completed`, `error_in_discover`, `gateway_aborted`, `controller_aborted`, `not_accepted`, `waiting_for_refine`, `validation_started`, `validation_completed`, `sent_to_collector`, `discovery_in_progress`, `validation_in_progress`, `error_in_validation`, `discovery_result_posted_with_error`, `discovery_result_posted_no_error`, `validation_result_posted_with_error`, `validation_result_posted_no_error`, `fact_collection_started`, `fact_collection_in_progress`, `fact_collection_completed`, `error_in_fact_collection`, `fact_validation_started`, `fact_validation_in_progress`, `fact_validation_completed`, `error_in_fact_validation`, `abort_task_request_received`, `error_in_abort_task_request`, `abort_task_request_completed`, `user_aborted`, `abort_task_request_failed`, `cve_validation_started`, `cve_validation_completed`, `cve_validation_error`, `eol_validation_started`, `eol_validation_completed`, `eol_validation_error`, `cve_regular_validation_started`, `cve_regular_validation_completed`, `cve_regular_validation_error`, `eol_regular_validation_started`, `eol_regular_validation_completed`, `eol_regular_validation_error`, `cert_validation_started`, `cert_validation_completed`, `cert_validation_error`, `cert_regular_validation_started`, `cert_regular_validation_completed`, `cert_regular_validation_error`, `remediation_started`, `remediation_in_progress`, `error_in_remediation`, `remediation_completed`, `inventory_started`, `inventory_in_progress`, `inventory_completed`, `error_in_inventory`, `inventory_completed_with_error`, `location_change_aborted`.
	* `task_status_msg` - (Optional, String) Stores the value of task_status_msg .
	* `task_task_type` - (Optional, String) Stores the value of task_task_type .
	  * Constraints: Allowable values are: `nop`, `discover`, `evidence`, `factcollection`, `script`, `tldiscover`, `subsetvalidate`, `factvalidation`, `aborttasks`, `cve_validation`, `eol_validation`, `cve_regular_validation`, `eol_regular_validation`, `cert_regular_validation`, `cert_validation`, `remediation`, `inventory`.
	* `task_updated_time` - (Optional, Integer) Stores the value of task_updated_time .

* `tld_credentail` - (Optional, List) Stores the value of ScopeDetailsCredential .
Nested scheme for **tld_credentail**:
	* `data` - (Optional, Map) Stores the value of credential_data .
	* `description` - (Optional, String) Stores the value of credential_description .
	* `display_fields` - (Optional, List) Details the fields on the credential. This will change as per credential type selected.
	Nested scheme for **display_fields**:
		* `auth_url` - (Optional, String) auth url of the Open Stack cloud.This is mandatory for Open Stack Credential type.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
		* `aws_arn` - (Optional, String) AWS arn value.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
		* `aws_client_id` - (Optional, String) AWS client Id.This is mandatory for AWS Cloud.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
		* `aws_client_secret` - (Optional, String) AWS client secret.This is mandatory for AWS Cloud.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.\\*,_\\s]*$/`.
		* `aws_region` - (Optional, String) AWS region.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
		* `azure_client_id` - (Optional, String) Azure client Id. This is mandatory for Azure Credential type.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
		* `azure_client_secret` - (Optional, String) Azure client secret.This is mandatory for Azure Credential type.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.\\*,_\\s]*$/`.
		* `azure_resource_group` - (Optional, String) Azure resource group.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
		* `azure_subscription_id` - (Optional, String) Azure subscription Id.This is mandatory for Azure Credential type.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
		* `database_name` - (Optional, String) Database name.This is mandatory for Database Credential type.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
		* `ibm_api_key` - (Optional, String) The IBM Cloud API Key. This is mandatory for IBM Credential Type.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.\\*,_\\s]*$/`.
		* `ms_365_client_id` - (Optional, String) The MS365 client Id.This is mandatory for Windows MS365 Credential type.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
		* `ms_365_client_secret` - (Optional, String) The MS365 client secret.This is mandatory for Windows MS365 Credential type.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
		* `ms_365_tenant_id` - (Optional, String) The MS365 tenantId.This is mandatory for Windows MS365 Credential type.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
		* `password` - (Optional, String) password of the user.This is mandatory for DataBase, Kerbros,OpenStack Credentials.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.\\*,_\\s]*$/`.
		* `pem_data` - (Optional, String) The base64 encoded data to associate with the PEM file.
		  * Constraints: The maximum length is `4000` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
		* `pem_file_name` - (Optional, String) The name of the PEM file.
		  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.\\*,_\\s]*$/`.
		* `project_domain_name` - (Optional, String) project domain name of the Open Stack cloud.This is mandatory for Open Stack Credential type.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
		* `project_name` - (Optional, String) Project name of the Open Stack cloud.This is mandatory for Open Stack Credential type.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
		* `user_domain_name` - (Optional, String) user domain name of the Open Stack cloud.This is mandatory for Open Stack Credential type.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
		* `username` - (Optional, String) username of the user.This is mandatory for DataBase, Kerbros,OpenStack Credentials.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
		* `winrm_authtype` - (Optional, String) Kerberos windows auth type.This is mandatory for Windows Kerberos Credential type.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
		* `winrm_port` - (Optional, String) Kerberos windows port.This is mandatory for Windows Kerberos Credential type.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9]*$/`.
		* `winrm_usessl` - (Optional, String) Kerberos windows ssl.This is mandatory for Windows Kerberos Credential type.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `gateway_key` - (Optional, String) Stores the value of credential_gateway_key .
	* `id` - (Optional, String) Stores the value of credential_id .
	* `is_enabled` - (Optional, Boolean) Stores the value of credential_is_enabled .
	* `name` - (Optional, String) Stores the value of credential_name .
	* `purpose` - (Optional, String) Stores the value of credential_purpose .
	* `type` - (Optional, String) Stores the value of credential_type .
	* `uuid` - (Optional, String) Stores the value of credential_uuid .
	* `version_timestamp` - (Optional, Map) Stores the value of credential_version_timestamp .

* `tld_credential_id` - (Optional, Integer) Stores the value of scope_tld_credential_id .Will be displayed only when value exists.

* `type` - (Optional, String) Stores the value of scope_type .Will be displayed only when value exists.
  * Constraints: Allowable values are: `validation`, `inventory`.

* `uuid` - (Optional, String) Stores the value of scope_uuid .Will be displayed only when value exists.

!> **Removal Notification** Resource Removal: Resource ibm_scc_posture_scope is deprecated and being removed.\n This resource will not be available from future release (v1.54.0).
