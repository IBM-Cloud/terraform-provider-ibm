---

layout: "ibm"
page_title: "IBM : ibm_pdr_event"
description: |-
Get information about pdr_event
subcategory: "HA and DR Automation for IBM® Power® Virtual Server API reference"
--------------------------------------------------------------------------------

# ibm_pdr_event

Retrieves the details of a specific event for the given service instance.

## Example Usage

```hcl
data "ibm_pdr_event" "pdr_event" {
	accept_language = "en-US"
	event_id = "00116b2a-9326-4024-839e-fb5364b76898"
	instance_id = "123456d3-1122-3344-b67d-4389b44b7bf9"
}
```
### Path Parameters

* `instance_id` - (Required, Forces new resource, String) The unique identifier of the DR service instance.
  * Constraints: The maximum length is `1048` characters. The minimum length is `36` characters. The value must match regular expression `/^.*$/`.

### Query Parameters

* `accept_language` - (Optional, Forces new resource, String) The language in which the response should be returned.
* `accepts_incomplete` - (Optional, Forces new resource, Boolean) Indicates whether the request can be accepted before the operation is complete.
* `event_id` - (Required, Forces new resource, String) Event ID.
 * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - (String) Unique identifier of the API key.

  * Constraints: The maximum length is `512` characters. The minimum length is `20` characters. The value must match regular expression `/^crn:v1:[a-zA-Z0-9\\-_]+:public:power-dr-automation:[a-zA-Z0-9\\-_]+:[a-zA-Z0-9\\-_\/]+:[a-zA-Z0-9\\-_]+::$/`.
* `action` - (String) Type of action for this event.
  * Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
* `api_source` - (String) Source of API when it being executed.
  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
* `event_id` - (String) ID of the Activity.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
* `href` - (String) Resource reference.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
* `level` - (String) Level of the event (notice, info, warning, error).
  * Constraints: Allowable values are: `notice`, `info`, `warning`, `error`.
* `message` - (String) The (translated) message of the event.
  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^[\\x20-\\x7E]*$/`.
* `message_data` - (Map) Any message data associated with the event.
* `metadata` - (Map) Any metadata associated with the event.
* `resource` - (String) Type of resource for this event.
  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
* `time` - (String) Time of activity in ISO 8601 - RFC3339.
  * Constraints: The maximum length is `1048` characters. The minimum length is `20` characters. The value must match regular expression `/^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[+\\-]\\d{2}:\\d{2})$/`.
* `timestamp` - (String) Time of activity in ISO 8601 - RFC3339.
  * Constraints: The maximum length is `1048` characters. The minimum length is `20` characters. The value must match regular expression `/^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[+\\-]\\d{2}:\\d{2})$/`.
* `user` - (Object) Container object holding a list of user event.

Nested schema for **user**:
* `email` - (String) Email of the User.
  * Constraints: The maximum length is `256` characters. The minimum length is `5` characters. The value must match regular expression `/^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$/`.
* `name` - (String) Name of the User.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[\\x20-\\x7E]*$/`.
* `user_id` - (String) ID of user who created/caused the event.
  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
