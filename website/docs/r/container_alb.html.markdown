---

subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: container_alb"
description: |-
  Manages IBM container Application Load Balancer.
---

# ibm_container_alb
Enable or disable an Ingres application load balancer (ALB) that is set up in your cluster. ALBs are used to set up HTTP or HTTPS load-balancing for containerized apps that are deployed into an IBM Cloud Kubernetes Service or Red Hat OpenShift on IBM Cloud cluster. For more information, about Ingress ALBs, see [about Ingress ALBs](https://cloud.ibm.com/docs/containers?topic=containers-ingress-about)

## Example usage

```terraform
resource "ibm_container_alb" "alb" {
  alb_id = "public-cr083d810e501d4c73b42184eab5a7ad56-alb"
  enable = true
}

```

## Timeouts

The `ibm_container_alb` provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 60 minutes) Used for creating Instance.
- **update** - (Default 60 minutes) Used for updating Instance.


## Argument reference
Review the argument references that you can specify for your resource. 
  
- `alb_id` - (Required, Forces new resource, String) The unique identifier of the ALB. To retrieve the ID, run `ibmcloud ks alb ls`.
- `disable_deployment` - (Optional, Forces new resource, Bool) If set to **true**, the default Ingress ALB in your cluster is disabled. If set to **false**, the default Ingress ALB is enabled in your cluster and configured with the IBM-provided Ingress subdomain. If you do not specify this option, you must specify the `enable` parameter.
- `enable` - (Optional, Bool) If set to **true**, the default Ingress ALB in your cluster is enabled and configured with the IBM-provided Ingress subdomain. If set to **false**, the default Ingress ALB is disabled in your cluster. If you do not specify this option, you must specify the `disable_deployment` parameter.
- `region` - (Optional, Forces new resource, String) The region where the Ingress ALB is provisioned.
- `user_ip` - (Optional, Forces new resource, String) For a private ALB only. The private ALB is deployed with an IP address from a user-provided private subnet. If no IP address is provided, the ALB is deployed with a random IP address from a private subnet in the IBM Cloud account.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `alb_type` - (String) The type of the ALB. Supported values are `public` and `private`.
- `cluster` - (String) The name of the cluster where the ALB is provisioned.
- `id` - (String) The unique identifier of the ALB. 
- `name` -  (String) The name of the ALB.
