---

subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_service_id"
description: |-
  Manages IBM IAM service ID.
---

# ibm_iam_service_id

Create, update, or delete an IAM service ID by using resource group and resource type.  For more information, about IAM role action, see [managing service ID API keys](https://cloud.ibm.com/docs/account?topic=account-serviceidapikeys).

## Example usage

```terraform
resource "ibm_iam_service_id" "serviceID" {
  name        = "test"
  description = "New ServiceID"
}
```

## Argument reference

Review the argument references that you can specify for your resource.

- `name` - (Required, String) The name of the service ID.
- `description`  (Optional, String) The description of the service ID.
- `tags` (Optional, Array of Strings)  A list of tags that you want to add to the service ID. **Note** The tags are managed locally and not stored on the IBM Cloud Service Endpoint at this moment.

## Attribute reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `crn`  - (String) The CRN of the service ID.
- `iam_id`-  (String) The IAM ID of the service ID.
- `id` - (String) The unique identifier of the service ID.
- `locked`- (Bool) The Service Id lock status
- `version`  - (String) The version of the service ID.
