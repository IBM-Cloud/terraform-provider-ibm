---
subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : ibm_iam_api_key"
sidebar_current: "docs-ibm-datasource-iam-api-key"
description: |-
  Get information about a IAM API key.
---

# ibm_iam_api_key

Retrieve information about an IAM API key data sources. For more information, about IAM API key, see [managing user API keys](https://cloud.ibm.com/docs/account?topic=account-userapikey&interface=ui).


## Example usage

```terraform
data "ibm_iam_api_key" "iam_api_key" {
	apikey_id = "id"
}
```

## Argument reference

Review the argument references that you can specify for your data source.

- `apikey_id` - (Required, String) Unique ID of the API key.

## Attribute reference

In addition to the argument reference list, you can access the following attribute references after your data source is created.

- `account_id` - (String) ID of the account that this API key authenticates for.
- `apikey_id` - (String) The unique identifier of the IBM-Cloud Api Key.
- `crn` - (String) Cloud Resource Name (CRN) of the item. For example Cloud Resource Name: `crn:v1:bluemix:public:iam-identity:us-south:a/myaccount::apikey:1234-9012-5678`.
- `created_at` - (Timestamp) If set contains a date time string of the creation date in ISO format.
- `created_by` - (String) IAM ID of the user or service which created the API key.
- `description` - (String) The optional description of the API key. The 'description' property is only available if a description is provided when you create an API key.
- `entity_tag` - (String) Version of the API Key details object. You need to specify this value when updating the API key to avoid stale updates.
- `iam_id` - (String) The `iam_id` that this API key authenticates.
- `locked` - (Bool) The API key cannot be changed if set to true.
- `expires_at` - (String) Date and time when the API key becomes invalid, ISO 8601 datetime in the format 'yyyy-MM-ddTHH:mm+0000'. WARNING An API key will be permanently and irrevocably deleted when both the expires_at and modified_at timestamps are more than ninety (90) days in the past, regardless of the key's locked status or any other state.
- `modified_at` - (Timestamp) If set contains a date time string of the last modification date in ISO format.
- `name` - (String) Name of the API key. The name is not checked for uniqueness. Therefore, multiple names with the same value can exist. Access is done by using the UUID of the API key.