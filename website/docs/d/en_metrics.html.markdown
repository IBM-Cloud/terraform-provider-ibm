---
layout: "ibm"
page_title: "IBM : ibm_en_metrics"
description: |-
  Get information about en_metrics
subcategory: "Event Notifications"
---

# ibm_en_metrics

Provides a read-only data source to retrieve information about en_metrics. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_en_metrics" "en_metrics" {
	destination_type = "smtp_custom"
	gte = "gte"
	instance_id = "instance_id"
	lte = "lte"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `destination_id` - (Optional, String) Unique identifier for Destination.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]/`.
* `destination_type` - (Required, String) Destination type. Allowed values are [smtp_custom].
  * Constraints: Allowable values are: `smtp_custom`.
* `email_to` - (Optional, String) Receiver email id.
  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9\\._%+\\-]+@[A-Za-z0-9\\.\\-]+\\.[A-Za-z]{2,}/`.
* `gte` - (Required, String) GTE (greater than equal), start timestamp in UTC.
  * Constraints: The maximum length is `28` characters. The minimum length is `1` character. The value must match regular expression `/[0-9]{1,4}-[0-9]{1,2}-[0-9]{1,2}T[0-9]{1,2}:[0-9]{1,2}:[0-9]{1,2}Z/`.
* `instance_id` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.
  * Constraints: The maximum length is `256` characters. The minimum length is `10` characters. The value must match regular expression `/[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]/`.
* `lte` - (Required, String) LTE (less than equal), end timestamp in UTC.
  * Constraints: The maximum length is `28` characters. The minimum length is `1` character. The value must match regular expression `/[0-9]{1,4}-[0-9]{1,2}-[0-9]{1,2}T[0-9]{1,2}:[0-9]{1,2}:[0-9]{1,2}Z/`.
* `notification_id` - (Optional, String) Notification Id.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}/`.
* `source_id` - (Optional, String) Unique identifier for Source.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/[a-zA-Z0-9-:_]*/`.
* `subject` - (Optional, String) Email subject.
  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/[a-zA-Z0-9]/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the en_metrics.
* `metrics` - (List) array of metrics.
  * Constraints: The maximum length is `5` items. The minimum length is `5` items.
Nested schema for **metrics**:
	* `doc_count` - (Integer) doc count.
	* `histogram` - (List) Payload describing histogram.
	Nested schema for **histogram**:
		* `buckets` - (List) List of buckets.
		  * Constraints: The maximum length is `48` items. The minimum length is `0` items.
		Nested schema for **buckets**:
			* `doc_count` - (Integer) Total count.
			* `key_as_string` - (String) Timestamp.
	* `key` - (String) key.
	  * Constraints: Allowable values are: `bounced`, `deferred`, `opened`, `success`, `submitted`. The maximum length is `255` characters. The minimum length is `1` character.

