---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_page_rules"
description: |-
  Get information on an IBM Cloud Internet Services page rules.
---

# ibm_cis_page_rules
Retrieve an information of an IBM Cloud Internet Services page rules resource. For more information, about IBM Cloud Internet Services page rules, see [using page rules](https://cloud.ibm.com/docs/cis?topic=cis-use-page-rules).

## Example usage

```terraform
data "ibm_cis_page_rules" "rules" {
  cis_id    = ibm_cis.instance.id
  domain_id = ibm_cis_domain.example.id
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance .
- `domain_id` - (Required, String) The ID of the domain.


## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `cis_page_rules` - (String) The page rules detail.

  Nested scheme for `cis_page_rules`:
	- `priority` - (String) The priority of the page rule.
	- `status` - (String)   The status of the page rule. Default value is `active`.
	- `rule_id` - (String) The page rule ID.
	- `targets` - (String)   The targets, of the added rule.

	  Nested scheme for `targets`:
		- `constraint` - (String) The constraint of the page rule.

		  Nested scheme for `constraint`:
			- `operator` - (String) The operation on page rule. Valid value is `matches`.
			- `value` - (String) The URL value on the applied page rule.
		- `target` - (String) The target type. Valid value is `url`.
	- `actions` - (String)   The actions to be performed on the URL.

	  Nested scheme for `actions`:
		- `id` - (String) The action ID. Valid values are `page rule action field map from console` to `API CF-UI map API`.
		- `value` - (String) The values for corresponding actions.
		Please refer table in `ibm_cis_page_rule` resource document for corresponding valid values of `id` and `value`.
	- `status_code` - (String) The status code to check for URL forwarding. The required attribute for `forwarding_url` action. Valid values are `301` and `302`. It returns `0` for all other actions.
	- `url` - (String) The forward rule URL, a required attribute for `forwarding_url` action.
	- `css` - (String) The required attribute for `minify` action. CSS supported values are `on` and `off`.
    - `html` - (String) The required attribute for `minify` action. HTML supported values are `on` and `off`.
    - `js` - (String) The required attribute for `minify` action. JS supported values are `on` and `off`.