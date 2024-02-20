---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_integration_cos'
description: |-
   Get information about COS integration.
---

# ibm_en_integration_cos

Provides a read-only data source for COS Integration. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example usage

```terraform
resource "ibm_en_integration_cos" "cos_integration" {
  instance_guid = ibm_resource_instance.en_terraform_test_resource.guid
  integration_id = ibm_en_integration_cos.cos_integration.integration_id
}
```

## Argument reference

Review the argument reference that you can specify for your resource.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `integration_id` - (Required, String) Unique identifier for Integration created with .

- `type` - (Required, String) The integration type collect_failed_events.

- `metadata` - (Required, List)

  Nested scheme for **params**:

  - `endpoint` - (Required, String) endpoint url for COS bucket region.
  - `crn` - (Required, String) CRN of the COS instance.
  - `bucket_name` - (Required, String) cloud object storage bucket name.


## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - (String) The unique identifier of the `cos_integration`.
- `updated_at` - (String) Last updated time.