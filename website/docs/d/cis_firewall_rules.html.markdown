---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_firewall_rules"
description: |-
  Get information on an IBM Cloud Internet Services firewall rules.
---

# ibm_cis_firewall_rules
Retrieve information about an existing IBM Cloud Internet Services instance. For more information, see [firewall rule actions](https://cloud.ibm.com/docs/cis?topic=cis-actions).

## Example usage

```terraform
data "ibm_cis_firewall_rules" "firewall_rules_instance" {
  cis_id    = ibm_cis.instance.id
  domain_id = ibm_cis_domain.example.id
}
```

## Argument reference
The following arguments are supported:

- `cis_id` - (Required, String) The ID of the CIS service instance.
- `domain_id` - (Required, String) The ID of the domain.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:

- `action` - (String) Create a firewall rules using these actions `log`, `allow`, `challenge`, `js_challenge`, `block` The firewall action to perform, "log" action is only available for enterprise plan instances.
- `description` - (String) The information about this firewall rules to help identify the purpose of it.
- `firewall_rules` - (List of Firewall Rules)
- `filter` - (Map) An existing filter which contains expression, paused and description.

  Nested scheme for `filter`:
    - `description` - (String) The information about this filter to help identify the purpose of it.
    - `expression` - (String) The expresson of the filter.
    - `paused` - (Boolean) Whether this filter is currently disabled.
    
- `id` - (String) The Firewall rules ID. It is a combination of <`firewall_rule_id`>,<`domain_id`>,<`cis_id`> attributes concatenated with ":"
- `paused` - (Boolean)  Whether this firewall rules is currently disabled.
  
   

