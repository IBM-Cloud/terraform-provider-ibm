---
layout: "ibm"
page_title: "IBM : ibm_scc_posture_scopes"
description: |-
  Get information about list_scopes
subcategory: "Security and Compliance Center"
---

# ibm_scc_posture_scopes

Provides a read-only data source for list_scopes. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_scc_posture_scopes" "list_scopes" {
}
```


## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the list_scopes.
* `first` - (List) The URL of a page.
Nested scheme for **first**:
	* `href` - (String) The URL of a page.

* `last` - (List) The URL of a page.
Nested scheme for **last**:
	* `href` - (String) The URL of a page.

* `limit` - (Integer) The number of scopes displayed per page.

* `next` - (List) The URL of a page.
Nested scheme for **next**:
	* `href` - (String) The URL of a page.

* `offset` - (Integer) The offset of the page.

* `previous` - (List) The URL of a page.
Nested scheme for **previous**:
	* `href` - (String) The URL of a page.

* `scopes` - (List) Scopes.
Nested scheme for **scopes**:
	* `collectors` - (List) Stores the value of collectors .Will be displayed only when value exists.
	Nested scheme for **collectors**:
		* `approved_internet_gateway_ip` - (String) The approved internet gateway ip of the collector. This field will be populated only when collector is installed.
		* `approved_local_gateway_ip` - (String) The approved local gateway ip of the collector. This field will be populated only when collector is installed.
		* `collector_version` - (String) The collector version. This field is populated when collector is installed.
		* `created_at` - (String) The ISO Date/Time the collector was created.
		* `created_by` - (String) The id of the user that created the collector.
		* `credential_public_key` - (String) The credential public key.
		* `description` - (String) The description of the collector.
		* `display_name` - (String) The user-friendly name of the collector.
		* `enabled` - (Boolean) Identifies whether the collector is enabled or not(deleted).
		* `failure_count` - (Integer) The number of times the collector has failed.
		* `hostname` - (String) The collector host name. This field will be populated when collector is installed.This will have fully qualified domain name.
		* `id` - (String) The id of the collector.
		* `image_version` - (String) The image version of the collector. This field is populated when collector is installed. ".
		* `install_path` - (String) The installation path of the collector. This field will be populated when collector is installed.The value will be folder path.
		* `is_public` - (Boolean) Determines whether the collector endpoint is accessible on a public network.If set to `true`, the collector connects to resources in your account over a public network. If set to `false`, the collector connects to resources by using a private IP that is accessible only through the IBM Cloud private network.
		* `is_ubi_image` - (Boolean) Determines whether the collector has a Ubi image.
		* `last_failed_internet_gateway_ip` - (String) The failed internet gateway ip of the collector.
		* `last_failed_local_gateway_ip` - (String) The failed local gateway ip. This field will be populated only when collector is installed.
		* `last_heartbeat` - (String) Stores the heartbeat time of a controller . This value exists when collector is installed and running.
		* `managed_by` - (String) The entity that manages the collector.
		  * Constraints: Allowable values are: `ibm`, `customer`.
		* `name` - (String) The name of the collector.
		* `public_key` - (String) The public key of the collector.Will be used for ssl communciation between collector and orchestrator .This will be populated when collector is installed.
		* `registration_code` - (String) The registration code of the collector.This is will be used for initial authentication during installation of collector.
		* `reset_reason` - (String) The reason for the collector reset .User resets the collector with a reason for reset. The reason entered by the user is saved in this field .
		* `reset_time` - (String) The ISO Date/Time of the collector reset. This value will be populated when a collector is reset. The data-time when the reset event is occured is captured in this field.
		* `status` - (String) The status of collector.
		  * Constraints: Allowable values are: `ready_to_install`, `core_downloaded`, `approval_required`, `approved_download_in_progress`, `approved_install_in_progress`, `install_in_progress`, `installed`, `installed_credentials_required`, `installed_assigning_credentials`, `active`, `unable_to_connect`, `waiting_for_upgrade`, `suspended`, `installation_failed`.
		* `status_description` - (String) The collector status.
		* `trial_expiry` - (String) The trial expiry. This holds the expiry date of registration_code. This field will be populated when collector is installed.
		* `type` - (String) The type of the collector.
		  * Constraints: Allowable values are: `restricted`, `unrestricted`.
		* `updated_at` - (String) The ISO Date/Time the collector was modified.
		* `updated_by` - (String) The id of the user that modified the collector.
		* `use_private_endpoint` - (Boolean) Whether the collector should use a public or private endpoint. This value is generated based on is_public field value during collector creation. If is_public is set to true, this value will be false.
	* `created_at` - (String) The time that the scope was created in UTC.
	* `created_by` - (String) The user who created the scope.
	* `credential_type` - (String) The environment that the scope is targeted to.
	  * Constraints: Allowable values are: `ibm`, `aws`, `azure`, `on_premise`, `hosted`, `services`, `openstack`, `gcp`.
	* `description` - (String) A detailed description of the scope.
	* `enabled` - (Boolean) Indicates whether scope is enabled/disabled.
	* `id` - (String) An auto-generated unique identifier for the scope.
	* `modified_by` - (String) The user who most recently modified the scope.
	* `name` - (String) A unique name for your scope.
	* `updated_at` - (String) The time that the scope was last modified in UTC.
	* `uuid` - (String) Stores the value of scope_uuid .

* `total_count` - (Integer) The total number of scopes. This value is 0 if no scopes are available and below fields will not be available in that case.

!> **Removal Notification** Resource Removal: Resource ibm_scc_posture_scopes is deprecated and being removed.\n This resource will not be available from future release (v1.54.0).
