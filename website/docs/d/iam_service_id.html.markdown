---
subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_service_id"
description: |-
  Manages IBM IAM Service ID.
---

# ibm_iam_service_id

Retrieve information about an IAM service ID. For more information, about IAM role action, see [managing service ID API keys](https://cloud.ibm.com/docs/account?topic=account-serviceidapikeys).

## Example Usage

```hcl
data "ibm_iam_service_id" "iam_service_id" {
  name = "sample"
}
```

## Argument Reference

You can specify the following arguments for this data source.

- `name` - (Required, String) The name of the service ID.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

- `service_ids` - (List of Objects)  A nested block list of IAM service IDs.
  - `bound_to`-  (String) The service the service ID is bound to. This attribute is Deprecated.
  - `crn`-  (String) The CRN of the service ID.
  - `description`-  (String) A description of the service ID.
  - `iam_id`-  (String) The IAM ID of the service ID.
  - `id` - (String) The unique identifier of the service ID.
  - `locked`- (Bool) If set to **true**, the service ID is locked.
  - `version`-  (String) The version of the service ID.
