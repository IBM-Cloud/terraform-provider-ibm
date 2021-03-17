---

subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : dns_glb_monitor"
description: |-
  Manages IBM Private DNS GLB monitor.
---

# ibm\_dns_glb_monitor

Provides a private dns GLB monitor resource. This allows dns GLB monitor to be created,updated and deleted.

## Example Usage

```hcl

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

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, string,ForceNew) The GUID of the private DNS. 
* `name` - (Required, string) The name of the load balancer monitor.
* `description` -  (Optional,string) Descriptive text of the load balancer monitor.
* `type` - (Optional, string) The protocol to use for the health check. Currently supported protocols are 'HTTP','HTTPS' and 'TCP'.Default Value:HTTP
* `port` - (Optional, int) Port number to connect to for the health check. Required for TCP checks. HTTP and HTTPS checks should only define the port when using a non-standard port (HTTP: default 80, HTTPS: default 443).
* `interval` - (Optional, int) The interval between each health check.
* `retries` - (Optional, int) The number of retries to attempt in case of a timeout before marking the origin as unhealthy.
* `timeout` - (Optional, int) The timeout (in seconds) before marking the health check as failed.
* `method` - (Optional, string) The method to use for the health check applicable to HTTP/HTTPS based checks, the default value is 'GET'.
* `path` - (Optional, string) The endpoint path to health check against. This parameter is only valid for HTTP and HTTPS monitors.
* `headers` - (Optional, set) The HTTP request headers to send in the health check. It is recommended you set a Host header by default. The User-Agent header cannot be overridden. This parameter is only valid for HTTP and HTTPS monitors.
  * `name` - (Required, string) The name of HTTP request header.
  * `value` - (Required, list of string) The value of HTTP request header.
* `allow_insecure` - (Optional, string) Do not validate the certificate when monitor use HTTPS. This parameter is currently only valid for HTTPS monitors.
* `expected_codes` - (Optional, string) The expected HTTP response code or code range of the health check. This parameter is only valid for HTTP and HTTPS monitors.Allowable values: [200,201,202,203,204,205,206,207,208,226,2xx]
* `expected_body` - (Optional, string) A case-insensitive sub-string to look for in the response body. If this string is not found, the origin will be marked as unhealthy. This parameter is only valid for HTTP and HTTPS monitors.



## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the private DNS Monitor. The id is composed of <instance_id>/<glb_monitor_id>.
* `created_on` - The time (Created On) of the DNS glb monitor. 
* `modified_on` - The time (Modified On) of the DNS glb monitor. 
* `monitor_id` - Monitor Id.

## Import

ibm_dns_glb_monitor can be imported using private DNS instance ID and GLB Monitor ID, eg

```
$ terraform import ibm_dns_glb_monitor.example 6ffda12064634723b079acdb018ef308/435da12064634723b079acdb018ef308
```