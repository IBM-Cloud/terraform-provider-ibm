---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_integrations'
description: |-
  List all the integrations
---

# ibm_en_destinations

Provides a read-only data source for integrations. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example usage

```terraform
data "ibm_en_integrations" "en_integrations" {
  instance_guid = ibm_resource_instance.en_terraform_test_resource.guid
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `search_key` - (Optional, String) Filter the integrations type.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the `en_integrations`.

- `integrations` - (List) List of destinations.

  - `id` - (String) Integration ID.

  - `type` - (String) Integration type kms/hs-crypto.

  - `updated_at` - (String) Last updated time.

- `total_count` - (Integer) Total number of destinations.
