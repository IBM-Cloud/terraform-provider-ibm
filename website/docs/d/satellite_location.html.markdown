---
subcategory: "Satellite"
layout: "ibm"
page_title: "IBM : satellite_location"
description: |-
  Get information about an IBM Cloud satellite location.
---

# ibm\_satellite_location

Import the details of an existing satellite location as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.


## Example Usage

```hcl
data "ibm_satellite_location" "location" {
  location  = var.location
}
```

## Argument Reference

The following arguments are supported:

* `location` - (Required, string) The name of the location to be created or pass existing location name.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id`  - The unique identifier of the location.
* `crn` - The CRN for this satellite location.
* `managed_from` - (Required, string) The IBM Cloud metro from which the Satellite location is managed. To list available multizone regions, run 'ibmcloud ks locations'. such as 'wdc04', 'wdc06' or 'lon04'.
* `description` - Description of the new Satellite location.
* `logging_account_id` -  The account ID for IBM Log Analysis with LogDNA log forwarding.
* `zones` - The names for the host zones. For high availability, allocate your hosts across these three zones based on your infrastructure provider zones. ex: [ us-east-1, us-east-2, us-east-3 ]
* `resource_group_id` - The ID of the resource group.
* `resource_group_name` - The name of the resource group.
* `host_attached_count` - The total number of hosts that are attached to the Satellite location.
* `host_available_count` - The available number of hosts that can be assigned to a cluster resource in the Satellite location.
* `created_on` - The created time of the satellite location.
* `ingress_hostname` - The Ingress hostname.
* `ingress_secret` - The Ingress secret.
* `tags` - List of tags associated with resource instance.

