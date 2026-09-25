---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_api_gateway_operation"
description: |-
  Provides an IBM Cloud Internet Services API Gateway Operation resource.
---

# ibm_cis_api_gateway_operation

Create or delete an API Gateway operation (endpoint) in IBM Cloud Internet Services (CIS) API Shield Endpoint Management. For more information, see [IBM Cloud Internet Services](https://cloud.ibm.com/docs/cis).

## Example usage

```terraform
resource "ibm_cis_api_gateway_operation" "example" {
  cis_id    = data.ibm_cis.cis.id
  domain_id = data.ibm_cis_domain.cis_domain.domain_id
  endpoint  = "/api/v1/chat"
  host      = "api.example.com"
  method    = "POST"
}
```

## Argument reference

Review the argument references that you can specify for your resource.

- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `domain_id` - (Required, String) The ID of the domain. The ID is a combination of the zone ID and CRN, separated by a `:`.
- `endpoint` - (Required, String) The API endpoint path, for example `/api/v1/chat`. Forces new resource.
- `host` - (Required, String) The API hostname, for example `api.example.com`. Forces new resource.
- `method` - (Required, String) The HTTP method. Valid values are `GET`, `POST`, `PUT`, `DELETE`, `PATCH`, `HEAD`, `OPTIONS`. Forces new resource.

## Attribute reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The record ID. It is a combination of `<operation_id>:<domain_id>:<cis_id>` attributes concatenated with `:`.
- `operation_id` - (String) The UUID of the created API Gateway operation.

## Import

The `ibm_cis_api_gateway_operation` resource can be imported using the `id`. The ID is formed from the operation UUID, the domain ID, and the CRN concatenated using `:` characters.

- **Operation ID** is a UUID string of the form: `f174e90a-fafe-4643-bbbc-4a0ed4fc8415`

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

**Syntax**

```
$ terraform import ibm_cis_api_gateway_operation.example <operation-id>:<domain-id>:<crn>
```

**Example**

```
$ terraform import ibm_cis_api_gateway_operation.example f174e90a-fafe-4643-bbbc-4a0ed4fc8415:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
