---
subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_api_key"
sidebar_current: "docs-ibm-resource-iam-api-key"
description: |-
  Manages iam_api_key.
---

# ibm\_iam_api_key

Provides a resource for iam_api_key. This allows iam_api_key to be created, updated and deleted.

## Example Usage

```hcl
resource "iam_api_key" "iam_api_key" {
  name = "name"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) Name of the API key. The name is not checked for uniqueness. Therefore multiple names with the same value can exist. Access is done via the UUID of the API key.
* `description` - (Optional, string) The optional description of the API key. The 'description' property is only available if a description was provided during a create of an API key.
* `apikey` - (Optional, string) You can optionally passthrough the API key value for this API key. If passed, NO validation of that apiKey value is done, i.e. the value can be non-URL safe. If omitted, the API key management will create an URL safe opaque API key value. The value of the API key is checked for uniqueness. Please ensure enough variations when passing in this value.
* `store_value` - (Optional, bool) Send true or false to set whether the API key value is retrievable in the future by using the Get details of an API key request. If you create an API key for a user, you must specify `false` or omit the value. We don't allow storing of API keys for users.
* `entity_lock` - (Optional, string) Indicates if the API key is locked for further write operations. False by default.

## Attribute Reference

The following attributes are exported:

* `apikey_id` - The unique identifier of the iam_api_key.
* `entity_tag` - Version of the API Key details object. You need to specify this value when updating the API key to avoid stale updates.
* `crn` - Cloud Resource Name of the item. Example Cloud Resource Name: 'crn:v1:bluemix:public:iam-identity:us-south:a/myaccount::apikey:1234-9012-5678'.
* `locked` - The API key cannot be changed if set to true.
* `created_at` - If set contains a date time string of the creation date in ISO format.
* `created_by` - IAM ID of the user or service which created the API key.
* `modified_at` - If set contains a date time string of the last modification date in ISO format.

## Import

ibm_iam_api_key can be imported using User API Key, eg:

```
$ terraform import ibm_iam__api_key.testacc_apiKey <ApiKey-UniqueId>
```
