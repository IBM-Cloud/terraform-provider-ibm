---

subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : "
description: |-
  Manages IBM Cloud Infrastructure Private Domain Name Service GLB monitors.
---

# ibm\_dns_glb_monitors

Import the details of an existing IBM Cloud Infrastructure private domain name service GLB monitors as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl

data "ibm_dns_glb_monitors" "ds_pdns_glb_monitors" {
  instance_id = "resource_instance_guid"
}

```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, string) The GUID of the private DNS. 



## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `dns_glb_monitors` - List of all private domain name service GLB monitors in the IBM Cloud Infrastructure.
  * `name` - The name of the load balancer monitor.
  * `description` -   Descriptive text of the load balancer monitor.
  * `type` -  The protocol to use for the health check. Currently supported protocols are 'HTTP','HTTPS' and 'TCP'.
  * `port` - Port number to connect to for the health check. Required for TCP checks. HTTP and HTTPS checks
  * `interval` - The interval between each health check.
  * `retries` - The number of retries to attempt in case of a timeout before marking the origin as unhealthy.
  * `timeout` - The timeout (in seconds) before marking the health check as failed.
  * `method` - The method to use for the health check applicable to HTTP/HTTPS based checks, the default value is 'GET'.
  * `path` - The endpoint path to health check against. This parameter is only valid for HTTP and HTTPS monitors.
  * `headers` - The HTTP request headers to send in the health check.
    * `name` - The name of HTTP request header.
    * `value` - The value of HTTP request header.
  * `allow_insecure` -  Do not validate the certificate when monitor use HTTPS.
  * `expected_codes` - The expected HTTP response code or code range of the health check.
  * `expected_body` - A case-insensitive sub-string to look for in the response body.
  * `monitor_id` - Monitor Id.

