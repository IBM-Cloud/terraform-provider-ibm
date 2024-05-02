---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_rulesets"
description: |-
  Get information on an IBM Cloud Internet Services ruleset.
---

# ibm_cis_rulesets

Retrieve information about IBM Cloud Internet Services Instance/Zone rulesets data sources. For more information, see [IBM Cloud Internet Services].

## Example usage

```terraform
data "ibm_cis_rulesets" "tests" {
    cis_id    = ibm_cis.instance.id
    domain_id = data.ibm_cis_domain.cis_domain.domain_id
    ruleset_id = data.ibm_cis_ruleset.cis_ruleset.ruleset_id
    }
```

## Argument reference
Review the argument references that you can specify for your data source.

- `cis_id` - (Required, String) The ID of the CIS service instance.
- `domain_id` - (Optional, String) The Domain/Zone ID of the CIS service instance. If domain_id is provided the request will be made at the zone/domain level, otherwise the request will be made at the instance level.  
- `ruleset_id` - (Optional, String) The ID of the ruleset. If ruleset_id is not provided then the request will be made to get the list of the rulesets. That list will not contain the information about the rules of the ruleset. If the ruleset_id is provided then you will get the information of the ruleset and the associated rules.

## Attributes reference 

In addition to the argument reference list, you can access the following attribute references after your data source is created.

Attribute references when `ruleset_id` is not provided.

- `result` - (List)
    - `id` - (string) Ruleset ID.
    - `description` - (string) Description of the ruleset.
    - `kind` - (string) The kind of the ruleset.
    - `Phase` - (string) Phase of the ruleset.
    - `name` - (string) Name of the ruleset.
    - `last updated` - (string) Last update date of the ruleset.
    - `version` - (string) Version of the ruleset.

Extra attribute references when `ruleset_id` is provided. 

- `rules` - (List) This list contains the information of rules associated with the `ruleset_id`.
  
  Nested scheme of `rules`
    - `id` (String). ID of the rule.
    - `version` (String). Version of the rule.
    - `action` (String). Action of the rule.
    - `description` (String) Description of the rule.
    - `enable` (Boolean) Enables/Disables the rule.
    - `expression` (String) Expression used by the rule to match the incoming request.
    - `ref` (String) ID of an referrenced rule.
    - `last_updated` (String) Date and time of the last update was made on the rule.
    - `categories` (List) List of categories.
    - `logging` (Map) 
      - `enabled` (Boolean) Logging is enabled or not.
    - `action_parameters` (List) Action Parameters of the rule.
    
      Nested scheme of `action_parameters`
      - `id` (String) ID of the managed ruleset to be deployed.
      - `overrides` (List) Provides the parameters which are overridden.

        Nested scheme of `overrides`
        - `action` (String) Action of the rule. Examples: log, block, skip.
        - `enabled` (Boolean) Enables/Disables the rule.
        - `sensitivity_level` (String) Defines the sensitivity level of the rule.
        - `rules` (Optional, List) List of details of the managed rules which are overridden.

          Nested scheme of `rules`
          - `id` (String) ID of the rule.
          - `enabled` (Boolean) Enables/Disables the rule.
          - `action` (String) Action of the rule.
          - `sensitivity_level` (String) Defines the sensitivity level of the rule.
        - `categories` (List)
          
          Nested scheme of `categories`
          - `category` (String) Category of the rule.
          - `enabled` (Boolean) Enables/Disables the rule.
          - `action` (String) Action of the rule.
      - `version` (String) Latest version.
      - `ruleset` (String) ID of the ruleset.
      - `rulesets` (List) IDs of the rulesets.
      - `response` (Map) Custom response from the API.
        - `content` (String) Content of the response.
        - `content_type` (string) Content type of the response.
        - `status_code` (Int) Status code returned by the API.
  

    
