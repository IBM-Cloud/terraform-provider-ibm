---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_integration_cos'
description: |-
  Manages Event Notifications COS Integration.
---

# ibm_en_integration

Manage COS integration using IBM Cloud™ Event Notifications.

## Example usage

```terraform
resource "ibm_en_integration_cos" "en_cos_integration" {
  instance_guid = ibm_resource_instance.en_terraform_test_resource.guid
  type = "collect_failed_events"
  metadata {
    endpoint = "https://s3.us-west.cloud-object-storage.test.appdomain.cloud",
		crn = "crn:v1:bluemix:public:cloud-object-storage:global:xxxx6db359a81a1dde8f44bxxxxxx:xxxx-1d48-xxxx-xxxx-xxxxxxxxxxxx::"
		bucket_name = "cloud-object-storage"
  }
}
```

## Argument reference

Review the argument reference that you can specify for your resource.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `type` - (Required, String) The integration type collect_failed_events.

- `metadata` - (Required, List)

  Nested scheme for **params**:

  - `endpoint` - (Required, String) endpoint url for COS bucket region.
  - `crn` - (Required, String) CRN of the COS instance.
  - `bucket_name` - (Required, String) cloud object storage bucket name.


## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the `en_cos_integration`.
- `updated_at` - (String) Last updated time.

## Import

You can import the `ibm_en_integration_cos` resource by using `id`.

The `id` property can be formed from `instance_guid`, and `integration_id` in the following format:

```
<instance_guid>/<integration_id>
```

- `instance_guid`: A string. Unique identifier for IBM Cloud Event Notifications instance.

- `integration_id`: A string. Unique identifier for Integration.

**Example**

```
$ terraform import ibm_en_integration_cos.en_cos_integration <instance_guid>/<integration_id>
```
