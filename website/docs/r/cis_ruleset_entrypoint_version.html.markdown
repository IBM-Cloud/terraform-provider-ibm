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
        action_parameters  {
          id = "efb7b8c949ac4650a09736fc376e9aee"
        } 
      }
    }
  }

# Create/Update entrypoint ruleset and deploy multiple managed ruleset.

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
        action_parameters  {
          id = "efb7b8c949ac4650a09736fc376e9aee"
        } 
      }
      rules {
        action =  "execute"
        description = "Deploy CIS OWASP core ruleset"
        enabled = true
        expression = "true"
        action_parameters  {
          id = "4814384a9e5d4991b9815dcfc25d2f1f"
        } 
      }
      rules {
        action =  "execute"
        description = "Deploy CIS exposed credentials check ruleset"
        enabled = true
        expression = "true"
        action_parameters  {
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
        action_parameters  {
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
- `rulesets` - (Required, List) Values that will be created or updated.

  Nested scheme of `rulesets`
  - `description` (Optional, String) Description of the ruleset
  - `rules` (Optional, List) Rules that are required to be added/modified.
  Nested scheme of `rules`
    - `action` (String). If you are deploying a rule, then action is required. The `execute` action is used for deploying the ruleset. If you are updating the rule, the action is optional.
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
          - `score_threshold` (Optional, Int) Score threshold of the rule. Allowed values are 25, 40, 60 for high, medium and low sensitivity respectively. 
        - `categories` (Optional, List)

          Nested scheme of `categories`
          - `category` (Required, String) Category of the rule.
          - `enabled` (Optional, Boolean) Enables/Disables the rule.
          - `action` (Optional, String) Action of the rule.

## Attribute reference

There are no attribute references in addition to the argument reference list.

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
