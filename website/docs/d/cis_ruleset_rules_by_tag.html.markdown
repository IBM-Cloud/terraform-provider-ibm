---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_ruleset_rules_by_tag"
description: |-
  Get information on an IBM Cloud Internet Services ruleset rules by tag.
---

# ibm_cis_ruleset_rules_by_tag

Retrieve information about IBM Cloud Internet Services rulesets rule by tag data sources. For more information, see [IBM Cloud Internet Services].

## Example usage

```terraform
data "ibm_cis_ruleset_rules_by_tag" "test"{
    cis_id    = ibm_cis.instance.id
    ruleset_id = "dcdec3fe0cbe41edac08619503da8de5"
    version = "2"
    rulesets_rule_tag = "wordpress"
}  

```

## Argument reference
Review the argument references that you can specify for your data source.

- `cis_id` - (Required, String) The ID of the CIS service instance.  
- `ruleset_id` - (Required, String) The ID of the ruleset. 
- `version` (Required, String) Version of the ruleset.
- `rulesets_rule_tag` (Required, String) The tag of the rule.

## Attributes reference 

In addition to the argument reference list, you can access the following attribute references after your data source is created.


- `result` - (Map)
    - `id` - (string) Ruleset ID.
    - `description` - (string) Description of the ruleset.
    - `kind` - (string) The kind of the ruleset.
    - `Phase` - (string) Phase of the ruleset.
    - `name` - (string) Name of the ruleset.
    - `last updated` - (string) Last update date of the ruleset.
    - `version` - (string) Version of the ruleset.



  - `rules` - (List) This list contains the information of rules associated with the `ruleset_id` with the given tag.
  
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
  

    
