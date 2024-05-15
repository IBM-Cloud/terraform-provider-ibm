---
layout: "ibm"
page_title: "IBM : ibm_logs_outgoing_webhook"
description: |-
  Get information about logs_outgoing_webhook
subcategory: "Cloud Logs"
---

~> **Beta:** This resource is in Beta, and is subject to change.

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
* `region` - (Optional, String) Cloud Logs Instance Region.
* `logs_outgoing_webhook_id` - (Required, Forces new resource, String) Outbound webhook ID.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the logs_outgoing_webhook.
* `created_at` - (String) The creation time of the outbound webhook.

* `external_id` - (Integer) The external ID of the outbound webhook.
  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.

* `ibm_event_notifications` - (List) The configuration of an IBM Event Notifications outbound webhook.
Nested schema for **ibm_event_notifications**:
	* `event_notifications_instance_id` - (String) The instance ID of the IBM Event Notifications configuration.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$/`.
	* `region_id` - (String) The region ID of the IBM Event Notifications configuration.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `4` characters. The value must match regular expression `/^[a-z]{2}-[a-z]+-\\d+$/`.

* `name` - (String) The name of the outbound webhook.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.

* `type` - (String) Outbound webhook type.
  * Constraints: The default value is `ibm_event_notifications`. Allowable values are: `ibm_event_notifications`.

* `updated_at` - (String) The update time of the outbound webhook.

* `url` - (String) The URL of the outbound webhook.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^https?:\/\/.*$/`.

