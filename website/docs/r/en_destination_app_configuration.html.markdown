---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_destination_app_configuration'
description: |-
  Manages Event Notification App Configuration destinations.
---

# ibm_en_destination_app_configuration

Create, update, or delete a App Configuration destination by using IBM Cloud™ Event Notifications.

## Example usage

```terraform
resource "ibm_en_destination_app_configuration" "ac_destination" {
  instance_guid         = ibm_resource_instance.en_terraform_test_resource.guid
  name                  = "App Configuration EN Destination"
  type                  = "app_configuration"
  collect_failed_events = false
  description           = "Destination App Configuration for event notification"
  config {
    params {
      type  = "features"
      crn = "crn:v1:bluemix:public:apprapp:us-south:a/4a74f2c31f554afc88156b73a1d577c6:dbxxxx93-0xxa-4xx5-axcf-c2faxxxd::"
      feature_id = "test"
      environment_id = "stage"
  }
}
}
```

## Argument reference

Review the argument reference that you can specify for your resource.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `name` - (Required, String) The Destintion name.

- `description` - (Optional, String) The Destination description.

- `type` - (Required, String) msteams.

- `collect_failed_events` - (boolean) Toggle switch to enable collect failed event in Cloud Object Storage bucket.

- `config` - (Optional, List) Payload describing a destination configuration.

  Nested scheme for **config**:

  - `params` - (Required, List)

  Nested scheme for **params**:

  - `type` - (Required, String) The App Configuration Destination type, the only supported type is **features** currently.
  - `crn` - (Required, String) CRN of the App Configuration instance.
  - `environment_id` - (Required, String) Environment ID of App Configuration.
  - `feature_id` - (Required, String) Feature ID of App Configuration.
## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the `ac_destination`.
- `destination_id` - (String) The unique identifier of the created destination.
- `subscription_count` - (Integer) Number of subscriptions.
  - Constraints: The minimum value is `0`.
- `subscription_names` - (List) List of subscriptions.
- `updated_at` - (String) Last updated time.

## Import

You can import the `ibm_en_destination_app_configuration` resource by using `id`.

The `id` property can be formed from `instance_guid`, and `destination_id` in the following format:

```
<instance_guid>/<destination_id>
```

- `instance_guid`: A string. Unique identifier for IBM Cloud Event Notifications instance.

- `destination_id`: A string. Unique identifier for Destination.

**Example**

```
$ terraform import ibm_en_destination_app_configuration.ac_destination <instance_guid>/<destination_id>
```
