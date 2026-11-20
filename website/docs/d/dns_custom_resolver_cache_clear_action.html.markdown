---
subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : Custom Resolver Cache Clear Action"
description: |-
  Retrieves status and details of a private DNS custom resolver cache clear action.
---

# ibm_dns_custom_resolver_cache_clear_action

Provides a read-only data source for a private DNS custom resolver cache clear action. You can reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, refer to [get-clear-cache-action](https://cloud.ibm.com/apidocs/dns-svcs#get-clear-cache-action).

## Example usage

```terraform
data "ibm_dns_custom_resolver_cache_clear_action" "test" {
  instance_id = "2be5d4a7-78f0-4c62-a957-41dc15342777"
  resolver_id = "ddbe7a53-7971-46dc-b021-420335c31562"
  action_id   = "1717236000000000"
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `instance_id` - (Required, String) The unique identifier of a private DNS service instance.
- `resolver_id` - (Required, String) The unique identifier of a custom resolver.
- `action_id` - (Required, String) The unique identifier of a cache clear action.

## Attribute reference

In addition to the argument references list, you can access the following attribute references after your data source is created.

- `id` - (String) The composite identifier in the format `instance_id/resolver_id/action_id`.
- `status` - (String) The status of the cache clear action (`pending`, `success`, `failed`).
- `created_on` - (String) The timestamp when the cache clear action was created.
- `modified_on` - (String) The timestamp when the cache clear action was modified.
- `expires_on` - (String) The timestamp when the cache clear action expires.
- `errors` - (List) List of error details if the cache clear action failed.

  Nested schema for `errors`:
  - `code` - (String) The error code identifier.
  - `message` - (String) The error message text.
  - `location_id` - (String) The unique identifier of the location that failed to clear its cache.
