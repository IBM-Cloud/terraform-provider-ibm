---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_ruleset_rule"
description: |-
  Provides an IBM CIS ruleset rule resource.
---

# ibm_cis_ruleset_rule
Provides an IBM Cloud Internet Services rulesets rule resource to create, update, and delete the ruleset rule of an instance or domain. For more information, about IBM Cloud Internet Services ruleset rule, see [ruleset instance]().

## Example usage

```terraform

resource "ibm_cis_ruleset_rule" "tests" {
    cis_id    = ibm_cis.instance.id
    domain_id = data.ibm_cis_domain.cis_domain.domain_id
    ruleset_id = "dcdec3fe0cbe41edac08619503da8de5"
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
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `cis_id` - (Required, String) The ID of the CIS service instance.
- `domain_id` - (Optional, String) The Domain/Zone ID of the CIS service instance. If domain_id is provided the request will be made at the zone/domain level otherwise the  request will be made at the instance level.
- `ruleset_id` - (Required, String) ID of the ruleset inside which rules will be created, updated, or deleted.
- `rules` (Optional, List) Rules which are required to be added/modified.
  
  Nested scheme of `rules`
    - `action` (String). If you are deploying a rule then action is required. The `execute` action is used for deploying the ruleset. If you are updating the rule we then action is optional.
    - `description` (Optional, String) Description of the rule.
    - `enable` (Optional, Boolean) Enables/Disables the rule.
    - `expression` (Optional, String) Expression used by the rule to match the incoming request.
    - `ref` (Optional, String) ID of an existing rule. If not provided it is populated by the ID of the created rule.
    - `action_parameters` (Optional, List) Parameters which are used to modify the rules.
    
      Nested scheme of `action parameters`
      - `id` (Required, String) ID of the managed ruleset to be deployed.
      - `overrides` (Optional, List) Provides the parameters which are to be overridden.

        Nested scheme of `overrides`
        - `action` (Optional, String) Action of the rule. Examples: log, block, skip.
        - `enabled` (Optional, Boolean) Enables/Disables the rule
        - `rules` (Optional, List) List of details of managed rules to be overridden. These rules are already present in the managed ruleset.

          Nested scheme of `rules`
          - `id` (Required, String) ID of the rule.
          - `enabled` (Optional, Boolean) Enables/Disables the rule.
          - `action` (Optional, String) Action of the rule.
        - `categories` (Optional, List)
          
          Nested scheme of `categories`
          - `category` (Required, String) Category of the rule.
          - `enabled` (Optional, Boolean) Enables/Disables the rule.
          - `action` (Optional, String) Action of the rule.
    - `position` (Optional, List). Provides the position when a new rule is added or updates the position of the current rule. If not provided the new rule will be added at the end. You can use only one of the before, after, and index fields at a time.
      - `index` (Optional, String) Index of the rule to be added. 
      - `before` (Optional, String) Id of the rule before which new rule will be added. 
      - `after` (Optional, String) Id of the rule after which new rule will be added.

        

## Attribute reference
In addition to the argument reference list, you can access the following attribute reference after your resource is created.

- `rule_id` - (String) ID of the rule. 


## Import
The `ibm_cis_ruleset_rule` resource is imported by using the ID. The ID is formed from the rule ID, the ruleset ID, the domain ID of the domain and the Cloud Resource Name (CRN)concatenated  using a `:` character.

The domain ID and CRN are located on the **Overview** page of the Internet Services instance of the domain heading of the console, or by using the `ibm cis` CLI commands.

- **Rule ID** is a 32-digit character string of the form: `dcdec3fe0cbe41edac08619503da8de5`.

- **Ruleset ID** is a 32-digit character string of the form: `489d96f0da6ed76251b475971b097205c`.

- **Domain ID** is a 32-digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`.

- **CRN** is a 120-digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`.

**Syntax**

```
$ terraform import ibm_cis_ruleset_rule.config <rule_id>:<ruleset_id>:<domain-id>:<crn>
```

**Example**

```
$ terraform import ibm_cis_ruleset_rule.config dcdec3fe0cbe41edac08619503da8de5:48996f0da6ed76251b475971b097205c:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
