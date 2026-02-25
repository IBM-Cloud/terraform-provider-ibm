---
layout: "ibm"
page_title: "IBM : ibm_en_bounce_metrics"
description: |-
  Get information about en_bounce_metrics
subcategory: "Event Notifications"
---

# ibm_en_bounce_metrics

Provides a read-only data source to retrieve information about en_bounce_metrics. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_en_bounce_metrics" "en_bounce_metrics" {
	destination_type = "smtp_custom"
	gte = "gte"
	instance_id = "instance_id"
	lte = "lte"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `destination_id` - (Optional, String) Unique identifier for Destination.
* `destination_type` - (Required, String) Destination type. Allowed values are [smtp_custom].
* `email_to` - (Optional, String) Receiver email id.
* `gte` - (Required, String) GTE (greater than equal), start timestamp in UTC.
* `instance_id` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.
* `lte` - (Required, String) LTE (less than equal), end timestamp in UTC.
* `notification_id` - (Optional, String) Notification Id.
* `source_id` - (Optional, String) Unique identifier for Source.
* `subject` - (Optional, String) Email subject.
* `subscription_id` - (Optional, String) Unique identifier for Subscription.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the en_bounce_metrics.
* `metrics` - (List) array of bounce metrics.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **metrics**:
	* `email_address` - (String) Email address.
	* `error_message` - (String) Error message.
	* `ip_address` - (String) IP address.
	* `subject` - (String) Subject.
	* `subscription_id` - (String) Subscription ID.
	* `timestamp` - (String) Bounced at.
* `total_count` - (Integer) total number of bounce metrics.
  * Constraints: The minimum value is `0`.

