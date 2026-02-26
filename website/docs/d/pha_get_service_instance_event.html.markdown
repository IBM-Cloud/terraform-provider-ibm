---
layout: "ibm"
page_title: "IBM : ibm_pha_get_service_instance_event"
description: |-
  Get information about pha_get_service_instance_event
subcategory: "PowerhaAutomation Service"
---

# ibm_pha_get_service_instance_event

Provides a read-only data source to retrieve information about a pha_get_service_instance_event. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_pha_get_service_instance_event" "pha_get_service_instance_event" {
	accept_language = "en-US"
	event_id = "00116b2a-9326-4024-839e-fb5364b76898"
	if_none_match = "abcdef"
	pha_instance_id = "8eefautr-4c02-0009-0086-8bd4d8cf61b6"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `accept_language` - (Optional, String) The language requested for the return document.
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_,;=.*]+$/`.
* `event_id` - (Required, Forces new resource, String) Event ID.
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-]+$/`.
* `if_none_match` - (Optional, String) ETag for conditional requests (optional).
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_,;=.*]+$/`.
* `pha_instance_id` - (Required, Forces new resource, String) instance id of instance to provision.
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-]+$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the pha_get_service_instance_event.
* `action` - (String) Type of action for this event.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `api_source` - (String) Source of API when it being executed.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `level` - (String) Level of the event (notice, info, warning, error).
  * Constraints: Allowable values are: `notice`, `info`, `warning`, `error`.
* `message` - (String) The (translated) message of the event.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._: -]+$/`.
* `message_data` - (Map) Dynamic key-value data related to the event message.
* `meta_data` - (Map) Metadata providing additional context for the event.
* `resource` - (String) Type of resource for this event.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `time` - (String) Time of activity in ISO 8601 - RFC3339.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `time_stamp` - (String) Timestamp indicating when the event occurred.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `user` - (List) events for pha user.
Nested schema for **user**:
* `email` - (String) Email of the User.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,}$/`.
* `name` - (String) Name of the User.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `user_id` - (String) ID of user who created/caused the event.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.

