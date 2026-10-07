---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_api_gateway_discovery_operations"
description: |-
  Provides an IBM Cloud Internet Services API Gateway Discovery Operations resource.
---

# ibm_cis_api_gateway_discovery_operations

Bulk-update the state (`saved`, `ignored`, `review`) of discovered API Gateway operations in IBM Cloud Internet Services (CIS). For more information, see [IBM Cloud Internet Services](https://cloud.ibm.com/docs/cis).

## Example usage

```terraform
resource "ibm_cis_api_gateway_discovery_operations" "example" {
  cis_id    = data.ibm_cis.cis.id
  domain_id = data.ibm_cis_domain.cis_domain.domain_id
  operation_states = {
    "f174e90a-fafe-4643-bbbc-4a0ed4fc8415" = "saved"
    "a1b2c3d4-0000-1111-2222-333344445555" = "ignored"
  }
}
```

## Argument reference

Review the argument references that you can specify for your resource.

- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `domain_id` - (Required, String) The ID of the domain. The ID is a combination of the zone ID and CRN, separated by a `:`.
- `operation_states` - (Required, Map of String) A map of discovered operation UUIDs to their desired state. Valid state values are `saved`, `ignored`, and `review`.

## Attribute reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The record ID. It is a combination of `<domain_id>:<cis_id>` attributes concatenated with `:`.
- `result` - (List) List of updated discovery operations returned by the API.
  - `id` - (String) UUID of the discovered operation.
  - `endpoint` - (String) Endpoint path.
  - `host` - (String) Host.
  - `method` - (String) HTTP method.
  - `state` - (String) Updated state (`saved`, `ignored`, or `review`).
  - `origin` - (List of String) Discovery engine origins.
  - `last_updated` - (String) Timestamp of last update in RFC3339 format.

## Import

The `ibm_cis_api_gateway_discovery_operations` resource can be imported using the `id`. The ID is formed from the domain ID and the CRN concatenated using a `:` character.

**Syntax**

```
$ terraform import ibm_cis_api_gateway_discovery_operations.example <domain-id>:<crn>
```

**Example**

```
$ terraform import ibm_cis_api_gateway_discovery_operations.example 9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
