---
layout: "ibm"
page_title: "IBM : ibm_pha_list_service_instance_events"
description: |-
  Get information about pha_list_service_instance_events
subcategory: "PowerhaAutomation Service"
---

# ibm_pha_list_service_instance_events

Provides a read-only data source to retrieve information about pha_list_service_instance_events. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_pha_list_service_instance_events" "pha_list_service_instance_events" {
	accept_language = "en-US"
	from_time = "2025-06-19T00:00:00Z"
	if_none_match = "abcdef"
	pha_instance_id = "8eefautr-4c02-0009-0086-8bd4d8cf61b6"
	time = "2025-06-19T23:59:59Z"
	to_time = "2025-06-19T23:59:59Z"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `accept_language` - (Optional, String) The language requested for the return document.
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_,;=.*]+$/`.
* `from_time` - (Optional, String) A from query time in either ISO 8601 or unix epoch format.
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-\\:T]+$/`.
* `if_none_match` - (Optional, String) ETag for conditional requests (optional).
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_,;=.*]+$/`.
* `pha_instance_id` - (Required, Forces new resource, String) instance id of instance to provision.
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-]+$/`.
* `time` - (Optional, String) (deprecated - use from_time) A time in either ISO 8601 or unix epoch format.
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9T:\\-Z]+$/`.
* `to_time` - (Optional, String) A to query time in either ISO 8601 or unix epoch format.
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9T:\\-Z]+$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the pha_list_service_instance_events.
* `events` - (List) Pha automation Events.
  * Constraints: The maximum length is `16` items. The minimum length is `0` items.
Nested schema for **events**:
	* `action` - (String) Type of action for this event.
	  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `api_source` - (String) Source of API when it being executed.
	  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `event_id` - (String) ID of the Activity.
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

