---
subcategory: "Satellite"
layout: "ibm"
page_title: "IBM : satellite_location"
description: |-
  Manages IBM Cloud Satellite location.
---

# ibm_satellite_location

Create, update, or delete [IBM Cloud Satellite Location](https://cloud.ibm.com/docs/satellite?topic=satellite-locations). Set up an IBM Cloud Satellite location to represent a data center that you fill with your own infrastructure resources, and start running IBM Cloud services on your own infrastructure.
Create, update, or delete [IBM Cloud Satellite Host](https://cloud.ibm.com/docs/satellite?topic=satellite-locations). Set up an IBM Cloud Satellite location to represent a data center in your infrastructure resources, and start running IBM Cloud services on your own infrastructure.


## Example usage

### Sample to create location

```terraform
data "ibm_resource_group" "group" {
    name = "Default"
}

resource "ibm_satellite_location" "create_location" {
  location          = var.location
  zones             = var.location_zones
  managed_from      = var.managed_from
  resource_group_id = data.ibm_resource_group.group.id
}

```

###  Sample to create location by using COS bucket

```terraform
resource "ibm_satellite_location" "create_location" {
  location      = var.location
  zones         = var.location_zones
  managed_from  = var.managed_from  

  cos_config {
    bucket  = var.cos_bucket
  }
}
```

## Timeouts

The `ibm_satellite_location` provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 30 minutes) Used for creating Instance.
- **update** - (Default 10 minutes) Used for updating Instance.
- **delete** - (Default 60 minutes) Used for deleting Instance.


## Argument reference
Review the argument references that you can specify for your resource. 

- `cos_config` - (Optional, List) The IBM Cloud Object Storage bucket configuration details. Nested cos_config blocks have the following structure.

  Nested scheme for `cos_config`:
  - `bucket`- (Optional, String) The name of the IBM Cloud Object Storage bucket that you want to use to back up the control plane data.
	- `endpoint` - (Optional, String) The IBM Cloud Object Storage bucket endpoint.
  - `region`- (Optional, String) The name of a region, such as `us-south` or `eu-gb`.
- `cos_credentials`- (Optional, List) The IBM Cloud Object Storage authorization keys. Nested `cos_credentials` blocks have the following structure.

  Nested scheme for `cos_credentials`:
  - `access_key-id` - (Required, String)The `HMAC` secret access key ID.
  - `secret_access_key`-  (Optional, String) The `HMAC` secret access key.
- `description` - (Optional, String)  A description of the new Satellite location.
- `is_location_exist`- (Optional, Bool) Determines the location has to be created or not.
- `location` - (Required, String) The name of the location to be created or pass existing location name.
- `logging_account_id` - (Optional, String) The account ID for IBM Log Analysis with LogDNA log forwarding.
- `managed_from` - (Required, String) The IBM Cloud regions that you can choose from to manage your Satellite location. To list available multizone regions, run `ibmcloud ks locations`. For more information, refer to [supported IBM Cloud locations](https://cloud.ibm.com/docs/satellite?topic=satellite-sat-regions).
- `zones`- Array of Strings - Optional- The names for the host zones. For high availability, allocate your hosts across these three zones based on your infrastructure provider zones. For example, `us-east-1`, `us-east-2`, `us-east-3` .


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `crn` - (String) The CRN for this satellite location.
- `created_on` - (Timestamp) The created time of the satellite location.
- `id` - (String) The unique identifier of the location.
- `ingress_hostname` - (String) The Ingress hostname.
- `ingress_secret` - (String) The Ingress secret.
- `host_attached_count` - (Timestamp) The total number of hosts that are attached to the Satellite location.
- `host_available_count` - (Timestamp) The available number of hosts that can be assigned to a cluster resource in the Satellite location.
- `resource_group_name` - (String) The name of the resource group.

## Import

The `ibm_satellite_location` resource can be imported by using the location ID or name.

**Syntax**

```
$ terraform import ibm_satellite_location.location location
```

**Example**

```
$ terraform import ibm_satellite_location.location satellite-location
```
