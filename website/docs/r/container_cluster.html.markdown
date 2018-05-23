---
layout: "ibm"
page_title: "IBM: container_cluster"
sidebar_current: "docs-ibm-resource-container-cluster"
description: |-
  Manages IBM container cluster.
---

# ibm\_container_cluster

Create, update, or delete a Kubernetes cluster. An existing subnet can be attached to the cluster by passing the subnet ID. A webhook can be registered to a cluster, and you can add multiple worker nodes with the `workers` option.
During the creation of cluster the workers are created with the kube version by default. During update user need to specify the version of worker to update the worker.

**Before updating the version of cluster and workers via terraform get the list of available updates and their pre and post update instructions at https://console.bluemix.net/docs/containers/cs_versions.html#version_types. Also please go through the instructions at https://console.bluemix.net/docs/containers/cs_cluster_update.html#update.
_Users must read these docs carefully before updating the version via terraform_.**

## Example Usage

In the following example, you can create a Kubernetes cluster:

```hcl
resource "ibm_container_cluster" "testacc_cluster" {
  name            = "test"
  datacenter      = "dal10"
  machine_type    = "free"
  isolation       = "public"
  public_vlan_id  = "vlan"
  private_vlan_id = "vlan"
  subnet_id       = ["1154643"]

  workers = [{
    name = "worker1"
    action = "add"
  }]

  webhook = [{
    level = "Normal"
    type = "slack"
    url = "https://hooks.slack.com/services/yt7rebjhgh2r4rd44fjk"
  }]

  org_guid     = "test"
  space_guid   = "test_space"
  account_guid = "test_acc"
}
```

Create the Kubernetes cluster using worker_num:

```hcl
resource "ibm_container_cluster" "testacc_cluster" {
  name            = "test"
  datacenter      = "dal10"
  machine_type    = "free"
  isolation       = "public"
  public_vlan_id  = "vlan"
  private_vlan_id = "vlan"
  subnet_id       = ["1154643"]

  worker_num = 2
  webhook = [{
    level = "Normal"
    type = "slack"
    url = "https://hooks.slack.com/services/yt7rebjhgh2r4rd44fjk"
  }]

  account_guid = "test_acc"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the cluster.
* `datacenter` - (Required, string)  The datacenter of the worker nodes. You can retrieve the value by running the `bluemix cs locations` command in the [IBM Cloud CLI](https://console.bluemix.net/docs/cli/reference/bluemix_cli/get_started.html#getting-started).
* `kube_version` - (Optional, string) The desired Kubernetes version of the created cluster. If present, at least major.minor must be specified.
* `org_guid` - (Optional, string) The GUID for the IBM Cloud organization associated with the cluster. You can retrieve the value from data source `ibm_org` or by running the `bx iam orgs --guid` command in the IBM Cloud CLI.
* `space_guid` - (Optional, string) The GUID for the IBM Cloud space associated with the cluster. You can retrieve the value from data source `ibm_space` or by running the `bx iam space <space-name> --guid` command in the IBM Cloud CLI.
* `account_guid` - (Required, string) The GUID for the IBM Cloud account associated with the cluster. You can retrieve the value from data source `ibm_account` or by running the `bx iam accounts` command in the IBM Cloud CLI.
* `workers` - (Optional, array) The worker nodes that you want to add to the cluster. Nested `workers` blocks have the following structure:
	* `action` - valid actions are add, reboot and reload.
	* `name` - Name of the worker.
	* `version` - worker version.
	**NOTE**: Conflicts with `worker_num`. 
* `worker_num` - (Optional, int)  The number of cluster worker nodes.
	**NOTE**: Conflicts with `workers`. 
* `machinetype` - (Optional, string) The machine type of the worker nodes. You can retrieve the value by running the `bx cs machine-types <data-center>` command in the IBM Cloud CLI.
* `billing` - (Optional, string) The billing type for the instance. Accepted values are `hourly` or `monthly`.
* `isolation` - (Optional, string) Accepted values are `public` or `private`. Use `private` if you want to have available physical resources dedicated to you only or `public` to allow physical resources to be shared with other IBM customers.
* `public_vlan_id`- (Optional, string) The public VLAN of the worker node. You can retrieve the value by running the `bx cs vlans <data-center>` command in the IBM Cloud CLI.
* `private_vlan_id` - (Optional, string) The private VLAN of the worker node. You can retrieve the value by running the `bx cs vlans <data-center>` command in the IBM Cloud CLI.
* `subnet_id` - (Optional, string) The existing subnet ID that you want to add to the cluster. You can retrieve the value by running the `bx cs subnets` command in the IBM Cloud CLI.
* `no_subnet` - (Optional, boolean) Set to `true` if you do not want to automatically create a portable subnet.
* `disk_encryption` - (Optional, boolean) Set to `false` to disable encryption on a worker.
* `webhook` - (Optional, string) The webhook that you want to add to the cluster.
* `wait_time_minutes` - (Optional, integer) The duration, expressed in minutes, to wait for the cluster to become available before declaring it as created. It is also the same amount of time waited for no active transactions before proceeding with an update or deletion. The default value is `90`.
* `tags` - (Optional, array of strings) Tags associated with the container cluster instance.
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the cluster.
* `name` - The name of the cluster.
* `server_url` - The server URL.
* `ingress_hostname` - The Ingress hostname.
* `ingress_secret` - The Ingress secret.
* `workers_info` - The worker nodes attached to this cluster.
* `subnet_id` - The subnets attached to this cluster.
* `workers` -  Exported attributes are:
	* `id` - The id of the worker