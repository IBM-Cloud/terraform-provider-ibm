---
layout: "ibm"
page_title: "IBM: ibm_cis_firewall"
sidebar_current: "docs-ibm-datasource-cis-firewall"
description: |-
  Get information on an IBM Cloud Internet Services Firewall.
---

# ibm_cis_firewall

Imports a read only copy of an existing Internet Services Firewall resource.

## Example Usage

```hcl
data "ibm_cis_firewall" "lockdown" {
  cis_id    = ibm_cis.instance.id
  domain_id = ibm_cis_domain.example.id
  firewall_type = "lockdowns"
}
```

**NOTE:** IBM terraform provider supports only lockdowns rules

## Argument Reference

The following arguments are supported:

- `cis_id` - (Required,string) The ID of the CIS service instance
- `domain_id` - (Required,string) The ID of the domain to add the Lockdown.
- `firewall_type` - (Required,string) The type of firewall. Allowable values are [`lockdowns`],[`access_rules`],[`ua_rules`].

**NOTE:**

1. [`access_rules`]: Access Rules are a way to allow, challenge, or block requests to your website. You can apply access rules to one domain only or all domains in the same service instance.
2. [`ua_rules`]: Perform access control when matching the exact UserAgent reported by the client. The access control mechanisms can be defined within a rule to help manage traffic from particular clients. This will enable you to customize the access to your site.
3. [`lockdowns`]: Lock access to URLs in this domain to only permitted addresses or address ranges.

## Attributes Reference

The following attributes are exported:

- `id` - The record ID. It is a combination of <`firewall_type`>,<`lockdown_id`>,<`domain_id`>,<`cis_id`> attributes concatenated with ":".
- `lockdown` - List of lockdown to be created. It is the data describing a lockdowns rule.
  - `lockdown_id` - The lockdown ID.
  - `paused` - Whether this rule is currently disabled.
  - `description` - Some useful information about this rule to help identify the purpose of it.
  - `priority` - The priority of the record.
  - `urls` - URLs included in this rule definition. Wildcards are permitted. The URL pattern entered here is escaped before use. This limits the URL to just simple wildcard patterns.
  - `configurations` - List of IP addresses or CIDR ranges to use for this rule. This can include any number of [`ip`] or [`ip_range`] configurations that can access the provided URLs.
    - `target` - The request property to target. Valid values: [`ip`], [`ip_range`].
    - `value` - IP addresses or CIDR.
- `access_rule` - Access rule to be created. It is the data describing access rule.
  - `access_rule_id` - The access rule ID.
  - `notes` - Free text for notes.
  - `mode` - The mode of access rule. The valid modes are [`block`], [`challenge`], [`whitelist`], [`js_challenge`].
  - `configuration` - The Configuration of firewall.
    - `target` - The request property to target. Valid values: [`ip`], [`ip_range`], [`asn`], [`country`].
    - `value` - IP address or CIDR or Autonomous or Country code.
- `ua_rule` - User Agent rule to be created. It is the data describing user agent rule.
  - `ua_rule_id` - The User Agent rule ID.
  - `description` - Free text.
  - `mode` - The mode of access rule. The valid modes are [`block`], [`challenge`], [`js_challenge`].
  - `paused` - Whether this rule is currently disabled.
  - `configuration` - The Configuration of firewall.
    - `target` - The request property to target. Valid value: [`ua`].
    - `value` - The exact User Agent string to match with this rule.

**NOTE:**

- Exactly one of [`lockdown`], [`access_rule`] and [`ua_rule`] will be set for respective firewall types [`lockdowns`], [`access_rules`], [`ua_rules`].
