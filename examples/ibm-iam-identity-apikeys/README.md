# Example for IamIdentityV1

This example illustrates how to use the IamIdentityV1

These types of resources are supported:

* iam_api_key

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## IamIdentityV1 resources

iam_api_key resource:

```hcl
resource "iam_api_key" "iam_api_key_instance" {
  name = var.iam_api_key_name
  description = var.iam_api_key_description
  apikey = var.iam_api_key_apikey
  store_value = var.iam_api_key_store_value
  entity_lock = var.iam_api_key_entity_lock
}
```

## IamIdentityV1 Data sources

iam_api_key data source:

```hcl
data "iam_api_key" "iam_api_key_instance" {
  apikey_id = var.iam_api_key_id
}
```

## Assumptions

1. TODO

## Notes

1. TODO

## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | 1.13.1 |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| name | Name of the API key. The name is not checked for uniqueness. Therefore multiple names with the same value can exist. Access is done via the UUID of the API key. | `string` | true |
| description | The optional description of the API key. The 'description' property is only available if a description was provided during a create of an API key. | `string` | false |
| apikey | You can optionally passthrough the API key value for this API key. If passed, NO validation of that apiKey value is done, i.e. the value can be non-URL safe. If omitted, the API key management will create an URL safe opaque API key value. The value of the API key is checked for uniqueness. Please ensure enough variations when passing in this value. | `string` | false |
| store_value | Send true or false to set whether the API key value is retrievable in the future by using the Get details of an API key request. If you create an API key for a user, you must specify `false` or omit the value. We don't allow storing of API keys for users. | `bool` | false |
| entity_lock | Indicates if the API key is locked for further write operations. False by default. | `string` | false |
| apikey_id | Unique ID of the API key. | `string` | true |

## Outputs

| Name | Description |
|------|-------------|
| iam_api_key | iam_api_key object |
