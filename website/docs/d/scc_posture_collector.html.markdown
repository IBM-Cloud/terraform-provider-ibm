---
layout: "ibm"
page_title: "IBM : ibm_scc_posture_collector"
description: |-
  Get information about collector
subcategory: "Security and Compliance Center"
---

# ibm_scc_posture_collector

Provides a read-only data source for collector. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_scc_posture_collector" "collector" {
	collector_id = "collector_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `collector_id` - (Required, Forces new resource, String) The id for the given API.
  * Constraints: The maximum length is `20` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9]*$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the collector.
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

!> **Removal Notification** Resource Removal: Resource ibm_scc_posture_collector is deprecated and being removed.\n This resource will not be available from future release (v1.54.0).
