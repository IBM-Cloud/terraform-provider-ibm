---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM : cdn"
description: |-
  Manages IBM cdn.
---

# ibm\_cdn

This resource is used to order a cdn domain mapping.

## Example Usage

```hcl
resource "ibm_cdn" "test_cdn1" {
  hostname = "www.default.com"
  vendor_name = "akamai"
  origin_address = "111.111.111.5"
  origin_type = "HOST_SERVER"
}
```

## Argument Reference

* `host_name` - (Required, Forces new resource, string) Hostname associated with the cdn domain mapping.
* `cname` - (Optional, Forces new resource, string) enter a unique cname for your cdn.
* `path` - (Optional, Forces new resource, string) enter the path for the cdn .
* `vendor_name` - (Required, Forces new resource, string) only "akamai" is supported for now.
* `origin_type` - (Required, Forces new resource, string) mention the type of storage. It can be "HOST_SERVER" or "OBJECT_STORAGE".
* `origin_address` - (Required,  string) Provide the IP address for domain mapping.
* `protocol` - (Optional, Forces new resource, string) "HTTP is taken as default".
* `http_port` - (Optional, Int) 80 is taken as default. **NOTE**: It can only be populated if protocol is set to "HTTP" or "HTTP_AND_HTTPS"
* `https_port` - (Optional, Int) 0 is taken as default. **NOTE**: It can only be populated if protocol is set to "HTTPS" or "HTTP_AND_HTTPS"
* `bucket_name` - (Required, string) required for "OBJECT_STORAGE" origin_type only.
* `certificate_type`: (Optional, Forces new resource) required for HTTPS protocol. SHARED_SAN_CERT or WILDCARD_CERT.
* `header`: (Optional, string) Provide Header for CDN.
* `respect_headers`: (Optional, bool) A boolean value that, if set to true, will cause TTL settings in the Origin to override CDN TTL settings.
* `file_extension` - (optional for Object Storage) File extensions that are allowed to be cached.
* `performance_configuration` - (Optional, Forces new resource, string) Provide the performance configuration. Default is "General web delivery"
* `cache_key_query_rule`: (Optional, string) The following options are available to configure Cache Key behavior:
    include-all - includes all query arguments default
    ignore-all - ignores all query arguments
    ignore: space separated query-args - ignores those specific query arguments. For example, ignore: query1 query2
    include: space separated query-args: includes those specific query arguments. For example, include: query1 query2

## Attribute Reference

The following attributes are exported:

* `id` - The unique internal identifier of the cdn domian mapping.
* `status` - The Status of the cdn domian mapping.