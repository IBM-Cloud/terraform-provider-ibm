---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_flow_log"
description: |-
  Manages IBM VPC flow log.
---

# ibm_is_flow_log
Create, update, delete and suspend the flow log resource. For more information, about VPC flow log, see [creating a flow log collector](https://cloud.ibm.com/docs/vpc?topic=vpc-ordering-flow-log-collector).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```


## Example usage

```terraform

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


## Argument reference
Review the argument references that you can specify for your resource. 

- `name` - (Required, String) The unique user-defined name for the flow log collector.No.
- `target` - (Required, Forces new resource, String) The ID of the target to collect flow logs. If the target is an instance, subnet, or VPC, flow logs is not collected for any network interfaces within the target that are more specific flow log collector.
- `storage_bucket` - (Required, Forces new resource, String) The name of the IBM Cloud Object Storage bucket where the collected flows will be logged. The bucket must exist and an IAM service authorization must grant IBM Cloud flow logs resources of VPC infrastructure services writer access to the bucket.
- `active` - (Optional, String) Indicates whether the collector is active. If **false**, this collector is created in inactive mode. Default value is true.
- `resource_group` - (Optional, Forces new resource, String) The resource group ID where the flow log is created.
- `tags` - (Optional, Array of Strings) The tags associated with the flow log.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `created_at`-  (String) The date and time that the flow log collector created.
- `crn` - (String) The CRN of the flow log collector.
- `href` - (String) The URL of the flow log collector.
- `id` - (String) The unique identifier of the flow log collector.
- `lifecycle_state` - (String) The lifecycle state of the flow log collector.
- `name`-  (String) The user-defined name of the flow log collector.
- `vpc` - (String) The VPC of the flow log collector that is associated.


## Import
The `ibm_is_flow_log` resource can be imported by using VPC flow log ID.

**Example**

```
$ terraform import ibm_is_flow_log.example d7bec597-4726-451f-8a53-e62e6f19c32c
```
