---
subcategory: "Satellite"
layout: "ibm"
page_title: "IBM : satellite_location"
description: |-
  Get information about an IBM Cloud Satellite location.
---

# ibm_satellite_location
Retrieve information of an existing Satellite location. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax. For more information, about IBM Cloud regions for Satellite see [Satellite regions](https://cloud.ibm.com/docs/satellite?topic=satellite-sat-regions).


## Example usage

```terraform
data "ibm_satellite_location" "location" {
  location  = var.location
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `location` - (Required, String) The name or ID of the Satellite location to  be created or pass existing location.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.
- `crn` - (String) The CRN for this satellite location.
- `created_on` - (Timestamp) The created time of the satellite location.
- `description` - (String) Description of the new Satellite location.
- `id` - (String) The unique identifier of the location.
- `ingress_hostname` - (String) The Ingress hostname.
- `ingress_secret` - (String) The Ingress secret.
- `host_attached_count` - (Integer) The total number of hosts that are attached to the Satellite location.
- `host_available_count` - (Integer) The available number of hosts that can be assigned to a cluster resource in the Satellite location.
- `logging_account_id` - (String) The account ID for IBM Cloud Log Analysis with IBM Cloud Log Analysis log forwarding.
- `managed_from` - (String) The IBM Cloud metro from which the Satellite location is managed. To list available multizone regions, run `ibmcloud ks locations`. such as `wdc04`, `wdc06`, or `lon04`.
- `resource_group_id` - (String) The ID of the resource group.
- `resource_group_name` - (String) The name of the resource group.
- `tags` - (String) List of tags associated with resource instance.
- `zones` - (String) The names for the host zones. For high availability, allocate your hosts across these three zones based on your infrastructure provider zones. For example, `us-east-1`, `us-east-2`, `us-east-3`.

