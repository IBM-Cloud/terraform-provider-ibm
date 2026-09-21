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
- `rulesets` - (Optional, List) Input block containing the values to update. Contains only user-configurable fields; read-only API output fields such as `name`, `kind`, `phase`, `version`, and rule `version` are not accepted here.

  Nested scheme of `rulesets`:
  - `description` - (Optional, String) Description of the ruleset.
  - `rules` - (Optional, List) Rules to add or modify.

    Nested scheme of `rules`:
    - `id` - (Optional, String) ID of an existing rule. Required when updating a specific rule.
    - `action` - (Optional, String) Action of the rule.
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
        - `action` - (Optional, String) Action of the rule. Examples: `log`, `block`, `skip`.
        - `enabled` - (Optional, Boolean) Enables/Disables the rule.
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

The `rulesets` block contains only the user-supplied input and is not overwritten with API values on read. Read-only fields returned by the CIS API (such as ruleset `name`, `kind`, `phase`, `version`, and rule-level `version`, `ref`, and `last_updated_at`) are not exposed as computed attributes on this resource. To read the live state of a ruleset, use the [`ibm_cis_rulesets`](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/cis_rulesets) data source.

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
