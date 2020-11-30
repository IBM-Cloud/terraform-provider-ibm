---
layout: "ibm"
page_title: "IBM: ibm_cis_range_app"
sidebar_current: "docs-ibm-resource-cis-range-app"
description: |-
  Provides a IBM CIS Range Application resource.
---

# ibm_cis_range_app

Provides a IBM CIS Range Application resource. This resource is associated with an IBM Cloud Internet Services instance and a CIS Domain resource. It allows to create, update, delete range app of a domain of a CIS instance

## Example Usage

```hcl
resource "ibm_cis_range_app" "app" {
	cis_id         = data.ibm_cis.cis.id
	domain_id      = data.ibm_cis_domain.cis_domain.id
	protocol       = "tcp/22"
	dns_type       = "CNAME"
	dns            = "ssh.example.com"
	origin_direct  = ["tcp://12.1.1.1:22"]
	ip_firewall    = true
	proxy_protocol = "v1"
	traffic_type   = "direct"
	tls            = "off"
}
```

## Argument Reference

The following arguments are supported:

- `cis_id` - (Required,string) The ID of the CIS service instance
- `domain_id` - (Required,string) The ID of the domain to add the range app.
- `protocol` - (Required,string) The Edge application protocol type. Valid values: `tcp`, `udp`. This attribute specified along with port number. Ex. `tcp/22`.
- `dns` - (Required,string) The name of DNS record for the range application.
- `dns_type` - (Required,string) The DNS record type.
- `origin_direct` - (Optional,list(string)) A list of destination addresses to the origin. IP address and port of the origin for Range application. If configuring a load balancer, use `origin_dns` and `origin_port`. This can not be combined with `origin_dns` and `origin_port`. Ex. `["tcp://192.0.2.1:22"]`
- `origin_dns` - (Optional,string) DNS record pointing to the origin for this Range application. This is used for configuring a load balancer. This requires `origin_port` and can not be combined with `origin_direct`. When specifying an individual IP address, use `origin_direct`. Ex. `origin.net`.
- `origin_port` - (Optional,integer) Port at the origin that listens to traffic from this Range application. Requires `origin_dns` and can not be combined with `origin_direct`.
- `ip_firewall` - (Optional,boolean) Enables the IP Firewall for this application. Only available for TCP applications.
- `proxy_protocol` - (Optional,string) Allows for the true client IP to be passed to the service. Valid values: `off`, `v1`, `v2`, `simple`. Default value is `off`.
- `edge_ips_type` - (Optional,string) The type of edge IP configuration. Valid value : `dynamic`. Default value is `dynamic`.
- `edge_ips_connectivity` - (Optional,string) Specified IP version. Valid value: `ipv4`, `ipv6`, `all`. Default value is `all`.
- `traffic_type` - (Optional,string) Configure how traffic is handled at the edge. If set to `direct` traffic is passed through to the service. In the case of `http` or `https` HTTP/s features at the edge are applied ot this traffic. Valid values: `direct`, `http`, `https`. Default value is `direct`.
- `tls` - (Optional,string) Configure if and how TLS connections are terminated at the edge. Valid values: `off`, `flexible`, `full`, `strict`. Default value is `off`.

## Attributes Reference

The following attributes are exported:

- `id` - The range application ID. It is a combination of <`app_id`>,<`domain_id`>,<`cis_id`> attributes concatenated with ":".
- `app_id` - The Range application id.

## Import

The `ibm_cis_range_app` resource can be imported using the `id`. The ID is formed from the `App ID`, the `Domain ID` of the domain and the `CRN` (Cloud Resource Name) concatentated using a `:` character.

The Domain ID and CRN will be located on the **Overview** page of the Internet Services instance under the **Domain** heading of the UI, or via using the `ibm cis` CLI commands.

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

- **App ID** is a 32 digit character string of the form: `489d96f0da6ed76251b475971b097205c`.

```
$ terraform import ibm_cis_range_app.myorg <app_id>:<domain-id>:<crn>

$ terraform import ibm_cis_range_app.myorg 48996f0da6ed76251b475971b097205c:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
