---
subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: ibm_container_vni_attachment"
description: |-
  Get information about a VNI attachment to a bare metal worker node.
---

# ibm_container_vni_attachment
Retrieve information about a Virtual Network Interface (VNI) attachment to a bare metal worker node in a Red Hat OpenShift on IBM Cloud cluster. This feature is currently supported for OpenShift clusters only.

## Example usage

```terraform
data "ibm_container_vni_attachment" "attachment" {
  cluster = "mycluster"
  worker  = "kube-c4u8l44d0hf4s8k25u90-mycluster-bm-000001"
  vni_id  = "r006-1234abcd-5678-90ef-1234-567890abcdef"
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `cluster` - (Required, String) The cluster ID or name where the worker is located.
- `worker` - (Required, String) The worker ID where the VNI is attached. To find the worker ID, run `ibmcloud ks worker ls --cluster <cluster_name_or_ID>`.
- `vni_id` - (Required, String) The ID of the VNI to retrieve information about.
- `resource_group_id` - (Optional, String) The ID of the resource group where the cluster is provisioned. To find the resource group, run `ibmcloud resource groups` or use the `ibm_resource_group` data source.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `id` - (String) The unique identifier of the VNI attachment (same as `vni_id`).
- `vlan_id` - (Integer) The VLAN ID for the bare metal worker.
- `status` - (String) The status of the attachment.
- `created_at` - (String) The timestamp when the attachment was created.
- `auto_delete` - (Bool) Whether the VNI will be deleted when the attachment is destroyed.
- `primary_ip_address` - (String) The primary IP address of the VNI.
- `mac_address` - (String) The MAC address of the VNI.
- `vni_name` - (String) The name of the VNI.
