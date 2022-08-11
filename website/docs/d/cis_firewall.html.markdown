---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_firewall"
description: |-
  Get information on an IBM Cloud Internet Services Firewall.
---

# ibm_cis_firewall
Retrieve information about an existing IBM Cloud Internet Services instance. For more information, see [firewall rule actions](https://cloud.ibm.com/docs/cis?topic=cis-actions).

## Example usage

```terraform
data "ibm_cis_firewall" "lockdown" {
  cis_id    = ibm_cis.instance.id
  domain_id = ibm_cis_domain.example.id
  firewall_type = "lockdowns"
}
```
**Note**
IBM Terraform provider supports only lock down rules.

## Argument reference
Review the argument references that you can specify for your data source. 

- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance where you want to create the firewall.
- `domain_id` - (Required, String) The ID of the domain where you want to add the lock down.
- `firewall_type` - (Required, String) The type of firewall that you want to create for your domain. Supported values are `lockdowns`, `access_rules`, and `ua_rules`. **Note** Consider the following information when choosing your firewall type: <ul><li><strong><code>access_rules</code></strong>: Access rules allow, challenge, or block requests to your website. You can apply access rules to one domain only or all domains in the same service instance.</li><li><strong><code>`lockdowns`</code></strong>: Allow access to your domain for specific IP addresses or IP address ranges only. If you choose this firewall type, you must define your firewall rules in the `lockdown` input parameter.</li><li><strong><code>ua_rules</code></strong>: Apply firewall rules only if the user agent that is used by the client matches the user agent that you defined. </li></ul>.

## Attribute reference
Review the attribute references that you can access after your data source is created. 

- `access_rule` - (String) Create the data describing the access rule.
	
	Nested scheme for `access_rule`:
	- `access_rule_id` - (String) The access rule ID.
	- `configuration` (List) The Configuration of firewall. (Maximum items is 1).

	  Nested scheme for `configuration`:
	  - `target` - (String) The request property to target. Valid values are `ip`, `ip_range`, `asn`, `country`.
	  - `value` - (String) IP address or CIDR or Autonomous or Country code.
	- `mode` - (String) The mode of access rule. The valid modes are `block`, `challenge`, `whitelist`, `js_challenge`.
	- `notes` - (String) The free text for notes.
- `id` - (String) The ID of the record. The ID is composed of `<firewall_type>,<lockdown_id/access_rule_id/ua_rule_id>,<domain_ID>,<cis_crn>`. Attributes are concatenated with `:`.
- `lockdown` (List) List of lock down to be created. The data describing a lock down rule.

  Nested scheme for `lockdown`:
	- `configurations`- (List) A list of IP address or CIDR ranges that you want to allow access to the URLs that you defined in `lockdown.urls`.

	  Nested scheme for `configurations`:
		- `target` - (String) Specify if you want to target an `IP` or `ip_range`.
		- `value` - (String) The IP addresses or CIDR.
	- `description` - (String) A description for your firewall rule.
	- `lockdown_id` (String) The lock down ID.
	- `paused`- (Bool) If set to **true**, the firewall rule is disabled. If set to **false**, the firewall rule is enabled.
	- `priority`- (Integer) The priority of the firewall rule. A low number is associated with a high priority.
	- `urls`-List of URLs-A list of URLs that you want to include in your firewall rule. You can specify wildcard URLs. The URL pattern is escaped before use.
- `ua_rule` - (String) Create the data describing the user agent rule.
	
   Nested scheme for `ua_rule`:
   - `configuration` (List) The Configuration of firewall.

     Nested scheme for `configuration`:
	 - `target` - (String) The request property to target. Valid values are `ua`.
	  - `value` - (String) The exact User Agent string to match the rule.
   - `description ` - (String) The free text for description.
   - `mode` - (String) The mode of access rule. The valid modes are `block`, `challenge`,  `js_challenge`.
   - `paused` - (String) Whether the rule is currently disabled.
   - `ua_rule_id` - (String) The user agent rule ID.

**Note**

Exactly one of `lockdown`, `access_rule`, and `ua_rule` is allowed for the firewall types `lockdowns`, `access_rules`, and `ua_rules`.

