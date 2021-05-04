---
subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_api_key"
sidebar_current: "docs-ibm-datasource-iam-api-key"
description: |-
  Get information about iam_api_key
---

# ibm\_iam_api_key

Provides a read-only data source for iam_api_key. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "iam_api_key" "iam_api_key" {
	apikey_id = "id"
}
```

## Argument Reference

The following arguments are supported:

* `apikey_id` - (Required, string) Unique ID of the API key.

## Attribute Reference

The following attributes are exported:

* `apikey_id` - The unique identifier of the iam_api_key.

* `entity_tag` - Version of the API Key details object. You need to specify this value when updating the API key to avoid stale updates.

* `crn` - Cloud Resource Name of the item. Example Cloud Resource Name: 'crn:v1:bluemix:public:iam-identity:us-south:a/myaccount::apikey:1234-9012-5678'.

* `locked` - The API key cannot be changed if set to true.

* `created_at` - If set contains a date time string of the creation date in ISO format.

* `created_by` - IAM ID of the user or service which created the API key.

* `modified_at` - If set contains a date time string of the last modification date in ISO format.

* `name` - Name of the API key. The name is not checked for uniqueness. Therefore multiple names with the same value can exist. Access is done via the UUID of the API key.

* `description` - The optional description of the API key. The 'description' property is only available if a description was provided during a create of an API key.

* `iam_id` - The iam_id that this API key authenticates.

* `account_id` - ID of the account that this API key authenticates for.

* `apikey` - The API key value. This property only contains the API key value for the following cases: create an API key, update a service ID API key that stores the API key value as retrievable, or get a service ID API key that stores the API key value as retrievable. All other operations don't return the API key value, for example all user API key related operations, except for create, don't contain the API key value.