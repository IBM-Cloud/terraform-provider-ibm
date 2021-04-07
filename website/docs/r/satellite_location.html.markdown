---
subcategory: "Satellite"
layout: "ibm"
page_title: "IBM : satellite_location"
description: |-
  Manages IBM Cloud satellite location.
---

# ibm\_satellite_location

Create, update, or delete [IBM Cloud Satellite Location](https://cloud.ibm.com/docs/satellite?topic=satellite-locations). Set up an IBM Cloud™ Satellite location to represent a data center that you fill with your own infrastructure resources, and start running IBM Cloud services on your own infrastructure.


## Example Usage

###  Create location

```hcl
resource "ibm_satellite_location" "create_location" {
  location      = var.location
  zones         = var.location_zones
  managed_from  = var.managed_from
}

```

###  Create location using COS bucket

```hcl
resource "ibm_satellite_location" "create_location" {
  location      = var.location
  zones         = var.location_zones
  managed_from  = var.managed_from  

  cos_config {
    bucket  = var.cos_bucket
    region  = var.ibm_region
  }
}
```

## Argument Reference

The following arguments are supported:

* `location` - (Required, string) The name of the location to be created or pass existing location name.
* `is_location_exist` - (Optional, bool) Determines if the location has to be created or not.
* `managed_from` - (Required, string) The IBM Cloud metro from which the Satellite location is managed. To list available multizone regions, run 'ibmcloud ks locations'. such as 'wdc04', 'wdc06' or 'lon04'.
* `description` - (Optional, string) A description of the new Satellite location.
* `logging_account_id` - (Optional, string) The account ID for IBM Log Analysis with LogDNA log forwarding.
* `cos_config` - (Optional, list) IBM Cloud Object Storage bucket configuration details. Nested `cos_config` blocks have the following structure:
    * `bucket` - The name of the IBM Cloud Object Storage bucket that you want to use to back up the control plane data.
    * `endpoint` - COS bucket endpoint.
    * `region` - Name of region, such as 'us-south' or 'eu-gb'.
* `cos_credentials` - (Optional, list) IBM Cloud Object Storage authorization keys. Nested `cos_credentials` blocks have the following structure:
    * `access_key-id` - The HMAC secret access key ID.
    * `secret_access_key` - The HMAC secret access key. 
* `zones` - (Optional, array of strings) The names for the host zones. For high availability, allocate your hosts across these three zones based on your infrastructure provider zones. ex: [ us-east-1, us-east-2, us-east-3 ]

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the location.
* `crn` - The CRN for this satellite location.

## Import

`ibm_satellite_location` can be imported using the location id or name.

Example:

```
$ terraform import ibm_satellite_location.location location

$ terraform import ibm_satellite_location.location satellite-location

```
