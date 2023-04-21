---
subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: ibm_container_ingress_instance"
description: |-
  Registers an IBM Cloud Secrets Manager instance with your cluster
---

# ibm_container_ingress_instance
Registers an IBM Cloud Secrets Manager instance with your IBM Cloud Kubernetes Service or Red Hat OpenShift on IBM Cloud cluster. For more information about how this instance can be used see [about Secrets Manager](https://cloud.ibm.com/docs/containers?topic=containers-secrets-mgr)

## Example usage

```terraform
resource "ibm_container_ingress_instance" "instance" {
  cluster="exampleClusterName"
  instance_crn = var.sm_instance_crn
  is_default = true
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `instance_crn` - (String) The CRN of the IBM Cloud Secrets Manager instance you want to register.
- `cluster` - (String) The name of the cluster where the ALB is going to be created.
- `secret_group_id` - (Optional, string) If set, the default ingress certificates for the instance will be uploaded into this secret group in the instance.
- `is_default` - (Optional, bool) Marks the instance as the default instance for your cluster. The default ingress certificates will be uploaded to the default instance.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `instance_name` - (String) The unique identifier of the instance. To retrieve run `ibmcloud ks ingress instance ls` for your cluster
- `secret_group_name` -  (String) The name of the secret group if set.
- `status` - (String) Indicates the status of the instance registration. 
- `instance_type` -  (String) Indicate whether the instance is of type Secrets Manager
- `user_managed` - (Bool) Indicates the user created and registered the instance.