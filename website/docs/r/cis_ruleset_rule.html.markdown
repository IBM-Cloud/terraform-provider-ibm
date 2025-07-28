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
# use http_ratelimit for ratelimit entrypoint ruleset

  data "ibm_cis_ruleset_entrypoint_versions" "test"{
    cis_id    = ibm_cis.instance.id
    domain_id = data.ibm_cis_domain.cis_domain.domain_id
    phase = "var.phase" 
  }   

# Add/Update rule for managed ruleset for deplying a ruleset.

  resource "ibm_cis_ruleset_rule" "config" {
    cis_id    = ibm_cis.instance.id
    domain_id = data.ibm_cis_domain.cis_domain.domain_id
    ruleset_id = data.ibm_cis_ruleset_entrypoint_versions.config.rulesets[0].ruleset_id
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
    ruleset_id = data.ibm_cis_ruleset_entrypoint_versions.config.rulesets[0].ruleset_id
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

# Add/Update a ratelimit rule. make sure phase is http_ratelimit

  resource ibm_cis_ruleset_rule "config" {
    cis_id    = ibm_cis.instance.id
    domain_id = data.ibm_cis_domain.cis_domain.domain_id
    ruleset_id = data.ibm_cis_ruleset_entrypoint_versions.config.rulesets[0].ruleset_id
    rule {
      action =  "block"
      description = "var.description"
      expression = "true"
      enabled = "false"
      ratelimit {
        characteristics = ['cf.colo.id', ...var.ratelimit.characteristics]
        mitigation_timeout = var.ratelimit.mitigation_timeout
        period = var.ratelimit.period
        requests_per_period = var.ratelimit.requests_per_period
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
      - `ruleset` (Optional, String)  Skips the remaining rules in the current ruleset. Allowed value is `current`.
      - `phases` (Optional, List) Skips the execution of one or more phases. Allowed values for phases are `http_ratelimit`, `http_request_sbfm`, `http_request_firewall_managed`.
      - `products` (Optional, List) Skips specific security products. Allowed values for products are `zoneLockdown`, `uaBlock`, `bic`, `hot`, `securityLevel`, `rateLimit`, `waf`.
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
    - `ratelimit` (Optional, Map) Ratelimit of the rule to be added(custom ruleset). entry point ruleset should be `http_ratelimit` and Ruleset action should not be `execute`
      - `characteristics` (StringList) Set of parameters defining how tracks the request rate for the rule. `cf.colo.id` is mandatory to be passed, regardless of any additional strings in the list.
      - `counting_expression` (Optional, String) Defines the criteria used for determining the request rate. By default, the counting expression is the same as the rule matching expression (defined in If incoming requests match).
      - `mitigation_timeout` (Integer) Once the rate is reached, the rate limiting rule applies the rule action to further requests for the period of time defined in this field (in seconds).
      - `period` (Integer) The period of time to consider (in seconds) when evaluating the request rate.
      - `requests_per_period` (Integer) The number of requests over the period of time that will trigger the rule.

## Attribute reference

In addition to the argument reference list, you can access the following attribute reference after your resource is created.

- `rule_id` - (String) ID of the rule.

## Import

Import is not possible, as there is no way to read the resource module.
