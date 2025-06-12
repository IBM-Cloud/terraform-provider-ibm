---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_firewall_rules"
description: |-
  Provides a IBM CIS Firewall Rules resource.
---

# ibm_cis_firewall_rules


Create, update, or delete a firewall rules for a domain that you included in your IBM Cloud Internet Services instance and a CIS domain resource. For more information, about CIS firewall rules resource, see [using fields, functions, and expressions](https://cloud.ibm.com/docs/cis?topic=cis-fields-and-expressions). Note - Deletion of Firewall Rules will result in deletion of the respective Filter too.

## Example usage

```terraform
# Add a firewall to the domain

resource "ibm_cis_filter" "test" {
  cis_id = ibm_cis.instance.id
  domain_id = ibm_cis_domain.example.id
  expression = "(ip.src eq 175.25.53.188 and http.request.uri.path eq \"^.*/wp-login[0-9].php$\")"
  paused =  true
  description = "Filter-creation"
}

resource "ibm_cis_firewall_rule" "firewall_rules_instance" {
  cis_id = ibm_cis.instance.id
  domain_id = ibm_cis_domain.example.id
  filter_id = ibm_cis_filter.test.filter_id
  action = "allow"
  priority = 5
  description = "Firewallrules-creation"

}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance where you want to create the firewall rules.
- `domain_id` - (Required, String) The ID of the domain where you want to apply the firewall rules.
- `action` - (Required, String) Create firewall rules by using these log, allow, challenge, js_challenge, block actions.
The firewall action to perform, log action is only available for the Enterprise plans instances.
- `description` - (Optional, String) The information about these firewall rules helps identify its purpose. 
- `filter_id` - (Required, String) The type of filter id from which you want to create firewall rules.
- `priority` - (Optional, Int) Create a firewall rules with priority.
- `paused` - (Optional, Bool) Whether this firewall rules is currently disabled.
 
  
## Attribute reference
In addition to all arguments above, the following attributes are exported:

- `filter_id` - (String) The filter ID.
- `id` - (String) The ID of the firewall rules. The ID is composed of `<filter_id>,<domain_ID>,<cis_crn>`. Attributes are concatenated with `:`.

## Import
The `ibm_cis_firewall_rules` resource is imported by using the `id`. The ID is formed from the `Filter ID`, the `Domain ID` of the domain and the `CRN` (Cloud Resource Name) concatentated usinga `:` character.

The Domain ID and CRN will be located on the **Overview** page of the Internet Services instance under the **Domain** heading of the UI, or via using the `ibmcloud cis` CLI commands.

- **Domain ID** is a 32 digit character string of the form: `029f08f2f2f0ab759fb28493b99f4df2`.

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/bcf1865e99742d38d2d5fc3fb80a5496:d428087d-3f36-48f4-8626-99c37aee95bc::`.

- **Firewall Rules ID** is a 32 digit character string of the form: `489d96f0da6ed76251b475971b097205c`.

**Syntax**

```
$ terraform import ibm_cis_firewall_rules.firewall_rules <firewall_rules_id>:<domain-id>:<crn>
```

**Example**

```
$ terraform import ibm_cis_firewall_rules.firewall_rules
d72c91492cc24d8286fb713d406abe91:0b30801280dc2dacac1c3960c33b9ccb:crn:v1:bluemix:public:internet-svcs-ci:global:a/01652b251c3ae2787110a995d8db0135:9054ad06-3485-421a-9300-fe3fb4b79e1d::
```

