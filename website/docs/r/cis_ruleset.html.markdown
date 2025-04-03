---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_ruleset"
description: |-
  Provides an IBM CIS ruleset resource.
---

# ibm_cis_ruleset

Provides an IBM Cloud Internet Services ruleset resource to update and delete the ruleset of an instance or domain. To deploy the managed rulesets see [entrypoint ruleset](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/cis_ruleset_entrypoint_version). For more information about IBM Cloud Internet Services rulesets, see [ruleset instance](https://cloud.ibm.com/docs/cis?topic=cis-managed-rules-overview).
**As there is no option to create a ruleset resource, it is required to use import module to generate the respective resource configurations([Reference](https://test.cloud.ibm.com/docs/cis?topic=cis-terraform-generating-configuration)) and use the import command to populate the state file, as stated at the end of this page.**

## Example usage

```terraform
# update ruleset of a domain or instance

resource "ibm_cis_ruleset" "config" {
    cis_id    = ibm_cis.instance.id
    domain_id = data.ibm_cis_domain.cis_domain.domain_id
    ruleset_id = "943c5da120114ea5831dc1edf8b6f769"
    rulesets {
      description = "Entry point ruleset"
      rules {
        id = var.rule.id
        action =  "execute"
        action_parameters {
          id = var.to_be_deployed_ruleset.id
          overrides {
            action = "log"
            enabled = true
            override_rules {
                rule_id = var.overriden_rule.id
                enabled = true
                action = "block"
                score_threshold = 60
            }
            categories {
                category = "wordpress"
                enabled = true
                action = "block"
            }
          }
        }
        description = var.rule.description
        enabled = false
        expression = "true"
        ref = var.reference_rule.id
      }
    }
  }

```

## Argument reference

Review the argument references that you can specify for your resource.

- `cis_id` - (Required, String) The ID of the CIS service instance.
- `domain_id` - (Optional, String) The Domain/Zone ID of the CIS service instance. If `domain_id` is provided, the request is made at the zone/domain level; otherwise, the request is made at the instance level.
- `ruleset_id` - (Required, String) ID of the ruleset.
- `rulesets` - (Required, List) Values that will be updated.

  Nested scheme of `rulesets`
  - `description` (optional, string) Description of the ruleset.
  - `rules` (optional, list) Rules that are required to be added/modified.
  Nested scheme of `rules`
    - `id` (Required, String) ID of the rule.
    - `action` (Required, String). Action of the rule.
    - `description` (Optional, String) Description of the rule.
    - `enable` (Optional, Boolean) Enables/Disables the rule.
    - `expression` (Optional, String) Expression used by the rule to match the incoming request.
    - `ref` (Optional, String) ID of an existing rule. If not provided, it is populated by the ID of the created rule.
    - `action_parameters` (Optional, List) Parameters that are used to modify the rules.
    Nested scheme of `action parameters`
      - `id` (Required, String) ID of the managed ruleset to be deployed.
      - `overrides` (Optional, List) Provides the parameters that are to be overridden.

        Nested scheme of `overrides`
        - `action` (Optional, String) Action of the rule. Examples: log, block, skip.
        - `enabled` (Optional, Boolean) Enables/Disables the rule.
        - `override_rules` (Optional, List) List of details of rules to be overridden. These rules are already present in the managed ruleset.

          Nested scheme of `override_rules`
          - `rule_id` (Required, String) ID of the rule.
          - `enabled` (Optional, Boolean) Enables/Disables the rule.
          - `action` (Optional, String) Action of the rule.
          - `score_thrshold` (Optional, Int) Score theshold of the rule. Allowed values are 25, 40, 60 for high, medium and low sensitivity respectively. 
        
        - `categories` (Optional, List)

          Nested scheme of `categories`
          - `category` (Required, String) Category of the rule.
          - `enabled` (Optional, Boolean) Enables/Disables the rule.
          - `action` (Optional, String) Action of the rule.

## Attribute reference

There are no attribute references in addition to the argument reference list.

## Import

The `ibm_cis_ruleset` resource is imported by using the ID. The ID is formed from the ruleset ID, the domain ID of the domain, and the CRN (Cloud Resource Name) concatenated  using a `:` character.

The domain ID and CRN are located on the **Overview** page of the Internet Services instance of the domain heading of the console, or by using the `ibm cis` CLI commands.

- **Ruleset ID** is a 32-digit character string of the form: `489d96f0da6ed76251b475971b097205c`.

- **Domain ID** is a 32-digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`.

- **CRN** is a 120-digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`.

### Syntax

``` terraform
terraform import ibm_cis_ruleset.config <ruleset_id>:<domain-id>:<crn>
```

### Example

``` terraform
terraform import ibm_cis_ruleset.config 48996f0da6ed76251b475971b097205c:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
