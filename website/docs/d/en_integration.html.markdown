---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_integration'
description: |-
   Get information about kms integration.
---

# ibm_en_integration

Provides a read-only data source for kms/hs-crypto Integration. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example usage

```terraform
resource "ibm_en_integration" "en_kms_integration" {
  instance_guid = ibm_resource_instance.en_terraform_test_resource.guid
  integration_id = ibm_en_integration.kms_integration.integration_id
}
```

## Argument reference

Review the argument reference that you can specify for your resource.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `integration_id` - (Required, String) Unique identifier for Integration created with .

- `type` - (Required, String) The integration type kms/hs-crypto.

- `metadata` - (Required, List)

  Nested scheme for **params**:

  - `endpoint` - (Required, String) key protect/hyper protect service endpoint.
  - `crn` - (Required, String) crn of key protect/ hyper protect instance.
  - `root_key_id` - (Required, String) Root key id.


## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - (String) The unique identifier of the `en_kms_integration`.
- `updated_at` - (String) Last updated time.