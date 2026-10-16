---
subcategory: "Container Registry"
layout: "ibm"
page_title: "IBM : ibm_cr_quota"
description: |-
  Manages IBM Cloud Container Registry storage and traffic quotas.
---

# ibm_cr_quota

Manages the account-level storage and pull traffic quotas of Container Registry in one region. Only the quotas you set are managed.

## Example usage

```terraform
resource "ibm_cr_quota" "quota" {
  endpoint          = "https://us.icr.io"
  storage_megabytes = 5120
  traffic_megabytes = 10240
}
```

## Argument reference

- `endpoint` - (Optional, Forces new resource, String) Regional Container Registry endpoint. Defaults to the provider endpoint.
- `storage_megabytes` - (Optional, Integer) Storage quota in MB. `-1` means unlimited.
- `traffic_megabytes` - (Optional, Integer) Pull traffic quota in MB. `-1` means unlimited.
- `reset_on_destroy` - (Optional, Boolean) Reset managed quotas to the defaults (500 MB storage, 5120 MB traffic) on destroy. Default `true`.

At least one of `storage_megabytes` or `traffic_megabytes` must be set.

## Attribute reference

- `id` - Account ID and endpoint.
- `storage_usage_bytes` - Current storage usage in bytes.
- `traffic_usage_bytes` - Current pull traffic usage in bytes.
