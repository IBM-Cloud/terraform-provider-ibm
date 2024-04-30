---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_ruleset"
description: |-
  Provides a IBM CIS ruleset resource.
---

# ibm_cis_rulesets
Provides an IBM Cloud Internet Services ruleset resource, to update and delete ruleset of an Instance or Domain. For more information, about IBM Cloud Internet Services ruleset, see [ruleset instance]().

## Example usage

```terraform
# update ruleset of a domain or instance

resource "ibm_cis_ruleset" "tests" {
    cis_id    = ibm_cis.instance.id
    domain_id = data.ibm_cis_domain.cis_domain.domain_id
    ruleset_id = "dcdec3fe0cbe41edac08619503da8de5"
    rulesets {
      description = "Entry Point Ruleset"
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
- `ruleset_id` - (Required, String) ID of the Ruleset.
- `rulesets` - (Required, List) Values that will be updated.

  Nested scheme of `rulesets`
  - `description` (optional, string) Description of the ruleset
  - `rules` (optional, list) Rules which are required to be added/modified.

  Nested scheme of `rules`
    - `action` (String). If we are deploying a rule then action is required. `execute` action is used for deploying the ruleset. If we are updating the rule then action is optional. For more understanding - [Deploy ruleset]()
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
There are no extra attribute refereneces in addition to the argument reference list.

