---
subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : Custom Resolvers"
description: |-
  Manages IBM Cloud Infrastructure Private DNS Custom Resolvers.
---

# ibm_dns_custom_resolvers

Provides a read-only data source for Custom Resolvers. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example usage

```terraform
data "ibm_dns_custom_resolvers" "example-cr" {
  instance_id = "resource_instance_guid"
}
```

## Argument reference
The following arguments are supported: 

- `instance_id` - (Required, String) The GUID of the private DNS service instance.

## Attribute reference
In addition to the argument references list, you can access the following attribute references after your data source is created. 

- `id`- (String) The unique identifier of the custom resolvers.
- `custom_resolvers` (List) List of all private DNS custom resolvers.
 
   Nested scheme for `custom_resolvers`:
   - `custom_resolver_id` - (String) Identifier of the  custom resolver.
   - `name` - (String) Name of the  custom resolver.
   - `description` - (String) Descriptive text of the custom resolver.
   - `enabled` - (String) Descriptive text of the custom resolver.
   - `health`- (String) The status of DNS Custom Resolver's health. Possible values are `DOWN`, `CRITICAL`, `HEALTHY`.
  
    Nested scheme for `locations`:
    - `healthy`- (String) The health status.
    - `dns_server_ip`- (String) The dns server ip.
    - `enabled`- (String) Whether the location is enabled.
    - `location_id`- (String) The location id.
    - `subnet_crn` - (String) The subnet crn
 	 
