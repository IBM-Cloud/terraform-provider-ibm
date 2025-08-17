---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_tenant"
description: |-
  Manages a tenant in the IBM Power Virtual Server cloud.
---

# ibm_pi_tenant

Retrieve information about the tenants that are configured for your Power Systems Virtual Server instance. For more information, about power virtual server tenants, see [network security](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-network-security).

## Example Usage

The following example retrieves all tenants for the Power Systems Virtual Server instance with the ID.

```terraform
data "ibm_pi_tenant" "ds_tenant" {
  pi_cloud_instance_id = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
}
```

### Notes

- Please find [supported Regions](https://cloud.ibm.com/apidocs/power-cloud#endpoint) for endpoints.
- If a Power cloud instance is provisioned at `lon04`, The provider level attributes should be as follows:
  - `region` - `lon`
  - `zone` - `lon04`

Example usage:

  ```terraform
    provider "ibm" {
      region    =   "lon"
      zone      =   "lon04"
    }
  ```
  
## Argument Reference

Review the argument references that you can specify for your data source.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `cloud_instances` - (Set) Set of regions and Power Systems Virtual Server instance IDs that the tenant owns.

  Nested scheme for `cloud_instances`:
  - `cloud_instance_id` - (String) The unique identifier of the cloud instance.
  - `region` - (String) The region of the cloud instance.
- `creation_date` - (String) Date of tenant creation.
- `enabled` - (Boolean) Indicates if the tenant is enabled for the Power Systems Virtual Server instance ID.
- `id` - (String) The ID of the tenant.
- `tenant_name` -  (String) The name of the tenant.
