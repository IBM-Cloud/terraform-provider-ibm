---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_edge_functions_actions"
description: |-
  Get information on an IBM Cloud Internet Services Edge Function Actions.
---

# ibm_cis_edge_functions_actions
Retrieve information about an IBM Cloud Internet Services edge function actions resource. For more information, about CIS edge functions action, see [working with Edge Functions actions](https://cloud.ibm.com/docs/cis?topic=cis-edge-functions-actions).

## Example usage

```terraform
data "ibm_cis_edge_functions_actions" "test_actions" {
    cis_id    = data.ibm_cis.cis.id
    domain_id = data.ibm_cis_domain.cis_domain.domain_id
}
```
## Argument reference
Review the argument references that you can specify for your data source. 

- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `domain_id` - (Required, String) The ID of the domain to add an edge functions action.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `created_on` - (String) An action created date.
- `etag` - (String) An action E-Tag.
- `handler` - (String) An action handler methods.
- `modified_on` - (String) An action modified date.
- `routes` - (String) An action route detail.

  Nested scheme for `routes`:
	- `action_name` - (String) An action route detail.
	- `pattern_url` - (String) The Route pattern. It is a domain name in which the action is performed.
	- `request_limit_fail_open` - (String) An action request limit fail open.
	- `trigger_id` - (String) The Trigger ID of an action.
