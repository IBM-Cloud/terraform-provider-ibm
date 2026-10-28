---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_api_gateway_operations"
description: |-
  Provides an IBM Cloud Internet Services API Gateway Operations resource.
---

# ibm_cis_api_gateway_operations

Bulk-create and manage a set of API Gateway operations in IBM Cloud Internet Services (CIS) API Shield Endpoint Management. Each operation defines an HTTP method, host, and endpoint path. The resource creates all declared operations in a single API call and tracks their server-assigned UUIDs. For more information, see [IBM Cloud Internet Services](https://cloud.ibm.com/docs/cis).

## Example usage

```terraform
resource "ibm_cis_api_gateway_operations" "example" {
  cis_id    = data.ibm_cis.cis.id
  domain_id = data.ibm_cis_domain.cis_domain.domain_id

  operations {
    method   = "POST"
    host     = "api.example.com"
    endpoint = "/v1/chat/completions"
  }

  operations {
    method   = "GET"
    host     = "api.example.com"
    endpoint = "/v1/models"
  }
}
```

## Argument reference

Review the argument references that you can specify for your resource.

- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `domain_id` - (Required, String) The ID of the domain. The ID is a combination of the zone ID and CRN, separated by a `:`. Forces new resource.
- `operations` - (Required, List) List of API Gateway operations to register in Endpoint Management. At least one block is required.
  - `method` - (Required, String) The HTTP method. Valid values are `GET`, `POST`, `PUT`, `DELETE`, `PATCH`, `HEAD`, `OPTIONS`.
  - `host` - (Required, String) The RFC 3986-compliant API hostname, for example `api.example.com`.
  - `endpoint` - (Required, String) The endpoint path. Must start with `/`, for example `/v1/chat/completions`.

## Attribute reference

In addition to all argument reference list, you can access the following attribute references after your resource is created.

- `id` - (String) The record ID. It is a combination of `<domain_id>:<cis_id>` attributes concatenated with `:`.
- `operations` - (List) The resolved list of operations after creation, with server-assigned UUIDs populated.
  - `operation_id` - (String) The server-assigned UUID for the operation.
  - `method` - (String) The HTTP method.
  - `host` - (String) The API hostname.
  - `endpoint` - (String) The endpoint path.

## Import

The `ibm_cis_api_gateway_operations` resource can be imported using the `id`. The ID is formed from the domain ID and the CRN concatenated using a `:` character.

**Note:** On import, the resource state is populated with all operations currently registered in the zone, which may include operations created outside of Terraform. After importing, run `terraform plan` to review the difference between the imported state and your configuration, then `terraform apply` to reconcile.

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

**Syntax**

```
$ terraform import ibm_cis_api_gateway_operations.example <domain-id>:<crn>
```

**Example**

```
$ terraform import ibm_cis_api_gateway_operations.example 9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
