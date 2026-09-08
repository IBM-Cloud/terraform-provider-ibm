---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_ruleset_entrypoint_version"
description: |-
  Provides an IBM CIS ruleset entrypoint version resource.
---

# ibm_cis_ruleset_entrypoint_version

Provides an IBM Cloud Internet Services ruleset entrypoint version resource to create and update the ruleset entrypoint of an instance or domain. This entrypoint version is also used to deploy the managed ruleset and to add custom rules. For more information, about the IBM Cloud Internet Services ruleset entrypoint version, see [ruleset entrypoint instance](https://cloud.ibm.com/docs/cis?topic=cis-managed-rules-overview). To manage rules individually, you can also use [ruleset rule](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/cis_ruleset_rule).

## Example usage

```terraform
# create entrypoint ruleset for a domain.

  resource "ibm_cis_ruleset_entrypoint_version" "test" {
    cis_id    = ibm_cis.instance.id
    domain_id = data.ibm_cis_domain.cis_domain.domain_id
    phase = "http_request_firewall_managed"
    rulesets {
      description = "Entrypoint ruleset for managed ruleset"
    }
  }

# Create/Update entrypoint ruleset and deploy managed ruleset.

  resource "ibm_cis_ruleset_entrypoint_version" "test" {
    cis_id    = ibm_cis.instance.id
    domain_id = data.ibm_cis_domain.cis_domain.domain_id
    phase = "http_request_firewall_managed"
    rulesets {
      description = "Entrypoint ruleset for managed ruleset"
      rules {
        action =  "execute"
        description = "Deploy CIS managed ruleset"
        enabled = true
        expression = "true"
        action_parameters {
          id = "efb7b8c949ac4650a09736fc376e9aee"
        }
      }
    }
  }

# Create/Update entrypoint ruleset and deploy multiple managed rulesets.

  resource "ibm_cis_ruleset_entrypoint_version" "test" {
    cis_id    = ibm_cis.instance.id
    domain_id = data.ibm_cis_domain.cis_domain.domain_id
    phase = "http_request_firewall_managed"
    rulesets {
      description = "Entrypoint ruleset for managed ruleset"
      rules {
        action =  "execute"
        description = "Deploy CIS managed ruleset"
        enabled = true
        expression = "true"
        action_parameters {
          id = "efb7b8c949ac4650a09736fc376e9aee"
        }
      }
      rules {
        action =  "execute"
        description = "Deploy CIS OWASP core ruleset"
        enabled = true
        expression = "true"
        action_parameters {
          id = "4814384a9e5d4991b9815dcfc25d2f1f"
        }
      }
      rules {
        action =  "execute"
        description = "Deploy CIS exposed credentials check ruleset"
        enabled = true
        expression = "true"
        action_parameters {
          id = "c2e184081120413c86c3ab7e14069605"
        }
      }
    }
  }

# Override rules and categories in a deployed managed ruleset

  resource "ibm_cis_ruleset_entrypoint_version" "test" {
    cis_id    = ibm_cis.instance.id
    domain_id = data.ibm_cis_domain.cis_domain.domain_id
    phase = "http_request_firewall_managed"
    rulesets {
      description = "Entrypoint ruleset for managed ruleset"
      rules {
        action =  "execute"
        description = "Deploy CIS managed ruleset"
        enabled = true
        expression = "true"
        action_parameters {
          id = "efb7b8c949ac4650a09736fc376e9aee"
          overrides {
            action = "block"
            enabled = true
            override_rules {
              rule_id = "var.overriden_rule.id"
              enabled = true
              action = "block"
            }
            categories {
              category = "wordpress"
              enabled = true
              action = "block"
            }
          }
        }
      }
    }
  }

#  Add custom rules. Rules can also be added using the ruleset rule resource.

  resource "ibm_cis_ruleset_entrypoint_version" "config" {
    cis_id    = "crn:v1:bluemix:public:internet-svcs:global:a/bcf1865e99742d38d2d5fc3fb80a5496:d428087d-3f36-48f4-8626-99c37aee95bc::"
    domain_id = "de8e5d94f7033a29b026166e5f7c6f96"
    phase = "http_request_firewall_custom"
    rulesets {
      description = "var.description"
      rules {
        action = "var.action"
        expression = "var.expression"
        description = "var.rule.description"
        enabled = "true"
      }
      rules {
        action = "var.action"
        expression = "var.expression"
        description = "var.rule.description"
        enabled = "true"
      }
    }
  }

```

**Note**: If an update is required in a particular rule, you must still provide the data for other rules. Otherwise, the new update overrides the previous configuration. To add or update an individual rule, see the resource [ruleset rule](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/cis_ruleset_rule).

## Argument reference

Review the argument references that you can specify for your resource.

- `cis_id` - (Required, String) The ID of the CIS service instance.
- `domain_id` - (Optional, String) The Domain/Zone ID of the CIS service instance. If `domain_id` is provided, the request is made at the zone/domain level; otherwise, the request is made at the instance level.
- `phase` - (Required, String) Phase of the ruleset. Currently, only `http_request_firewall_managed` phase is supported.
- `rulesets` - (Optional, List, MaxItems: 1) Input block containing the ruleset description and rules to deploy.

  Nested scheme of `rulesets`:
  - `description` - (Optional, String) Description of the ruleset.
  - `rules` - (Optional, List) Rules to add or modify.

    Nested scheme of `rules`:
    - `id` - (Optional, String) ID of an existing rule. Required when updating a specific rule; omit when creating a new rule.
    - `action` - (Optional, String) Action of the rule. Use `execute` to deploy a managed ruleset.
    - `description` - (Optional, String) Description of the rule.
    - `enabled` - (Optional, Boolean) Enables/Disables the rule.
    - `expression` - (Optional, String) Expression used by the rule to match the incoming request.
    - `ref` - (Optional, String) Reference ID of an existing rule. If not provided, it is populated by the ID of the created rule.
    - `action_parameters` - (Optional, List, MaxItems: 1) Parameters used to configure the rule action.

      Nested scheme of `action_parameters`:
      - `id` - (Optional, String) ID of the managed ruleset to execute.
      - `ruleset` - (Optional, String) Skips the remaining rules in the current ruleset. Allowed value: `current`.
      - `phases` - (Optional, List) Skips the execution of one or more phases. Allowed values: `http_ratelimit`, `http_request_sbfm`, `http_request_firewall_managed`.
      - `products` - (Optional, List) Skips specific security products. Allowed values: `zoneLockdown`, `uaBlock`, `bic`, `hot`, `securityLevel`, `rateLimit`, `waf`.
      - `rulesets` - (Optional, List) List of ruleset IDs to apply the action to.
      - `response` - (Optional, Set) Custom response returned by the API.
        - `content` - (Optional, String) Response body content.
        - `content_type` - (Optional, String) Response content type.
        - `status_code` - (Optional, Integer) HTTP status code to return.
      - `overrides` - (Optional, Set) Override parameters for the managed ruleset.

        Nested scheme of `overrides`:
        - `action` - (Optional, String) Default action to apply to all rules in the managed ruleset. Examples: `log`, `block`, `skip`.
        - `enabled` - (Optional, Boolean) Enables/Disables all rules in the managed ruleset.
        - `override_rules` - (Optional, List) Per-rule overrides for specific rules already present in the managed ruleset.

          Nested scheme of `override_rules`:
          - `rule_id` - (Required, String) ID of the rule to override.
          - `enabled` - (Optional, Boolean) Enables/Disables the rule.
          - `action` - (Optional, String) Action to apply to the rule.
          - `score_threshold` - (Optional, Integer) Score threshold of the rule. Allowed values: `25` (high), `40` (medium), `60` (low sensitivity).
          - `sensitivity_level` - (Optional, String) Sensitivity level of the rule.
        - `categories` - (Optional, List) Category-level overrides.

          Nested scheme of `categories`:
          - `category` - (Required, String) Category name.
          - `enabled` - (Optional, Boolean) Enables/Disables rules in the category.
          - `action` - (Optional, String) Action to apply to rules in the category.
    - `position` - (Optional, Set) Position of the rule within the ruleset. Only one of `before`, `after`, or `index` may be set.
      - `before` - (Optional, String) Place this rule before the rule with this ID.
      - `after` - (Optional, String) Place this rule after the rule with this ID.
      - `index` - (Optional, Integer) Place this rule at this index position.

## Attribute reference

In addition to the argument references listed above, you can access the following computed attribute after the resource is applied.

- `rulesets_response` - (Computed, List) Full ruleset response as returned by the CIS API after apply. This block is read-only and reflects the live state of the entrypoint ruleset.

  Nested scheme of `rulesets_response`:
  - `ruleset_id` - (String) ID of the ruleset.
  - `description` - (String) Description of the ruleset.
  - `kind` - (String) Kind of the ruleset.
  - `name` - (String) Name of the ruleset.
  - `phase` - (String) Phase of the ruleset.
  - `version` - (String) Version of the ruleset.
  - `last_updated` - (String) Timestamp of the last update to the ruleset.
  - `rules` - (List) Rules associated with the ruleset.

    Nested scheme of `rules`:
    - `id` - (String) ID of the rule.
    - `version` - (String) Version of the rule.
    - `action` - (String) Action of the rule.
    - `description` - (String) Description of the rule.
    - `enabled` - (Boolean) Whether the rule is enabled.
    - `expression` - (String) Expression used by the rule to match the incoming request.
    - `ref` - (String) Reference ID of the rule.
    - `last_updated_at` - (String) Timestamp of the last update to the rule.
    - `categories` - (List) List of categories the rule belongs to.
    - `logging` - (Map) Logging configuration. Contains `enabled` (Boolean).
    - `action_parameters` - (Set) Action parameters of the rule.

      Nested scheme of `action_parameters`:
      - `id` - (String) ID of the managed ruleset deployed by this rule.
      - `version` - (String) Version of the managed ruleset.
      - `ruleset` - (String) Ruleset identifier.
      - `phases` - (List) Phases targeted by this rule.
      - `products` - (List) Security products targeted by this rule.
      - `rulesets` - (List) List of ruleset IDs.
      - `rules_to_skip` - (List) Ruleset-to-rule-IDs mappings for rules to skip.
        - `ruleset_id` - (String) Ruleset identifier.
        - `rule_ids` - (List of String) Rule IDs to skip within the ruleset.
      - `response` - (Set) Custom response configuration.
        - `content` - (String) Response body content.
        - `content_type` - (String) Response content type.
        - `status_code` - (Integer) HTTP status code.
      - `overrides` - (Set) Override configuration applied to the managed ruleset.
        - `action` - (String) Default action override.
        - `enabled` - (Boolean) Enable/disable all rules override.
        - `override_rules` - (List) Per-rule overrides.
          - `rule_id` - (String) ID of the overridden rule.
          - `enabled` - (Boolean) Enable/disable the rule.
          - `action` - (String) Action override.
          - `score_threshold` - (Integer) Score threshold override.
        - `categories` - (List) Category overrides.
          - `category` - (String) Category name.
          - `enabled` - (Boolean) Enable/disable the category.
          - `action` - (String) Action override for the category.

## Import

The `ibm_cis_ruleset_entrypoint_version` resource is imported by using the ID. The ID is formed from the ruleset phase, the domain ID of the domain, and the Cloud Resource Name (CRN) concatenated using a `:` character.

The domain ID and CRN are located on the **Overview** page of the Internet Services instance of the domain heading of the console, or by using the `ibm cis` CLI commands.

- **Ruleset Phase** is a string of the form: `http_request_firewall_managed`.

- **Domain ID** is a 32-digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`.

- **CRN** is a 120-digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`.

### Syntax

``` terraform
terraform import ibm_cis_ruleset_entrypoint_version.config <phase>:<domain-id>:<crn>
```

### Example

``` terraform
terraform import ibm_cis_ruleset_entrypoint_version.config http_request_firewall_managed:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
