---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_alerts"
description: |-
  Get information on an IBM Cloud Internet Services alerts.
---

# ibm_cis_alerts

Retrieve information about an IBM Cloud Internet Services alerts data sources. For more information, see [IBM Cloud Internet Services](https://cloud.ibm.com/docs/cis?topic=cis-about-ibm-cloud-internet-services-cis).

## Example usage

```terraform
data "ibm_cis_alerts" "tests" {
	cis_id    = ibm_cis.instance.id
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `cis_id` - (Required, String) The ID of the CIS service instance.


## Attributes reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `id` - (String) The Webhook ID. It is a combination of <`policy_id`>,<`cis_id`> attributes concatenated with ":"
- `alert_policies` - (List)
   - `policy_id` - (String) The Alert Policy ID.
   - `name` - (String) The name of Alert policies.
   - `description` - (String) Description of the Alert Policies.
   - `enabled` - (Boolean) Whether this alert policies is active or not.
   - `alert_type` - (String) Type of the Alert Policy.
   - `mechanisms` - (List) 	Delivery mechanisms for the alert.
   - `filters` - (String) Optional filters depending for the alert type.
   - `conditions` - (String) Optional conditions depending for the alert type

