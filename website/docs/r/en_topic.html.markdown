---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_topic'
description: |-
  Manages Event Notifications topics.
---

# ibm_en_topic

Create, update, or delete a topic by using IBM Cloud™ Event Notifications.

## Example usage

Subscription example for Topic for other than scheduler source

```terraform
resource "ibm_en_topic" "en_topic" {
  instance_guid = ibm_resource_instance.en_terraform_test_resource.guid
  name          = "e2e topic"
  description   = "Topic for EN events routing"
  sources {
		id = ibm_resource_instance.cloud_logs_instance.crn
		
	rules {
		enabled =  true
		event_type_filter = "$.*"
		notification_filter = "$.data.alert_definition.name == '[${var.environment}]'"
	}
	}
}
```

Subscription example for Topic periodic-timer source

```terraform
resource "ibm_en_topic" "en_topic" {
  instance_guid = ibm_resource_instance.en_terraform_test_resource.guid
  name          = "Scheduler Topic"
  description   = "Topic for periodic-timer as source"
  sources {
		id = [for s in toset(data.ibm_en_sources.en_sources.sources): s.id if s.type == "periodic-timer"].0
		
	rules {
		enabled =  true
		event_schedule_filter {
      starts_at = "2024-12-23T05:00:00.000Z"
      ends_at = "2024-12-23T22:50:00.000Z"
      expression = "* * * * *"
    }
	}
	}
}
```

## Argument reference

Review the argument reference that you can specify for your resource.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `name` - (Required, String) Name of the topic.

- `description` - (Optional, String) Description of the topic.

- `sources` - (Optional, List) List of sources.
  Nested scheme for **sources**:

  - `id` - (Required, String) ID of the source.

  - `rules` - (Required, List) List of rules.
    Nested scheme for **rules**:

  - `enabled` - (Boolean) Whether the rule is enabled or not. The default value is `true`.

  - `event_type_filter` - (Optional, String) Event type filter. The default value is `$.*`. The maximum length is `255`characters. The minimum length is`3`characters. The value must match regular expression`/[a-zA-Z 0-9-_$.=']_/`.

  - `notification_filter` - (Optional, String) Notification filter. The minimum length is`0`characters. The value must match regular 
  expression`/[a-zA-Z 0-9-_$.=']-/`.

  * `event_schedule_filter` - (Optional, List) Event schedule filter attributes.
		Nested schema for **event_schedule_filter**:
			* `ends_at` - (Optional, String) event schedule end time.
			* `expression` - (Optional, String) cron schedule expression.
			  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/.*/`.
			* `starts_at` - (Optional, String) event schedule start time.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the `en_topic`.
- `topic_id` - (String) The unique identifier of the created topic.
- `source_count` - (Required, Integer) Number of sources.
- `subscription_count` - (Required, Integer) Number of subscriptions.
- `subscriptions` - (Required, List) List of subscriptions.
  Nested scheme for **subscriptions**:

  - `id`- (Required, String) Subscription ID.

  - `name`- (Required, String) Subscription name.

  - `description`- (Required, String) Subscription description.

  - `destination_id` - (Required, String) The destination ID.

  - `destination_type`- (Required, String) The type of destination.

  - `topic_id` - (Required, String) Topic ID.

  - `updated_at` - (Required, String) Last updated time.

- `updated_at` - (Required, String) Last time the topic was updated.

## Import

You can import the `ibm_en_topic` resource by using `id`.
The `id` property can be formed from `instance_guid`, and `topic_id` in the following format:

```
<instance_guid>/<topic_id>
```

- `instance_guid`: A string. Unique identifier for IBM Cloud Event Notifications instance.
- `topic_id`: A string. Unique identifier for Topic.

**Example**

```
$ terraform import ibm_en_topic.en_topic <instance_guid>/<topic_id>

```
