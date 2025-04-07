---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_ruleset_rule"
description: |-
  Provides an IBM CIS ruleset rule resource.
---

# ibm_cis_ruleset_rule

Provides an IBM Cloud Internet Services rulesets rule resource to create, update, and delete the ruleset rule of an instance or domain. For more information about the IBM Cloud Internet Services ruleset rule, see [ruleset instance](https://cloud.ibm.com/docs/cis?topic=cis-managed-rules-overview).

## Example usage

```terraform

# Get the ruleset it for the entrypoint ruleset in which the rule will be added
# use http_request_firewall_custom for custom entrypoint ruleset
# use http_request_firewall_managed for managed entrypoint ruleset

  data "ibm_cis_ruleset_entrypoint_versions" "test"{
    cis_id    = ibm_cis.instance.id
    domain_id = data.ibm_cis_domain.cis_domain.domain_id
    phase = "var.phase" 
  }   

# Add/Update rule for managed ruleset for deplying a ruleset.

  resource "ibm_cis_ruleset_rule" "config" {
    cis_id    = ibm_cis.instance.id
    domain_id = data.ibm_cis_domain.cis_domain.domain_id
    ruleset_id = "data.ibm_cis_ruleset_entrypoint_versions.ruleset_id"
      rule {
        action =  "execute"
        description = var.rule.description
        enabled = true
        expression = "true"
        ref = var.reference_rule.id
        action_parameters  {
          id = var.to_be_deployed_ruleset.id
          overrides {
            action =  "block"
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
        position {
          index = 1
          after = <id of any existing rule>
          before = <id of any existing rule>
        }
      }
}

# Add/Update a custom rule.

  resource ibm_cis_ruleset_rule "config" {
    cis_id    = ibm_cis.instance.id
    domain_id = data.ibm_cis_domain.cis_domain.domain_id
    ruleset_id = "data.ibm_cis_ruleset_entrypoint_versions.ruleset_id"
    rule {
      action =  "block"
      description = "var.description"
      expression = "true"
      enabled = "false"
      action_parameters {
        response {
          status_code = var.status_code
          content =  var.content
          content_type = "text/plain"
        }
      }
      position {
        index = var.index
        after = <id of any existing rule>
        before = <id of any existing rule>
      }
    }
  }

```

## Argument reference

Review the argument references that you can specify for your resource.

- `cis_id` - (Required, String) The ID of the CIS service instance.
- `domain_id` - (Optional, String) The Domain/Zone ID of the CIS service instance. If `domain_id` is provided, the request is made at the zone/domain level; otherwise, the request is made at the instance level.
- `ruleset_id` - (Required, String) ID of the ruleset inside which rules will be created, updated, or deleted.
- `rule` (Optional, List) Rule that is required to be added/modified.
  
  Nested scheme of `rule`
  - `action` (Required, String). If you are deploying a managed rule, then the `execute` action is used. If you are adding a custom rule, then any action can be used other then `execute`.
  - `description` (Optional, String) Description of the rule.
  - `enable` (Required, Boolean) Enables/Disables the rule.
  - `expression` (Required, String) Expression used by the rule to match the incoming request.
    - `ref` (Optional, String) ID of an existing rule. If not provided, it is populated by the ID of the created rule.
    - `action_parameters` (Optional, List) Parameters that are used to modify the rules.
    Nested scheme of `action parameters`
      - `id` (Optional, String) ID of the managed ruleset to be deployed. It is not required in custom rule.
      - `response` (Optional, Map). Custom response used for custom rules.

        Nested scheme of `response`

        - `status_code` (Optional, Integer) Status code of the response.
        - `content` (Optional, String) Content of the response.
        - `content_type` (Optional, String) Content type of the response.
      - `overrides` (Optional, List) Provides the parameters that are to be overridden.

        Nested scheme of `overrides`
        - `action` (Optional, String) Action of the rule. Examples: log, block, skip.
        - `enabled` (Optional, Boolean) Enables/Disables the rule.
        - `override_rules` (Optional, List) List of details of managed rules to be overridden. These rules are already present in the managed ruleset.

          Nested scheme of `override_rules`
          - `rule_id` (Required, String) ID of the rule.
          - `enabled` (Optional, Boolean) Enables/Disables the rule.
          - `action` (Optional, String) Action of the rule.
          - `score_threshold` (Optional, Int) Score threshold of the rule. Allowed values are 25, 40, 60 for high, medium and low sensitivity respectively. 
        - `categories` (Optional, List)

          Nested scheme of `categories`
          - `category` (Required, String) Category of the rule.
          - `enabled` (Optional, Boolean) Enables/Disables the rule.
          - `action` (Optional, String) Action of the rule.
    - `position` (Optional, List). You can use only one of the before, after, and index fields at a time. It is used to update the positing of the existing rule.
      - `index` (Optional, String) Index of the rule to be added.
      - `before` (Optional, String) ID of the rule before which the new rule will be added.
      - `after` (Optional, String) ID of the rule after which the new rule will be added.

## Attribute reference

In addition to the argument reference list, you can access the following attribute reference after your resource is created.

- `rule_id` - (String) ID of the rule.

## Import

Import is not possible, as there is no way to read the resource module.
