---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM : cdn"
description: |-
  Manages IBM Cloud CDN.
---

# ibm_cdn
Create, update, or delete a Content Delivery Networks (CDN) mapping. For more information, about IBM Cloud CDN, see [about Content Delivery Networks](https://cloud.ibm.com/docs/CDN?topic=CDN-about-content-delivery-networks-cdn-).

# DEPRECATED
CDN has now deprecated, backend services will no longer available after 28th March 2025. This docs will be removed in coming release.

## Example usage

```terraform
resource "ibm_cdn" "test_cdn1" {
  hostname = "www.default.com"
  vendor_name = "akamai"
  origin_address = "111.111.111.5"
  origin_type = "HOST_SERVER"
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `bucket_name` - (Required, String) If `origin_type` is set to `OBJECT_STORAGE`, you must provide the name of the bucket to use.
- `cache_key_query_rule` - (Optional, String) The rule for caching keys. Valid values are **include-all** - (includes all query arguments), **ignore-all** - (ignores all query arguments), **ignore: space separated query-args** - (ignores specific query arguments). The default value is **include-all**.
- `certificate_type` - (Conditional, Forces new resource, String) The type of certificate to use. This value is required if `protocol` is set to `HTTPS`. Valid values are `SHARED_SAN_CERT` or `WILDCARD_CERT`.
- `cname` - (Optional, Forces new resource, String) The CNAME for your CDN.
- `file_extension` - (Optional, String) If `origin_type` is set to `OBJECT_STORAGE`, you can specify the file extensions that you want to cache.
- `header` - (Optional, String) The header for the CDN.
- `host_name` - (Required, Forces new resource, String) The host name that is associated with the CDN domain mapping.
- `http_port` - (Optional, Integer) The port to be opened up. Default value is 80. This option can be set only if you use `HTTP` or `HTTPS` as the `protocol`.
- `https_port` - (Optional, Integer) The HTTPS port. Default value is 0. This option can be set only if you use `HTTP` or `HTTPS` as the `protocol`.
- `origin_address` - (Required, String) The IP address or hostname for the domain mapping. If **origin_type=HOST_SERVER** provide the hostname or IP address. If **origin_type=OBJECT_STORAGE** provide your [COS endpoints](https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-endpoints). 
- `origin_type` - (Required, Forces new resource, String) The type of storage to use. Valid values are `HOST_SERVER` or `OBJECT_STORAGE`.
- `path` - (Optional, Forces new resource, String) The path for the CDN.
- `performance_configuration` - (Optional, String) The performance configuration. Default is `General web delivery`.
- `protocol` - (Optional, Forces new resource, String) The protocol to use. Default value is `HTTP`.
- `respect_headers` - (Optional, Bool)  If set to **true**, the TTL settings in the origin override CDN TTL settings.
- `vendor_name` - (Required, Forces new resource, String)  Only `akamai` is supported for now.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique internal identifier of the CDN domain mapping.
- `status` - (String) The Status of the CDN domain mapping.
