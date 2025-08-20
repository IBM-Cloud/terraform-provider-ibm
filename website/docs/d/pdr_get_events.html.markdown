---
layout: "ibm"
page_title: "IBM : ibm_pdr_get_events"
description: |-
  Get information about pdr_get_events
subcategory: "DrAutomation Service"
---

# ibm_pdr_get_events

Provides a read-only data source to retrieve information about pdr_get_events. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_pdr_get_events" "pdr_get_events" {
	from_time = "2025-06-19T00:00:00Z"
	provision_id = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"
	time = "2025-06-19T23:59:59Z"
	to_time = "2025-06-19T23:59:59Z"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `accept_language` - (Optional, String) The language requested for the return document.
* `from_time` - (Optional, String) A from query time in either ISO 8601 or unix epoch format.
* `if_none_match` - (Optional, String) ETag for conditional requests (optional).
* `provision_id` - (Required, Forces new resource, String) provision id.
* `time` - (Optional, String) (deprecated - use from_time) A time in either ISO 8601 or unix epoch format.
* `to_time` - (Optional, String) A to query time in either ISO 8601 or unix epoch format.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the pdr_get_events.
* `events` - (List) Events.
Nested schema for **events**:
	* `action` - (String) Type of action for this event.
	* `api_source` - (String) Source of API when it being executed.
	* `event_id` - (String) ID of the Activity.
	* `level` - (String) Level of the event (notice, info, warning, error).
	  * Constraints: Allowable values are: `notice`, `info`, `warning`, `error`.
	* `message` - (String) The (translated) message of the event.
	* `message_data` - (Map)
	* `metadata` - (Map) Any metadata associated with the event.
	* `resource` - (String) Type of resource for this event.
	* `time` - (String) Time of activity in ISO 8601 - RFC3339.
	* `timestamp` - (String) Time of activity in unix epoch.
	* `user` - (List)
	Nested schema for **user**:
		* `email` - (String) Email of the User.
		* `name` - (String) Name of the User.
		* `user_id` - (String) ID of user who created/caused the event.

