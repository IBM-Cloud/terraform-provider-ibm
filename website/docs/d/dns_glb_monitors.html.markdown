---
subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : "
description: |-
  Manages IBM Cloud Infrastructure Private Domain Name Service GLB monitors.
---

# ibm_dns_glb_monitors

Retrieve the details of an existing IBM Cloud infrastructure private DNS Global Load Balancers monitors as a read-only data source. For more information, see [viewing Global Load Balancer events](https://cloud.ibm.com/docs/dns-svcs?topic=dns-svcs-health-check-events).


## Example usage

```terraform
data "ibm_dns_glb_monitors" "ds_pdns_glb_monitors" {
  instance_id = "resource_instance_guid"
}
```

## Argument reference
Review the argument reference that you can specify for your data source. 

- `instance_id` - (Required, String) The GUID of the private DNS service instance.

## Attribute reference
In addition to the argument references list, you can access the following attribute references after your data source is created. 

- `dns_glb_monitors` (List) List of all private DNS Load balancer monitors in the IBM Cloud infrastructure.
 
   Nested scheme for `dns_glb_monitors`:
   - `allow_insecure` - (String) Do not validate the certificate when monitor use HTTPS.
   - `description` - (String) The descriptive text of the DNS Load balancer monitor.
   - `expected_codes` - (String) The expected HTTP response code or code range of the health check.
   - `expected_body` - (String) A case insensitive substring to look for in the response body.
   - `headers` - (String) The HTTP request headers to send in the health check.
      
     Nested scheme for `headers`:
     - `name` - (String) The name of the HTTP request header.
     - `value` - (String) The value of the HTTP request header.
   - `interval` - (String) The interval between each health check.
   - `method` - (String) The method to use for the health check applicable to HTTP, HTTPS based checks, the default value is `GET`.
   - `monitor_id` - (String) The monitor ID.
   - `name` - (String) The name of the DNS Load balancer monitor.
   - `path` - (String) The endpoint path to health check against. This parameter is only valid for HTTP and HTTPS monitors.
   - `port` - (String)  Port number to connect to for the health check. Required for TCP checks, HTTP, and HTTPS checks.
   - `retries`(String) The number of retries to attempt in case of a timeout before marking the origin as unhealthy.
   - `timeout` - (String) The timeout (in seconds) before marking the health check as failed.
   - `type` - (String) The protocol to use for the health check. Currently supported protocols are `HTTP`, `HTTPS`, and `TCP`.
   
