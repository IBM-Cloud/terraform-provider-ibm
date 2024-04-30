---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_rulesets_entrypoint_version"
description: |-
  Provides a IBM CIS rulesets entrypoint version resource.
---

# ibm_cis_rulesets_entrypoint_version
Provides an IBM Cloud Internet Services rulesets entrypoint version resource, to create and update ruleset entrypoint of an instance or domain. For more information, about IBM Cloud Internet Services rulesets entrypoint version, see [ruleset entrypoint instance]().

## Example usage

```terraform
# create/update entrypoint ruleset of a domain or instance

resource "ibm_cis_rulesets_entrypoint_version" "tests" {
    cis_id    = ibm_cis.instance.id
    domain_id = data.ibm_cis_domain.cis_domain.domain_id
    ruleset_phase = "http_request_firewall_managed"
    rulesets {
  
      description = "Entry Point ruleset"
      rules {
      {
        action =  "execute"
        action_parameters  {
          id : <id of ruleset to be deployed>
          overrides  {
            action = "log"
            enabled = true
            rules {
              {
                id = <id of rule to be overriden>
                enabled = true
                action = "log"
              }
            }
            categories {
              {
                category = "wordpress"
                enabled = true
                action = "log"
              }
            }
          }
        }
        description = "<description of rule>"
        enabled = true
        expression = "ip.src ne 1.1.1.1"
        ref = <reference to another rule>
        position  {
          index = 1
          after = <id of any existing rule>
          before = <id of any existing rule>
        }
      }
    }
  }
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `cis_id` - (Required, String) The ID of the CIS service instance.
- `domain_id` - (Optional, String) The Domain/Zone ID of the CIS service instance. If domain_id is provided request will be made at zone/domain level else request will be made at instance level.
- `ruleset_phase` - (Required, String) Phase of the ruleset. Currently, only `http_request_firewall_managed` phase is supported.
- `rulesets` - (Required, List) Values that will be created or updated.

  Nested scheme of `rulesets`
  - `description` (optional, string) Description of the ruleset
  - `rules` (optional, list) Rules which are required to be added/modified.

  Nested scheme of `rules`
    - `action` (String). If we are deploying a rule then action is required. `execute` action is used for deploying the ruleset. If we are updating the rule we then action is optional.
    - `description` (Optional, String) Description of the rule.
    - `enable` (Optional, Boolean) Enables/Disables the rule.
    - `expression` (Optional, String) Expression used by the rule to match the incoming request.
    - `ref` (Optional, String) Id of an existing rule. If not provided it is populated by the id if the new created rule.
    - `action_parmeters` (Optional, list) Parameters which are used to modify the rules.

      Nested scheme of `action parameters`
      - `id` (Required, String) Id of the managed ruleset to be deploted.
      - `overrides` (Optional, list) provides the parameter which are to be overridden.

        Nested scheme of `overrides`
        - `action` (Optional, String) Action of the rule. Examples: log, block, skip.
        - `enabled` (Optional, Boolean) Enables/Disables the rule
        - `rules` (Optional, list) list of details of rules to be overridden.

          Nested scheme of `rules`
          - `id` (Required, String) id of the rule.
          - `enabled` (Optional, Boolean) Enables/Disables the rule.
          - `action` (Optional, String) Action of the rule.
        - `categories` (Optional, list)

          Nested scheme of `categories`
          - `category` (Required, String) Category of the rule.
          - `enabled` (Optional, Boolean) Enables/Disables the rule.
          - `action` (Optional, String) Action of the rule.
    - `position` (Optional, list). Provides the postion when a new rule is added or updated the position of the current rule. If not provided new rule will be added at the last. You can only use one of the before, after, and index fields at a time.
      - `index` (Optional, String) Index of the rule to be added. 
      - `before` (Optional, String) Id of the rule before which new rule will be added. 
      - `after` (Optional, String) Id of the rule after which new rule will be added.
        
## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The entrypoint ruleset ID.


## Import
The `ibm_cis_rulesets_entrypoint_version` resource is imported by using the ID. The ID is formed from the ruleset phase, the domain ID of the domain and the CRN (Cloud Resource Name) concatenated  using a `:` character.

The domain ID and CRN is located on the **overview** page of the internet services instance of the domain heading of the console, or by using the `ibm cis` command line commands.

- **Ruleset Phase** is a string of the form: `http_request_firewall_managed`.

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`.

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`.

**Syntax**

```
$ terraform import ibm_cis_rulesets_entrypoint_version.config <ruleset_phase>:<domain-id>:<crn>
```

**Example**

```
$ terraform import ibm_cis_rulesets_entrypoint_version.config http_request_firewall_managed:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
