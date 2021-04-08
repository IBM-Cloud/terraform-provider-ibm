---
subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: container_bind_service"
description: |-
  Get information about existing service attached to IBM container cluster .
---

# ibm\_container_bind_service

Import the details of a service attached to IBM Cloud cluster as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

In the following example, you can get service info attached to a cluster:

```hcl
data "ibm_container_bind_service" "bind_service" {
  cluster_name_id       = "cluster_name"
  service_instance_name = "service_name"
  namespace_id          = "default"
}
```

## Argument Reference

The following arguments are supported:

* `cluster_name_id` - (Required, string) The name or ID of the cluster.
* `namespace_id` - (Required, string) The Kubernetes namespace.
* `service_instance_name` - (Optional, string) The name of the service that is attached to the cluster. Conflicts with `service_instance_id`.
* `service_instance_id` - (Optional, string) The ID of the service that is attached to the cluster. Conflicts with `service_instance_name`.


## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the bind service resource. The id is composed of \<cluster_name_id\>/\<service_instance_name or service_instance_id\>/\<namespace_id/>
* `service_key_name` - The service key name.
