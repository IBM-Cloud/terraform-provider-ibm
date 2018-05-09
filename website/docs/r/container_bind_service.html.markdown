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

In the following example, you can bind a service to a cluster:

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

* `cluster_name_id` - (Required, string) The name or ID of the cluster.
* `service_instance_space_guid` - (Required, string) The space GUID associated with the service instance.
* `service_instance_name_id` - (Required, string) The name or ID of the service that you want to bind to the cluster.
* `namespace_id` - (Required, string) The Kubernetes namespace.
* `org_guid` - (Optional, string) The GUID for the IBM Cloud organization associated with the cluster. You can retrieve the value from data source `ibm_org` or by running the `bx iam orgs --guid` command in the [IBM Cloud CLI](https://console.bluemix.net/docs/cli/reference/bluemix_cli/get_started.html#getting-started).
* `space_guid` - (Optional, string) The GUID for the IBM Cloud space associated with the cluster. You can retrieve the value from data source `ibm_space` or by running the `bx iam space <space-name> --guid` command in the IBM Cloud CLI.
* `account_guid` - (Optional, string) The GUID for the IBM Cloud account associated with the cluster. You can retrieve the value from data source `ibm_account` or by running the `bx iam accounts` command in the IBM Cloud CLI.
* `tags` - (Optional, array of strings) Tags associated with the container bind service instance.
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

The following attributes are exported:

* `service_instance_name_id` - The name or ID of the service that is bound to the cluster.
* `namespace_id` -  The Kubernetes namespace.
* `space_guid` - The IBM Cloud space GUID.
* `secret_name` - The secret name.
