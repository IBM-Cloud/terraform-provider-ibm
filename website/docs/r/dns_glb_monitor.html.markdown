---
subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : dns_glb_monitor"
description: |-
  Manages IBM Private DNS GLB monitor.
---

# ibm_dns_glb_monitor

Provides a private DNS GLB (GLB) monitor resource. This allows DNS (GLB)) monitor to create, update, and delete. For more information, see [Viewing GLB events](https://cloud.ibm.com/docs/dns-svcs?topic=dns-svcs-health-check-events). 


## Example usage

```terraform
resource "ibm_dns_glb_monitor" "test-pdns-monitor" {
  depends_on     = [ibm_dns_zone.test-pdns-zone]
  name           = "test-pdns-glb-monitor"
  instance_id    = ibm_resource_instance.test-pdns-instance.guid
  description    = "test monitor description"
  interval       = 63
  retries        = 3
  timeout        = 8
  port           = 8080
  type           = "HTTP"
  expected_codes = "200"
  path           = "/health"
  method         = "GET"
  expected_body  = "alive"
  headers {
    name  = "headerName"
    value = ["example", "abc"]
  }
}
```

## Argument reference
Review the argument reference that you can specify for your resource. 

- `allow_insecure` - (Optional, String) Do not validate the certificate when monitor use HTTPS. This parameter is currently only valid for HTTPS monitors.
- `description` - (Optional, String)  Descriptive text of the Load Balancer monitor.
- `expected_body` - (Optional, String) A case-insensitive sub-string to look in the response body. If the string is not found, the origin will be marked as unhealthy. This parameter is only valid for HTTP and HTTPS monitors.
- `expected_codes` - (Optional, String) The expected HTTP response code or code range of the health check. This parameter is only valid for HTTP and HTTPS monitors. Allowable values are `200, 201, 202, 203, 204, 205, 206, 207, 208, 226, 2xx, 3xx, 4xx, 5xx`.
- `headers` - (Optional, Set) The HTTP request headers to send in the health check. It is recommended you set a host header by default. The `User-Agent` header cannot be overridden. This parameter is only valid for HTTP and HTTPS monitors.

  Nested scheme for `headers`:
	- `name` - (Required, String) The name of the HTTP request header.
  - `value`- (Required, List of string) The value of HTTP request header.
- `interval` - (Optional, Integer) The interval between each health check.
- `instance_id` - (Required, Forces new resource, String) The GUID of the private DNS.
- `method` - (Optional, String) The method to use for the health check applicable to HTTP or HTTPS based checks, the default value is `GET`.
- `name` - (Required, String) The name of the Load Balancer monitor.
- `path` - (Optional, String) The endpoint path to health check against. This parameter is only valid for HTTP and HTTPS monitors.
- `port` - (Optional, Integer) The port number to connect to for the health check. Required for TCP checks. HTTP and HTTPS checks should only define the port when using a non-standard port. For example, HTTP  default is `80`, and HTTPS default is `443`).
- `retries` - (Optional, Integer) The number of retries to attempt in case of a timeout before marking the origin as unhealthy.
- `timeout` - (Optional, Integer) The timeout (in seconds) before marking the health check as failed.
- `type` - (Optional, Forces new resource, String) The protocol to use for the health check. Currently supported protocols are `HTTP`,`HTTPS` and `TCP`. Default Value is `HTTP`.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your resource is created. 

- `created_on` - (Timestamp) The time (created on) of the DNS GLB monitor. 
- `id` - (String) The unique ID of the private DNS Monitor. The ID is composed of `<instance_id>/<glb_monitor_id>`. 
- `modified_on` - (Timestamp) The time (modified on) of the DNS GLB monitor.
- `monitor_id`- (String) The monitor ID.

## Import
The `ibm_dns_glb_monitor` can be imported by using private DNS instance ID, and GLB Monitor ID.

**Example**

```
$ terraform import ibm_dns_glb_monitor.example 6ffda12064634723b079acdb018ef308/435da12064634723b079acdb018ef308
```
