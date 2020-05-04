# Example for API Gateway resources

This example illustrates how to use the API Gateway Endpoint and Subscription resources to create an endpoint for a given OpenAPI definition; and to create a subscription for this endpoint.It allows the user to input a single openAPI document or a directory of documents.

These types of resources are supported:

* [API Gateway Endpoint](https://cloud.ibm.com/docs/terraform?topic=terraform-api-gateway-resources#api-gw-endpoint)
* [API Gateway Endpoint Subscription](https://cloud.ibm.com/docs/terraform?topic=terraform-api-gateway-resources#api-gw-endpoint-subscript)

## Terraform versions

Terraform 0.12. Pin module version to `~> v1.4.0`. Branch - `master`.

Terraform 0.11. Pin module version to `~> v0.25.0`. Branch - `terraform_v0.11.x`.

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## API Gateway Resources

API Gateway Endpoint resource with single OpenAPI document:

```hcl
resource "ibm_api_gateway_endpoint" "endpoint" {
  service_instance_crn = ibm_resource_instance.apigateway.id
  name                 = "test-endpoint"
  managed              = false
  open_api_doc_name    = "/path"
  type                 = "share" //required only when updating action
}
```
API Gateway Endpoint resource with directory of OpenAPI documents:
```hcl
resource "ibm_api_gateway_endpoint" "endpoint" {
  for_each             = fileset(var.dir_path, "*.json")
  service_instance_crn = ibm_resource_instance.apigateway.id
  managed              = false
  name                 = replace("endpoint-${each.key}", ".json", "")
  open_api_doc_name    = format("%s%s", var.dir_path, each.key)
  type                 = "share" //required only when updating action
}
```
API Gateway Endpoint Subscription Resource:
```hcl
resource "ibm_api_gateway_endpoint_subscription" "subs" {
  artifact_id   = data.ibm_api_gateway.endpoint.endpoints[0].endpoint_id
  client_id     = "testapikey"
  name          = "testname"
  type          = "external"
  client_secret = "testsecret"
  //generate_secret=var.gen_secret //conflicts with client_secret
}
```
##  API Gateway Data Source
Lists all endpoints and its subscriptions of an API Gateway Instance.

```hcl
data "ibm_api_gateway" "endpoint"{
    service_instance_crn =ibm_resource_instance.apigateway.id
}
```

## Assumptions

1. It's recommended to use subscription resource by making the endpoint online. i.e manged attribute of endpoint resource should be true.
2. To view the subscriptions it is required to enable any of the two options of `Application authentication via API key` under `Define and Secure` page and save the endpoint in API Gateway service page.
3. The `client ID` of a particular subscription is available as an `API key` in the `Manage and Sharing` page of an endpoint of the API Gateway service.

## Notes

1. Terraform IBM provider v1.4.0 (via Terraform 0.12) supports "Autogeneration of Client ID i.e API key and Client Secret for the endpoint subscription".

## Examples

* [API Gateway Endpoint and Subscription resources](https://github.com/IBM-Cloud/terraform-provider-ibm/tree/master/examples/ibm-api-gateway)

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

Single OpenAPI document or directory of documents.

## Providers

| Name | Version |
|------|---------|
| ibm | n/a |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| region | THe region where the resource has to be provisioned. Default: `us-south`| `string` | yes |
| service\_name | The name of the API Gateway Service Instance. | `string` | yes |
| endpoint_name | The name of the API Gateway Endpoint resource. | `string` | yes |
| managed | Indicates whether endpoint is online or not. Default: false | `bool` | yes |
| routes | Invokable routes for an endpoint | `list` | no |
| file\_path | The API document name that represents the endpoint. It is required when a single endpoint is created| `string` | yes |
| dir_path | The directory name of API documents that represents multiple endpoint. It is required when a multipple endpoints are created| `string` | no |
| action\_type | The type of action that is performed on the API endpoint. Supported values are [`share`], [`unshare`], [`manage`], and [`unmanage`].To manage API to offline and online action\_type has to be set. The default value is [`unshare`]. Note that endpoint actions are performed by using the type parameter after the endpoint is created. As a consequence, endpoint actions are invoked during an endpoint update only. | `string` | required while managing actions. |
| subscription\_name | The name of the subscription resource indicates the name for an API key. | `string` | yes |
| client\_id | The API key to generate an API key for the subscription. The generated API key represents the ID of a subscription. If not provided it is auto generated. | `string` | yes |
| subscription\_type | The type of the subscription resource indicates the type of API key sharing. Supported values are [`external`], and [`internal`]. | `string` | yes |
| secret | The secret of the API key. | `string` | yes |
| generate_secret | It conflicts with secret. If `generate_secret`- `true`, secret is auto generated. | `bool` | no |

## Outputs

| Name | Description |
|------|-------------|
| endpointID | Endpoint ID or Artifact ID|
| clientID | Client ID or subscription ID |

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
