---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_firewall"
description: |-
  Provides a IBM CIS Firewall resource.
---

# ibm_cis_firewall

Provides a IBM CIS Firewall resource. This resource is associated with an IBM Cloud Internet Services instance and a CIS Domain resource. It allows to create, update, delete firewall of a domain of a CIS instance

## Example Usage

```hcl
# Add a firewall to the domain

resource "ibm_cis_firewall" "lockdown" {
  cis_id    = ibm_cis.instance.id
  domain_id = ibm_cis_domain.example.id
  firewall_type = "lockdowns"
  lockdown {
    paused      = "false"
    description = "test"
    urls = ["www.cis-terraform.com"]
    configurations {
      target = "ip"
      value  = "127.0.0.2"
    }
    priority=1
  }
}

resource "ibm_cis_firewall" "access_rules" {
  cis_id    = ibm_cis.instance.id
  domain_id = ibm_cis_domain.example.id
  firewall_type = "access_rules"
  access_rule {
    mode  = "block"
    notes = "access rule notes"
    configuration {
      target = "asn"
      value  = "AS12346"
    }
  }
}

resource "ibm_cis_firewall" "ua_rules" {
  cis_id    = ibm_cis.instance.id
  domain_id = ibm_cis_domain.example.id
  firewall_type = "ua_rules"
  ua_rule {
    mode = "challenge"
    configuration {
      target = "ua"
      value  = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4"
    }
  }
}
```

## Argument Reference

The following arguments are supported:

- `cis_id` - (Required,string) The ID of the CIS service instance
- `domain_id` - (Required,string) The ID of the domain to add the Lockdown.
- `firewall_type` - (Required,string) The type of firewall. Allowable values are [`lockdowns`],[`access_rules`],[`ua_rules`].

**NOTE:**

1. [`access_rules`]: Access Rules are a way to allow, challenge, or block requests to your website. You can apply access rules to one domain only or all domains in the same service instance.
2. [`ua_rules`]: Perform access control when matching the exact UserAgent reported by the client. The access control mechanisms can be defined within a rule to help manage traffic from particular clients. This will enable you to customize the access to your site.
3. [`lockdowns`]: Lock access to URLs in this domain to only permitted addresses or address ranges.

- `lockdown` - (Optional,list) (MinItems: 1) List of lockdown to be created. It is the data describing a lockdowns rule.
  - `paused` - (Optional,boolean). Whether this rule is currently disabled.
  - `description` - (Optional,string). Some useful information about this rule to help identify the purpose of it.
  - `priority` - (Optional,int) The priority of the record.
  - `urls` - (Optional,list). URLs included in this rule definition. Wildcards are permitted. The URL pattern entered here is escaped before use. This limits the URL to just simple wildcard patterns.
  - `configurations` - (Optional,list). List of IP addresses or CIDR ranges to use for this rule. This can include any number of [`ip`] or [`ip_range`].configurations that can access the provided URLs. This can not be modified once it is created.
    - `target` - (Optional,string). The request property to target. Valid values: [`ip`], [`ip_range`].
    - `value` - (Optional,string). IP addresses or CIDR.
- `access_rule` - (Optional) (MaxItem: 1) Access rule to be created. It is the data describing access rule.
  - `notes` - (Optional, string) Free text for notes.
  - `mode` - (Required, string) The mode of access rule. The valid modes are [`block`], [`challenge`], [`whitelist`], [`js_challenge`].
  - `configuration` - (Required, List) (MaxItems: 1) The Configuration of firewall.
    - `target` - (Required, string) The request property to target. Valid values: [`ip`], [`ip_range`], [`asn`], [`country`].
    - `value` - (Required, string) IP address or CIDR or Autonomous or Country code.
- `ua_rule` - (Optional) (MaxItem: 1) User Agent rule to be created. It is the data describing user agent rule.
  - `description` - (Optional, string) Free text.
  - `mode` - (Required, string) The mode of access rule. The valid modes are [`block`], [`challenge`], [`js_challenge`].
  - `paused` - (Optional, boolean) Whether this rule is currently disabled.
  - `configuration` - (Required, List) (MaxItems: 1) The Configuration of firewall.
    - `target` - (Required, string) The request property to target. Valid values: [`ua`].
    - `value` - (Required, string) The exact User Agent string to match with this rule.

**NOTE:**

- Exactly one of [`lockdown`], [`access_rule`] and [`ua_rule`] is allowed to be given for input for respective firewall types [`lockdowns`], [`access_rules`], [`ua_rules`].

## Attributes Reference

The following attributes are exported:

- `id` - The firewall ID. It is a combination of <`firewall_type`>,<`lockdown_id/access_rul_id/ua_rule_id`>,<`domain_id`>,<`cis_id`> attributes concatenated with ":".
- `lockdown`
  - `lockdown_id` - The lockdown ID.
- `access_rule`
  - `access_rule_id` - The access rule ID.
- `ua_rule`
  - `ua_rule_id` - The User Agent rule ID.

## Import

The `ibm_cis_firewall` resource can be imported using the `id`. The ID is formed from the `Firewall Type`,the `Firewall ID`, the `Domain ID` of the domain and the `CRN` (Cloud Resource Name) concatentated using a `:` character.

The Domain ID and CRN will be located on the **Overview** page of the Internet Services instance under the **Domain** heading of the UI, or via using the `ibm cis` CLI commands.

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

- **Firewall ID** is a 32 digit character string of the form: `489d96f0da6ed76251b475971b097205c`.

- **Firewall Type** is a string. It can be either of [`lockdowns`],[`access_rules`],[`ua_rules`].

```
$ terraform import ibm_cis_firewall.myorg <firewall_type>:<firewall_id>:<domain-id>:<crn>

$ terraform import ibm_cis_firewall.myorg lockdowns lockdowns:48996f0da6ed76251b475971b097205c:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
