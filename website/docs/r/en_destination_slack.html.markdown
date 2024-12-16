---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_destination_slack'
description: |-
  Manages Event Notification Slack destinations.
---

# ibm_en_destination_slack

Create, update, or delete a Slack destination by using IBM Cloud™ Event Notifications.

## Example usage

Destination example for Slack destination type `incoming_webhook`

```terraform
resource "ibm_en_destination_slack" "slack_en_destination" {
  instance_guid         = ibm_resource_instance.en_terraform_test_resource.guid
  name                  = "My Slack Destination"
  type                  = "slack"
  collect_failed_events = false
  description           = "Destination slack for event notification"
  config {
    params {
      type = "incoming_webhook"
      url  = "https://hooks.slack.com/services/G0gyhsush/TYodsjhs/GHTbfidsimkk"
  }
}
}
```
Destination example for Slack destination type `direct_message`

```terraform
resource "ibm_en_destination_slack" "slack_en_destination" {
  instance_guid         = ibm_resource_instance.en_terraform_test_resource.guid
  name                  = "My Slack Destination"
  type                  = "slack"
  collect_failed_events = false
  description           = "Destination slack for event notification"
  config {
    params {
      type = "direct_message"
      url  = "dveufewldfwefewlfdweckewweldewfdewd"
  }
}
}
```
## Argument reference

Review the argument reference that you can specify for your resource.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `name` - (Required, String) The Destintion name.

- `description` - (Optional, String) The Destination description.

- `type` - (Required, String) slack.

- `collect_failed_events` - (boolean) Toggle switch to enable collect failed event in Cloud Object Storage bucket.

- `config` - (Optional, List) Payload describing a destination configuration.

  Nested scheme for **config**:

  - `params` - (Required, List)

  Nested scheme for **params**:

  - `type` - (Required, String) The Slack destination type. The supported type are incoming_webhook and direct_message
  - `url` - (Optional, String) Slack Webhook url.
  - `token` - (Optional, String) Token of slack application.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the `webhook_en_destination`.
- `destination_id` - (String) The unique identifier of the created destination.
- `subscription_count` - (Integer) Number of subscriptions.
  - Constraints: The minimum value is `0`.
- `subscription_names` - (List) List of subscriptions.
- `updated_at` - (String) Last updated time.

## Import

You can import the `ibm_en_destination_slack` resource by using `id`.

The `id` property can be formed from `instance_guid`, and `destination_id` in the following format:

```
<instance_guid>/<destination_id>
```

- `instance_guid`: A string. Unique identifier for IBM Cloud Event Notifications instance.

- `destination_id`: A string. Unique identifier for Destination.

**Example**

```
$ terraform import ibm_en_destination_slack.slack_en_destination <instance_guid>/<destination_id>
```
