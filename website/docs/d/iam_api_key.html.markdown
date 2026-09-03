---
subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : ibm_iam_api_key"
sidebar_current: "docs-ibm-datasource-iam-api-key"
description: |-
  Get information about an IAM API key.
---

# ibm_iam_api_key

Provides a read-only data source to retrieve information about an IAM API key. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax. For more information about IAM API keys, see [managing user API keys](https://cloud.ibm.com/docs/account?topic=account-userapikey&interface=ui).

## Example Usage

```hcl
data "ibm_iam_api_key" "iam_api_key" {
	apikey_id = "id"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `apikey_id` - (Required, Forces new resource, String) Unique ID of the API key.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `account_id` - (String) ID of the account that this API key authenticates for.
* `apikey_id` - (String) The unique identifier of the IBM-Cloud Api Key.
* `created_at` - (String) If set contains a date time string of the creation date in ISO format.
* `created_by` - (String) IAM ID of the user or service which created the API key.
* `crn` - (String) Cloud Resource Name of the item. Example Cloud Resource Name: `crn:v1:bluemix:public:iam-identity:us-south:a/myaccount::apikey:1234-9012-5678`.
* `description` - (String) The optional description of the API key. The 'description' property is only available if a description was provided during a create of an API key.
* `entity_tag` - (String) Version of the API Key details object. You need to specify this value when updating the API key to avoid stale updates.
* `expires_at` - (String) Date and time when the API key becomes invalid, ISO 8601 datetime in the format 'yyyy-MM-ddTHH:mm+0000'. **WARNING** An API key will be permanently and irrevocably deleted when both the expires_at and modified_at timestamps are more than ninety (90) days in the past, regardless of the key's locked status or any other state.
* `iam_id` - (String) The `iam_id` that this API key authenticates.
* `locked` - (Boolean) The API key cannot be changed if set to true.
* `modified_at` - (String) If set contains a date time string of the last modification date in ISO format.
* `name` - (String) Name of the API key. The name is not checked for uniqueness. Therefore multiple names with the same value can exist. Access is done via the UUID of the API key.
