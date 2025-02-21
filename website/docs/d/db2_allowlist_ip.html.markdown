---
subcategory: "Db2 SaaS"
layout: "ibm"
page_title: "IBM : ibm_db2_allowlist_ip"
description: |-
  Get Information about allowed list of IPs of IBM Db2 instance.
---

# ibm_db2_allowlist_ip

Retrieve information about allowed list of IPs of an existing [IBM Db2 Instance](https://cloud.ibm.com/docs/Db2onCloud).

## Example Usage

```hcl
data "ibm_db2_allowlist_ip" "db2_allowlistips" {
    x_deployment_id = "<crn>"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `x_deployment_id` - (Required, String) CRN of the instance this allowed list of IPs relates to.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.
* `ip_addresses` - (string) A List of IP addresses.
Nested scheme for **ip_addresses**:
    * `address` - (String) The IP address, in IPv4 format.
    * `description` - (String) Description of the IP address.