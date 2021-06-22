---

subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: container_cluster_feature"
description: |-
  Manages IBM container cluster feature.
---

# ibm_container_cluster_feature

Enable or disable a feature in your IBM Cloud Kubernetes Service cluster. For more information, about IBM container cluster feature, see [Security for IBM Cloud Kubernetes Service](https://cloud.ibm.com/docs/containers?topic=containers-security). Supported features include: 
- Public service endpoint
- Private service endpoint

## Example usage
The following example enables the private service endpoint feature for a cluster that is named `mycluster`.

```terraform
resource "ibm_container_cluster_feature" "feature" {
  cluster                  = "test1"
  private_service_endpoint = "true"
}

```

## Timeouts

The `ibm_container_cluster_feature` provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create**: The enablement of the feature is considered `failed` if no response is received for 90 minutes.
- **update**: The update of the feature is considered `failed` if no response is received for 90 minutes. 


## Argument reference
Review the argument references that you can specify for your resource. 
 
- `cluster` - (Required, Forces new resource, String) The name or ID of the cluster for which you want to enable or disabled a feature. To find the name or ID, use the `ibmcloud ks cluster ls` command.
- `public_service_endpoint` - (Optional, Bool) Enable(**true**) or disable (**false**) the public service endpoint for your cluster. You can use the public service endpoint to access the Kubernetes master from the public network. To use service endpoints, your account must be enabled for [Virtual Routing and Forwarding (VRF)](https://cloud.ibm.com/docs/account?topic=account-vrf-service-endpoint#vrf). For more information, see [Worker-to-master and user-to-master communication: Service endpoints](https://cloud.ibm.com/docs/containers?topic=containers-plan_clusters#workeruser-master).
- `private_service_endpoint` - (Optional, Bool) Enable (**true**) or disable (**false**) the private service endpoint for your cluster. When the private service endpoint is enabled, communication between the Kubernetes and the worker nodes is established over the private network. If you enable the private service endpoint, you cannot disable it later. To use service endpoints, your account must be enabled for [Virtual Routing and Forwarding (VRF)](https://cloud.ibm.com/docs/account?topic=account-vrf-service-endpoint#vrf). For more information, see [Worker-to-master and user-to-master communication: Service endpoints](https://cloud.ibm.com/docs/containers?topic=containers-plan_clusters#workeruser-master).
- `refresh_api_servers` - (Optional, Bool) If set to **true**, the Kubernetes master of the cluster is refreshed to apply the changes of your feature. If set to **false**, no refresh of the Kubernetes master is performed.
- `reload_workers` -  (Optional, Bool) If set to **true**, your worker nodes are reloaded after the feature is enabled. If set to **false**, no reload of the worker nodes is performed.
- `resource_group_id` - (Optional, String) The ID of the resource group that your cluster belongs to. You can retrieve the resource group by using the `ibm_resource_group` data source.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The ID of the cluster feature. 
- `public_service_endpoint_url` - (String) The URL to the public service endpoint of your cluster. 
- `private_service_endpoint_url` - (String) The URL to the private service endpoint of your cluster. 
