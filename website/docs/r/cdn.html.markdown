---
layout: "ibm"
page_title: "IBM : cdn"
sidebar_current: "docs-ibm-resource-cdn"
description: |-
  Manages IBM cdn.
---

# ibm\_cdn

This iresource is used to order a cdn domain mapping.

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

* `hostname` - (Required,  string) Hostname associated with the cdn domain mapping.
* `cname` - (Optional,  string) enter a unique cname for your cdn.
* `path` - (Optional,  string) enter the path for the cdn .
* `vendor_name` - (Required,  string) only “akamai” is supported for now.
* `origin_type` - (Required,  string) mention the type of storage. It can be “HOST_SERVER” or “OBJECT_STORAGE”.
* `origin_address` - (Required,  string) Provide the IP address for domain mapping.
* `protocol` - (Optional, string) “HTTP is taken as default”.
* `httpport` - (Optional, Int) 80 is taken as default. **NOTE**: It can only be populated if protocol is set to “HTTP” or “HTTP_AND_HTTPS”
* `httpsport` - (Optional, Int) 0 is taken as default. **NOTE**: It can only be populated if protocol is set to “HTTPS” or “HTTP_AND_HTTPS”
* `bucketname` - (Required, string) required for “OBJECT_STORAGE” origin_type only.

## Attribute Reference

The following attributes are exported:

* `id` - The unique internal identifier of the cdn domian mapping.
