---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_range_app"
description: |-
  Provides a IBM CIS range application resource.
---

# ibm_cis_range_app
Create, update, or delete range application an IBM Cloud Internet Services domain. For more information, about range, see [protecting TCP traffic](https://cloud.ibm.com/docs/cis?topic=cis-cis-range).

## Example usage

```terraform
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

## Argument reference
Review the argument references that you can specify for your resource. 

- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `dns` - (Required, String) The name of DNS record for the range application.
- `dns_type` - (Required, String) The DNS record type. 
- `domain_id` - (Required, String) The ID of the domain where you want  to add the range app. 
- `edge_ips_type` - (Optional, String) The type of edge IP configuration. Valid value is `dynamic`. Default value is `dynamic`.
- `edge_ips_connectivity` - (Optional, String) Specified IP version. Valid values are `ipv4`, `ipv6`, `all`. Default value is `all`.
- `ip_firewall` - (Optional, Bool) Enables the IP firewall for the application. Only available for TCP applications.
- `origin_direct`-List of string-Optional-A list of destination addresses to the origin. IP address and port of the origin for range application. If configuring a Load Balancer, use `origin_dns` and `origin_port`. This cannot be combined with `origin_dns` and `origin_port`. For example, `tcp://192.0.2.1:22`.
- `origin_dns` - (Optional, String)  DNS record pointing to the origin for the range application. This is used for configuring a Load Balancer. This requires `origin_port` and cannot be combined with `origin_direct`. When specifying an individual IP address, use `origin_direct`. For example, `origin.net`.
- `origin_port` - (Optional, Integer) Port at the origin that listens to traffic from the range application. Requires `origin_dns` and cannot be combined with `origin_direct`.
- `protocol` - (Required, String) The edge application protocol type. Valid values are `tcp`, `udp`. This attribute specified along with port number. For example, `tcp/22`.
- `proxy_protocol` - (Optional, String)  Allows for the true client IP to be passed to the service. Valid values are `off`, `v1`, `v2`, `simple`. Default value is `off`. 
- `traffic_type` - (Optional, String) Configure how traffic is handled at the edge. If set to direct traffic is passed through to the service. In the case of HTTP or HTTPS, HTTPS features at the edge are applied to this traffic. Valid values are `direct`, `http`, `https`. Default value is `direct`.
- `tls` - (Optional, String) Configure how TLS connections are terminated at the edge. Valid values are `off`, `flexible`, `full`, `strict`. Default value is `off`.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The ID of the range application in the format `<app_id>:<domain_id>:<cis_id>`.
- `app_id` - (String) The range application ID.

## Import
The `ibm_cis_range_app` resource can be imported using the ID. The ID is formed from the application ID, the Domain ID of the domain and the Cloud Resource Name (CRN) concatenated  by using a `:` character.

The Domain ID and CRN will be located on the overview page of the Internet Services instance from the Domain heading of the console, or by using the `ibm cis` command line commands.

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

- **App ID** is a 32 digit character string of the form: `489d96f0da6ed76251b475971b097205c.`

**Syntax**

```
$ terraform import ibm_cis_range_app.myorg <app_id>:<domain-id>:<crn>
```


**Example**

```
$ terraform import ibm_cis_range_app.myorg 48996f0da6ed76251b475971b097205c:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```

