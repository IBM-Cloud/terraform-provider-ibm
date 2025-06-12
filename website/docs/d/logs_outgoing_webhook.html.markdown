---
layout: "ibm"
page_title: "IBM : ibm_logs_outgoing_webhook"
description: |-
  Get information about logs_outgoing_webhook
subcategory: "Cloud Logs"
---


# ibm_logs_outgoing_webhook

Provides a read-only data source to retrieve information about a logs_outgoing_webhook. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_logs_outgoing_webhook" "logs_outgoing_webhook_instance" {
  instance_id              = ibm_logs_outgoing_webhook.logs_outgoing_webhook_instance.instance_id
  region                   = ibm_logs_outgoing_webhook.logs_outgoing_webhook_instance.region
  logs_outgoing_webhook_id =ibm_logs_outgoing_webhook.logs_outgoing_webhook_instance.webhook_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `instance_id` - (Required, String)  Cloud Logs Instance GUID.
* `region` - (Optional, String) Cloud Logs Instance Region.* `logs_outgoing_webhook_id` - (Required, Forces new resource, String) The ID of the Outbound Integration to delete.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the logs_outgoing_webhook.
* `created_at` - (String) The creation time of the Outbound Integration.

* `external_id` - (Integer) The external ID of the Outbound Integration, for connecting with other parts of the system.
  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.

* `ibm_event_notifications` - (List) The configuration of the IBM Event Notifications Outbound Integration.
Nested schema for **ibm_event_notifications**:
	* `event_notifications_instance_id` - (String) The ID of the selected IBM Event Notifications instance.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
	* `region_id` - (String) The region ID of the selected IBM Event Notifications instance.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `4` characters. The value must match regular expression `/^[a-z]{2}-[a-z]+$/`.
	* `endpoint_type` - (String) The endpoint type of integration.
  * `source_id` - (String) The ID of the created source in the IBM Event Notifications instance. Corresponds to the Cloud Logs instance crn. Not required when creating an Outbound Integration.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
	* `source_name` - (String) The name of the created source in the IBM Event Notifications instance. Not required when creating an Outbound Integration.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.

* `name` - (String) The name of the Outbound Integration.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.

* `type` - (String) The type of the deployed Outbound Integrations to list.
  * Constraints: Allowable values are: `ibm_event_notifications`.

* `updated_at` - (String) The update time of the Outbound Integration.

* `url` - (String) The URL of the Outbound Integration. Null for IBM Event Notifications integration.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.

