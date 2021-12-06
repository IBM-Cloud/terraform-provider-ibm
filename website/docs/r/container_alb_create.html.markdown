---

subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: container_alb_create"
description: |-
  Creates new IBM container Application Load Balancer.
---

# ibm_container_alb
Creates a new Ingres application load balancer (ALB) that is set up in your cluster. ALBs are used to set up HTTP or HTTPS load-balancing for containerized apps that are deployed into an IBM Cloud Kubernetes Service or Red Hat OpenShift on IBM Cloud cluster. For more information, about Ingress ALBs, see [about Ingress ALBs](https://cloud.ibm.com/docs/containers?topic=containers-ingress-about)

## Example usage

```terraform
resource "ibm_container_alb_create" "alb" {
  cluster="exampleClusterName"
  enable = "true"
  alb_type = "private"
  vlan_id = "123456"
  zone = "dal10"
}

```

## Timeouts

The `ibm_container_alb_create` provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 60 minutes) Used for creating Instance.
- **update** - (Default 60 minutes) Used for updating Instance.


## Argument reference
Review the argument references that you can specify for your resource. 

- `alb_type` - (String) The type of the ALB. Supported values are `public` and `private`.
- `cluster` - (String) The name of the cluster where the ALB is going to be created.
- `enable` - (Optional, Bool) If set to **true**, the default Ingress ALB in your cluster is enabled and configured with the IBM-provided Ingress subdomain. If set to **false**, the default Ingress ALB is enabled in your cluster.
- `ingress_image` - (Optional,ForceNew,String) The type of Ingress image that you want to use for your ALB deployment.
- `ip` - (Optional,String) The IP address that you want to assign to the ALB.
- `nlb_version` - (Optional,String) The version of the network load balancer that you want to use for the ALB.
- `vlan_id` - (String) ID of the cluster's vlan where the ALB is going to be created.
- `zone` - (String) The name of cluster's zone where the ALB is going to be created.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `alb_id` - (String) The unique identifier of the ALB. To retrieve the ID, run `ibmcloud ks alb ls`.
- `name` -  (String) The name of the ALB.
- `replicas` - (String) Desired number of ALB replicas. 
- `resize` -  (Bool) Indicate whether resizing should be done
- `user_ip` - (String) Private ALB only. The private ALB is deployed with an IP address from a user-provided private subnet.
- `disable_deployment` - (Optional, Bool) If set to **true**, the default Ingress ALB in your cluster is disabled. If set to **false**, the default Ingress ALB is enabled in your cluster and configured with the IBM-provided Ingress subdomain.