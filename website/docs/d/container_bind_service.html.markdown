---
subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: container_bind_service"
description: |-
  Get information about existing service attached to IBM container cluster .
---

# ibm_container_bind_service
Retrieve information of a service attached to IBM Cloud cluster. For more information, about service binding, see [Adding services by using IBM Cloud service binding](https://cloud.ibm.com/docs/containers?topic=containers-service-binding).

## Example usage
The following example retrieves service information attached to a cluster.

```terraform
data "ibm_container_bind_service" "bind_service" {
  cluster_name_id       = "cluster_name"
  service_instance_name = "service_name"
  namespace_id          = "default"
}
```

## Argument reference
Review the argument references that you can specify for your data source.
 
- `cluster_name_id` - (Required, String) The name or ID of the cluster.
- `namespace_id` - (Required, String) The Kubernetes namespace.
- `service_instance_name` - (Optional, String) The name of the service that is attached to the cluster. This conflicts with the `service_instance_id` parameter.
- `service_instance_id` - (Optional, String) The ID of the service that is attached to the cluster. This conflicts with the `service_instance_name` parameter.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique Id of the bind service resource. The ID is composed of `<cluster_name_id>/<service_instance_name or service_instance_id>/<namespace_id/>`.
- `service_key_name` - (String) The service key name.
