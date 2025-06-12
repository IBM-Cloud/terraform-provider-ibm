---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_destination_cos'
description: |-
  Manages Event Notification IBM Cloud Object Storage destinations.
---

# ibm_en_destination_cos

Create, update, or delete a IBM Cloud Object Storage destination by using IBM Cloud™ Event Notifications.

## Example usage

```terraform
resource "ibm_en_destination_cos" "cos_en_destination" {
  instance_guid         = ibm_resource_instance.en_terraform_test_resource.guid
  name                  = "COS Test Destination"
  type                  = "ibmcos"
  collect_failed_events = true
  description           = "IBM Cloud Object Storage Destination for event notification"
  config {
    params {
      bucket_name     = "cos-test-bucket"
      instance_id     = "1f7avhy78-3ehu-4d02-b123-8297333e0748399"
      endpoint        = "https://s3.us-east.cloud-object-storage.appdomain.cloud"
    }
  }
}
```

## Argument reference

Review the argument reference that you can specify for your resource.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `name` - (Required, String) The Destintion name.

- `description` - (Optional, String) The Destination description.

- `type` - (Required, String) ibmcos.

- `collect_failed_events` - (boolean) Toggle switch to enable collect failed event in Cloud Object Storage bucket.

- `config` - (Optional, List) Payload describing a destination configuration.

  Nested scheme for **config**:

  - `params` - (Required, List)

  Nested scheme for **params**:

  - `bucket_name` - (Required, string) The bucket name in IBM cloud object storage instance.
  - `instance_id` - (Required, string) The instance id for IBM Cloud object storage instance.
  - `endpoint`   - (Required, string) The endpoint for bucket region.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the `cos_en_destination`.
- `destination_id` - (String) The unique identifier of the created destination.
- `subscription_count` - (Integer) Number of subscriptions.
  - Constraints: The minimum value is `0`.
- `subscription_names` - (List) List of subscriptions.
- `updated_at` - (String) Last updated time.

## Import

You can import the `ibm_en_destination_cos` resource by using `id`.

The `id` property can be formed from `instance_guid`, and `destination_id` in the following format:

```
<instance_guid>/<destination_id>
```

- `instance_guid`: A string. Unique identifier for IBM Cloud Event Notifications instance.

- `destination_id`: A string. Unique identifier for Destination.

**Example**

```
$ terraform import ibm_en_destination_cos.cos_en_destination <instance_guid>/<destination_id>
```
