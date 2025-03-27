---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_ruleset_entrypoint_version"
description: |-
  Provides an IBM CIS ruleset entrypoint version resource.
---

# ibm_cis_ruleset_entrypoint_version
Provides an IBM Cloud Internet Services ruleset entrypoint version resource, to create and update the ruleset entrypoint of an instance or domain. For more information, about IBM Cloud Internet Services ruleset entrypoint version, see [ruleset entrypoint instance](https://cloud.ibm.com/docs/cis?topic=cis-managed-rules-overview).
**As there is no option to create a ruleset entry point resource, it is required to use import module to generate the respective resource configurations([Reference](https://test.cloud.ibm.com/docs/cis?topic=cis-terraform-generating-configuration)) and use the import command to populate the state file, as stated at the end of this page.**

## Example usage

```terraform
# create/update entrypoint ruleset of a domain or instance


resource "ibm_cis_ruleset_entrypoint_version" "config" {
    cis_id    = ibm_cis.instance.id
    domain_id = data.ibm_cis_domain.cis_domain.domain_id
    phase = "http_request_firewall_managed"
    rulesets {
      description = "Entry Point ruleset"
      rules {
        action =  "execute"
        action_parameters  {
          id = var.to_be_deployed_ruleset.id
          overrides  {
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
                action = "log"
            }
          }
        }
        description = var.rule.description
        enabled = true
        expression = "ip.src ne 1.1.1.1"
        ref = var.reference_rule.id
      }
    }
  }



```

## Argument reference
Review the argument references that you can specify for your resource. 

- `cis_id` - (Required, String) The ID of the CIS service instance.
- `domain_id` - (Optional, String) The Domain/Zone ID of the CIS service instance. If domain_id is provided the request will be made at the zone/domain level otherwise the request will be made at the instance level.
- `phase` - (Required, String) Phase of the ruleset. Currently, only `http_request_firewall_managed` phase is supported.
- `rulesets` - (Required, List) Values that will be created or updated.

  Nested scheme of `rulesets`
  - `description` (Optional, String) Description of the ruleset
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
        - `override_rules` (Optional, List) List of details of rules to be overridden. These rules are already present in the managed ruleset.

          Nested scheme of `override_rules`
          - `rule_id` (Required, String) Id of the rule.
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
The `ibm_cis_ruleset_entrypoint_version` resource is imported by using the ID. The ID is formed from the ruleset phase, the domain ID of the domain and the Cloud Resource Name (CRN) concatenated  using a `:` character.

The domain ID and CRN are located on the **Overview** page of the Internet Services instance of the domain heading of the console, or by using the `ibm cis` CLI commands.

- **Ruleset Phase** is a string of the form: `http_request_firewall_managed`.

- **Domain ID** is a 32-digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`.

- **CRN** is a 120-digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`.

**Syntax**

```
$ terraform import ibm_cis_ruleset_entrypoint_version.config <phase>:<domain-id>:<crn>
```

**Example**

```
$ terraform import ibm_cis_ruleset_entrypoint_version.config http_request_firewall_managed:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
