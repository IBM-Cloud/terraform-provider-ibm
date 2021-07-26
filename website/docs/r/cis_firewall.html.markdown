---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_firewall"
description: |-
  Provides a IBM CIS Firewall resource.
---

# ibm_cis_firewall


Create, update, or delete a firewall for a domain that you included in your IBM Cloud Internet Services instance and a CIS domain resource. For more information, about CIS firewall resource, see [using fields, functions, and expressions](https://cloud.ibm.com/docs/cis?topic=cis-fields-and-expressions).

## Example usage

```terraform
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

## Argument reference
Review the argument references that you can specify for your resource. 

- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance where you want to create the firewall.
- `domain_id` - (Required, String) The ID of the domain where you want to apply the firewall rules.
- `firewall_type` - (Required, String) The type of firewall that you want to create for your domain. Supported values are `lockdowns`, `access_rules`, and `ua_rules`. Consider the following information when choosing your firewall type: <ul><li><strong><code>access_rules</code></strong>: Access rules allow, challenge, or block requests to your website. You can apply access rules to one domain only or all domains in the same service instance.</li><li><strong><code>ua_rules</code></strong>: Apply firewall rules only if the user agent that is used by the client matches the user agent that you defined. </li><li><strong><code>lockdowns</code></strong>: Allow access to your domain for specific IP addresses or IP address ranges only. If you choose this firewall type, you must define your firewall rules in the `lockdown` input parameter.</li></ul>.
- `access_rule` - (Optional, String)  Create the data the describing access rule. (Maximum item is 1).
 
  Nested scheme for `access_rules`:	
  - `configuration` - (Required, List)  The Configuration of firewall. (Maximum items is 1).
  
    Nested scheme for `configuration`: 
    - `target` - (Required, String) The request property to target. Valid values are `ip`, `ip_range`, `asn`, `country`.
    - `value` - (Required, String)  IP address or CIDR or Autonomous or Country code.
  - `mode` - (Required, String) The mode of access rule. The valid modes are `block`, `challenge`, `whitelist`, `js_challenge`.
  - `notes` - (Optional, String) The free text for notes.
- `lockdown`- (Required, List) A list of firewall rules that you want to create for your `lockdowns` firewall. You can specify one item in this list only.

  Nested scheme for `lockdown`:
  - `configurations`- (Optional, List) A list of IP address or CIDR ranges that you want to allow access to the URLs that you defined in `urls`.

    Nested scheme for `configurations`:
    - `target` - (Optional, String) Specify if you want to target an `IP` or `ip_range`.
    - `value` - (Optional, String) The IP address or IP address range that you want to target. Make sure that the value that you enter here matches the type of target that you specified in `lockdown.configurations.target`.
  - `description` - (Optional, String) A description for your firewall rule.
  - `paused`- (Optional, Bool) If set to **true**, the firewall rule is disabled. If set to **false**, the firewall rule is enabled.
  - `priority` - (Optional, Integer) The priority of the firewall rule. A low number is associated with a high priority.
  - `urls`- (Optional, List) A list of URLs that you want to include in your firewall rule. You can specify wildcard URLs. The URL pattern is escaped before use.
- `ua_rule` - (Optional, String) Create the data describing the user agent rule. (Maximum item is 1).

  Nested scheme for `ua_rule`:
  - `configuration` - (Required, List)  The Configuration of firewall. (Maximum item is 1).
  
    Nested scheme for `configuration`:
    - `target` - (Required, String) The request property to target. Valid values are `ua`.
    - `value` - (Required, String) The exact user agent string to match the rule.
  - `description ` - (Optional, String) The free text for description.
  - `mode` - (Optional, String) The mode of access rule. The valid modes are `block`, `challenge`,  `js_challenge`.
  - `paused` - (Optional, String) Whether the rule is currently disabled.
  

**Note**

Exactly one of `lockdown`, `access_rule`, and `ua_rule` is allowed for the firewall types `lockdowns`, `access_rules`, and `ua_rules`.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `access_rule_id` - (String) The access rule ID.
- `id` - (String) The ID of the record. The ID is composed of `<firewall_type>,<lockdown_id/access_rule_id/ua_rule_id>,<domain_ID>,<cis_crn>`. Attributes are concatenated with `:`.
- `lockdown_id` - (String) The lock down ID.
- `ua_rule_id` - (String) The user agent rule ID.

## Import
The `ibm_cis_firewall` resource is imported by using the ID. The ID is formed from the firewall type, the firewall ID, the domain ID of the domain and the CRN (Cloud Resource Name) concatenated  using a `:` character.

The domain ID and CRN is located on the **overview** page of the internet services instance of the domain heading of the console, or by using the `ibm cis` command line commands.

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`.

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`.

- **Firewall ID** is a 32 digit character string of the form: `489d96f0da6ed76251b475971b097205c`.

- **Firewall type** is a string. It can be either of `lockdowns`, `access_rules`, `ua_rules`.

**Syntax**

```
$ terraform import ibm_cis_firewall.myorg <firewall_type>:<firewall_id>:<domain-id>:<crn>
```

**Example**

```
$ terraform import ibm_cis_firewall.myorg lockdowns lockdowns:48996f0da6ed76251b475971b097205c:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```

