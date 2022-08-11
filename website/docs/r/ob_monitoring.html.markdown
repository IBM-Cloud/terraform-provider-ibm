---
subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: ibm_ob_monitoring"
description: |-
  Create a monitoring configuration for your cluster to automatically collect pod logs and send them to IBM Log Analysis.
---

# ibm_ob_monitoring
Create, update, or delete a monitoring instance. This resource creates a monitoring configuration for your IBM Cloud Kubernetes Service cluster to automatically collect pod metrics, and send these metrics to your monitoring service instance. For more information, about Observability monitoring, see [setting up loggin with IBM Cloud Log Analysis](https://cloud.ibm.com/docs/containers?topic=containers-istio-health).

## Example usage
The following example enables you to create a monitoring configuration. 

```terraform

data "ibm_resource_group" "testacc_ds_resource_group" {
  name = "Default"
}

resource "ibm_container_cluster" "testacc_cluster" {
  name              = "TestCluster"
  datacenter        = "dal10"
  resource_group_id = data.ibm_resource_group.testacc_ds_resource_group.id
  default_pool_size = 1
  wait_till         = "MasterNodeReady"
  hardware          = "shared"
  machine_type      = "%s"
  timeouts {
    create = "720m"
    update = "720m"
  }
}

resource "ibm_resource_instance" "instance" {
  name     = "TestMonitoring"
  service  = "sysdig-monitor"
  plan     = "graduated-tier"
  location = "us-south"
}

resource "ibm_resource_key" "resourceKey" {
  name                 = "TestKey"
  resource_instance_id = ibm_resource_instance.instance.id
  role                 = "Manager"
}

resource "ibm_ob_monitoring" "test2" {
  depends_on  = [ibm_resource_key.resourceKey]
  cluster     = ibm_container_cluster.testacc_cluster.id
  instance_id = ibm_resource_instance.instance.guid
}

```

## Timeouts

ibm_ob_monitoring provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create**: The enablement of the feature is considered `failed` if no response is received for 45 minutes.
- **delete**: The delete of the feature is considered `failed` if no response is received for 10 minutes. 
- **update**: The update of the feature is considered `failed` if no response is received for 10 minutes. 


## Argument reference
Review the argument references that you can specify for your resource. 

- `cluster` - (Required, String) The name or ID of the cluster.
- `instance_id` - (Required, String) The GUID of the monitoring instance.
- `sysdig_access_key` - (Optional, String) The monitoring ingestion key that you want to use for your configuration.
- `private_endpoint`- (Optional, String)  Add this option to connect to your monitoring service instance through the private service endpoint.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `agent_key` - (String) The IBM Cloud Activity Tracker agent key.
- `agent_namespace` - (String) The IBM Cloud Activity Tracker agent namespace.
- `crn` - (String) The CRN of the instance attach.
- `daemonset_name` - Name of the deamonset.
- `discovered_agent` - (String) The name of the discovered agent.
- `id` - (String) The unique identifier of the logging instance attach to cluster. The ID is composed of `<cluster_name_id>/< logging_instance_id>`.
- `instance_name` - (String) The name of the logging instance.
- `namespace` - (String) The name of namespace.

## Import
The `ibm_ob_monitoring` can be imported by using `cluster_name_id`, `monitoring_instance_id`.

**Example**

```
$ terraform import ibm_ob_monitoring.example mycluster/5c4f4d06e0dc402084922dea70850e3b-7cafe35
```
