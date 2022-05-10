---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_destination_firefox'
description: |-
  Manages Event Notification firefox destinations.
---

# ibm_en_destination_webhook

Create, update, or delete a firefox destination by using IBM Cloud™ Event Notifications.

## Example usage

```terraform
resource "ibm_en_destination_firefox" "firefox_en_destination" {
  instance_guid = ibm_resource_instance.en_terraform_test_resource.guid
  name          = "My Firefox Destination"
  type          = "push_firefox"
  description   = "Destination firefox for event notification"
  config {
    params {
      website_url  = "https://testwebsite.com"
    }
  }
}
```

## Argument reference

Review the argument reference that you can specify for your resource.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `name` - (Required, String) The Destintion name.

- `description` - (Optional, String) The Destination description.

- `type` - (Required, String) push_firefox.

- `config` - (Optional, List) Payload describing a destination configuration.

  Nested scheme for **config**:

  - `params` - (Required, List)

  Nested scheme for **params**:

  - `website_url` - (Required, String) URL of the website project.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the `firefox_en_destination`.
- `destination_id` - (String) The unique identifier of the created destination.
- `subscription_count` - (Integer) Number of subscriptions.
  - Constraints: The minimum value is `0`.
- `subscription_names` - (List) List of subscriptions.
- `updated_at` - (String) Last updated time.

## Import

You can import the `ibm_en_destination_firefox` resource by using `id`.

The `id` property can be formed from `instance_guid`, and `destination_id` in the following format:

```
<instance_guid>/<destination_id>
```

- `instance_guid`: A string. Unique identifier for IBM Cloud Event Notifications instance.

- `destination_id`: A string. Unique identifier for Destination.

**Example**

```
$ terraform import ibm_en_destination_firefox.firefox_en_destination <instance_guid>/<destination_id>
```
