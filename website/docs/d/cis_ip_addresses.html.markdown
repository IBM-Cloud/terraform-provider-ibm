---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_ip_addresses"
description: |-
  List the IP addresses used by name servers by Cloud Internet Services. Required for setting whitelist addresses for internet facing application ports.
---

# ibm_cis_ip_addresses
Retrieve information of all IP addresses that the CIS proxy uses. The CIS proxy uses these IP addresses for both `client-to-proxy` and `proxy-to-origin` communication. You can reference the IP addresses by using Terraform interpolation syntax to configure and allowed IP addresses in firewalls, network ACLs, and security groups. For more information, about CIS IP addressess, see [best practices for CIS setup](https://cloud.ibm.com/docs/cis?topic=cis-best-practices-for-cis-setup).

## Example usage
The following example retrieves information about IP addresses that IBM Cloud Internet Services uses for name servers.

```terraform
data "ibm_cis_ip_addresses" "ip_addresses" {}
```

## Argument reference
No arguments are required. All CIS instances on an account use the same range of name servers.

## Attribute reference
You can access the following attribute references after your data source is created. 

- `ipv4_cidrs` - (String) The IPv4 address ranges that the CIS proxy uses and that you can reference to configure and allowed IP addresses in firewalls, network ACLs, and security groups.
- `ipv6_cidrs` - (String) The IPv6 address ranges that the CIS proxy uses and that you can reference to configure and allowed IP addresses in firewalls, network ACLs, and security groups.
