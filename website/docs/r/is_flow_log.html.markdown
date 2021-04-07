---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_flow_log"
description: |-
  Manages IBM VPC Flow Log.
---

# ibm\_is_flow_log

Provides a flow log resource. This allows flow log to be created, updated, deleted and suspended.


## Example Usage

```hcl

resource "ibm_is_instance" "testacc_instance" {
  name    = "testinstance"
  image   = "7eb4e35b-4257-56f8-d7da-326d85452591"
  profile = "b-2x8"

  primary_network_interface {
    port_speed = "1000"
    subnet     = "70be8eae-134c-436e-a86e-04849f84cb34"
  }

  vpc  = "01eda778-b822-43a2-816d-d30713df5e13"
  zone = "us-south-1"
  keys = ["eac87f33-0c00-4da7-aa66-dc2d972148bd"]
}


data "ibm_resource_group" "instance_group" {
  name = var.resource_group
}

resource "ibm_resource_instance" "instance1" {
  name              = "cos-instance"
  resource_group_id = data.ibm_resource_group.instance_group.id
  service           = "cloud-object-storage"
  plan              = "standard"
  location          = "global"
}

resource "ibm_cos_bucket" "bucket1" {
   bucket_name          = "us-south-bucket-vpc1"
   resource_instance_id = ibm_resource_instance.instance1.id
   region_location = var.region
   storage_class = "standard"
}

resource ibm_is_flow_log test_flowlog {
  depends_on = [ibm_cos_bucket.bucket1]
  name = "test-instance-flow-log"
  target = ibm_is_instance.testacc_instance.id
  active = true
  storage_bucket = ibm_cos_bucket.bucket1.bucket_name
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The unique user-defined name for this flow log collector.
* `target` - (Required, Forces new resource, string) The id of the target this collector is to collect flow logs for. If the target is an instance, subnet, or VPC, flow logs will not be collected for any network interfaces within the target that are themselves the target of a more specific flow log collector.
* `storage_bucket` - (Required, Forces new resource, string) The name of the Cloud Object Storage bucket where the collected flows will be logged. The bucket must exist and an IAM service authorization must grant IBM Cloud Flow Logs resources of VPC Infrastructure Services writer access to the bucket.
* `active` - (Optional, string) Indicates whether this collector is active. If false, this collector is created in inactive mode. Default is true. 
* `resource_group` - (Optional, Forces new resource, string) The resource group ID where the flow log is to be created.
* `tags` - (Optional, array of strings) Tags associated with the Flow log.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `created_at` - The date and time that the flow log collector was created. 
* `crn` - The CRN for this flow log collector.
* `href` - The URL for this flow log collector.
* `id` - The unique identifier for this flow log collector.
* `lifecycle_state` - The lifecycle state of the flow log collector.
* `name` - The user-defined name for this flow log collector.
* `vpc` - The VPC this flow log collector is associated with.

## Import

ibm_is_flow_log can be imported using VPC Flow log ID, eg

```
$ terraform import ibm_is_flow_log.example d7bec597-4726-451f-8a53-e62e6f19c32c
```
 
