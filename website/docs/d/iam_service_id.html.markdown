---
subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_service_id"
description: |-
  Manages IBM IAM Service ID.
---

# ibm_iam_service_id

Retrieve information about an IAM service ID. For more information, about IAM role action, see [managing service ID API keys](https://cloud.ibm.com/docs/account?topic=account-serviceidapikeys).

## Example usage

```terraform
data "ibm_iam_service_id" "ds_serviceID" {
  name = "sample"
}

```

## Argument reference

Review the argument references that you can specify for your data source.

- `name` - (Required, String) The name of the service ID.

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created.

- `service_ids` - (List of Objects)  A nested block list of IAM service IDs.
  - `bound_to`-  (String) The service the service ID is bound to.
  - `crn`-  (String) The CRN of the service ID.
  - `description`-  (String) A description of the service ID.
  - `iam_id`-  (String) The IAM ID of the service ID.
  - `id` - (String) The unique identifier of the service ID.
  - `locked`- (Bool) If set to **true**, the service ID is locked.
  - `version`-  (String) The version of the service ID.

  
