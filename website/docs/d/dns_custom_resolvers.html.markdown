---
subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : Custom Resolvers"
description: |-
  Manages IBM Cloud Infrastructure private DNS custom resolvers.
---

# ibm_dns_custom_resolvers

Provides a read-only data source for custom resolvers. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information about custom resolver, see [List custom resolvers](https://cloud.ibm.com/apidocs/dns-svcs#list-custom-resolvers).

## Example usage

```terraform
data "ibm_dns_custom_resolvers" "test-cr" {
  instance_id = ibm_dns_custom_resolver.test.instance_id
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `instance_id` - (Required, String) The GUID of the private DNS service instance.

## Attribute reference

In addition to the argument references list, you can access the following attribute references after your data source is created.

- `custom_resolvers` (List) List of all private DNS custom resolvers.

   Nested scheme for `custom_resolvers`:
  - `custom_resolver_id` - (String) Identifier of the custom resolver.
  - `description` - (String) Descriptive text of the custom resolver.
  - `enabled` - (String) Whether the custom resolver is enabled or disabled.
  - `health`- (String) The status of DNS custom resolver's health. Supported values are `CRITICAL`, `DEGRADED`, `HEALTHY`.
  - `name` - (String) Name of the custom resolver.
  - `profile` - (String) The profile name of the custom resolver. Supported values are `ESSENTIAL`, `ADVANCED`, `PREMIER`.
  - `allow_disruptive_updates` - (Boolean) Whether a disruptive update is allowed for the custom resolver.
  - `locations` (List) The list of locations within the custom resolver.
  
    Nested scheme for `locations`:
    - `dns_server_ip`- (String) The DNS server IP.
    - `enabled`- (String) Whether the location is enabled or disabled.
    - `healthy`- (String) The health status.
    - `location_id`- (String) The location ID.
    - `subnet_crn` - (String) The subnet CRN.
