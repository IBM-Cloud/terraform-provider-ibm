---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_ruleset_entrypoint_versions"
description: |-
  Get information on an IBM Cloud Internet Services ruleset version.
---

# ibm_cis_ruleset_entrypoint_versions

Retrieve information about an IBM Cloud Internet Services Instance/Zone Entry Point ruleset's versions data sources. For more information, see [IBM Cloud Internet Services].

## Example usage

```terraform

data "ibm_cis_ruleset_entrypoint_versions" "test"{
    cis_id    = ibm_cis.instance.id
    domain_id= data.ibm_cis_domain.cis_domain.domain_id
    phase = "http_request_firewall_managed"
    version = "2"
    list_all = false
}  
```

## Argument reference
Review the argument references that you can specify for your data source.

- `cis_id` - (Required, String) The ID of the CIS service instance.
- `domain_id` - (Optional, String) The Domain/Zone ID of the CIS service instance. If domain_id is provided the request will be made at the zone/domain level otherwise the request will be made at the instance level.  
- `phase` - (Required, String) The phase of the ruleset.
- `list_all` - (Optional, boolean) If you provide `list_all` as true then you will get a list which wil contain the  information of all the ruleset's version. In this case you will not get the information of the rules associated with the rulesets. If you do not provide `list_all` argument or mark it as false then you will get the information of the latest version of the ruleset along with the information of associated rules. 
- `version` - (Optional, String) If `version` of the Entry Point ruleset is not provided then will get the information of the latest version of the ruleset along with the information of associated rules. If the `version` is provided then you will get the information of that particular version of the Entry Point ruleset along with the rules associated with it. If `list_all` is marked as true then you do not need to provide `version`. Even if you provide the value of `version` it won't make any effect on the request. 


## Attributes reference
In addition to the argument reference list, you can access the following attribute references after your data source is created.

Attribute references when `version` is not provided.

- `result` - (list)
    - `id` - (string) Ruleset ID.
    - `description` - (string) Description of the ruleset.
    - `kind` - (string) The kind of the ruleset.
    - `Phase` - (string) Phase of the ruleset.
    - `name` - (string) Name of the ruleset.
    - `last updated` - (string) Last update date of the ruleset.
    - `version` - (string) Version of the ruleset.

Extra attributes when `version` is provide.

- `rules` - (List) This list contains the information of rules associated with the Entry Point ruleset's version.
  
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
          - `score_threshold` (Int) Defines the score threshold of the rule.
        - `categories` (List)
          
          Nested scheme of `categories`
          - `category` (String) Category of the rule.
          - `enabled` (Boolean) Enables/Disables the rule.
          - `action` (String) Action of the rule.
      - `version` (String) Latest version.
      - `ruleset` (String) ID of the ruleset.
      - `phases` (List) Phases of the rule.
      - `products` (List) Products of the rule.
      - `rulesets` (List) IDs of the rulesets.
      - `response` (Map) Custom response from the API.
        - `content` (String) Content of the response.
        - `content_type` (string) Content type of the response.
        - `status_code` (Int) Status code returned by the API.
