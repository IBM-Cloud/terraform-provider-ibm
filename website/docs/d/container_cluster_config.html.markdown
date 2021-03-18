---
subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: ibm_container_cluster_config"
description: |-
  Get the cluster configuration for Kubernetes on IBM Cloud.
---

# ibm\_container_cluster_config


Download a configuration for Kubernetes clusters on IBM Cloud. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl
data "ibm_container_cluster_config" "cluster_foo" {
  cluster_name_id = "FOO"
  config_dir      = "/home/foo_config"
}
```
## Example Usage for connecting to kubernetes provider for classic or vpc kubernetes cluster with admin certificates
```hcl
data "ibm_container_cluster_config" "cluster_foo" {
  cluster_name_id = "FOO"
  admin           = true
}

provider "kubernetes" {
  load_config_file       = "false"
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
## Example Usage for connecting to kubernetes provider for classic or vpc kubernetes cluster with host and token
```hcl
data "ibm_container_cluster_config" "cluster_foo" {
  cluster_name_id = "FOO"
}

provider "kubernetes" {
  load_config_file       = "false"
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
## Example Usage for connecting to kubernetes provider for classic openshift cluster with admin certificates
```hcl
data "ibm_container_cluster_config" "cluster_foo" {
  cluster_name_id = "FOO"
  admin           = true
}

provider "kubernetes" {
  load_config_file       = "false"
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
## Example Usage for connecting to kubernetes provider for classic openshift cluster with host and token
```hcl
data "ibm_container_cluster_config" "cluster_foo" {
  cluster_name_id = "FOO"
}

provider "kubernetes" {
  load_config_file       = "false"
  host                   = data.ibm_container_cluster_config.cluster_foo.host
  token                  = data.ibm_container_cluster_config.cluster_foo.token
}

resource "kubernetes_namespace" "example" {
  metadata {
    name = "terraform-example-namespace"
  }
}
```

## Argument Reference

The following arguments are supported:

* `cluster_name_id` - (Required, string) The name or ID of the cluster.
* `config_dir` - (Required, string) The directory where you want the cluster configuration to download.
* `admin` - (Optional, boolean) Set the value to `true` to download the configuration for the administrator. The default value is `false`.
* `download` - (Optional, boolean) Set the value to `false` to skip downloading the configuration for the administrator. The default value is `true`. Because it is part of a data source, by default the configuration is downloaded for every Terraform call. For a particular cluster name or ID, the configuration is guaranteed to be downloaded to the same path for a given `config_dir`.
* `org_guid` - (Deprecated, string) The GUID for the IBM Cloud organization associated with the cluster. You can retrieve the value from the `ibm_org` data source or by running the `ibmcloud iam orgs --guid` command in the [IBM Cloud CLI](https://cloud.ibm.com/docs/cli?topic=cloud-cli-getting-started).
* `space_guid` - (Deprecated, string) The GUID for the IBM Cloud space associated with the cluster. You can retrieve the value from the `ibm_space` data source or by running the `ibmcloud iam space <space-name> --guid` command in the IBM Cloud CLI.
* `account_guid` - (Deprecated, string) The GUID for the IBM Cloud account associated with the cluster. You can retrieve the value from the `ibm_account` data source or by running the `ibmcloud iam accounts` command in the IBM Cloud CLI.
* `region` - (Deprecated, string) The region where the cluster is provisioned. If the region is not specified it will be defaulted to provider region(IC_REGION/IBMCLOUD_REGION). To get the list of supported regions please access this [link](https://containers.bluemix.net/v1/regions) and use the alias.
* `network` - (Optional, boolean) Set the value to `true` to download the configuration for the Calico network config with the Admin config. The default value is `false`.
* `resource_group_id` - (Optional, string) The ID of the resource group.  You can retrieve the value from data source `ibm_resource_group`. If not provided defaults to default resource group.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the cluster configuration.
* `admin_key`- (Sensitive) The admin key of the cluster configuration.
* `admin_certificate`- The admin certificate of the cluster configuration.
* `ca_certificate`- The cluster ca certificate of the cluster configuration.
* `host`- The Host of the cluster configuration.
* `token`- The token of the cluster configuration.
* `config_file_path` - The path to the cluster configuration file. This is typically the Kubernetes YAML configuration file.
* `calico_config_file_path` - The path to the cluster calico configuration file.