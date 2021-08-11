---
subcategory: "Satellite"
layout: "ibm"
page_title: "IBM: ibm_satellite_location_nlb_dns"
description: |-
  Get information about registered NLB subdomains of Satellite location.
---

# ibm_satellite_location_nlb_dns
List registered NLB subdomains in a Satellite location.


## Example usage
The following example retrieves information about NLB subdomains of a cluster that is named `mycluster`. 

```terraform
data ibm_satellite_location_nlb_dns dns {
    name ="mycluster"
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `location` - (Required, String) The name or ID of the Satellite Location.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `nlb_config` - List of objects 

  Nested scheme for `nlb_config`:
  - `dns_type` -  (String) Type of DNS.
  - `lb_hostname` - (String) Host Name of load Balancer.
  - `nlb_ips` - (List(String)) NLB IPs.
  - `nlb_sub_domain`- (String) NLB Sub-Domain.
  - `secret_name` - (String) Name of the secret.
  - `secret_namespace` - (String) Namespace of Secret.
  - `secret_status` - (String) Status of Secret.
  - `type` -  (String)  Nlb Type.
