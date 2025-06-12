---

subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: container_bind_service"
description: |-
  Manages IBM container cluster.
---

# ibm_container_bind_service

> [!CAUTION]
> This resource will be deprecated, please check [this guide](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/guides/binding-services) on how to bind services.

Bind an IBM Cloud service to an IBM Cloud Kubernetes Service cluster. Service binding is a quick way to create service credentials for an IBM Cloud service by using its public service endpoint and storing these credentials in a Kubernetes secret in your cluster. The Kubernetes secret is automatically encrypted in etcd to protect your data.

To bind a service to your cluster, you must provision an instance of the service first. For more information, about service binding, see [Adding services by using IBM Cloud service binding](https://cloud.ibm.com/docs/containers?topic=containers-service-binding).

## Example usage
In the following example, you can bind a service to a cluster:

```terraform
resource "ibm_container_bind_service" "bind_service" {
  cluster_name_id       = "cluster_name"
  service_instance_name = "service_name"
  namespace_id          = "default"
}
```


## Argument reference
Review the argument references that you can specify for your resource. 
  
- `cluster_name_id` - (Required, Forces new resource, String) The name or ID of the cluster to which you want to bind an IBM Cloud service. To find the cluster name or ID, run `ibmcloud ks cluster ls`.
- `key` - (Optional, Forces new resource, String) The name or guid of existing service credentials that you want to use for the service. If you do not provide this option, service credentials are automatically created as part of the service binding process.
- `namespace_id` - (Required, Forces new resource, String) The Kubernetes namespace where you want to create the Kubernetes secret that holds the service credentials of the service that you want to bind to the cluster.
- `resource_group_id` - (Optional, Forces new resource, String) The ID of the resource group where your IBM Cloud service is provisioned into. To list resource groups, run `ibmcloud resource groups`.
- `role` - (Optional, Forces new resource, String) The IAM service access role that you want to use to create the service credentials for the IBM Cloud service instance. If you specified existing service credentials in the `key` parameter, settings for the `role` parameter are ignored.
- `service_instance_id` - (Optional, Forces new resource, String) The ID of the service that you want to bind to the cluster. If you specify this parameter, do not specify `service_instance_name` at the same time.
- `service_instance_name` - (Optional, Forces new resource, String) The name of the service that you want to bind to the cluster. If you specify this parameter, do not specify `service_instance_id` at the same time.
- `tags` - (Optional, Array of string)  A list of tags that you want to associate with the IBM Cloud service instance that you bind to the cluster. **Note** `Tags` are managed locally and are not stored on the IBM Cloud service endpoint.

**Deprecated reference**

- `account_guid` - (Deprecated, Forces new resource, string) The GUID for the IBM Cloud account associated with the cluster. You can retrieve the value from data source `ibm_account` or by running the `ibmcloud iam accounts` command in the IBM Cloud CLI.
- `org_guid` - (Deprecated, Forces new resource, string) The GUID for the IBM Cloud organization associated with the cluster. You can retrieve the value from data source `ibm_org` or by running the `ibmcloud iam orgs --guid` command in the [IBM Cloud CLI](https://cloud.ibm.com/docs/cli?topic=cloud-cli-getting-started).
- `region` - (Deprecated, Forces new resource, string) The region where the cluster is provisioned. If the region is not specified it will be defaulted to provider region(IC_REGION/IBMCLOUD_REGION). To get the list of supported regions please access this [link](https://containers.bluemix.net/v1/regions) and use the alias.
-  `space_guid` - (Deprecated, Forces new resource, string) The GUID for the IBM Cloud space associated with the cluster. You can retrieve the value from data source `ibm_space` or by running the `ibmcloud iam space <space-name> --guid` command in the IBM Cloud CLI.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the service bind resource in your cluster in the format `<cluster_name_ID>/<service_instance_name>` or `<service_instance_id>/<namespace_id>`.
- `namespace_id` - (String) The namespace in your cluster where the Kubernetes secret is located that holds the credentials to access your IBM Cloud service instance.
- `secret_name` - (String) The name of the Kubernetes secret that holds the credentials to access your IBM Cloud service instance.

## Import
The `ibm_container_bind_service` can be imported by using cluster_name_id, service_instance_name or service_instance_id and namespace_id.

**Example**

```
$ terraform import ibm_container_bind_service.example mycluster/myservice/default
```
