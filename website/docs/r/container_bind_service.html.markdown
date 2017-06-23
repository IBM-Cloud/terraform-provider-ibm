---
layout: "ibm"
page_title: "IBM: container_bind_service"
sidebar_current: "docs-ibm-resource-container-bind-service"
description: |-
  Manages IBM container cluster.
---

# ibm\_container_bind_service

Bind an IBM service to a Kubernetes namespace. With this resource, you can attach an existing service to an existing Kubernetes cluster. 

## Example Usage

In the following example, you can bind a service to a cluster.

```hcl
resource "ibm_container_bind_service" "bind_service" {
  cluster_name_id             = "cluster_name"
  service_instance_space_guid = "space_guid"
  service_instance_name_id    = "service_name"
  namespace_id                = "default"
  org_guid                    = "test"
  space_guid                  = "test_space"
  account_guid                = "test_account"
}
```

## Argument Reference

The following arguments are supported:

* `cluster_name_id` - (Required) Name or ID of the cluster.
* `service_instance_space_guid` - (Required) The space GUID the service instance is associated with.
* `service_instance_name_id` - (Required) The name or ID of the service that you want to bind to the cluster.
* `namespace_id` - (Required) The Kubernetes namespace.
* `org_guid` - (Required) The GUID for the Bluemix organization that the cluster is associated with. The values can be retrieved from data source `ibm_org`, or by running the `bx iam orgs --guid` command in the [Bluemix CLI](https://console.ng.bluemix.net/docs/cli/reference/bluemix_cli/index.html#getting-started).
* `space_guid` - (Required) The GUID for the Bluemix space that the cluster is associated with. The values can be retrieved from data source `ibm_space`, or by running the `bx iam space <space-name> --guid` command in the Bluemix CLI.
* `account_guid` - (Optional) The GUID for the Bluemix account that the cluster is associated with. The values can be retrieved from data source `ibm_account`, or by running the `bx iam accounts` command in the Bluemix CLI.
    
## Attributes Reference

The following attributes are exported:

* `service_instance_name_id` - The name or ID of the service that is bound to the cluster.
* `namespace_id` -  The Kubernetes namespace.
* `space_guid` - The Bluemix space GUID. 
* `secret_name` - The secret name.
