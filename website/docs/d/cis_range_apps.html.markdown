---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_range_apps"
description: |-
  Get information of IBM Cloud Internet Services range applicatiions.
---

# ibm_cis_range_apps
Retrieve an information of an IBM Cloud Internet Services range applications. For more information, about CIS range application, see [getting started with range](https://cloud.ibm.com/docs/cis?topic=cis-cis-range).

## Example usage

```terraform
data "ibm_cis_range_apps" "apps" {
  cis_id    = ibm_cis.instance.id
  domain_id = ibm_cis_domain.example.id
}
```
## Argument reference
Review the argument references that you can specify for your data source.  

- `app_id` -  (String) The Range application id.
- `cis_id` -  (String) The ID of the CIS service instance.
- `domain_id` - (String) The ID of the domain to add the range application.
- `dns` -  (String) The name of DNS record for the range application.
- `dns_type` -  (String) The DNS record type.
- `edge_ips_type` -  (String) The type of edge IP configuration. Valid value and default value is `dynamic`.
- `edge_ips_connectivity` - (String)  Specified IP version. Valid values are `ipv4`, `ipv6`, `all`. Default value is `all`.
- `id` - (String) The range application ID. It is a combination of `<app_id>,<domain_id>,<cis_id>` attributes are concatenated with `:` character.
- `ip_firewall` - (String)  Enables the IP firewall for the application. Only available for `TCP` applications.
- `origin_direct` -  (String) A list of destination addresses to the origin. IP address and port of the origin for Range application. If configuring a Load Balancer, use `origin_dns` and `origin_port`. This cannot be combined with `origin_dns` and `origin_port`. For example, `tcp://192.0.2.1:22`.
- `protocol` -  (String) The Edge application protocol type. Valid values are `tcp`, `udp`. This attribute specified along with port number. For example, `tcp/22`.
- `proxy_protocol` -  (String) Allows for the true client IP to be passed to the service. Valid values are `off`, `v1`, `v2`, `simple`. Default value is `off`.
- `traffic_type` -  (String) Configure how traffic is handled at the edge. If set to direct traffic is passed through to the service. In the case of HTTP or HTTPS, HTTPS features at the edge are applied to this traffic. Valid values are `direct`, `http`, `https`. Default value is `direct`.
- `tls` -  (String) Configure how TLS connections are terminated at the edge. Valid values are `off`, `flexible`, `full`, `strict`. Default value is `off`.
