---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_destination_huawei'
description: |-
  Manages Event Notifications Huawei destination.
---

# ibm_en_destination_huawei

Create, update, or delete a  Huawei destination by using IBM Cloud™ Event Notifications.

## Example usage

```terraform
resource "ibm_en_destination_huawei" "huawei_en_destination" {
  instance_guid         = ibm_resource_instance.en_terraform_test_resource.guid
  name                  = "Huawei Destination"
  type                  = "push_huawei"
  collect_failed_events = false
  description           = "The Huawei Destination"
  config {
    params {
      client_id   = "5237288990"
      client_secret  = "36228ghutwervhudokmksiegfevssavdvywvwww" // pragma: allowlist secret
      pre_prod = false
    }
  }
}
```
  
## Argument reference

Review the argument reference that you can specify for your resource.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `name` - (Required, String) The Destintion name.

- `description` - (Optional, String) The Destination description.

- `type` - (Required, String) push_huawei.

- `collect_failed_events` - (boolean) Toggle switch to enable collect failed event in Cloud Object Storage bucket.

- `config` - (Optional, List) Payload describing a destination configuration.

  Nested scheme for **config**:

  - `params` - (Required, List)

  Nested scheme for **params**:

  - `client_id` - (String) Project Id value for Huawei project.
  - `client_secret` - (String) Private Key value for Huawei project
  - `pre_prod` - (Optional, bool) The flag to set your destination as pre prod destination or Prod Destination. The option is only available with Standard plan

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the `huawei_en_destination`.
- `destination_id` - (String) The unique identifier of the created destination.
- `subscription_count` - (Integer) Number of subscriptions.
  - Constraints: The minimum value is `0`.
- `subscription_names` - (List) List of subscriptions.
- `updated_at` - (String) Last updated time.

## Import

You can import the `ibm_en_destination_huawei` resource by using `id`.

The `id` property can be formed from `instance_guid`, and `huawei_id` in the following format:

```
<instance_guid>/<destination_id>
```

- `instance_guid`: A string. Unique identifier for IBM Cloud Event Notifications instance.

- `destination_id`: A string. Unique identifier for Destination.

**Example**

```
$ terraform import ibm_en_destination_huawei.huawei_en_destination <instance_guid>/<destination_id>
```
