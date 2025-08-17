---
layout: "ibm"
page_title: "IBM : ibm_logs_outgoing_webhook"
description: |-
  Manages logs_outgoing_webhook.
subcategory: "Cloud Logs"
---


# ibm_logs_outgoing_webhook

Create, update, and delete logs_outgoing_webhooks with this resource.

## Example Usage

```hcl
resource "ibm_logs_outgoing_webhook" "logs_outgoing_webhook_instance" {
  instance_id = ibm_resource_instance.logs_instance.guid
  region      = ibm_resource_instance.logs_instance.location
  name        = "example-webhook"
  type        = "ibm_event_notifications"
  ibm_event_notifications {
    event_notifications_instance_id = "6b33da73-28b6-4201-bfea-b2054bb6ae8a"
    region_id                       = "us-south"
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `instance_id` - (Required, Forces new resource, String)  Cloud Logs Instance GUID.
* `region` - (Optional, Forces new resource, String) Cloud Logs Instance Region.
* `endpoint_type` - (Optional, String) Cloud Logs Instance Endpoint type. Allowed values `public` and `private`.
* `ibm_event_notifications` - (Optional, List) The configuration of the IBM Event Notifications Outbound Integration.
Nested schema for **ibm_event_notifications**:
	* `event_notifications_instance_id` - (Required, String) The ID of the selected IBM Event Notifications instance.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
	* `region_id` - (Required, String) The region ID of the selected IBM Event Notifications instance.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `4` characters. The value must match regular expression `/^[a-z]{2}-[a-z]+$/`.
	* `endpoint_type` - (Optional, String) The endpoint type of integration. Allowed values: `private` and `public`. Default is `public`.
	* `source_id` - (Optional, String) The ID of the created source in the IBM Event Notifications instance. Corresponds to the Cloud Logs instance crn. Not required when creating an Outbound Integration.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
	* `source_name` - (Optional, String) The name of the created source in the IBM Event Notifications instance. Not required when creating an Outbound Integration.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
* `name` - (Required, String) The name of the Outbound Integration.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
* `type` - (Required, String) The type of the deployed Outbound Integrations to list.
  * Constraints: Allowable values are: `ibm_event_notifications`.
* `url` - (Optional, String) The URL of the Outbound Integration. Null for IBM Event Notifications integration.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the logs_outgoing_webhook resource.
* `webhook_id` - The unique identifier of the logs_outgoing_webhook.
* `created_at` - (String) The creation time of the Outbound Integration.
* `external_id` - (Integer) The external ID of the Outbound Integration, for connecting with other parts of the system.
  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
* `updated_at` - (String) The update time of the Outbound Integration.


## Import

You can import the `ibm_logs_outgoing_webhook` resource by using `id`. `id` combination of `region`, `instance_id` and `webhook_id`.

# Syntax
<pre>
$ terraform import ibm_logs_outgoing_webhook.logs_outgoing_webhook < region >/< instance_id >/< webhook_id >;
</pre>

# Example
```
$ terraform import ibm_logs_outgoing_webhook.logs_outgoing_webhook eu-gb/3dc02998-0b50-4ea8-b68a-4779d716fa1f/585bea36-bdd1-4bfb-9a26-51f1f8a12660
```
