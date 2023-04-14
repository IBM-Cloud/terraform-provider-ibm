---
subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: ibm_container_ingress_instance"
description: |-
  Get information about a registered Secrets Manager instance
---

# ibm_container_ingress_instance
Get details about a registered IBM Cloud Secrets Manager instance for your cluster.


## Example usage
The following example retrieves information about the registered Secrets Manager instance that is named `myinstance` of a cluster that is named `mycluster`. 

```terraform
data ibm_container_ingress_instance instance {
    cluster ="mycluster"
    instance_name = "myinstance"
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `cluster` - (Required, String) The name or ID of the cluster.
- `instance_name` - (Required, string) The name of the instance registration.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 
- `instance_crn` - (String) The unique identifier of the instance.
- `secret_group_name` -  (String) The name of the secret group if set.
- `secret_group_id` -  (String) The ID of the secret group if set.
- `is_default` - (Bool) Indicates whether the instance is the registered default for the cluster.
- `status` - (String) Indicates the status of the instance registration. 
- `instance_type` -  (String) Indicate whether the instance is of type Secrets Manager
- `user_managed` - (Bool) Indicates the user created and registered the instance.