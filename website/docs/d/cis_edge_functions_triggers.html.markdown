---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_edge_functions_triggers"
description: |-
  Get information on an IBM Cloud Internet Services Edge Function Triggers.
---

# ibm_cis_edge_functions_triggers
Retrieve information about an IBM Cloud Internet Services edge function triggers resource. For more information, about CIS edge functions trigger, see [working with triggers](https://cloud.ibm.com/docs/cis?topic=cis-edge-functions-actions#triggers).

## Example usage
The following example retrieves information about an IBM Cloud Internet Services edge function actions resource.

```terraform
data "ibm_cis_edge_functions_triggers" "test_triggers" {
    cis_id    = data.ibm_cis.cis.id
    domain_id = data.ibm_cis_domain.cis_domain.domain_id
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `cis_id` - (Required, String) The ID of the IBM CCIS instance.
- `domain_id` - (Required, String) The ID of the domain to add an edge functions triggers.


## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `action_name` - (String) An action script for execution.
- `pattern_url` - (String) The Route pattern. It is a domain name in which the action is performed.
- `request_limit_fail_open` - (String) An action request limit fail open.
- `trigger_id` - (String) The route ID of an action trigger.
