---
layout: "ibm"
page_title: "IBM : iam_service_api_key"
sidebar_current: "docs-ibm-resource-iam-service-id_api_key"
description: |-
  Manages IBM IAM Service API Key.
---

# ibm\_iam_service_api_key

Provides a resource for IAM Service API Key. This allows Service API Key to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_iam_service_id" "serviceID" {
  name = "servicetest"
}

resource "ibm_iam_service_api_key" "testacc_apiKey" {
  name = "testapikey"
  iam_service_id = ibm_iam_service_id.serviceID.iam_id
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) Name of the Service API Key.
* `description` - (Optional, string) Description of the Service API Key.
* `iam_service_id` - (Required, string) IAM ID of the service.
* `apikey` - (Optional, string) The API key value.T his property only contains the API key value for the following cases: create an API key, update a Service API key that stores the API key value as retrievable, or get a Service API key that stores the API key value as retrievable. All other operations don't return the API key value, for example all user API key related operations, except for create, don't contain the API key value 
* `locked` - (Optional, bool) The API key cannot be changed if set to true.
* `store_value` - (Optional, bool) Boolean value deciding whether API key value is retrievable in the future.
* `file` - (Optional, string) The File name where api key is to be stored.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the Service API Key.
* `account_id` - The account ID of the API key.
* `entity_tag` - Version OR Entity tag of the Service API Key.
* `crn` - crn of the Service API Key.
* `created_by` - IAM ID of the service which created the API key
* `created_at` - The date and time Service API Key was created
* `modified_at` - The date and time Service API Key was modified.

## Import

ibm_iam_service_api_key can be imported using Service API Key, eg:

```
$ terraform import ibm_iam_service_api_key.testacc_apiKey ApiKey-9d1958af-5f42-41c2-a541-7b0be37c3da0
```