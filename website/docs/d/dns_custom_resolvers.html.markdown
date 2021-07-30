---
subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : Custom Resolvers"
description: |-
  Manages IBM Cloud Infrastructure Private DNS Custom Resolvers.
---

# ibm_dns_custom_resolvers

Provides a read-only data source for Custom Resolvers. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.For more information about Custom Resolver, refer to [List custom resolvers](https://cloud.ibm.com/apidocs/dns-svcs#list-custom-resolvers).


## Example usage

```terraform
data "ibm_dns_custom_resolvers" "test-custom-resolver" {
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
   - `custom_resolver_id` - (String) Identifier of the  custom resolver.
   - `description` - (String) Descriptive text of the custom resolver.
   - `enabled` - (String) Whether custom resolver is enabled.
   - `health`- (String) The status of DNS Custom Resolver's health. Possible values are `CRITICAL`, `DEGRADED`, `HEALTHY`.
   - `name` - (String) Name of the  custom resolver.
   - `locations` (List) The list of locations within the custom resolver. 
    
      Nested scheme for `locations`:
       - `dns_server_ip`- (String) The dns server ip.
       - `enabled`- (String) Whether the location is enabled.
       - `healthy`- (String) The health status.
       - `location_id`- (String) The location id.
       - `subnet_crn` - (String) The subnet crn
 	 
