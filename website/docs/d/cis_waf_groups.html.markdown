---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_waf_groups"
description: |-
  List the WAF Rule Groups in Cloud Internet Services.
---

# ibm_cis_waf_groups
Retrieve information of an existing IBM Cloud Internet Services WAF rule groups. For more information, about WAF refer to [Web Application Firewall concepts](https://cloud.ibm.com/docs/cis?topic=cis-waf-q-and-a).

## Example usage

```terraform
data "ibm_cis_waf_groups" "waf_groups" {
  cis_id     = data.ibm_cis.cis.id
  domain_id  = data.ibm_cis_domain.cis_domain.id
  package_id = "c504870194831cd12c3fc0284f294abb"
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `cis_id` - (Required, String) The resource CRN ID of the CIS on which zones were created.
- `domain_id` - (Required, String) The ID of the domain to retrieve the Load Balancers.
- `package_id` - (Required, String) The WAF Rule Package ID.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `description`  - (String) The WAF rule group description.
- `group_id` - (String) The WAF group ID.
- `modified_rules_count` - (Integer)  Number of rules modified in WAF Group.
- `mode`  - (String) The `on`, or `off` mode setting of the WAF rule group.
- `name` - (String) The name of the  WAF rule group.
- `rules_count` - (Integer)   Number of rules in WAF Group.
