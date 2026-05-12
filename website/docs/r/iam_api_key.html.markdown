---
subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : ibm_iam_api_key"
sidebar_current: "docs-ibm-resource-iam-api-key"
description: |-
  Manages IAM API Key.
---

# ibm_iam_api_key

Create, modify, or delete an IAM API key resources.  For more information, about IAM API Key, see [managing user API keys](https://cloud.ibm.com/docs/account?topic=account-userapikey).

## Example usage

```terraform
resource "ibm_iam_api_key" "iam_api_key" {
  name = "name"
}
```

## Argument reference

Review the argument references that you can specify for your resource.

- `apikey` - (Optional, String) You can passthrough an API key value for this API key. If passed, that API key value is not validated, means, the value can be non URL safe. If omitted, the API key management creates an URL safe opaque API key value. The value of the API key is checked for uniqueness. Please ensure enough variations when passing the value.
- `description` - (Optional, String) The description of the API key. The `description` property is only available if a description was provided during API key creation.
- `expires_at` - (Optional, String) Date and time when the API key becomes invalid, ISO 8601 datetime in the format 'yyyy-MM-ddTHH:mm+0000'. WARNING An API key will be permanently and irrevocably deleted when both the expires_at and modified_at timestamps are more than ninety (90) days in the past, regardless of the key's locked status or any other state.
- `entity_lock` - (Optional, Bool) Indicates the API key is locked for further write operations. Default value is `false`.
- `file` - (Optional, String) The file name where API key is to be stored.
- `name` - (Required, String) The name of the API key. The name is not checked for uniqueness. Therefore, multiple names with the same value can exist. Access is done through the UUID of the API key.
- `store_value` - (Optional, Bool) Use `true` or `false` to set whether the API key value is retrievable in the future by using the `Get` details of an API key request. If you create an API key for a user, you must specify `false` or omit the value. Users cannot store the API key.


## Attribute reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `apikey_id` - (String) The unique identifier of the `ibm_iam_api_key`.
- `created_at` -  (Timestamp) If set contains the creation date time string in an ISO format.
- `created_by` - (String) The IAM ID of the user or service that creates the API key.
- `crn` - (String) The Cloud Resource Name (CRN) of an item. For example, CRN =  `crn:v1:bluemix:public:iam-identity:us-south:a/myaccount::apikey:1234-9012-1111`.
- `entity_tag` - (String) The version of the API Key details object. You need to specify this value when updating the API key to avoid stale updates.
- `locked` - (String) The API key cannot be changed if set to `true`.
- `modified_at` - (Timestamp) If set contains the last modification date in an ISO format.

## Import

The `ibm_iam_api_key` resource that can be imported by using user API Key.

**Syntax**

```
$ terraform import ibm_iam__api_key.iam_api_key <ApiKey-UniqueId>
```
