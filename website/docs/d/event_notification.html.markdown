---
layout: "ibm"
page_title: "IBM : ibm_event_notification"
description: |-
  Get information about Get Event Notifications Integration response
subcategory: "Projects API Specification"
---

# ibm_event_notification

Provides a read-only data source for Get Event Notifications Integration response. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_event_notification" "event_notification" {
	id = "id"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `id` - (Required, Forces new resource, String) The ID of the project, which uniquely identifies it.
  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the Get Event Notifications Integration response.
* `description` - (String) A description of the instance of the event.
  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.

* `enabled` - (Boolean) The status of instance of the event.

* `name` - (String) The name of the instance of the event.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s).+\\S$/`.

* `topic_count` - (Integer) The topic count of the instance of the event.

* `topic_names` - (List) The topic names of the instance of the event.
  * Constraints: The list items must match regular expression `/^(?!\\s).+\\S$/`. The maximum length is `10000` items. The minimum length is `0` items.

* `type` - (String) The type of the instance of event.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s).+\\S$/`.

* `updated_at` - (String) A date/time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date-time format as specified by RFC 3339.

