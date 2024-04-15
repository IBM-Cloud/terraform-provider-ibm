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
  ibm_event_notifications {
		event_notifications_instance_id = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
		region_id = "region_id"
  }
  name = "name"
  type = "ibm_event_notifications"
  url = "url"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `instance_id` - (Required, Forces new resource, String)  Cloud Logs Instance GUID.
* `region` - (Optional, Forces new resource, String) Cloud Logs Instance Region.
* `endpoint_type` - (Optional, String) Cloud Logs Instance Endpoint type. Allowed values `public` and `private`.
* `ibm_event_notifications` - (Optional, List) The configuration of an IBM Event Notifications outbound webhook.
Nested schema for **ibm_event_notifications**:
	* `event_notifications_instance_id` - (Required, String) The instance ID of the IBM Event Notifications configuration.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$/`.
	* `region_id` - (Required, String) The region ID of the IBM Event Notifications configuration.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `4` characters. The value must match regular expression `/^[a-z]{2}-[a-z]+-\\d+$/`.
* `name` - (Required, String) The name of the outbound webhook.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
* `type` - (Required, String) Outbound webhook type.
  * Constraints: The default value is `ibm_event_notifications`. Allowable values are: `ibm_event_notifications`.
* `url` - (Required, String) The URL of the outbound webhook.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^https?:\/\/.*$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the logs_outgoing_webhook resource.
* `webhook_id` - The unique identifier of the logs_outgoing_webhook.
* `created_at` - (String) The creation time of the outbound webhook.
* `external_id` - (Integer) The external ID of the outbound webhook.
  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
* `updated_at` - (String) The update time of the outbound webhook.


## Import

You can import the `ibm_logs_outgoing_webhook` resource by using `id`. `id` combination of `region`, `instance_id` and `webhook_id`.

# Syntax
<pre>
$ terraform import ibm_logs_outgoing_webhook.logs_outgoing_webhook <region>/<instance_id>/<webhook_id>;
</pre>
