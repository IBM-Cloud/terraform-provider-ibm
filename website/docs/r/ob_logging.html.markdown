---
subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: ibm_ob_logging"
description: |-
  Create a logging configuration for your cluster to automatically collect pod logs and send them to IBM Log Analysis.
---

# ibm_ob_logging
Create, update, or delete a logging instance. This resource creates a logging configuration for your cluster to automatically collect pod logs and send them to IBM Cloud Log Analysis. For more information, about Observability plug-in, see [managing  Observability logging commands](https://cloud.ibm.com/docs/containers?topic=containers-observability_cli).

## Example usage
In the following example enables you to create a logging configuration:

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
  name     = "TestLogging"
  service  = "logdna"
  plan     = "7-day"
  location = "us-south"
}

resource "ibm_resource_key" "resourceKey" {
  name                 = "TestKey"
  resource_instance_id = ibm_resource_instance.instance.id
  role                 = "Manager"
}

resource "ibm_ob_logging" "test2" {
  depends_on  = [ibm_resource_key.resourceKey]
  cluster     = ibm_container_cluster.testacc_cluster.id
  instance_id = ibm_resource_instance.instance.guid
}

```

## Timeouts

The `ibm_ob_logging` provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create**: The enablement of the feature is considered `failed` if no response is received for 10 minutes.
- **delete**: The delete of the feature is considered `failed` if no response is received for 10 minutes. 
- **update**: The update of the feature is considered `failed` if no response is received for 10 minutes. 


## Argument reference
Review the argument references that you can specify for your resource. 

- `cluster` - (Required, String) The name or ID of the cluster.
- `instance_id` - (Required, String) The GUID of the monitoring instance.
- `logdna_ingestion_key` - (Optional, String) The LogDNA ingestion key that you want to use for your configuration.
- `private_endpoint` - (Optional, String) Add this option to connect to your logging service instance through the private service endpoint.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `agent_key` - (String) The IBM Cloud Activity Tracker agent key.
- `agent_namespace` - (String) The IBM Cloud Activity Tracker agent namespace.
- `daemonset_name` - (String) The name of the `daemon` set
- `crn` - (String) The CRN of the IBM Cloud Activity Tracker instance attach.
- `id` - (String) The unique identifier of the logging instance attach to cluster. The ID is composed of `<cluster_name_id>/< logging_instance_id>`.
- `instance_name` - (String) Name of the logging instance.
- `namespace` - (String) The name of namespace.


## Import
The `ibm_ob_logging` can be imported by using `cluster_name_id`, `logging_instance_id`.

**Example**

```
$ terraform import ibm_ob_logging.example mycluster/5c4f4d06e0dc402084922dea70850e3b-7cafe35
```

