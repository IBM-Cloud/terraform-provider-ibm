---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_destination_chrome'
description: |-
  Manages Event Notification Webhook destinations.
---

# ibm_en_destination_chrome

Create, update, or delete a Chrome destination by using IBM Cloud™ Event Notifications.

## Example usage

```terraform
resource "ibm_en_destination_chrome" "chrome_en_destination" {
  instance_guid = ibm_resource_instance.en_terraform_test_resource.guid
  name          = "My Chrome Destination"
  type          = "push_chrome"
  description   = "Destination Chrome for event notification"
  config {
    params {
      api_key     = "The FCM api key for project"
      website_url = "https://testevents.com"
    }
  }
}
```

## Argument reference

Review the argument reference that you can specify for your resource.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `name` - (Required, String) The Destintion name.

- `description` - (Optional, String) The Destination description.

- `type` - (Required, String) push_chrome.

- `config` - (Optional, List) Payload describing a destination configuration.

  Nested scheme for **config**:

  - `params` - (Required, List)

  Nested scheme for **params**:

  - `website_url` - (Required, string) URL of website.
  - `api_key` - (Required, string) The apikey for website project.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the `chrome_en_destination`.
- `destination_id` - (String) The unique identifier of the created destination.
- `subscription_count` - (Integer) Number of subscriptions.
  - Constraints: The minimum value is `0`.
- `subscription_names` - (List) List of subscriptions.
- `updated_at` - (String) Last updated time.

## Import

You can import the `ibm_en_destination_chrome` resource by using `id`.

The `id` property can be formed from `instance_guid`, and `destination_id` in the following format:

```
<instance_guid>/<destination_id>
```

- `instance_guid`: A string. Unique identifier for IBM Cloud Event Notifications instance.

- `destination_id`: A string. Unique identifier for Destination.

**Example**

```
$ terraform import ibm_en_destination_chrome.chrome_en_destination <instance_guid>/<destination_id>
```
