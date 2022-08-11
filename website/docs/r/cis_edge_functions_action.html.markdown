---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_edge_functions_action"
description: |-
  Provides a IBM CIS Edge Functions Action resource.
---

# ibm_cis_edge_functions_action
Create, update, or delete an edge functions action for a domain to include in your CIS edge functions action resource. For more information, about CIS edge functions action, see [working with Edge functions actions](https://cloud.ibm.com/docs/cis?topic=cis-edge-functions-actions).

## Example usage

```terraform
# Add a Edge Functions Action to the domain
resource "ibm_cis_edge_functions_action" "test_action" {
  cis_id      = data.ibm_cis.cis.id
  domain_id   = data.ibm_cis_domain.cis_domain.domain_id
  action_name = "sample-script"
  script      = file("./script.js")
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `action_name` - (Required, String) The action name of an edge functions action.
- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `domain_id` - (Required, String) The ID of the domain to add the edge functions action.
- `script` - (Required, String) The script of an edge functions action.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The action ID with a combination of `<action_name>`,`<domain_id>`,`<cis_id>` attributes concatenate with colon (`:`).

## Import
The `ibm_cis_edge_functions_action` resource can be imported by using the ID. The ID is composed from an edge functions action name or script name, the domain ID of the domain and the CRN (Cloud Resource Name) is concatenated with colon (`:`).

The domain ID and CRN are located on the **overview** page of the Internet Services instance in the domain heading of the console, or by using the IBM Cloud CIS command line commands.

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`.

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`.

- **Edge functions action name or script name** is a string: `sample_script`.


**Syntax**

```
$ terraform import ibm_cis_edge_functions_action.test_action <action_name>:<domain-id>:<crn>
```


**Example**

```
$ terraform import ibm_cis_edge_functions_action.test_action sample_script:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
