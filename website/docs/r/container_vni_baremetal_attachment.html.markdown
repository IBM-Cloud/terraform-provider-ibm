---

subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: container_vni_baremetal_attachment"
description: |-
  Manages IBM container VNI bare metal attachment.
---

# ibm_container_vni_baremetal_attachment
Attach a Virtual Network Interface (VNI) to a bare metal worker node in a Red Hat OpenShift on IBM Cloud cluster. VNIs provide network connectivity for bare metal workers. This feature is currently supported for OpenShift clusters only.

## Example usage

### Attach VNI to a specific worker

```terraform
resource "ibm_container_vni_baremetal_attachment" "attachment" {
  vni_id  = "r006-1234abcd-5678-90ef-1234-567890abcdef"
  vlan_id = 100
  worker  = "kube-c4u8l44d0hf4s8k25u90-mycluster-bm-000001"
  auto_delete = false
}
```

### Attach VNI to any available worker in a cluster

```terraform
resource "ibm_container_vni_baremetal_attachment" "attachment" {
  vni_id  = "r006-1234abcd-5678-90ef-1234-567890abcdef"
  vlan_id = 100
  cluster = "mycluster"
  auto_delete = false
}
```

## Timeouts

The `ibm_container_vni_baremetal_attachment` provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 10 minutes) Used for attaching VNI.
- **delete** - (Default 10 minutes) Used for detaching VNI.

## Argument reference
Review the argument references that you can specify for your resource.

- `vni_id` - (Required, Forces new resource, String) The ID of the VNI to attach to the bare metal worker.
- `vlan_id` - (Required, Forces new resource, Integer) The VLAN ID for the bare metal worker. Valid range is 1-500.
- `cluster` - (Optional, Forces new resource, String) The cluster ID or name to attach the VNI to any available worker. Exactly one of `cluster` or `worker` must be specified.
- `worker` - (Optional, Forces new resource, String) The worker ID to attach the VNI to a specific worker. Exactly one of `cluster` or `worker` must be specified.
- `auto_delete` - (Optional, Forces new resource, Bool) Whether to delete the VNI when the attachment is destroyed. Default is `false`.
- `resource_group_id` - (Optional, Forces new resource, String) The ID of the resource group where the cluster is provisioned. To find the resource group, run `ibmcloud resource groups` or use the `ibm_resource_group` data source.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the VNI attachment (same as `vni_id`).
- `worker_id` - (String) The ID of the worker where the VNI is attached.
- `status` - (String) The status of the attachment.
- `created_at` - (String) The timestamp when the attachment was created.

## Import
The `ibm_container_vni_baremetal_attachment` resource can be imported by using the VNI ID.

```
terraform import ibm_container_vni_baremetal_attachment.attachment r006-1234abcd-5678-90ef-1234-567890abcdef
```
