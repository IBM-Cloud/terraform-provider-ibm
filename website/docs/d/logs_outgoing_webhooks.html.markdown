---
layout: "ibm"
page_title: "IBM : ibm_logs_outgoing_webhooks"
description: |-
  Get information about logs_outgoing_webhooks
subcategory: "Cloud Logs"
---

~> **Beta:** This resource is in Beta, and is subject to change.

# ibm_logs_outgoing_webhooks

Provides a read-only data source to retrieve information about logs_outgoing_webhooks. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_logs_outgoing_webhooks" "logs_outgoing_webhooks_instance" {
  instance_id = ibm_resource_instance.logs_instance.guid
  region      = ibm_resource_instance.logs_instance.location
  type        = "ibm_event_notifications"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `instance_id` - (Required, String)  Cloud Logs Instance GUID.
* `region` - (Optional, String) Cloud Logs Instance Region.
* `type` - (Optional, String) Outbound webhook type.
  * Constraints: The default value is `ibm_event_notifications`. Allowable values are: `ibm_event_notifications`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the logs_outgoing_webhooks.
* `outgoing_webhooks` - (List) List of deployed outbound webhooks.
  * Constraints: The maximum length is `4096` items. The minimum length is `1` item.
Nested schema for **outgoing_webhooks**:
	* `created_at` - (String) The creation time of the outbound webhook.
	* `external_id` - (Integer) The external ID of the outbound webhook.
	  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
	* `id` - (String) The ID of the outbound webhook.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$/`.
	* `name` - (String) The type of the outbound webhook.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
	* `updated_at` - (String) The update time of the outbound webhook.
	* `url` - (String) The URL of the outbound webhook.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.

