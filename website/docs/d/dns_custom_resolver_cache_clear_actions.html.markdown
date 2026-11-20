---
subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : Custom Resolver Cache Clear Actions"
description: |-
  Retrieves active or recent cache clear action records of a private DNS custom resolver.
---

# ibm_dns_custom_resolver_cache_clear_actions

Provides a read-only data source for listing cache clear action records of a private DNS custom resolver. For more information, refer to [list-clear-cache-actions](https://cloud.ibm.com/apidocs/dns-svcs#list-clear-cache-actions).

## Example usage

```terraform
data "ibm_dns_custom_resolver_cache_clear_actions" "test" {
  instance_id = "2be5d4a7-78f0-4c62-a957-41dc15342777"
  resolver_id = "ddbe7a53-7971-46dc-b021-420335c31562"
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `instance_id` - (Required, String) The unique identifier of a private DNS service instance.
- `resolver_id` - (Required, String) The unique identifier of a custom resolver.

## Attribute reference

In addition to the argument references list, you can access the following attribute references after your data source is created.

- `id` - (String) The composite identifier in the format `instance_id/resolver_id`.
- `total_count` - (Integer) The total number of cache clear actions returned.
- `actions` - (List) List of cache clear action records.

  Nested schema for `actions`:
  - `action_id` - (String) The unique identifier of the cache clear action.
  - `status` - (String) The status of the cache clear action (`pending`, `success`, `failed`).
  - `created_on` - (String) The timestamp when the cache clear action was created.
  - `modified_on` - (String) The timestamp when the cache clear action was modified.
  - `expires_on` - (String) The timestamp when the cache clear action expires.
