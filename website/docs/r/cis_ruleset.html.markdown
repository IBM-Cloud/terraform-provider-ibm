---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_ruleset"
description: |-
  Provides a IBM CIS ruleset resource.
---

# ibm_cis_rulesets
Provides an IBM Cloud Internet Services ruleset resource, to update and delete the ruleset of an Instance or Domain. For more information about IBM Cloud Internet Services ruleset, see [ruleset instance]().
**As there is no option to create a ruleset resource, it is required to use import module to generate the respective resource configurations([Reference](https://test.cloud.ibm.com/docs/cis?topic=cis-terraform-generating-configuration)) and use the import command to populate the state file, as stated at the end of this page.**

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
- `domain_id` - (Optional, String) The Domain/Zone ID of the CIS service instance. If domain_id is provided the request will be made at the zone/domain level otherwise the request will be made at the instance level.
- `ruleset_id` - (Required, String) ID of the Ruleset.
- `rulesets` - (Required, List) Values that will be updated.

  Nested scheme of `rulesets`
  - `description` (optional, string) Description of the ruleset.
  - `rules` (optional, list) Rules which are required to be added/modified.

  Nested scheme of `rules`
    - `action` (String). If you are deploying a rule then action is required. `execute` action is used for deploying the ruleset. If you are updating the rule then action is optional. For more information, see - [Deploy ruleset]().
    - `description` (Optional, String) Description of the rule.
    - `enable` (Optional, Boolean) Enables/Disables the rule.
    - `expression` (Optional, String) Expression used by the rule to match the incoming request.
    - `ref` (Optional, String) ID of an existing rule. If not provided, it is populated by the ID of the created rule.
    - `action_parameters` (Optional, List) Parameters which are used to modify the rules.
    
      Nested scheme of `action parameters`
      - `id` (Required, String) ID of the managed ruleset to be deployed.
      - `overrides` (Optional, List) provides the parameters which are to be overridden.

        Nested scheme of `overrides`
        - `action` (Optional, String) Action of the rule. Examples: log, block, skip.
        - `enabled` (Optional, Boolean) Enables/Disables the rule
        - `rules` (Optional, List) List of details of rules to be overridden. These rules are already present in the managed ruleset.

          Nested scheme of `rules`
          - `id` (Required, String) ID of the rule.
          - `enabled` (Optional, Boolean) Enables/Disables the rule.
          - `action` (Optional, String) Action of the rule.
          - `categories` (Optional, List)
          
          Nested scheme of `categories`
          - `category` (Required, String) Category of the rule.
          - `enabled` (Optional, Boolean) Enables/Disables the rule.
          - `action` (Optional, String) Action of the rule.

    - `position` (Optional, List). Provides the position when a new rule is added, or updates the position of the current rule. If not provided the new rule will be added at the end. You can use only one of the before, after, and index fields at a time.
      - `index` (Optional, String) Index of the rule to be added. 
      - `before` (Optional, String) ID of the rule before which the new rule will be added. 
      - `after` (Optional, String) Id of the rule after which the new rule will be added.

        

## Attribute reference
There are no attribute references in addition to the argument reference list.


## Import
The `ibm_cis_ruleset` resource is imported by using the ID. The ID is formed from the ruleset ID, the domain ID of the domain and the CRN (Cloud Resource Name) concatenated  using a `:` character.

The domain ID and CRN are located on the **Overview** page of the Internet Services instance of the domain heading of the console, or by using the `ibm cis` CLI commands.

- **Ruleset ID** is a 32-digit character string of the form: `489d96f0da6ed76251b475971b097205c`.

- **Domain ID** is a 32-digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`.

- **CRN** is a 120-digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`.

**Syntax**

```
$ terraform import ibm_cis_ruleset.config <ruleset_id>:<domain-id>:<crn>
```

**Example**

```
$ terraform import ibm_cis_ruleset.config 48996f0da6ed76251b475971b097205c:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
