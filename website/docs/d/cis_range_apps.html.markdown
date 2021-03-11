---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_range_apps"
description: |-
  Get information of IBM Cloud Internet Services Range Applicatiions.
---

# ibm_cis_range_apps

Imports a read only copy of an existing Internet Services Range Applications.

## Example Usage

```hcl
data "ibm_cis_range_apps" "apps" {
  cis_id    = ibm_cis.instance.id
  domain_id = ibm_cis_domain.example.id
}
```

## Argument Reference

The following arguments are supported:

- `id` - The range application ID. It is a combination of <`app_id`>,<`domain_id`>,<`cis_id`> attributes concatenated with ":".
- `app_id` - The Range application id.
- `cis_id` - The ID of the CIS service instance
- `domain_id` - The ID of the domain to add the range app.
- `protocol` - The Edge application protocol type. Valid values: `tcp`, `udp`. This attribute specified along with port number. Ex. `tcp/22`.
- `dns` - The name of DNS record for the range application.
- `dns_type` - The DNS record type.
- `origin_direct` - A list of destination addresses to the origin. IP address and port of the origin for Range application. If configuring a load balancer, use `origin_dns` and `origin_port`. This can not be combined with `origin_dns` and `origin_port`. Ex. `["tcp://192.0.2.1:22"]`
- `ip_firewall` - (Optional,boolean) Enables the IP Firewall for this application. Only available for TCP applications.
- `proxy_protocol` - Allows for the true client IP to be passed to the service. Valid values: `off`, `v1`, `v2`, `simple`. Default value is `off`.
- `edge_ips_type` - The type of edge IP configuration. Valid value : `dynamic`. Default value is `dynamic`.
- `edge_ips_connectivity` - Specified IP version. Valid value: `ipv4`, `ipv6`, `all`. Default value is `all`.
- `traffic_type` - Configure how traffic is handled at the edge. If set to `direct` traffic is passed through to the service. In the case of `http` or `https` HTTP/s features at the edge are applied ot this traffic. Valid values: `direct`, `http`, `https`. Default value is `direct`.
- `tls` - Configure if and how TLS connections are terminated at the edge. Valid values: `off`, `flexible`, `full`, `strict`. Default value is `off`.
