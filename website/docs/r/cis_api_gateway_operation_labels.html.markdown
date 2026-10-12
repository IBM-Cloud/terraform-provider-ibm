---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_api_gateway_operation_labels"
description: |-
  Provides an IBM Cloud Internet Services API Gateway Operation Labels resource.
---

# ibm_cis_api_gateway_operation_labels

Apply user-defined and managed labels to a set of API Gateway operations in an IBM Cloud Internet Services (CIS) instance. For more information, see [IBM Cloud Internet Services](https://cloud.ibm.com/docs/cis).

## Example usage

```terraform
resource "ibm_cis_api_gateway_operation_labels" "example" {
  cis_id       = data.ibm_cis.cis.id
  domain_id    = data.ibm_cis_domain.cis_domain.domain_id
  operation_ids = [
    ibm_cis_api_gateway_operation.op1.operation_id,
    ibm_cis_api_gateway_operation.op2.operation_id,
  ]
  managed_labels = ["cf-llm"]
  user_labels    = ["my-ai-api"]
}
```

## Argument reference

Review the argument references that you can specify for your resource.

- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `domain_id` - (Required, String) The ID of the domain. The ID is a combination of the zone ID and CRN, separated by a `:`.
- `operation_ids` - (Required, Set of String) The list of API Gateway operation UUIDs to apply labels to.
- `managed_labels` - (Optional, Set of String) Managed label strings to apply, for example `cf-llm`.
- `user_labels` - (Optional, Set of String) User-defined label strings to apply to the selected operations.

## Attribute reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The record ID. It is a combination of `<domain_id>:<cis_id>` attributes concatenated with `:`.
- `result` - (List) Per-operation label assignment results returned by the API.
  - `operation_id` - (String) The UUID of the operation.
  - `labels` - (List of String) Labels assigned to the operation.

## Import

The `ibm_cis_api_gateway_operation_labels` resource can be imported using the `id`. The ID is formed from the domain ID and the CRN concatenated using a `:` character.

**Syntax**

```
$ terraform import ibm_cis_api_gateway_operation_labels.example <domain-id>:<crn>
```

**Example**

```
$ terraform import ibm_cis_api_gateway_operation_labels.example 9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
