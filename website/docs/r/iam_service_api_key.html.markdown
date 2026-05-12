---

subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_service_api_key"
description: |-
  Manages IBM IAM service API key.
---

# ibm_iam_service_api_key

Create, update, or delete an IAM service API key by using resource group and resource type.For more information, about IAM service API key, see [managing IAM acces, API keys](https://cloud.ibm.com/docs/cli?topic=cli-ibmcloud_commands_iam).

## Example usage

```terraform
resource "ibm_iam_service_id" "serviceID" {
  name = "servicetest"
}

resource "ibm_iam_service_api_key" "testacc_apiKey" {
  name = "testapikey"
  iam_service_id = ibm_iam_service_id.serviceID.iam_id
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `apikey`  (Optional, String) The API key value. This property only contains the API key value for the following cases: `create an API key`, `update a Service API key that stores the API key value as retrievable`, or `get a service API key that stores the API key value as retrievable`. All other operations do not return the API key value. For example, all user API key related operations, except for create, do not contain the API key value.
- `description`  (Optional, String) The description of the service API key.
- `file` - (Optional, String) The file name where API key is to be stored.
- `iam_service_id`  - (Required, String) The IAM ID of the service.
- `locked`- (Optional, Bool) The API key cannot be changed if set to **true**.
- `name` - (Required, String) The name of the service API key.
- `store_value`- (Optional, Bool) The boolean value whether API key value is retrievable in the future.
- `expires_at` - (Optional, String) Date and time when the API key becomes invalid, ISO 8601 datetime in the format 'yyyy-MM-ddTHH:mm+0000'. WARNING An API key will be permanently and irrevocably deleted when both the expires_at and modified_at timestamps are more than ninety (90) days in the past, regardless of the key's locked status or any other state.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `account_id`  - (String) The account Id of the API key.
- `entity_tag `-  (String) The version or entity tag of the service API key.
- `crn`  - (String) The `CRN` of the service API key.
- `created_at` - (Timestamp) The date and time service API key was created.
- `created_by` - (String) The IAM ID of the service that is created by the API key.
- `id` - (String) The unique identifier of the API key.
- `modified_at` - (String) The date and time service API key was modified.

## Import
The `ibm_iam_service_api_key` resource can be imported by using service API Key.

**Syntax**

```
$ terraform import ibm_iam_service_api_key.testacc_apiKey <service API key>
```

**Example**

```
$ terraform import ibm_iam_service_api_key.testacc_apiKey ApiKey-9d12342134f-41c2-a541-7b0be37c3da0
```
