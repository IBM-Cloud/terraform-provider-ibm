---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_edge_functions_triggers"
description: |-
  Provides a IBM CIS Edge Functions trigger resource.
---

# ibm_cis_edge_functions_trigger

Create, update, or delete an edge functions trigger for a domain to include in your CIS edge functions trigger resource. For more information, about CIS edge functions trigger, see [working with triggers](https://cloud.ibm.com/docs/cis?topic=cis-edge-functions-actions#triggers).

## Example usage
The example to add an edge functions trigger to the domain.

```terraform
# Add a Edge Functions Trigger to the domain
resource "ibm_cis_edge_functions_trigger" "test_trigger" {
  cis_id      = ibm_cis_edge_functions_action.test_action.cis_id
  domain_id   = ibm_cis_edge_functions_action.test_action.domain_id
  action_name = ibm_cis_edge_functions_action.test_action.action_name
  pattern_url = "example.com/*"
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `action_name` - (Optional, String) An action name of the edge functions action on which the trigger associates. If it is not specified, then the trigger will be disabled.
- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `domain_id` - (Required, String) The ID of the domain to add the edge functions trigger.
- `pattern_url` - (Required, String) The domain name pattern on which the edge function action trigger should be executed.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `action_name` - (String) An edge functions action Script name.
- `id` - (String) The action ID with a combination of `<trigger_id>`,`<domain_id>`,`<cis_id>` attributes concatenate with colon (`:`).
- `pattern_url` - (String) The Route pattern. It is a domain name on which the action is performed.
- `request_limit_fail_open` - (String) An action request limit fail open.
- `trigger_id` - (String) The route ID of an action trigger.

## Import
The `ibm_cis_edge_functions_trigger` resource can be imported by using the ID. The ID is composed from an edge functions trigger route ID, the domain ID of the domain and the CRN (Cloud Resource Name) is concatenated with colon (`:`).


The domain ID and CRN are located on the overview page of the Internet Services instance in the domain heading of the console, or by using the IBM Cloud CIS command line commands.

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`.

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`.
Edge functions trigger route ID is a 32 digit character string of the form: `48996f0da6ed76251b475971b097205c`.


**Syntax**

```
$ terraform import ibm_cis_edge_functions_trigger.test_trigger <trigger_id>:<domain-id>:<crn>
```

**Example**

```
$ terraform import ibm_cis_edge_functions_trigger.test_trigger 48996f0da6ed76251b475971b097205c:9caf68812ae9b3f0377fdf986751a78f:crn:v1:ibmcloud:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
