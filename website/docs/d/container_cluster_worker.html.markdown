---
subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: ibm_container_cluster_worker"
description: |-
  Get information about a worker node that is attached to a Kubernetes cluster on IBM Cloud.
---

# ibm_container_cluster_worker
Retrieve information about the worker nodes of your IBM Cloud Kubernetes Service cluster. For more information, about cluster worker, see [updating clusters, worker nodes, and cluster components](https://cloud.ibm.com/docs/containers?topic=containers-update).


## Example usage

```terraform
data "ibm_container_cluster_worker" "cluster_foo" {
  worker_id = "dev-mex10-pa70c4414695c041518603bfd0cd6e333a-w1"
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `resource_group_id` - (Optional, String) The ID of the resource group where your cluster is provisioned into. To find the resource group, run `ibmcloud resource groups` or use the `ibm_resource_group` data source. If this parameter is not provided, the `default` resource group is used.
- `worker_id` - (Required, String) The ID of the worker node for which you want to retrieve information. To find the ID, run `ibmcloud ks worker ls cluster <cluster_name_or_ID>`. 

**Deprecated reference**

- `account_guid` - (Deprecated, string) The GUID for the IBM Cloud account that the cluster is associated with. You can retrieve the value from the `ibm_account` data source or by running the `ibmcloud iam accounts` command in the IBM Cloud CLI.
- `org_guid` - (Deprecated, string) The GUID for the IBM Cloud organization that the cluster is associated with. You can retrieve the value from the `ibm_org` data source or by running the `ibmcloud iam orgs --guid` command in the [IBM Cloud CLI](https://cloud.ibm.com/docs/cli?topic=cloud-cli-getting-started).
- `region` - (Deprecated, string) The region where the worker is provisioned. If the region is not specified it will be defaulted to provider region(IC_REGION/IBMCLOUD_REGION). To get the list of supported regions please access this [link](https://containers.bluemix.net/v1/regions) and use the alias.
- `space_guid` - (Deprecated, string) The GUID for the IBM Cloud space that the cluster is associated with. You can retrieve the value from the `ibm_space` data source or by running the `ibmcloud iam space <space-name> --guid` command in the IBM Cloud CLI.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `private_ip` - (String) The private IP address that is assigned to the worker node.
- `private_vlan` - (String) The ID of the private VLAN that the worker node is attached to.
- `public_ip` - (String) The public IP address that is assigned to the worker node. 
- `public_vlan` - (String) The ID of the public VLAN that the worker node is attached to.
- `state` - (String) The state of the worker node. 
- `status` - (String) The status of the worker node.
