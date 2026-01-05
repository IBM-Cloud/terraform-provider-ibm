---
layout: "ibm"
page_title: "IBM : ibm_pdr_get_event"
description: |-
  Get information about pdr_get_event
subcategory: "DrAutomation Service"
---

# ibm_pdr_get_event
Retrieves the details of a specific event for the given service instance provision ID.

## Example Usage

```hcl
data "ibm_pdr_get_event" "pdr_get_event" {
	event_id = "00116b2a-9326-4024-839e-fb5364b76898"
	instance_id = "123456d3-1122-3344-b67d-4389b44b7bf9"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `accept_language` - (Optional, String) The language requested for the return document.(ex., en,it,fr,es,de,ja,ko,pt-BR,zh-HANS,zh-HANT)
* `event_id` - (Required, Forces new resource, String) Event ID.
* `instance_id` - (Required, Forces new resource, String) ID of the service instance.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the pdr_get_event.
* `action` - (String) Type of action for this event.
* `api_source` - (String) Source of API when it being executed.
* `level` - (String) Level of the event (notice, info, warning, error).
  * Constraints: Allowable values are: `notice`, `info`, `warning`, `error`.
* `message` - (String) The (translated) message of the event.
* `message_data` - (Map) A flexible schema placeholder to allow any JSON value (aligns with interface{} in Go).
* `metadata` - (Map) A flexible schema placeholder to allow any JSON value (aligns with interface{} in Go).
* `resource` - (String) Type of resource for this event.
* `time` - (String) Time of activity in ISO 8601 - RFC3339.
* `timestamp` - (String) Time of activity in unix epoch.
* `user` - (List) Information about a user associated with an event.
Nested schema for **user**:
	* `email` - (String) Email of the User.
	* `name` - (String) Name of the User.
	* `user_id` - (String) ID of user who created/caused the event.
