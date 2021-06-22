---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM : lbaas"
description: |-
  Attach server instances to IBM Cloud load balancer.
---

# ibm_lbaas_server_instance_attachment
Create, delete, and update to attach the server instance to IBM cloud load balancer. This allows attach, detach, and update server instances as LoadBalancer members to IBM Cloud Load Balancer. A `depends_on` statement is required for the associated load balancer to ensure that attach and detach occur after and before load balancer creation and deletion. If you do not specify the `depends_on` parameter, intermittent attach failures will occur on creation and load balancer deletion will fail. Typically when apply or destroy is rerun the operation will be successful. For more information, about attaching a service instance to IBM Cloud load balancer, see [selecting the service and configuring basic parameters](https://cloud.ibm.com/docs/loadbalancer-service?topic=loadbalancer-service-configuring-ibm-cloud-load-balancer-basic-parameters).


## Example usage

```terraform

resource "ibm_compute_vm_instance" "vm_instances" {
  count = "2"
  
}

resource "ibm_lbaas" "lbaas" {
  name        = "terraformLB"
  description = "delete this"
  subnets     = [1511875]

  protocols {
    frontend_protocol     = "HTTPS"
    frontend_port         = 443
    backend_protocol      = "HTTP"
    backend_port          = 80
    load_balancing_method = "round_robin"
    tls_certificate_id    = 11670
  }
  protocols {
    frontend_protocol     = "HTTP"
    frontend_port         = 80
    backend_protocol      = "HTTP"
    backend_port          = 80
    load_balancing_method = "round_robin"
  }
}

resource "ibm_lbaas_server_instance_attachment" "lbaas_member" {
  count = 2
  private_ip_address = element(
    ibm_compute_vm_instance.vm_instances.*.ipv4_address_private,
    count.index,
  )
  weight     = 40
  lbaas_id   = ibm_lbaas.lbaas.id
  depends_on = [ibm_lbaas.lbaas]
}

```

## Argument reference 
Review the argument references that you can specify for your resource.

- `depends_on`- (Required, String) The UUID of a load balancer.
- `lbaas_id`- (Required, Forces new resource, String) The UUID of a load balancer.
- `private_ip_address` - (Required, Forces new resource, String) The private IP address of a load balancer member.
- `weight` - (Optional, Integer) The weight of a load balancer member.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `uuid`- (String) The unique identifier of the load balancer member.
