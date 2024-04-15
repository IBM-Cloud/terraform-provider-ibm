---
layout: "ibm"
page_title: "IBM : ibm_logs_outgoing_webhooks"
description: |-
  Get information about logs_outgoing_webhooks
subcategory: "Cloud Logs"
---

# ibm_logs_outgoing_webhooks

Provides a read-only data source to retrieve information about logs_outgoing_webhooks. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_logs_outgoing_webhooks" "logs_outgoing_webhooks" {
}
```

## Argument Reference

You can specify the following arguments for this data source.

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

