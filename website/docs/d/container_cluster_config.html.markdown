---
subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: ibm_container_cluster_config"
description: |-
  Get the cluster configuration for Kubernetes on IBM Cloud.
---

# ibm_container_cluster_config
Retrieve information about all the Kubernetes configuration files and certificates to access your cluster. For more information, about cluster configuration, see [accessing clusters](https://cloud.ibm.com/docs/containers?topic=containers-access_cluster).

If you plan to read a cluster that you also create with terraform and referencing its id, you may have to use wait_till field in the cluster resource with the value `Normal`.

## Example usage1

```terraform
data "ibm_container_cluster_config" "cluster_foo" {
  cluster_name_id = "FOO"
  config_dir      = "/home/foo_config"
}
```

## Example usage2
Example for connecting to Kubernetes provider for classic or VPC Kubernetes cluster with admin certificates

```terraform
data "ibm_container_cluster_config" "cluster_foo" {
  cluster_name_id = "FOO"
  admin           = true
}

provider "kubernetes" {
  host                   = data.ibm_container_cluster_config.cluster_foo.host
  client_certificate     = data.ibm_container_cluster_config.cluster_foo.admin_certificate
  client_key             = data.ibm_container_cluster_config.cluster_foo.admin_key
  cluster_ca_certificate = data.ibm_container_cluster_config.cluster_foo.ca_certificate
}

resource "kubernetes_namespace" "example" {
  metadata {
    name = "terraform-example-namespace"
  }
}
```
## Example usage3
Example for connecting to Kubernetes provider for classic or VPC Kubernetes cluster with host and token.

```terraform
data "ibm_container_cluster_config" "cluster_foo" {
  cluster_name_id = "FOO"
}

provider "kubernetes" {
  host                   = data.ibm_container_cluster_config.cluster_foo.host
  token                  = data.ibm_container_cluster_config.cluster_foo.token
  cluster_ca_certificate = data.ibm_container_cluster_config.cluster_foo.ca_certificate
}

resource "kubernetes_namespace" "example" {
  metadata {
    name = "terraform-example-namespace"
  }
}
```
## Example usage4
Example for connecting to Kubernetes provider for classic OpenShift cluster with admin certificates.

```terraform
data "ibm_container_cluster_config" "cluster_foo" {
  cluster_name_id = "FOO"
  admin           = true
}

provider "kubernetes" {
  host                   = data.ibm_container_cluster_config.cluster_foo.host
  client_certificate     = data.ibm_container_cluster_config.cluster_foo.admin_certificate
  client_key             = data.ibm_container_cluster_config.cluster_foo.admin_key
}

resource "kubernetes_namespace" "example" {
  metadata {
    name = "terraform-example-namespace"
  }
}
```
## Example usage5
Example usage for connecting to Kubernetes provider for classic OpenShift cluster with host and token.

```terraform
data "ibm_container_cluster_config" "cluster_foo" {
  cluster_name_id = "FOO"
}

provider "kubernetes" {
  host                   = data.ibm_container_cluster_config.cluster_foo.host
  token                  = data.ibm_container_cluster_config.cluster_foo.token
}

resource "kubernetes_namespace" "example" {
  metadata {
    name = "terraform-example-namespace"
  }
}
```

## Example usage6
Example for getting kubeconfig for VPC Kubernetes cluster with admin certificates and with VPE Gateway as server URL

```terraform
data "ibm_container_cluster_config" "cluster_foo" {
  cluster_name_id = "FOO"
  config_dir      = "/home/foo_config"
  admint          = "true"
  endpoint_type   = "vpe"
}
```


## Argument reference
Review the argument references that you can specify for your data source. 

- `admin` - (Optional, Bool) If set to **true**, the Kubernetes configuration for cluster administrators is downloaded. The default is **false**.
- `cluster_name_id` - (Required, String) The name or ID of the cluster that you want to log in to. 
- `config_dir` - (Required, String) The directory on your local machine where you want to download the Kubernetes config files and certificates.
- `download` - (Optional, Bool) Set the value to **false** to skip downloading the configuration for the administrator. The default value is **true**. The configuration files and certificates are downloaded to the directory that you specified in `config_dir` every time that you run your infrastructure code.
- `network` - (Optional, Bool) If set to **true**, the Calico configuration file, TLS certificates, and permission files that are required to run `calicoctl` commands in your cluster are downloaded in addition to the configuration files for the administrator. The default value is **false**. 
- `resource_group_id` - (Optional, String) The ID of the resource group where your cluster is provisioned into. To find the resource group, run `ibmcloud resource groups` or use the `ibm_resource_group` data source. If this parameter is not provided, the `default` resource group is used.
- `endpoint_type` - (Optional, String) The server URL for the cluster context. If you do not include this parameter, the default cluster service endpoint is used. Available options: `private`, `link` (Satellite), `vpe` (VPC). For Satellite clusters, the `link` endpoint is the default. When the public service endpoint is disabled in Red Hat OpenShift on IBM Cloud clusters, the `endpoint_type` parameter will also influence the communication method used by the provider plugin with the cluster when generating the cluster config. If you set it to `private`, the plugin will utilize the cluster's Private Service Endpoint URL for communication, while setting it to `vpe` will make it use the cluster's Virtual Private Endpoint gateway URL for communication purposes.

**Deprecated reference**

- `account_guid` - (Deprecated, String) The GUID for the IBM Cloud account associated with the cluster. You can retrieve the value from the `ibm_account` data source or by running the `ibmcloud iam accounts` command in the IBM Cloud CLI.
- `org_guid` - (Deprecated, String) The GUID for the IBM Cloud organization associated with the cluster. You can retrieve the value from the `ibm_org` data source or by running the `ibmcloud iam orgs --guid` command in the [IBM Cloud CLI](https://cloud.ibm.com/docs/cli?topic=cloud-cli-getting-started).
- `region` - (Deprecated, String) The region where the cluster is provisioned. If the region is not specified it will be defaulted to provider region (IC_REGION/IBMCLOUD_REGION). To get the list of supported regions please access this [link](https://containers.bluemix.net/v1/regions) and use the alias.
- `space_guid` - (Deprecated, String) The GUID for the IBM Cloud space associated with the cluster. You can retrieve the value from the `ibm_space` data source or by running the `ibmcloud iam space <space-name> --guid` command in the IBM Cloud CLI.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `calico_config_file_path` - (String) The path on your local machine where your Calico configuration files and certificates are downloaded to.
- `config_file_path` - (String) The path on your local machine where the cluster configuration file and certificates are downloaded to. 
- `id` - (String) The unique identifier of the cluster configuration.
- `admin_key` - (String) The admin key of the cluster configuration. Note that this key is case-sensitive.
- `admin_certificate` - (String) The admin certificate of the cluster configuration.
- `ca_certificate` - (String) The cluster CA certificate of the cluster configuration.
- `host` - (String) The host name of the cluster configuration.
- `token` - (String) The token of the cluster configuration.
