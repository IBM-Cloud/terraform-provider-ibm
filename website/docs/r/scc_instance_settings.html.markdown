---
layout: "ibm"
page_title: "IBM : ibm_scc_instance_settings"
description: |-
  Manages scc_instance_settings.
subcategory: "Admin"
---

# ibm_scc_instance_settings

Create, update, and delete scc_instance_settingss with this resource.

## Example Usage

```hcl
resource "ibm_scc_instance_settings" "scc_instance_settings_instance" {
  event_notifications {
		instance_crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/ff88f007f9ff4622aac4fbc0eda36255:7199ae60-a214-4dd8-9bf7-ce571de49d01::"
		updated_on = "2021-01-31T09:44:12Z"
		source_id = "crn:v1:bluemix:public:event-notifications:us-south:a/ff88f007f9ff4622aac4fbc0eda36255:b8b07245-0bbe-4478-b11c-0dce523105fd::"
		source_description = "source_description"
		source_name = "source_name"
  }
  object_storage {
		instance_crn = "instance_crn"
		bucket = "bucket"
		bucket_location = "bucket_location"
		bucket_endpoint = "bucket_endpoint"
		updated_on = "2021-01-31T09:44:12Z"
  }
  x_correlation_id = "1a2b3c4d-5e6f-4a7b-8c9d-e0f1a2b3c4d5"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `event_notifications` - (Optional, List) The Event Notifications settings.
Nested schema for **event_notifications**:
	* `instance_crn` - (Optional, String) The Event Notifications instance CRN.
	  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}|$/`.
	* `source_description` - (Optional, String) The description of the source of the Event Notifications.
	  * Constraints: The default value is `This source is used for integration with IBM Cloud Security and Compliance Center.`. The maximum length is `512` characters. The minimum length is `1` character.
	* `source_id` - (Optional, String) The connected Security and Compliance Center instance CRN.
	  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/([A-Za-z0-9]+(:[A-Za-z0-9]+)+)/`.
	* `source_name` - (Optional, String) The name of the source of the Event Notifications.
	  * Constraints: The default value is `compliance`. The maximum length is `512` characters. The minimum length is `1` character.
	* `updated_on` - (Optional, String) The date when the Event Notifications connection was updated.
* `object_storage` - (Optional, List) The Cloud Object Storage settings.
Nested schema for **object_storage**:
	* `bucket` - (Optional, String) The connected Cloud Object Storage bucket name.
	  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z]+|/`.
	* `bucket_endpoint` - (Optional, String) The connected Cloud Object Storage bucket endpoint.
	  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/([A-Za-z0-9-]+)/`.
	* `bucket_location` - (Optional, String) The connected Cloud Object Storage bucket location.
	  * Constraints: The maximum length is `32` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z]+/`.
	* `instance_crn` - (Optional, String) The connected Cloud Object Storage instance CRN.
	  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}|$/`.
	* `updated_on` - (Optional, String) The date when the bucket connection was updated.
* `settings_id` - (Optional, String) 

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the scc_instance_settings.


## Import

You can import the `ibm_scc_instance_settings` resource by using `settings_id`. The unique identifier of the scc_instance_settings.

# Syntax
```
$ terraform import ibm_scc_instance_settings.scc_instance_settings <settings_id>
```
