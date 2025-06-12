---
subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: ibm_container_cluster_versions"
description: |-
  List supported kubernetes versions on IBM Cloud.
---

# ibm_container_cluster_versions

Retrieve information about supported Kubernetes versions in IBM Cloud Kubernetes Service clusters. To find a list of supported Kubernetes versions, see the [IBM Cloud Kubernetes Service documentation](https://cloud.ibm.com/docs/containers?topic=containers-cs_versions)


## Example usage
The following example shows how to retrieve information about supported Kubernetes versions for the resource group `11222333111abc111`.

```terraform
data "ibm_container_cluster_versions" "cluster_versions" {
  resource_group_id          = "11222333111abc111"
}

data "ibm_container_cluster_versions" "cluster_versions" {
  region = "eu-de"
}
```
## Argument reference
Review the argument references that you can specify for your data source. 

- `resource_group_id` - (Optional, String) The ID of the resource group where your cluster is provisioned into. To find the resource group, run `ibmcloud resource groups` or use the `ibm_resource_group` data source. If this parameter is not provided, the `default` resource group is used.

**Deprecated reference**

- `account_guid` - (Deprecated, String) The GUID for the IBM Cloud account associated with the cluster. You can retrieve the value from the `ibm_account` data source or by running the `ibmcloud iam accounts` command in the IBM Cloud CLI.
- `org_guid` - (Deprecated, String) The GUID for the IBM Cloud organization associated with the cluster. You can retrieve the value from the `ibm_org` data source or by running the `ibmcloud iam orgs --guid` command in the [IBM Cloud CLI](https://cloud.ibm.com/docs/cli?topic=cloud-cli-getting-started).
- `region` - (Deprecated, String) The region to target. If the region is not specified it will be defaulted to provider region(IC_REGION/IBMCLOUD_REGION). To get the list of supported regions please access this [link](https://containers.bluemix.net/v1/regions) and use the alias.
- `space_guid` - (Deprecated, String) The GUID for the IBM Cloud space associated with the cluster. You can retrieve the value from the `ibm_space` data source or by running the `ibmcloud iam space <space-name> --guid` command in the IBM Cloud CLI.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `id` - (String) The unique identifier of the cluster. 
- `valid_kube_versions` - (String) The supported Kubernetes version in IBM Cloud Kubernetes Service clusters. 
- `valid_openshift_versions` - (String) The supported OpenShift Container Platform version in Red Hat OpenShift on IBM Cloud clusters.
- `default_kube_version` - (String) The default Kubernetes version in IBM Cloud Kubernetes Service clusters. 
- `default_openshift_version` - (String) The default OpenShift Container Platform version in Red Hat OpenShift on IBM Cloud clusters.
